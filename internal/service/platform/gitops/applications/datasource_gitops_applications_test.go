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
	clusterServer := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_SERVER_APP")
	clusterId := os.Getenv("HARNESS_TEST_GITOPS_CLUSTER_ID")
	repoId := os.Getenv("HARNESS_TEST_GITOPS_REPO_ID")
	repo := "https://github.com/argoproj/argocd-example-apps"
	clusterName := id
	namespace := "test"
	resourceName := "harness_platform_gitops_applications.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsApplication(id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func testAccDataSourceGitopsApplication(id string, accountId string, name string, agentId string, clusterName string, namespace string, clusterServer string, clusterId string, repo string, repoId string) string {
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
				cluster_id = "%[8]s"
				repo_id = "%[10]s"
				agent_id = "%[4]s"
		}
		data "harness_platform_gitops_applications" "test"{
			identifier = harness_platform_gitops_applications.test.id
			account_id = "%[2]s"
			project_id = harness_platform_project.test.id
			org_id = harness_platform_organization.test.id
			agent_id = "%[4]s"
			repo_id = "%[10]s"
		}
		`, id, accountId, name, agentId, clusterName, namespace, clusterServer, clusterId, repo, repoId)

}
