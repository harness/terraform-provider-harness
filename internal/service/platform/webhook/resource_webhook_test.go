package webhook_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitopsProjectAccLevel(t *testing.T) {
	resourceName := "harness_platform_gitx_webhook.test"
	accountId := "rXUXvbFqRr2XwcjBu3Oq-Q"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsProjectAccountLevel("agentId", accountId, "14a3dc9eeee999", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", "agentId"),
				),
			},
			{
				Config: testAccResourceGitopsProjectAccountLevel("agentId", accountId, "14a3dc9eeee999", "rollouts"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_id", "agentId"),
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

func testAccResourceGitopsProjectAccountLevel(agentId string, accountId string, name string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "t1"
			name = "t1"
			project_id = "shivam"
			org_id = "default"
			account_id = "%[1]s"
			repo_name =  "GitTest"
			connector_ref = "account.github_Account_level_connector"
			webhook_identifier = "We2"
			webhook_name = "We2"
		}
	`, accountId, agentId, name, namespace)
}
