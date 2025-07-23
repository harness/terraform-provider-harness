package registry_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceVirtualRegistry(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_registry")
	resourceName := "data.harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	println("accountId" + ":: " + accountId)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccDataSourceVirtualRegistry(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccDataSourceVirtualRegistry(id string, accId string) string {
	return fmt.Sprintf(`

	 resource "harness_platform_har_registry" "test" {
	   identifier   = "%[1]s"
	   package_type = "NUGET"
	   config {
		type = "VIRTUAL"
	   }
		org_id = "har_auto_orghDfwPNTfPS"
        project_id = "har_auto_projectnQRdqPHOUj"
	 }

	data "harness_platform_har_registry" "test" {
			identifier = harness_platform_har_registry.test.identifier
            org_id = "har_auto_orghDfwPNTfPS"
            project_id = "har_auto_projectnQRdqPHOUj"
	}
`, id, accId)
}
