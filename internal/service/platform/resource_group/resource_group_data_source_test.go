package resource_group_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceResourceGroup(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_resource_group.test"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
	)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceResourceGroup(name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "included_scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_filter.0.include_all_resources", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceResourceGroupProjectLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_resource_group.test"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
	)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceResourceGroupProjectLevel(name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "included_scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_filter.0.include_all_resources", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceResourceGroupOrgLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_resource_group.test"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
	)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceResourceGroupOrgLevel(name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "included_scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_filter.0.include_all_resources", "false"),
				),
			},
		},
	})
}

func testAccDataSourceResourceGroup(name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_resource_group" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		account_id = "%[2]s"
		allowed_scope_levels =["account"]
		included_scopes {
			filter = "EXCLUDING_CHILD_SCOPES"
			account_id = "%[2]s"
		}
		resource_filter {
			include_all_resources = false
			resources {
				resource_type = "CONNECTOR"
				attribute_filter {
					attribute_name = "category"
					attribute_values = ["value"]
				}
			}
		}
	}

		data "harness_platform_resource_group" "test" {
			identifier = harness_platform_resource_group.test.identifier
			
		}
	`, name, accountId)
}

func testAccDataSourceResourceGroupProjectLevel(name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}
	
		resource "harness_platform_resource_group" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			account_id = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			allowed_scope_levels =["project"]
			included_scopes {
				filter = "EXCLUDING_CHILD_SCOPES"
				account_id = "%[2]s"
				org_id = harness_platform_project.test.org_id
				project_id = harness_platform_project.test.id
			}
			resource_filter {
				include_all_resources = false
				resources {
					resource_type = "CONNECTOR"
					attribute_filter {
						attribute_name = "category"
						attribute_values = ["value"]
					}
				}
			}
		}
		data "harness_platform_resource_group" "test" {
			identifier = harness_platform_resource_group.test.identifier
			org_id = harness_platform_resource_group.test.org_id
			project_id = harness_platform_resource_group.test.project_id
		}
	`, name, accountId)
}

func testAccDataSourceResourceGroupOrgLevel(name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}
	resource "harness_platform_resource_group" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		account_id = "%[2]s"
		org_id = harness_platform_organization.test.id
		
		allowed_scope_levels =["organization"]
		included_scopes {
			filter = "EXCLUDING_CHILD_SCOPES"
			account_id = "%[2]s"
			org_id = harness_platform_organization.test.id
			
		}
		resource_filter {
			include_all_resources = false
			resources {
				resource_type = "CONNECTOR"
				attribute_filter {
					attribute_name = "category"
					attribute_values = ["value"]
				}
			}
		}
	}

		data "harness_platform_resource_group" "test" {
			identifier = harness_platform_resource_group.test.identifier
			org_id = harness_platform_resource_group.test.org_id
		}
	`, name, accountId)
}
