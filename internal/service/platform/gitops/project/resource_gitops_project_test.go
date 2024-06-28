package project_test

import (
	"fmt"
	"testing"

	"strings"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsRepositoryOrgLevel(t *testing.T) {

	// Org level
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := "account.rollouts"
	resourceName := "harness_platform_gitops_repository.test"
	accountId := "1bvyLackQK-Hapk25-Ry4w"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepositoryOrgLevel(id, name, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsAgentOrgLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccResourceGitopsRepositoryOrgLevel(id string, name string, agentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_gitops_project" "test" {
			account_id = "%[4]s"
			agent_id = "%[3]s"
			project {
				metadata {
					generation = 1
					name = "14a3dc9eee"
					namespace = "rollouts"
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "*"
					}
					destinations {
						namespace = "*"
						server = "*"
					}
					source_repos = ["*"]
				}
			}
		}
	`, id, name, agentId, accountId)
}

// func TestResourceCreate(t *testing.T) {
// 	// Mock API key and test data
// 	apiKey := "YOUR_API_KEY_HERE"
// 	testData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project": map[string]interface{}{
// 			"metadata": map[string]interface{}{
// 				"name":      "test-project",
// 				"namespace": "default",
// 			},
// 			"spec": map[string]interface{}{
// 				"clusterResourceWhitelist": []interface{}{
// 					map[string]interface{}{
// 						"group": "apps",
// 						"kind":  "Deployment",
// 					},
// 				},
// 				"destinations": []interface{}{
// 					map[string]interface{}{
// 						"namespace": "prod",
// 						"server":    "prod-server",
// 					},
// 				},
// 				"sourceRepos": []interface{}{
// 					"git@github.com:test/repo.git",
// 				},
// 			},
// 		},
// 	}

// 	// Create a new instance of the resource
// 	resource := resourceCreateProject()

// 	// Mock HTTP server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Verify request method and headers
// 		if r.Method != "POST" {
// 			t.Errorf("Expected POST request, got %s", r.Method)
// 		}
// 		if r.Header.Get("Content-Type") != "application/json" {
// 			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
// 		}
// 		if r.Header.Get("x-api-key") != apiKey {
// 			t.Errorf("Expected x-api-key header to be %s, got %s", apiKey, r.Header.Get("x-api-key"))
// 		}

// 		// Verify URL path and query parameters
// 		expectedURL := fmt.Sprintf("/gitops/api/v1/agents/%s/projects?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
// 			testData["agent_identifier"], testData["account_identifier"], testData["org_identifier"], testData["project_identifier"])
// 		if r.URL.String() != expectedURL {
// 			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
// 		}

// 		// Read request body and verify project data
// 		var requestBody map[string]interface{}
// 		err := json.NewDecoder(r.Body).Decode(&requestBody)
// 		if err != nil {
// 			t.Fatalf("Error decoding request body: %s", err)
// 		}
// 		if !isEqual(requestBody, testData["project"]) {
// 			t.Errorf("Expected request body %+v, got %+v", testData["project"], requestBody)
// 		}

// 		// Write mock response
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"id": "test-project-id"}`))
// 	}))

// 	defer server.Close()

// 	// Replace base URL in resource data with mock server URL
// 	resourceData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project":            testData["project"],
// 	}
// 	resourceData["project"].(map[string]interface{})["metadata"].(map[string]interface{})["generation"] = 1 // Adjust generation for create

// 	d := schema.TestResourceDataRaw(t, resource.Schema, resourceData)

// 	// Configure client with mock server URL
// 	client := &http.Client{Transport: &http.Transport{}}

// 	// Set the resourceCreateProject function with our mocked client
// 	err := resourceProjectCreate(d, client)

// 	// Check if there were any errors
// 	if err != nil {
// 		t.Fatalf("Error: %s", err)
// 	}

// 	// Check if the resource id has been set
// 	if d.Id() == "" {
// 		t.Fatal("Expected resource id to be set")
// 	}

// 	// Verify resource id is set properly
// 	expectedId := "test-project-id"
// 	if d.Id() != expectedId {
// 		t.Fatalf("Expected resource id %s, got %s", expectedId, d.Id())
// 	}
// }

// func TestResourceRead(t *testing.T) {
// 	// Mock API key and test data
// 	apiKey := "YOUR_API_KEY_HERE"
// 	testData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"query_name":         "test-project-id",
// 		"project": map[string]interface{}{
// 			"metadata": map[string]interface{}{
// 				"name":      "test-project",
// 				"namespace": "default",
// 			},
// 			"spec": map[string]interface{}{},
// 		},
// 	}

// 	// Create a new instance of the resource
// 	resource := resourceCreateProject()

// 	// Mock HTTP server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Verify request method and headers
// 		if r.Method != "GET" {
// 			t.Errorf("Expected GET request, got %s", r.Method)
// 		}
// 		if r.Header.Get("x-api-key") != apiKey {
// 			t.Errorf("Expected x-api-key header to be %s, got %s", apiKey, r.Header.Get("x-api-key"))
// 		}

// 		// Verify URL path and query parameters
// 		expectedURL := fmt.Sprintf("/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
// 			testData["agent_identifier"], testData["query_name"], testData["account_identifier"], testData["org_identifier"], testData["project_identifier"])
// 		if r.URL.String() != expectedURL {
// 			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
// 		}

// 		// Write mock response
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"id": "test-project-id", "name": "test-project", "namespace": "default"}`))
// 	}))

// 	defer server.Close()

// 	// Replace base URL in resource data with mock server URL
// 	resourceData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"query_name":         "test-project-id",
// 		"project":            testData["project"],
// 	}
// 	d := schema.TestResourceDataRaw(t, resource.Schema, resourceData)

// 	// Configure client with mock server URL
// 	client := &http.Client{Transport: &http.Transport{}}

// 	// Set the resourceProjectRead function with our mocked client
// 	err := resourceProjectRead(d, client)

// 	// Check if there were any errors
// 	if err != nil {
// 		t.Fatalf("Error: %s", err)
// 	}

// 	// Verify resource id is set properly
// 	expectedId := "test-agent/test-project-id"
// 	if d.Id() != expectedId {
// 		t.Fatalf("Expected resource id %s, got %s", expectedId, d.Id())
// 	}
// }

// func TestResourceUpdate(t *testing.T) {
// 	// Mock API key and test data
// 	apiKey := "YOUR_API_KEY_HERE"
// 	testData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project": map[string]interface{}{
// 			"metadata": map[string]interface{}{
// 				"name":      "test-project",
// 				"namespace": "default",
// 			},
// 			"spec": map[string]interface{}{
// 				"clusterResourceWhitelist": []interface{}{
// 					map[string]interface{}{
// 						"group": "apps",
// 						"kind":  "Deployment",
// 					},
// 				},
// 				"destinations": []interface{}{
// 					map[string]interface{}{
// 						"namespace": "prod",
// 						"server":    "prod-server",
// 					},
// 				},
// 				"sourceRepos": []interface{}{
// 					"git@github.com:test/repo.git",
// 				},
// 			},
// 		},
// 	}

// 	// Create a new instance of the resource
// 	resource := resourceCreateProject()

// 	// Mock HTTP server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Verify request method and headers
// 		if r.Method != "PUT" {
// 			t.Errorf("Expected PUT request, got %s", r.Method)
// 		}
// 		if r.Header.Get("Content-Type") != "application/json" {
// 			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
// 		}
// 		if r.Header.Get("x-api-key") != apiKey {
// 			t.Errorf("Expected x-api-key header to be %s, got %s", apiKey, r.Header.Get("x-api-key"))
// 		}

// 		// Verify URL path and query parameters
// 		expectedURL := fmt.Sprintf("/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
// 			testData["agent_identifier"], testData["project.metadata.name"], testData["account_identifier"], testData["org_identifier"], testData["project_identifier"])
// 		if r.URL.String() != expectedURL {
// 			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
// 		}

// 		// Read request body and verify project data
// 		var requestBody map[string]interface{}
// 		err := json.NewDecoder(r.Body).Decode(&requestBody)
// 		if err != nil {
// 			t.Fatalf("Error decoding request body: %s", err)
// 		}
// 		if !isEqual(requestBody, testData["project"]) {
// 			t.Errorf("Expected request body %+v, got %+v", testData["project"], requestBody)
// 		}

// 		// Write mock response
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"id": "test-project-id"}`))
// 	}))

// 	defer server.Close()

// 	// Replace base URL in resource data with mock server URL
// 	resourceData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project":            testData["project"],
// 	}
// 	resourceData["project"].(map[string]interface{})["metadata"].(map[string]interface{})["generation"] = 1 // Adjust generation for update

// 	d := schema.TestResourceDataRaw(t, resource.Schema, resourceData)

// 	// Configure client with mock server URL
// 	client := &http.Client{Transport: &http.Transport{}}

// 	// Set the resourceProjectUpdate function with our mocked client
// 	err := resourceProjectUpdate(d, client)

// 	// Check if there were any errors
// 	if err != nil {
// 		t.Fatalf("Error: %s", err)
// 	}

// 	// Verify resource id is set properly
// 	expectedId := "test-project-id"
// 	if d.Id() != expectedId {
// 		t.Fatalf("Expected resource id %s, got %s", expectedId, d.Id())
// 	}
// }

// func TestResourceDelete(t *testing.T) {
// 	// Mock API key and test data
// 	apiKey := "YOUR_API_KEY_HERE"
// 	testData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project": map[string]interface{}{
// 			"metadata": map[string]interface{}{
// 				"name":      "test-project",
// 				"namespace": "default",
// 			},
// 			"spec": map[string]interface{}{},
// 		},
// 	}

// 	// Create a new instance of the resource
// 	resource := resourceCreateProject()

// 	// Mock HTTP server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Verify request method and headers
// 		if r.Method != "DELETE" {
// 			t.Errorf("Expected DELETE request, got %s", r.Method)
// 		}
// 		if r.Header.Get("x-api-key") != apiKey {
// 			t.Errorf("Expected x-api-key header to be %s, got %s", apiKey, r.Header.Get("x-api-key"))
// 		}

// 		// Verify URL path and query parameters
// 		expectedURL := fmt.Sprintf("/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
// 			testData["agent_identifier"], testData["project.metadata.name"], testData["account_identifier"], testData["org_identifier"], testData["project_identifier"])
// 		if r.URL.String() != expectedURL {
// 			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
// 		}

// 		// Write mock response
// 		w.WriteHeader(http.StatusNoContent)
// 	}))

// 	defer server.Close()

// 	// Replace base URL in resource data with mock server URL
// 	resourceData := map[string]interface{}{
// 		"agent_identifier":   "test-agent",
// 		"account_identifier": "test-account",
// 		"org_identifier":     "test-org",
// 		"project_identifier": "test-project",
// 		"project":            testData["project"],
// 	}
// 	d := schema.TestResourceDataRaw(t, resource.Schema, resourceData)

// 	// Configure client with mock server URL
// 	client := &http.Client{Transport: &http.Transport{}}

// 	// Set the resourceProjectDelete function with our mocked client
// 	err := resourceProjectDelete(d, client)

// 	// Check if there were any errors
// 	if err != nil {
// 		t.Fatalf("Error: %s", err)
// 	}

// 	// Verify resource id is cleared after deletion
// 	if d.Id() != "" {
// 		t.Fatalf("Expected resource id to be empty after deletion, got %s", d.Id())
// 	}
// }

// // Utility function to check equality of two maps
// func isEqual(map1, map2 map[string]interface{}) bool {
// 	if len(map1) != len(map2) {
// 		return false
// 	}
// 	for k, v1 := range map1 {
// 		if v2, ok := map2[k]; !ok || v1 != v2 {
// 			return false
// 		}
// 	}
// 	return true
// }
