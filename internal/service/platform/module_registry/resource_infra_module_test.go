package module_registry_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccResourceModule(t *testing.T) {
	resourceName := "harness_platform_infra_module.test"
	name := strings.ToLower(t.Name())
	system := strings.ToLower(t.Name())
	updatedName := fmt.Sprintf("%s_updated", name)
	updatedSystem := fmt.Sprintf("%supdated", system)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceModule(name, system),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "system", system),
				),
			},
			{
				Config: testAccResourceModule(updatedName, updatedSystem),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "system", updatedSystem),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceWorkspaceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		mod, _ := testAccGetPlatformModule(resourceName, state)
		if mod != nil {
			return fmt.Errorf("module not found: %s %s", mod.Name, mod.System)
		}
		return nil
	}
}

func testAccGetPlatformModule(resourceName string, state *terraform.State) (*nextgen.ModuleResource2, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	module, resp, err := c.ModuleRegistryApi.ModuleRegistryListModulesById(
		ctx,
		id,
		c.AccountId,
	)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &module, nil
}

func testAccResourceModule(name, system string) string {

	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}
		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

			resource "harness_platform_infra_module" "test" {
			name = "%[1]s"
			system = "%[2]s"
			description = "description"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
  		}	
	`, name, system)
}
