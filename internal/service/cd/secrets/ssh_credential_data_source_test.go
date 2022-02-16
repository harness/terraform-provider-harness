package secrets_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSshCredentialByName(t *testing.T) {

	// Setup
	var (
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "data.harness_ssh_credential.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSshCredentialByName(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceSshCredentialById(t *testing.T) {

	// Setup
	var (
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "data.harness_ssh_credential.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSshCredentialById(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceSshCredential(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "secret_manager" {
			default = true
		}
		
		resource "harness_encrypted_text" "test" {
			name              = "%[1]s"
			value             = "somefakesshkey"
			secret_manager_id = data.harness_secret_manager.secret_manager.id
		}
		
		resource "harness_ssh_credential" "test" {
			name = "%[1]s"
		
			ssh_authentication {
				port     = 22
				username = "git"
				inline_ssh {
					ssh_key_file_id = harness_encrypted_text.test.id
				}
			}

			lifecycle {
				ignore_changes = [
					"ssh_authentication"
				]
			}
		}
	`, name)
}

func testAccDataSourceSshCredentialByName(name string) string {

	credential := testAccDataSourceSshCredential(name)

	return fmt.Sprintf(`		
		%[1]s

		data "harness_ssh_credential" "test" {
			name = harness_ssh_credential.test.name
		}
	`, credential)
}

func testAccDataSourceSshCredentialById(name string) string {

	credential := testAccDataSourceSshCredential(name)

	return fmt.Sprintf(`		
		%[1]s

		data "harness_ssh_credential" "test" {
			id = harness_ssh_credential.test.id
		}
	`, credential)
}
