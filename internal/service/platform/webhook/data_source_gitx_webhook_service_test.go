package webhook_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceGitxWebhookProjectLevel(t *testing.T) {
	resourceName := "data.harness_platform_gitx_webhook.test"
	accountId := "rXUXvbFqRr2XwcjBu3Oq-Q"
	webhook_identifier := "WebhookTest"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGitXProjectLevel(webhook_identifier, accountId, webhook_identifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", webhook_identifier),
				),
			},
		},
	})

}

func testDataSourceGitXProjectLevel(webhook_identifier string, accountId string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[2]s"
			name = "%[3]s"
			project_id = "shivam"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
		data "harness_platform_gitx_webhook" "test" {
			identifier = harness_platform_gitx_webhook.test.identifier
			name = harness_platform_gitx_webhook.test.name
			project_id = harness_platform_gitx_webhook.test.project_id
			org_id = harness_platform_gitx_webhook.test.org_id
		}
	`, accountId, webhook_identifier, webhook_name)
}
