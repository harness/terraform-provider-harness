package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecretText_inline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_text_inline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				Config: testAccResourceSecret_text_inline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceSecretText_reference(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_text_reference(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "azureSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", "secret"),
				),
			},
			{
				Config: testAccResourceSecret_text_reference(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "azureSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", "secret"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceSecret_text_inline(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "org" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "project" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.org.id
		color = "#0063F7"
	}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.org.id
			project_id = harness_platform_project.project.id

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
			lifecycle {
				ignore_changes = [
					value,
				]
			}
		}
`, id, name)
}

func testAccResourceSecret_text_reference(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "azureSecretManager"
			value_type = "Reference"
			value = "secret"
		}
`, id, name)
}
