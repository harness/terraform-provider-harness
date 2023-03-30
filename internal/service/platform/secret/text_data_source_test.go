package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSecretText(t *testing.T) {
	var (
		name        = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		secretValue = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

		resourceName = "data.harness_platform_secret_text.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_text(name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
		},
	})
}
func TestAccDataSourceSecretTextProjectLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		secretValue  = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName = "data.harness_platform_secret_text.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_textProjectLevel(name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
		},
	})
}
func TestAccDataSourceSecretTextOrgLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		secretValue  = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName = "data.harness_platform_secret_text.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_textOrgLevel(name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
		},
	})
}

func testAccDataSourceSecret_text(name string, secretValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "%[2]s"
		}

		data "harness_platform_secret_text" "test"{
			identifier = harness_platform_secret_text.test.identifier
			name = harness_platform_secret_text.test.name
		}
`, name, secretValue)
}

func testAccDataSourceSecret_textProjectLevel(name string, secretValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.id
	}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id= "%[1]s"
			project_id= "%[1]s"
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "%[2]s"
			depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "4s" 
	}

		data "harness_platform_secret_text" "test"{
			identifier = harness_platform_secret_text.test.identifier
			name = harness_platform_secret_text.test.name
			org_id = "%[1]s"
			project_id = "%[1]s"
		}
`, name, secretValue)
}

func testAccDataSourceSecret_textOrgLevel(name string, secretValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = "%[1]s"
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "%[2]s"
			depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "4s" 
	}

		data "harness_platform_secret_text" "test"{
			identifier = harness_platform_secret_text.test.identifier
			name = harness_platform_secret_text.test.name
			org_id = "%[1]s"
		}
`, name, secretValue)
}
