package applications_test

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsApplication(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")
	clusterName := id
	namespace := "test"
	repo := "https://github.com/willycoll/argocd-example-apps.git"
	namespaceUpdated := namespace + "_updated"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsApplication(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterToken, repo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceGitopsApplication(id, accountId, name, agentId, clusterName, namespaceUpdated, clusterServer, clusterToken, repo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "application.0.spec.0.destination.0.namespace", namespaceUpdated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.GitopsAgentAccountLevelResourceImportStateIdFunc(resourceName),
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
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	agentIdentifier := r.Primary.Attributes["agent_id"]
	orgIdentifier := r.Primary.Attributes["org_id"]
	projectIdentifier := r.Primary.Attributes["project_id"]
	queryName := r.Primary.Attributes["identifier"]
	repoIdentifier := r.Primary.Attributes["repo_id"]

	resp, _, err := c.ApplicationsApiService.AgentApplicationServiceGet(ctx, agentIdentifier, queryName, c.AccountId, orgIdentifier, projectIdentifier, &nextgen.ApplicationsApiAgentApplicationServiceGetOpts{
		QueryRepo: optional.NewString(repoIdentifier),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func testAccResourceGitopsApplication(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterToken string, repo string) string {
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
      		name = "%[2]s"
      		org_id = harness_platform_project.test.org_id
      		project_id = harness_platform_project.test.id
    	}
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  		}
		resource "harness_platform_gitops_cluster" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

 			request {
				upsert = true
				cluster {
					server = "%[7]s"
					name = "%[5]s"
					config {
						bearer_token = "%[8]s"
						tls_client_config {
							insecure = true
						}
						cluster_connection_type = "SERVICE_ACCOUNT"
					}

				}
			}
			lifecycle {
				ignore_changes = [
					request.0.upsert, request.0.cluster.0.config.0.bearer_token,
				]
			}
		}
		
		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo {
					repo = "%[9]s"
        			name = "%[2]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
			update_mask {
				paths = ["name"]
			}
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
				cluster_id = harness_platform_gitops_cluster.test.id
				repo_id = harness_platform_gitops_repository.test.id
				agent_id = "%[4]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterToken, repo)

}
