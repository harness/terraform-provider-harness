//go:build e2e

package registry

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// E2E test for remote_url_suffix against a live Harness QA environment.
// Run with:
//
//	HARNESS_ENDPOINT=https://qa.harness.io/gateway \
//	HARNESS_ACCOUNT_ID=<account> \
//	HARNESS_PLATFORM_API_KEY=<pat> \
//	go test -tags=e2e -run TestE2EUpstreamPythonRegistryRemoteUrlSuffix -v ./internal/service/har/registry/...
func TestE2EUpstreamPythonRegistryRemoteUrlSuffix(t *testing.T) {
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	endpoint := os.Getenv("HARNESS_ENDPOINT")
	if accountID == "" || apiKey == "" || endpoint == "" {
		t.Skip("set HARNESS_ACCOUNT_ID, HARNESS_PLATFORM_API_KEY, HARNESS_ENDPOINT to run")
	}

	org := envOr("HARNESS_ORG_ID", "default")
	project := envOr("HARNESS_PROJECT_ID", "jatin_test")
	spaceRef := fmt.Sprintf("%s/%s/%s", accountID, org, project)
	registryID := fmt.Sprintf("tf_e2e_tfprov_%d", time.Now().Unix())

	clientCfg := har.NewConfiguration()
	clientCfg.AccountId = accountID
	clientCfg.ApiKey = apiKey
	clientCfg.BasePath = endpoint + "/har/api/v1"
	clientCfg.BasePathV3 = endpoint + "/har/api/v3"
	clientCfg.DefaultHeader = map[string]string{"X-Api-Key": apiKey}
	client := har.NewAPIClient(clientCfg)
	ctx := context.WithValue(context.Background(), har.ContextAPIKey, har.APIKey{Key: apiKey})

	d := schema.TestResourceDataRaw(t, resourceRegistrySchema(false), map[string]interface{}{
		"identifier":   registryID,
		"space_ref":    spaceRef,
		"parent_ref":   spaceRef,
		"package_type": "PYTHON",
		"is_public":    false,
		"config": []interface{}{map[string]interface{}{
			"type":              "UPSTREAM",
			"auth_type":         "Anonymous",
			"source":            "Custom",
			"url":               "https://pypi.example.com",
			"remote_url_suffix": "simple",
		}},
	})

	t.Cleanup(func() {
		registryRef := spaceRef + "/" + registryID
		_, _, _ = client.RegistriesApi.DeleteRegistry(ctx, registryRef)
	})

	// CREATE via provider buildRegistry -> HAR API
	registry := buildRegistry(d)
	resp, httpResp, err := client.RegistriesApi.CreateRegistry(ctx, &har.RegistriesApiCreateRegistryOpts{
		Body:     optional.NewInterface(registry),
		SpaceRef: optional.NewString(spaceRef),
	})
	if err != nil {
		t.Fatalf("create failed: %v (status=%v)", err, httpRespStatus(httpResp))
	}
	if resp.Data == nil || resp.Data.Config == nil || resp.Data.Config.RemoteUrlSuffix != "simple" {
		t.Fatalf("create: expected remoteUrlSuffix=simple, got %+v", resp.Data)
	}

	// READ via provider readRegistry
	d.SetId(registryID)
	readRegistry(d, resp.Data)
	cfg := d.Get("config").([]interface{})[0].(map[string]interface{})
	if cfg["remote_url_suffix"] != "simple" {
		t.Fatalf("readRegistry after create: expected simple, got %v", cfg["remote_url_suffix"])
	}

	// UPDATE
	d.Set("config", []interface{}{map[string]interface{}{
		"type":              "UPSTREAM",
		"auth_type":         "Anonymous",
		"source":            "Custom",
		"url":               "https://pypi.example.com",
		"remote_url_suffix": "custom-index",
	}})
	registry = buildRegistry(d)
	registryRef := spaceRef + "/" + registryID
	modResp, httpResp, err := client.RegistriesApi.ModifyRegistry(ctx, registryRef, &har.RegistriesApiModifyRegistryOpts{
		Body: optional.NewInterface(registry),
	})
	if err != nil {
		t.Fatalf("update failed: %v (status=%v)", err, httpRespStatus(httpResp))
	}
	readRegistry(d, modResp.Data)
	cfg = d.Get("config").([]interface{})[0].(map[string]interface{})
	if cfg["remote_url_suffix"] != "custom-index" {
		t.Fatalf("readRegistry after update: expected custom-index, got %v", cfg["remote_url_suffix"])
	}

	// IMPORT simulation: fresh state, read from API
	imported := schema.TestResourceDataRaw(t, resourceRegistrySchema(false), map[string]interface{}{
		"identifier": registryID,
		"space_ref":  spaceRef,
		"parent_ref": spaceRef,
	})
	getResp, httpResp, err := client.RegistriesApi.GetRegistry(ctx, registryRef)
	if err != nil {
		t.Fatalf("get failed: %v (status=%v)", err, httpRespStatus(httpResp))
	}
	readRegistry(imported, getResp.Data)
	cfg = imported.Get("config").([]interface{})[0].(map[string]interface{})
	if cfg["remote_url_suffix"] != "custom-index" {
		t.Fatalf("import read: expected custom-index, got %v", cfg["remote_url_suffix"])
	}

	t.Logf("E2E passed for registry %s with remote_url_suffix round-trip", registryID)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func httpRespStatus(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}
