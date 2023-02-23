package repository_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceGitopsRepository(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	repo := "https://github.com/willycoll/argocd-example-apps.git"
	repoName := id
	agentId := os.Getenv("HARNESS_TEST_GITOPS_AGENT_ID")
	resourceName := "harness_platform_gitops_repository.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitopsRepository(id, name, repo, repoName, agentId, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "repo.0.name", repoName),
				),
			},
		},
	})
}

func testAccDataSourceGitopsRepository(id string, name string, repo string, repoName string, agentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_gitops_repository" "test" {
			identifier = "%[1]s"
			account_id = "%[6]s"
			agent_id = "%[5]s"
			repo {
					repo = "%[3]s"
        			name = "%[4]s"
        			insecure = true
        			connection_type = "HTTPS_ANONYMOUS"
			}
			upsert = true
			update_mask {
				paths = ["name"]
			}

		}
		
		data "harness_platform_gitops_repository" "test" {
			depends_on = [harness_platform_gitops_repository.test]	
			identifier = harness_platform_gitops_repository.test.id
			account_id = "%[3]s"
			agent_id = harness_platform_gitops_repository.test.agent_id
		}
	`, id, name, repo, repoName, agentId, accountId)
}
