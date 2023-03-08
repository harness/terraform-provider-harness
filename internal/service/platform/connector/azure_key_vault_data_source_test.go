package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAzureKeyVault(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_azure_key_vault.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_azureKeyVault(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tenant_id", "b229b2bb-5f33-4d22-bce0-730f6474e906"),
					resource.TestCheckResourceAttr(resourceName, "subscription", "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"),
					resource.TestCheckResourceAttr(resourceName, "vault_name", "Aman-test"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
		},
	})
}

func testAccDataSourceConnector_azureKeyVault(name string) string {
	return fmt.Sprintf(`

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.Azure_Secret_Key_Do_Not_Delete"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"
		is_default = false

		azure_environment_type = "AZURE"
	}

	data "harness_platform_connector_azure_key_vault" "test" {
		identifier = harness_platform_connector_azure_key_vault.test.id
	}
`, name)
}
