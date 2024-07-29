package webhook_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGitxWebhookProjectAccLevel(t *testing.T) {
	resourceName := "harness_platform_gitx_webhook.test"
	accountId := "rXUXvbFqRr2XwcjBu3Oq-Q"
	webhook_identifier := "WebhookTest"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccGitXProjectAccountLevel(webhook_identifier, accountId, webhook_identifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", webhook_identifier),
				),
			},
			{
				Config: testAccGitXProjectAccountLevel(webhook_identifier, accountId, "WebhookNew2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "WebhookNew2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"upsert", "update_mask", "repo.0.type_"},
				ImportStateIdFunc:       acctest.GitopsWebhookImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccGitXProjectAccountLevel(webhook_identifier string, accountId string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[2]s"
			name = "%[3]s"
			project_id = "shivam"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
	`, accountId, webhook_identifier, webhook_name)
}
