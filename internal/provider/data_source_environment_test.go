package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEnvironment_Id(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_environment.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentById(name, cac.EnvironmentTypes.Prod),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", string(cac.EnvironmentTypes.Prod)),
				),
			},
		},
	})
}
func TestAccDataSourceEnvironment_Name(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_environment.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentByName(name, cac.EnvironmentTypes.Prod),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", string(cac.EnvironmentTypes.Prod)),
				),
			},
		},
	})
}

func testAccDataSourceEnvironmentById(name string, envType cac.EnvironmentType) string {
	return fmt.Sprintf(`
	resource "harness_application" "test" {
		name = "%[1]s"
	}

	resource "harness_environment" "test" {
		app_id = harness_application.test.id
		name = "%[1]s"
		type = "%[2]s"
	}

	data "harness_environment" "test" {
		app_id = harness_application.test.id
		id = harness_environment.test.id
	}

	`, name, envType)
}

func testAccDataSourceEnvironmentByName(name string, envType cac.EnvironmentType) string {
	return fmt.Sprintf(`
	resource "harness_application" "test" {
		name = "%[1]s"
	}

	resource "harness_environment" "test" {
		app_id = harness_application.test.id
		name = "%[1]s"
		type = "%[2]s"
	}

	data "harness_environment" "test" {
		app_id = harness_application.test.id
		name = "%[1]s"
	}

	`, name, envType)
}
