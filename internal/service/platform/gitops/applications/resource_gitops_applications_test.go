package applications_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"encoding/json"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
)

func TestAccResourceGitopsApplication_AllTypes(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	repo := os.Getenv("HARNESS_TEST_GITOPS_REPO")
	helmRepoURL := os.Getenv("HARNESS_TEST_GITOPS_HELM_REPO_URL")
	helmChart := os.Getenv("HARNESS_TEST_GITOPS_HELM_REPO_CHART")
	namespace := "test"
	namespaceUpdated := namespace + "_updated"

	// Resource names for each application type
	helmAppResource := "harness_platform_gitops_applications.helm_app"
	kustomizeAppResource := "harness_platform_gitops_applications.kustomize_app"
	multiSourceAppResource := "harness_platform_gitops_applications.multisource_app"
	helmChartAppResource := "harness_platform_gitops_applications.helm_chart_app"
	skipRepoAppResource := "harness_platform_gitops_applications.skip_repo_app"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccResourceGitopsApplicationDestroy(helmAppResource),
			testAccResourceGitopsApplicationDestroy(kustomizeAppResource),
			testAccResourceGitopsApplicationDestroy(multiSourceAppResource),
			testAccResourceGitopsApplicationDestroy(helmChartAppResource),
			testAccResourceGitopsApplicationDestroy(skipRepoAppResource),
		),
		Steps: []resource.TestStep{
			{
				// Step 1: Create all applications
				Config: testAccResourceGitopsApplicationAllTypes(id, accountId, agentId, clusterServer, clusterId, clusterToken, repo, helmRepoURL, helmChart, namespace),
				Check: resource.ComposeTestCheckFunc(
					// Helm app checks
					resource.TestCheckResourceAttr(helmAppResource, "id", id+"helm"),
					resource.TestCheckResourceAttr(helmAppResource, "application.0.spec.0.destination.0.namespace", namespace),
					// Kustomize app checks
					resource.TestCheckResourceAttr(kustomizeAppResource, "id", id+"kustomize"),
					resource.TestCheckResourceAttr(kustomizeAppResource, "application.0.spec.0.destination.0.namespace", namespace),
					// Multi-source app checks
					resource.TestCheckResourceAttr(multiSourceAppResource, "name", id+"multisource"),
					resource.TestCheckResourceAttr(multiSourceAppResource, "application.0.spec.0.destination.0.namespace", namespace),
					// Helm chart app checks (skip_repo_validation)
					resource.TestCheckResourceAttr(helmChartAppResource, "id", id+"helmchart"),
					resource.TestCheckResourceAttr(helmChartAppResource, "skip_repo_validation", "true"),
					// Skip repo validation app checks
					resource.TestCheckResourceAttr(skipRepoAppResource, "id", id+"skiprepo"),
					resource.TestCheckResourceAttr(skipRepoAppResource, "skip_repo_validation", "true"),
				),
			},
			{
				// Step 2: Update namespace on all applications
				Config: testAccResourceGitopsApplicationAllTypes(id, accountId, agentId, clusterServer, clusterId, clusterToken, repo, helmRepoURL, helmChart, namespaceUpdated),
				Check: resource.ComposeTestCheckFunc(
					// Verify namespace updated on all apps
					resource.TestCheckResourceAttr(helmAppResource, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
					resource.TestCheckResourceAttr(kustomizeAppResource, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
					resource.TestCheckResourceAttr(multiSourceAppResource, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
					resource.TestCheckResourceAttr(helmChartAppResource, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
					resource.TestCheckResourceAttr(skipRepoAppResource, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
				),
			},
			{
				// Step 3: Import state for helm app
				ResourceName:      helmAppResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(helmAppResource),
			},
			{
				// Step 4: Import state for kustomize app
				ResourceName:      kustomizeAppResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(kustomizeAppResource),
			},
			{
				// Step 5: Import state for multi-source app
				ResourceName:      multiSourceAppResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentProjectLevelResourceImportStateIdFunc(multiSourceAppResource),
			},
		},
	})
}

func TestAccResourceGitopsApplication_DetectDrift(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	appName := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	repoId := id
	clusterName := id
	namespace := "test"
	repo := "https://github.com/harness-apps/hosted-gitops-example-apps"
	resourceName := "harness_platform_gitops_applications.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Step 1: Create the GitOps application
				Config: testAccResourceGitopsApplicationHelm(id, accountId, appName, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("%s not found in state", resourceName)
						}
						t.Logf("Application created with ID: %s", rs.Primary.ID)

						return nil
					},
				),
			},
			{
				// Step 2: Modify the application externally via API call to add a new label
				PreConfig: func() {
					// Allow some time for resource to be fully created
					time.Sleep(2 * time.Second)

					// Log what we're about to do
					t.Logf("Making API call to modify application %s", id)

					// Get the API client from the test acctest utilities
					c, ctx := acctest.TestAccGetPlatformClientWithContext()

					// First, get the current application to modify it
					resp, _, err := c.ApplicationsApiService.AgentApplicationServiceGet(
						ctx,
						agentId,   // agent identifier
						id,        // application name/id
						accountId, // account identifier
						id,        // org identifier (using same id)
						id,        // project identifier (using same id)
						&nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
							QueryRepo: optional.NewString(repoId),
						},
					)

					if err != nil {
						t.Logf("Failed to get application: %s", err)
						return
					}

					t.Logf("Successfully retrieved application")

					// Get the application to check labels BEFORE update
					log.Println("==== GITOPS DRIFT TEST: Checking application labels BEFORE update...")
					var beforeLabels map[string]string
					if resp.App != nil && resp.App.Metadata != nil && resp.App.Metadata.Labels != nil {
						beforeLabels = resp.App.Metadata.Labels
						labelsJSON, _ := json.Marshal(beforeLabels)
						log.Printf("==== GITOPS DRIFT TEST: BEFORE Labels: %s", string(labelsJSON))
					} else {
						log.Println("==== GITOPS DRIFT TEST: BEFORE: No labels present in application metadata")
						beforeLabels = make(map[string]string)
					}

					// Verify the drift-detection-test label is NOT present before update
					if _, exists := beforeLabels["drift-detection-test"]; exists {
						t.Fatalf("ERROR: drift-detection-test label already exists before update!")
					} else {
						log.Println("==== GITOPS DRIFT TEST: VERIFIED: drift-detection-test label does not exist before update")
					}

					// Add the drift detection label to the application metadata
					if resp.App != nil && resp.App.Metadata != nil {
						if resp.App.Metadata.Labels == nil {
							resp.App.Metadata.Labels = make(map[string]string)
						}
						// Add the label
						resp.App.Metadata.Labels["drift-detection-test"] = "true"
						log.Println("==== GITOPS DRIFT TEST: ADDED LABEL: Added 'drift-detection-test' label to application metadata")

						// Dump the modified metadata structure for debugging
						labelsJSON, _ := json.Marshal(resp.App.Metadata.Labels)
						log.Printf("==== GITOPS DRIFT TEST: LABELS AFTER MODIFICATION: %s", string(labelsJSON))
					}

					// Create the update request with the modified application
					// Make sure we're properly including the labels
					metadata := resp.App.Metadata
					spec := resp.App.Spec

					// Log the metadata before creating the request
					metadataJSON, _ := json.Marshal(metadata)
					fmt.Println("METADATA JSON:", string(metadataJSON))

					updateRequest := nextgen.ApplicationsApplicationUpdateRequest{
						// Create a new ApplicationsApplication with the proper fields
						Application: &nextgen.ApplicationsApplication{
							// Only use the fields that exist in the ApplicationsApplication struct
							Metadata: metadata,
							Spec:     spec,
						},
						Validate: false,
					}

					// Log the complete update request
					requestJSON, _ := json.Marshal(updateRequest)
					fmt.Println("REQUEST:", string(requestJSON))

					// Call the update API
					updatedResp, httpResp, err := c.ApplicationsApiService.AgentApplicationServiceUpdate(
						ctx,
						updateRequest,
						accountId, // account identifier
						id,        // org identifier
						id,        // project identifier
						agentId,   // agent identifier
						id,        // application name
						&nextgen.ApplicationsApiAgentApplicationServiceUpdateOpts{
							ClusterIdentifier: optional.NewString(clusterId),
							RepoIdentifier:    optional.NewString(repoId),
						},
					)

					if err != nil {
						t.Logf("Failed to update application: %s", err)
						if httpResp != nil {
							t.Logf("HTTP Response: %d", httpResp.StatusCode)
							responseBody, _ := io.ReadAll(httpResp.Body)
							t.Logf("Response body: %s", string(responseBody))
						}
						return
					}

					t.Logf("Successfully updated application")

					// Log the complete response for debugging
					responseJSON, _ := json.Marshal(updatedResp)
					fmt.Println("RESPONSE:", string(responseJSON))

					// Verify the update worked by explicitly checking for the label
					if updatedResp.App != nil && updatedResp.App.Metadata != nil && updatedResp.App.Metadata.Labels != nil {
						// Get labels after update
						afterLabels := updatedResp.App.Metadata.Labels
						labelsJSON, _ := json.Marshal(afterLabels)
						log.Printf("==== GITOPS DRIFT TEST: AFTER Labels: %s", string(labelsJSON))

						// Check if our label was added
						if val, ok := afterLabels["drift-detection-test"]; ok {
							log.Printf("==== GITOPS DRIFT TEST: SUCCESS: Label successfully added with value: %s", val)

							// Compare before and after to confirm the change
							log.Printf("==== GITOPS DRIFT TEST: DRIFT CONFIRMED: Label 'drift-detection-test' was added via API outside of Terraform")
						} else {
							log.Printf("==== GITOPS DRIFT TEST: WARNING: Label doesn't appear in the response! Labels received: %v", afterLabels)
						}
					} else {
						log.Printf("==== GITOPS DRIFT TEST: WARNING: Cannot verify label in response - missing metadata or labels")
						// Log the structure of the response for debugging
						if updatedResp.App == nil {
							log.Println("==== GITOPS DRIFT TEST: Response App is nil")
						} else if updatedResp.App.Metadata == nil {
							log.Println("==== GITOPS DRIFT TEST: Response App.Metadata is nil")
						} else if updatedResp.App.Metadata.Labels == nil {
							log.Println("==== GITOPS DRIFT TEST: Response App.Metadata.Labels is nil")
						}
					}

					// Verify the change using testAccGetApplication
					s := terraform.State{
						Modules: []*terraform.ModuleState{
							{
								Path: []string{"root"},
								Resources: map[string]*terraform.ResourceState{
									resourceName: {
										Type: "harness_platform_gitops_applications",
										Primary: &terraform.InstanceState{
											ID: id,
											Attributes: map[string]string{
												"id":         id,
												"identifier": id,
											},
										},
									},
								},
							},
						},
					}

					app, err := testAccGetApplication(resourceName, &s)
					if err != nil {
						t.Logf("Warning: Failed to get application: %s", err)
					} else if app != nil {
						t.Logf("Application after update: %v", app)

						// Check if our label was added
						if app.App.Metadata != nil && app.App.Metadata.Labels != nil {
							if _, ok := app.App.Metadata.Labels["drift-detection-test"]; ok {
								t.Logf("Label successfully added")
							} else {
								t.Logf("Warning: Label doesn't appear to be added")
							}
						}
					}
				},
				// Step 3: Verify TF detects change - check plan only without applying
				PlanOnly:           true,
				Config:             testAccResourceGitopsApplicationHelm(id, accountId, appName, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				ExpectNonEmptyPlan: true, // This indicates that Terraform detected drift
			},
			{
				// Step 4: Apply TF and see if operation succeeded
				Config: testAccResourceGitopsApplicationHelm(id, accountId, appName, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
		},
	})
}

func testAccResourceGitopsApplicationDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		application, _ := testAccGetApplication(resourceName, state)
		if application != nil {
			return fmt.Errorf("Found Application: %s", application.Name)
		}
		return nil
	}
}

func testAccGetApplication(resourceName string, state *terraform.State) (*nextgen.Servicev1Application, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	if r == nil {
		return nil, nil
	}
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_id"]
	orgIdentifier := r.Primary.Attributes["org_id"]
	projectIdentifier := r.Primary.Attributes["project_id"]
	queryName := r.Primary.Attributes["name"]
	repoIdentifier := r.Primary.Attributes["repo_id"]

	resp, _, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRepo: optional.NewString(repoIdentifier),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsApplicationHelm(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "https://github.com/harness-apps/hosted-gitops-example-apps"
					name = "%[1]s"
					insecure = true
					connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_repository.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "helm-guestbook"
						
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id = "%[8]s"
			repo_id = "%[10]s"
			agent_id = "%[4]s"
			name = "%[3]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId)
}

func testAccResourceGitopsApplicationHelmMultiSource(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string, clusterToken string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "https://github.com/harness-apps/hosted-gitops-example-apps"
					name = "%[1]s"
					insecure = true
					connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[8]s"
			account_id = "%[2]s"
			agent_id = "%[4]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			request {
				upsert = true
				cluster {
					server = "%[7]s"
					name = "%[8]s"
					config {
						bearer_token = "%[11]s"
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "SERVICE_ACCOUNT"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info,
				]
			}
		}

		resource "harness_platform_gitops_applications" "testmultisource" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test, harness_platform_service.test, harness_platform_environment.test, harness_platform_project.test, harness_platform_organization.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					sources {
							repo_url = "%[9]s"
							target_revision = "master"
							ref = "val"
					}
					sources {
							repo_url = "%[9]s"
							target_revision = "master"
							path = "helm-guestbook"
							helm {
							  value_files = [
								"$val/helm-guestbook/values.yaml"
							  ]
							}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			name = "%[3]s"
			cluster_id = "%[8]s"
			repo_ids = [
				"%[1]s",
				"%[1]s"
			]
			agent_id = "%[4]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId, clusterToken)
}

func testAccResourceGitopsApplicationHelmSkipRepoValidation(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repoURL string, chart string, skipRepoValidation bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_gitops_applications" "test" {
			application {
				metadata {
					annotations = {}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "18.0.1"
						repo_url = "%[9]s"
						chart = "%[10]s"
						
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id = "%[8]s"
			agent_id = "%[4]s"
			name = "%[3]s"
            skip_repo_validation = %[11]t
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repoURL, chart, skipRepoValidation)
}

func testAccResourceGitopsApplicationKustomize(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_service" "test" {
      		identifier = "%[1]s"
      		name = "%[3]s"
      		org_id = harness_platform_project.test.org_id
      		project_id = harness_platform_project.test.id
    	}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  		}
		
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "%[9]s"
        			name = "%[1]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_repository.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "kustomize-guestbook"
						kustomize {
							images = [
									"gcr.io/heptio-images/ks-guestbook-demo:0.1"
									]
						}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id =  "%[8]s"
			repo_id = "%[10]s"
			agent_id = "%[4]s"
			name = "%[3]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId)
}

func testAccResourceGitopsApplicationGitSkipRepoValidation(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, skipRepoValidation bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_service" "test" {
      		identifier = "%[1]s"
      		name = "%[3]s"
      		org_id = harness_platform_project.test.org_id
      		project_id = harness_platform_project.test.id
    	}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[3]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  		}
		

		resource "harness_platform_gitops_applications" "test" {
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]s"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[9]s"
						path = "kustomize-guestbook"
						kustomize {
							images = [
									"gcr.io/heptio-images/ks-guestbook-demo:0.1"
									]
						}
					}
					destination {
						namespace = "%[6]s"
						server = "%[7]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]s"
			cluster_id =  "%[8]s"
			agent_id = "%[4]s"
			name = "%[3]s"
            skip_repo_validation = %[10]t
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, skipRepoValidation)
}

func testAccResourceGitopsApplicationAllTypes(id string, accountId string, agentId string, clusterServer string, clusterId string, clusterToken string, repo string, helmRepoURL string, helmChart string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[3]s"
			repo {
				repo = "%[7]s"
				name = "%[1]s"
				insecure = true
				connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		resource "harness_platform_gitops_repository" "testhelmchart" {
			identifier = "%[1]shelmchart"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[3]s"
			repo {
				repo = "%[8]s"
				name = "%[1]shelmchart"
				insecure = true
				connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
		}

		

		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			agent_id = "%[3]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			request {
				upsert = true
				cluster {
					server = "%[4]s"
					name = "%[1]s"
					config {
						bearer_token = "%[6]s"
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "SERVICE_ACCOUNT"
					}
				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token, request.0.cluster.0.info, request.0.cluster.0.config.0.cluster_connection_type,
				]
			}
		}

		# Application 1: Helm app with path-based source
		resource "harness_platform_gitops_applications" "helm_app" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]shelm"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[7]s"
						path = "helm-guestbook"
					}
					destination {
						namespace = "%[10]s"
						server = "%[4]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]shelm"
			cluster_id = harness_platform_gitops_cluster.test.id
			repo_id = harness_platform_gitops_repository.test.id
			agent_id = "%[3]s"
			name = "%[1]shelm"
		}

		# Application 2: Kustomize app
		resource "harness_platform_gitops_applications" "kustomize_app" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]skustomize"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "%[7]s"
						path = "kustomize-guestbook"
						kustomize {
							images = [
								"gcr.io/heptio-images/ks-guestbook-demo:0.1"
							]
						}
					}
					destination {
						namespace = "%[10]s"
						server = "%[4]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]skustomize"
			cluster_id = harness_platform_gitops_cluster.test.id
			repo_id = harness_platform_gitops_repository.test.id
			agent_id = "%[3]s"
			name = "%[1]skustomize"
		}

		# Application 3: Multi-source app
		resource "harness_platform_gitops_applications" "multisource_app" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]smultisource"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					sources {
						repo_url = "%[7]s"
						target_revision = "master"
						ref = "val"
					}
					sources {
						repo_url = "%[7]s"
						target_revision = "master"
						path = "helm-guestbook"
						helm {
							value_files = [
								"$val/helm-guestbook/values.yaml"
							]
						}
					}
					destination {
						namespace = "%[10]s"
						server = "%[4]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			name = "%[1]smultisource"
			cluster_id = harness_platform_gitops_cluster.test.id
			repo_ids = [
				harness_platform_gitops_repository.test.id,
				harness_platform_gitops_repository.test.id
			]
			agent_id = "%[3]s"
		}

		# Application 4: Helm chart app with skip_repo_validation
		resource "harness_platform_gitops_applications" "helm_chart_app" {
			depends_on = [harness_platform_gitops_cluster.test]
			application {
				metadata {
					annotations = {}
					name = "%[1]shelmchart"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "2.6.0"
						repo_url = "https://grafana.github.io/helm-charts"
						chart = "fluent-bit"
					}
					destination {
						namespace = "%[10]s"
						server = "%[4]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]shelmchart"
			cluster_id = harness_platform_gitops_cluster.test.id
			agent_id = "%[3]s"
			name = "%[1]shelmchart"
			skip_repo_validation = true
		}

		# Application 5: Skip repo validation app (git-based)
		resource "harness_platform_gitops_applications" "skip_repo_app" {
			depends_on = [harness_platform_gitops_cluster.test]
			application {
				metadata {
					annotations = {}
					labels = {
						"harness.io/serviceRef" = harness_platform_service.test.id
						"harness.io/envRef" = harness_platform_environment.test.id
					}
					name = "%[1]sskiprepo"
				}
				spec {
					sync_policy {
						sync_options = [
							"PrunePropagationPolicy=undefined",
							"CreateNamespace=false",
							"Validate=false",
							"skipSchemaValidations=false",
							"autoCreateNamespace=false",
							"pruneLast=false",
							"applyOutofSyncOnly=false",
							"Replace=false",
							"retry=false"
						]
					}
					source {
						target_revision = "master"
						repo_url = "https://test.com"
						path = "kustomize-guestbook"
						kustomize {
							images = [
								"gcr.io/heptio-images/ks-guestbook-demo:0.1"
							]
						}
					}
					destination {
						namespace = "%[10]s"
						server = "%[4]s"
					}
				}
			}
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			account_id = "%[2]s"
			identifier = "%[1]sskiprepo"
			cluster_id = harness_platform_gitops_cluster.test.id
			agent_id = "%[3]s"
			name = "%[1]sskiprepo"
			skip_repo_validation = true
		}
		`, id, accountId, agentId, clusterServer, clusterId, clusterToken, repo, helmRepoURL, helmChart, namespace)
}
