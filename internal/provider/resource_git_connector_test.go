package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
)

func TestAccResourceGitConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "harness_git_connector.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitConnector(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "url", "https://github.com/micahlmartin/harness-demo"),
					resource.TestCheckResourceAttr(resourceName, "branch", "master"),
					resource.TestCheckResourceAttr(resourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(resourceName, "url_type", client.GitUrlTypes.Repo),
				),
			},
			{
				Config: testAccResourceGitConnector(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "url", "https://github.com/micahlmartin/harness-demo"),
					resource.TestCheckResourceAttr(resourceName, "branch", "master"),
					resource.TestCheckResourceAttrSet(resourceName, "webhook_url"),
					resource.TestCheckResourceAttr(resourceName, "url_type", client.GitUrlTypes.Repo),
				),
			},
		},
	})
}

func testAccResourceGitConnector(name string, generateWebhook bool) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			generate_webhook_url = %[2]t
			password_secret_id = harness_encrypted_text.test.id
			url_type = "REPO"
			username = "someuser"
		}	
`, name, generateWebhook)
}

// func testAccResourceEncryptedText_UsageScopes(name string) string {
// 	// nonprod :=
// 	return fmt.Sprintf(`
// 	resource "harness_encrypted_text" "usage_scope_test" {
// 		name = "%s"
// 		value = "someval"

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
// 		}

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
// 		}
// 	}
// 	`, name)
// }

// func testAccResourceEncryptedText_usageScopes_update(name string) string {
// 	// nonprod :=
// 	return fmt.Sprintf(`
// 	resource "harness_encrypted_text" "usage_scope_test" {
// 		name = "%s"
// 		value = "someval"

// 		usage_scope {
// 			application_filter_type = "ALL"
// 			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
// 		}
// 	}
// 	`, name)
// }

// func testAccResourceEncryptedText_update_secretmanagerid(name string, value string, secretManagerId string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_encrypted_text" "test" {
// 			name = "%s"
// 			value = "%s"
// 			secret_manager_id = "%s"
// 		}
// `, name, value, secretManagerId)
// }
