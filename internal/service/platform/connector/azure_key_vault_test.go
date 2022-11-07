package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorAzureKeyVault(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_key_vault.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzureKeyVault(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tenant_id", "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "subscription", "subscription"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "client_id"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				Config: testAccResourceConnectorAzureKeyVault(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tenant_id", "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "subscription", "subscription"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "client_id"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnectorAzureKeyVault(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "0ac78b03-57f0-b205-e85b-4931fb814366"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		client_id = "client_id"
		secret_key = 
		tenant_id = "tenant_id"
		vault_name = 
		subscription = "subscription"
		is_default = false

		azure_environment_type = "AZURE"
		delegate_selectors = ["harness-delegate"]
	}
`, id, name)
}
