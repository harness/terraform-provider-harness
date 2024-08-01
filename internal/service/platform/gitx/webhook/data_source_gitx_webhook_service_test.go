package webhook_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceGitxWebhookProjectLevel(t *testing.T) {
	resourceName := "data.harness_platform_gitx_webhook.test"
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//CheckDestroy:      testAccResourceGitopsRepositoryDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGitXProjectLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func TestDataSourceGitxWebhookAccLevel(t *testing.T) {
	resourceName := "data.harness_platform_gitx_webhook.test"
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGitXAccLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func TestDataSourceGitxWebhookOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resourceName := "harness_platform_gitx_webhook.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGitXOrgLevel(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})

}

func testDataSourceGitXAccLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
		data "harness_platform_gitx_webhook" "test" {
			identifier = harness_platform_gitx_webhook.test.identifier
			name = harness_platform_gitx_webhook.test.name
		}
	`, webhook_identifier, webhook_name)
}

func testDataSourceGitXOrgLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
			org_id = "default"
			repo_name =  "GitXTest3"
			connector_ref = "account.github_Account_level_connector"
		}
		data "harness_platform_gitx_webhook" "test" {
			identifier = harness_platform_gitx_webhook.test.identifier
			name = harness_platform_gitx_webhook.test.name
			org_id = harness_platform_gitx_webhook.test.org_id
		}
	`, webhook_identifier, webhook_name)
}

func testDataSourceGitXProjectLevel(webhook_identifier string, webhook_name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitx_webhook" "test" {
			identifier= "%[1]s"
			name = "%[2]s"
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
	`, webhook_identifier, webhook_name)
}
