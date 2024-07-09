package project_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitopsProjectAccLevel(t *testing.T) {
	resourceName := "harness_platform_gitops_project.test"
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsRepositoryOrgLevel(agentId, accountId, "14a3dc9eeee999", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
		},
	})

}

func testAccDataSourceGitopsRepositoryOrgLevel(agentId string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_project" "test" {
			account_id = "%[1]s"
			agent_id = "%[2]s"
			upsert = true
			project {
				metadata {
					generation = "1"
					name = "%[3]s"
					namespace = "rollouts"
				}
				spec {
					cluster_resource_whitelist {
						group = "*"
						kind = "*"
					}
					destinations {
						namespace = "%[4]s"
						server = "*"
					}
					source_repos = ["*"]
				}
			}
		}

		data "harness_platform_gitops_project" "test" {
			depends_on = [harness_platform_gitops_project.test]	
			account_id = "%[1]s"
			agent_id = harness_platform_gitops_project.test.agent_id
		}	
	`, accountId, agentId, name, namespace)
}
