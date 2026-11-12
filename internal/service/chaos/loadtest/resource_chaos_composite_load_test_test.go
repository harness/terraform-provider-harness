package loadtest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestCompositeIdentifierMatchesPipelineRule(t *testing.T) {
	if err := compositeIdentifierError("checkout"); err != nil {
		t.Fatal(err)
	}
	if err := compositeIdentifierError("_checkout$1"); err != nil {
		t.Fatal(err)
	}
	if err := compositeIdentifierError("a" + strings.Repeat("b", 118)); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"1checkout", "checkout-v2", strings.Repeat("a", 120)} {
		if err := compositeIdentifierError(id); err == nil {
			t.Fatalf("%s was accepted", id)
		}
	}
}

func TestChaosCompositeLoadTestLifecycle(t *testing.T) {
	store := &compositeStore{}
	server := httptest.NewServer(store.handler(t))
	defer server.Close()

	config := compositeConfig("checkout_k6")
	updated := compositeConfig("browse_k6")
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: compositeProviderFactories(server),
		CheckDestroy: func(*terraform.State) error {
			if !store.saw("DELETE") {
				return fmt.Errorf("destroy did not call delete")
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("harness_chaos_composite_load_test.checkout", "id", "checkout"),
					resource.TestCheckResourceAttr("harness_chaos_composite_load_test.checkout", "pipeline_identifier", "checkout_pipeline"),
					resource.TestCheckResourceAttr("harness_chaos_composite_load_test.checkout", "load_test_ref", "checkout_k6"),
					resource.TestCheckResourceAttr("harness_chaos_composite_load_test.checkout", "probe.0.identity", "checkout_probe"),
				),
			},
			{Config: config, PlanOnly: true, ExpectNonEmptyPlan: false},
			{
				Config: updated,
				Check:  resource.TestCheckResourceAttr("harness_chaos_composite_load_test.checkout", "load_test_ref", "browse_k6"),
			},
		},
	})
	if !store.saw("POST /v1/composite-load-tests") {
		t.Fatal("create did not hit /v1/composite-load-tests")
	}
	if !store.saw("PUT") {
		t.Fatal("changing load_test_ref did not update")
	}
}

func TestCompositePlaceholdersDoNotPlan(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPut {
			t.Errorf("placeholder values caused an update")
		}
		if r.Method == http.MethodPost {
			_, _ = w.Write([]byte(`{"identifier":"checkout","stageIdentifier":"checkout","pipelineIdentifier":"checkout_pipeline","name":"Checkout"}`))
			return
		}
		_, _ = w.Write([]byte(`{"identifier":"checkout","pipelineIdentifier":"checkout_pipeline","name":"Checkout","loadTest":{"loadTestRef":"checkout_k6"},"probe":{"identity":"checkout_probe"}}`))
	}))
	defer server.Close()
	config := `
provider "harness" {}
resource "harness_chaos_composite_load_test" "checkout" {
  org_id        = "default"
  project_id    = "demo"
  identifier    = "checkout"
  name          = "Checkout"
  load_test_ref = "checkout_k6"
  probe {
    identity        = "checkout_probe"
    infra_reference = "<+input>"
    duration        = "10m"
  }
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: compositeProviderFactories(server),
		Steps: []resource.TestStep{
			{Config: config},
			{Config: config, PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestChaosCompositeLoadTestReadNotFoundDropsResource(t *testing.T) {
	store := &compositeStore{}
	server := httptest.NewServer(store.handler(t))
	defer server.Close()

	config := compositeConfig("checkout_k6")
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: compositeProviderFactories(server),
		Steps: []resource.TestStep{
			{Config: config},
			{
				PreConfig:          func() { store.missing = true },
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func compositeConfig(loadTestRef string) string {
	return `
provider "harness" {}
resource "harness_chaos_composite_load_test" "checkout" {
  org_id        = "default"
  project_id    = "demo"
  identifier    = "checkout"
  name          = "Checkout"
  description   = "checkout path"
  objective     = "saturate checkout"
  load_test_ref = "` + loadTestRef + `"
  tags = {
    team = "payments"
  }
  probe {
    identity        = "checkout_probe"
    infra_reference = "infra1"
    duration        = "5m"
  }
}
`
}

type compositeStore struct {
	mu      sync.Mutex
	missing bool
	body    map[string]interface{}
	calls   []string
}

func (s *compositeStore) saw(part string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, call := range s.calls {
		if strings.Contains(call, part) {
			return true
		}
	}
	return false
}

func (s *compositeStore) handler(t *testing.T) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path
		if !strings.Contains(path, "/v1/composite-load-tests") {
			t.Errorf("unexpected path %s", path)
			http.NotFound(w, r)
			return
		}
		item := strings.Contains(path, "/v1/composite-load-tests/")
		s.mu.Lock()
		defer s.mu.Unlock()
		switch {
		case r.Method == http.MethodPost && !item:
			s.calls = append(s.calls, "POST /v1/composite-load-tests")
			raw, _ := io.ReadAll(r.Body)
			var body map[string]interface{}
			_ = json.Unmarshal(raw, &body)
			s.body = body
			_, _ = w.Write([]byte(`{"identifier":"checkout","stageIdentifier":"checkout","pipelineIdentifier":"checkout_pipeline","name":"Checkout"}`))
		case r.Method == http.MethodGet && item:
			s.calls = append(s.calls, "GET "+path)
			if s.missing {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"message":"not found"}`))
				return
			}
			_, _ = w.Write(s.getBody())
		case r.Method == http.MethodPut && item:
			s.calls = append(s.calls, "PUT "+path)
			raw, _ := io.ReadAll(r.Body)
			var body map[string]interface{}
			_ = json.Unmarshal(raw, &body)
			s.body = body
			_, _ = w.Write(s.getBody())
		case r.Method == http.MethodDelete && item:
			s.calls = append(s.calls, "DELETE "+path)
			s.body = nil
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			t.Errorf("unexpected %s %s", r.Method, path)
			http.NotFound(w, r)
		}
	}
}

func (s *compositeStore) getBody() []byte {
	loadTestRef := "checkout_k6"
	name := "Checkout"
	description := "checkout path"
	objective := "saturate checkout"
	probeIdentity := "checkout_probe"
	infra := "infra1"
	duration := "5m"
	if s.body != nil {
		if value, ok := s.body["name"].(string); ok && value != "" {
			name = value
		}
		if value, ok := s.body["description"].(string); ok {
			description = value
		}
		if value, ok := s.body["objective"].(string); ok {
			objective = value
		}
		if loadTest, ok := s.body["loadTest"].(map[string]interface{}); ok {
			if value, ok := loadTest["loadTestRef"].(string); ok && value != "" {
				loadTestRef = value
			}
		}
		if probe, ok := s.body["probe"].(map[string]interface{}); ok {
			if value, ok := probe["identity"].(string); ok && value != "" {
				probeIdentity = value
			}
			if value, ok := probe["infraReference"].(string); ok {
				infra = value
			}
			if value, ok := probe["duration"].(string); ok {
				duration = value
			}
		}
	}
	raw, _ := json.Marshal(map[string]interface{}{
		"identifier":         "checkout",
		"pipelineIdentifier": "checkout_pipeline",
		"name":               name,
		"description":        description,
		"objective":          objective,
		"tags":               map[string]string{"team": "payments"},
		"loadTest":           map[string]string{"loadTestRef": loadTestRef},
		"probe": map[string]string{
			"identity":       probeIdentity,
			"infraReference": infra,
			"duration":       duration,
		},
	})
	return raw
}

func compositeProviderFactories(server *httptest.Server) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"harness": func() (*schema.Provider, error) {
			client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{
				BasePath:   server.URL + "/loadTest/manager/api",
				AccountId:  "acc",
				HTTPClient: server.Client(),
			})
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"harness_chaos_composite_load_test": ResourceChaosCompositeLoadTest(),
				},
				DataSourcesMap: map[string]*schema.Resource{
					"harness_chaos_composite_load_test": DataSourceChaosCompositeLoadTest(),
				},
				ConfigureContextFunc: func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
					return &internal.Session{LoadTestClient: client}, nil
				},
			}, nil
		},
	}
}
