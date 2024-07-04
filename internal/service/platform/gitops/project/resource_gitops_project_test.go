package project_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsProjectAccLevel(t *testing.T) {
	agentId := "account.rollouts"
	resourceName := "harness_platform_gitops_project.test"
	accountId := "1bvyLackQK-Hapk25-Ry4w"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsRepositoryOrgLevel(agentId, accountId, "14a3dc9eeee999", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				Config: testAccResourceGitopsRepositoryOrgLevel(agentId, accountId, "14a3dc9eeee999", "rollouts"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsProjectImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccResourceGitopsRepositoryOrgLevel(agentId string, accountId string, name string, namespace string) string {
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
	`, accountId, agentId, name, namespace)
}
