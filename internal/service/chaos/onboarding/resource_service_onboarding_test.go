package onboarding

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestServiceOnboardingSchemaValid(t *testing.T) {
	resource := ResourceChaosServiceOnboarding()
	if err := resource.InternalValidate(resource.Schema, true); err != nil {
		t.Fatal(err)
	}
}

func TestServiceOnboardingCreateSelectedServices(t *testing.T) {
	var created chaos.V2OnboardingV2OnboardingRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("accountIdentifier") != "acc" || r.URL.Query().Get("projectIdentifier") != "proj" {
			t.Errorf("scope = %s", r.URL.RawQuery)
		}
		if onboardingPrelude(w, r) {
			return
		}
		if r.Method != http.MethodPost || r.URL.Path != "/rest/v2/onboarding" {
			t.Fatalf("create must not continue into discovery: %s %s", r.Method, r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&created); err != nil {
			t.Fatal(err)
		}
		writeOnboarding(w, "ob-1", "QUEUED", "", "")
	}))
	defer server.Close()

	d := onboardingData(t, map[string]interface{}{
		"service_ids": []interface{}{"svc-b", "svc-a"},
	})
	diags := resourceServiceOnboardingCreate(t.Context(), d, onboardingSession(server, server.URL))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if created.EnvironmentRef != "prod" || created.InfrastructureRef != "payments" {
		t.Fatalf("refs = %s %s", created.EnvironmentRef, created.InfrastructureRef)
	}
	if created.Mode == nil || *created.Mode != chaos.MANUAL_V2OnboardingOnboardingMode ||
		created.ManualMethod == nil || *created.ManualMethod != chaos.SERVICE_V2OnboardingManualMethod {
		t.Fatalf("mode = %v method = %v", created.Mode, created.ManualMethod)
	}
	if d.Id() != "ob-1" || d.Get("install_command") == "" {
		t.Fatalf("id=%s install=%q", d.Id(), d.Get("install_command"))
	}
}

func TestServiceOnboardingUpdateOnboardsAllDiscoveredServices(t *testing.T) {
	var onboarded []string
	confirmed := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "", "payments")
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/payments/discoveredservices":
			if r.URL.Query().Get("all") != "true" || r.URL.Query().Get("environmentIdentifier") != "prod" {
				t.Fatalf("query = %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"items":[{"id":"svc-a"},{"id":"gone","removed":true},{"id":""}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding-onboard-services/ob-1":
			var body struct {
				ExternalServiceIDs []string `json:"externalServiceIds"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			onboarded = body.ExternalServiceIDs
			_, _ = w.Write([]byte(`{"onboardedCount":1,"onboardedServiceIds":["svc-a"]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/rest/v2/onboarding-confirm-service-onboarding/ob-1":
			confirmed = true
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/services"):
			_, _ = w.Write([]byte(`{"services":[{"externalServiceId":"svc-a"}],"pagination":{"totalPages":1}}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingUpdate(t.Context(), d, onboardingSession(server, server.URL+"/chaos/manager/api"))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !confirmed || len(onboarded) != 1 || onboarded[0] != "svc-a" || d.Get("onboarded_count") != 1 {
		t.Fatalf("confirmed=%v onboarded=%v count=%v", confirmed, onboarded, d.Get("onboarded_count"))
	}
}

func TestServiceOnboardingCreateDiscoveryErroredKeepsID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if onboardingPrelude(w, r) {
			return
		}
		if r.Method != http.MethodGet || r.URL.Path != "/rest/v2/onboarding/ob-1" {
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
		writeOnboarding(w, "ob-1", "ERRORED", "", "")
	}))
	defer server.Close()

	d := onboardingData(t, map[string]interface{}{"service_ids": []interface{}{"svc-a"}})
	d.SetId("ob-1")
	diags := resourceServiceOnboardingUpdate(t.Context(), d, onboardingSession(server, server.URL))
	if !diags.HasError() || !strings.Contains(diagText(diags), "discovery for onboarding ob-1 is ERRORED") {
		t.Fatalf("diags = %#v", diags)
	}
	if d.Id() != "ob-1" {
		t.Fatalf("id=%s diags=%s", d.Id(), diagText(diags))
	}
}

func TestServiceOnboardingCreateNoDiscoveredServices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if onboardingPrelude(w, r) {
			return
		}
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/rest/v2/onboarding"):
			writeOnboarding(w, "ob-1", "COMPLETED", "", "agent-1")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/onboarding/ob-1"):
			writeOnboarding(w, "ob-1", "COMPLETED", "", "agent-1")
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/discoveredservices"):
			_, _ = w.Write([]byte(`{"items":[{"id":"gone","removed":true}]}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingUpdate(t.Context(), d, onboardingSession(server, server.URL+"/chaos/manager/api"))
	if !diags.HasError() || !strings.Contains(diagText(diags), "no services") {
		t.Fatalf("diags = %#v", diags)
	}
	if d.Id() != "ob-1" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestServiceOnboardingReadRejectsOtherMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := map[string]interface{}{
			"onboardingID": "ob-1", "manualMethod": "NetworkMap", "status": "PENDING",
		}
		_ = json.NewEncoder(w).Encode(body)
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingRead(t.Context(), d, onboardingSession(server, server.URL))
	if !diags.HasError() || !strings.Contains(diagText(diags), "NetworkMap") {
		t.Fatalf("diags = %#v", diags)
	}
}

func TestServiceOnboardingReadNotFoundClearsState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"not found"}`))
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingRead(t.Context(), d, onboardingSession(server, server.URL))
	if diags.HasError() || d.Id() != "" {
		t.Fatalf("id=%q diags=%#v", d.Id(), diags)
	}
}

func TestServiceOnboardingReadMissingSessionClearsState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"mongo: no documents in result"}`))
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingRead(t.Context(), d, onboardingSession(server, server.URL))
	if diags.HasError() || d.Id() != "" {
		t.Fatalf("id=%q diags=%#v", d.Id(), diags)
	}
}

func TestServiceOnboardingReadServerErrorKeepsState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"mongo down"}`))
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingRead(t.Context(), d, onboardingSession(server, server.URL))
	if !diags.HasError() || d.Id() != "ob-1" {
		t.Fatalf("id=%q diags=%#v", d.Id(), diags)
	}
}

func TestServiceOnboardingUpdateOnboardsNewSelection(t *testing.T) {
	var onboarded []string
	confirmed := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/v2/onboarding/ob-1":
			writeOnboarding(w, "ob-1", "COMPLETED", "COMPLETED", "agent-1")
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding-onboard-services/ob-1":
			var body struct {
				ExternalServiceIDs []string `json:"externalServiceIds"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			onboarded = body.ExternalServiceIDs
			_, _ = w.Write([]byte(`{"onboardedCount":1,"onboardedServiceIds":["svc-c"]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/rest/v2/onboarding-confirm-service-onboarding/ob-1":
			confirmed = true
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/services"):
			_, _ = w.Write([]byte(`{"services":[{"externalServiceId":"svc-c"}],"pagination":{"totalPages":1}}`))
		case r.Method == http.MethodGet && strings.Contains(r.URL.Path, "/rest/v2/infrastructure/"):
			_, _ = w.Write([]byte(`{"identity":"payments","infraNamespace":"chaos","serviceAccount":"litmus","environmentID":"prod"}`))
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/infrastructure/yaml"):
			_, _ = w.Write([]byte(`{"yaml":"apiVersion: v1\nkind: List\n"}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := onboardingData(t, map[string]interface{}{"service_ids": []interface{}{"svc-a"}})
	d.SetId("ob-1")
	if err := d.Set("service_ids", []interface{}{"svc-c"}); err != nil {
		t.Fatal(err)
	}
	if !d.HasChange("service_ids") {
		t.Fatal("expected service_ids change")
	}
	diags := resourceServiceOnboardingUpdate(t.Context(), d, onboardingSession(server, server.URL))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !confirmed || len(onboarded) != 1 || onboarded[0] != "svc-c" {
		t.Fatalf("confirmed=%v onboarded=%v", confirmed, onboarded)
	}
}

func TestServiceOnboardingDeleteMarksSessionRemoved(t *testing.T) {
	deleted := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/rest/v2/onboarding/ob-1" {
			t.Fatalf("destroy called %s %s", r.Method, r.URL.Path)
		}
		deleted = true
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	d := onboardingData(t, nil)
	d.SetId("ob-1")
	diags := resourceServiceOnboardingDelete(t.Context(), d, onboardingSession(server, server.URL))
	if !deleted || diags.HasError() || d.Id() != "" {
		t.Fatalf("deleted=%v id=%q diags=%#v", deleted, d.Id(), diags)
	}
}

func TestServiceOnboardingImport(t *testing.T) {
	d := onboardingData(t, nil)
	d.SetId("org/proj/ob-1")
	got, err := resourceServiceOnboardingImport(t.Context(), d, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Id() != "ob-1" || got[0].Get("org_id") != "org" || got[0].Get("project_id") != "proj" {
		t.Fatalf("imported = %#v", got[0].State())
	}

	d.SetId("org/ob-1")
	if _, err := resourceServiceOnboardingImport(t.Context(), d, nil); err == nil {
		t.Fatal("expected invalid import id")
	}
}

func onboardingData(t *testing.T, extra map[string]interface{}) *schema.ResourceData {
	t.Helper()
	raw := map[string]interface{}{
		"org_id":            "org",
		"project_id":        "proj",
		"environment_id":    "prod",
		"infrastructure_id": "payments",
	}
	for key, value := range extra {
		raw[key] = value
	}
	return schema.TestResourceDataRaw(t, ResourceChaosServiceOnboarding().Schema, raw)
}

func onboardingSession(server *httptest.Server, base string) *internal.Session {
	client := chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   base,
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	return &internal.Session{ChaosClient: client}
}

func onboardingPrelude(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path
	switch {
	case r.Method == http.MethodGet && strings.HasSuffix(path, "/rest/v2/onboarding"):
		_, _ = w.Write([]byte(`{"data":[]}`))
	case r.Method == http.MethodGet && strings.Contains(path, "/rest/v2/infrastructure/"):
		_, _ = w.Write([]byte(`{"identity":"payments","infraNamespace":"chaos","serviceAccount":"litmus","environmentID":"prod"}`))
	case r.Method == http.MethodPost && strings.HasSuffix(path, "/rest/v2/infrastructure/yaml"):
		_, _ = w.Write([]byte(`{"yaml":"apiVersion: v1\nkind: List\n"}`))
	case r.Method == http.MethodPost && strings.Contains(path, "/onboarding-confirm-discovery/"):
		_, _ = w.Write([]byte(`{"discoveryCompleted":true}`))
	default:
		return false
	}
	return true
}

func writeOnboarding(w http.ResponseWriter, id, discovery, service, agent string) {
	body := map[string]interface{}{
		"onboardingID":      id,
		"environmentRef":    "prod",
		"infrastructureRef": "payments",
		"mode":              "Manual",
		"manualMethod":      "Service",
		"status":            "PENDING",
		"discoveredServiceStatus": map[string]string{
			"status":          discovery,
			"sDAgentIdentity": agent,
		},
		"onboardingChaosServiceStatus": map[string]interface{}{
			"status":            service,
			"chaosServiceCount": 1,
		},
	}
	_ = json.NewEncoder(w).Encode(body)
}

func diagText(diags diag.Diagnostics) string {
	parts := make([]string, 0, len(diags))
	for _, item := range diags {
		parts = append(parts, item.Summary)
		if item.Detail != "" {
			parts = append(parts, item.Detail)
		}
	}
	return strings.Join(parts, " ")
}
