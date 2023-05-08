package apiKey_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceApiKey(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	resourceName := "data.harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApiKey(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
				),
			},
		},
	})
}

func TestAccDataSourceApiKeyOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	resourceName := "data.harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApiKeyOrgLevel(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceApiKeyProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	account_id := "SNM_3IzhRa6SFPz6DIV7aA"
	parent_id := "4PuRra9dTOCbT7RnG3-PRw"

	apiKeyName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	resourceName := "data.harness_platform_apiKey.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApiKeyProjectLevel(id, apiKeyName, parent_id, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", apiKeyName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceApiKey(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}

		data "harness_platform_apiKey" "test" {
			identifier = harness_platform_apiKey.test.identifier
			parent_id = harness_platform_apiKey.test.parent_id
			apikey_type = harness_platform_apiKey.test.apikey_type
			account_id = harness_platform_apiKey.test.account_id
			name = harness_platform_apiKey.test.name
		}
	`, id, name, parentId, accountId)
}

func testAccDataSourceApiKeyOrgLevel(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`

		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}

		data "harness_platform_apiKey" "test" {
			identifier = harness_platform_apiKey.test.identifier
			parent_id = harness_platform_apiKey.test.parent_id
			apikey_type = harness_platform_apiKey.test.apikey_type
			account_id = harness_platform_apiKey.test.account_id
			org_id = harness_platform_apiKey.test.org_id
			name = harness_platform_apiKey.test.name
		}
	`, id, name, parentId, accountId)
}

func testAccDataSourceApiKeyProjectLevel(id string, name string, parentId string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_apiKey" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description="Test Description"
			parent_id = "%[3]s"
			account_id = "%[4]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			apikey_type = "USER"
			default_time_to_expire_token = 1000
		}

		data "harness_platform_apiKey" "test" {
			identifier = harness_platform_apiKey.test.identifier
			parent_id = harness_platform_apiKey.test.parent_id
			apikey_type = harness_platform_apiKey.test.apikey_type
			account_id = harness_platform_apiKey.test.account_id
			org_id = harness_platform_apiKey.test.org_id
			project_id = harness_platform_apiKey.test.project_id
			name = harness_platform_apiKey.test.name
		}
	`, id, name, parentId, accountId)
}
