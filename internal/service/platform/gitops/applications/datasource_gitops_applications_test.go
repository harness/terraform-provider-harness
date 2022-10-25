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
	agentId := "account.terraformagent1"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	clusterName := id
	namespace := "test"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsApplication(id, accountId, name, agentId, clusterName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func testAccDataSourceGitopsApplication(id string, accountId string, name string, agentId string, clusterName string, namespace string) string {
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
					server = "https://34.121.144.229"
					name = "%[5]s"
					config {
						bearer_token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImpsQzlJUENCTllITzBBMXg1Rzl3bXgzUWtJRk1yVERDYlIxY1BHTGdtSHcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4tNXBxNHMiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjEzZmFhNjVhLTBjNmUtNDI2MC05MTFhLWE4MTMwNGQxNzZiYSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.lVyymRAZJLsfrO6jZNJHsZxYfx75nd8mf0cEzTL7Djr_A1ChuimFm2GjnAte6O7yT7vMwf3ITbnuCTxYO-qqtQbCWBh6DFdM9PLq2lTVkKFx-6hv8J7D9poXTCUhDWYdh98Od8eg5JkL9Zz0Xf2M1p4p-QAKs_TmjhDALR2X8DNqfztB7JuPirykyPu0DroIsEkMlcsDDvn9SD0nFg_pKHLgB1AEAPGApwyzf5A37wHCOrFAsyIJ2OSyKa-ul5ZW8hM9HOOjxofOQUJEaWUOXauS1wFdTSYxDfWPuLD5njRPG21E0OxEwl7jSzZ49t7niA3jJiIvGSEO-T0ANGOJwQ"
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
                			repo_url = "https://github.com/willycoll/argocd-example-apps.git"
                			path = "helm-guestbook"
                			
            			}
            			destination {
                			namespace = "%[6]s"
                			server = "https://34.121.144.229"
            			}
        			}
    			}
    			project_id = harness_platform_project.test.id
    			org_id = harness_platform_organization.test.id
    			account_id = "%[2]s"
				identifier = "%[1]s"
				cluster_id = harness_platform_gitops_cluster.test.id
				repo_id = "account.testrepo"
				agent_id = "%[4]s"
		}
		data "harness_platform_gitops_applications" "test"{
			identifier = harness_platform_gitops_applications.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo_id = "account.testrepo"
		}
		`, id, accountId, name, agentId, clusterName, namespace)

}
