package service_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceService(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_service.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceService(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testAccDataSourceService(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_service_kubernetes" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			helm_version = "V2"
		}

		data "harness_service" "test" {
			app_id = harness_application.test.id
			id = harness_service_kubernetes.test.id
		}
	`, name)
}
