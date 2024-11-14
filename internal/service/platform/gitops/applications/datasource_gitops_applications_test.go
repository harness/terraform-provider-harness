package applications_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsApplication(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	//	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	//	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	repoId := os.Getenv("HARNESS_TEST_GITOPS_REPO_ID")
	repo := "https://github.com/harness-apps/hosted-gitops-example-apps"

	clusterName := id
	clusterId := id
	//	clusterResourceName := "data.harness_platform_gitops_cluster.test"
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER")
	clusterToken := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_TOKEN")

	namespace := "test"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsApplication(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId, clusterToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})

}

func testAccDataSourceGitopsApplication(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string, clusterToken string) string {
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

		data "harness_platform_gitops_cluster" "test" {
			depends_on = [harness_platform_gitops_cluster.test]
			identifier = harness_platform_gitops_cluster.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"

		}

		resource "harness_platform_gitops_applications" "test" {
			depends_on = [harness_platform_gitops_repository.test, harness_platform_gitops_cluster.test]
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
			name = "%[3]s"
			cluster_id = "%[8]s"
			repo_id = "%[1]s"
			agent_id = "%[4]s"
		}

		data "harness_platform_gitops_applications" "test"{
			depends_on = [harness_platform_gitops_applications.test]
			identifier = harness_platform_gitops_applications.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			name = "%[3]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId, clusterToken)

}
