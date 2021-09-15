package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUserGroup_Id(t *testing.T) {

	var (
		groupName    = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_user_group.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupById(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", groupName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceUserGroup_Name(t *testing.T) {

	var (
		groupName    = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_user_group.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupByName(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", groupName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceUserGroup_NotFound(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserGroupById_NotFound(),
			},
		},
	})
}

func testAccDataSourceUserGroupById_NotFound() string {
	return `
		data "harness_user_group" "test" {
			id = "somebadid"
		}
	`
}

func testAccDataSourceUserGroupById(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%[1]s"
		}
		
		data "harness_user_group" "test" {
			id = harness_user_group.test.id
		}
	`, name)
}

func testAccDataSourceUserGroupByName(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%[1]s"
		}
		
		data "harness_user_group" "test" {
			name = harness_user_group.test.name
		}
	`, name)
}
