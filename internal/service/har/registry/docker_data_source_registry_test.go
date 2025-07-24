package registry_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccountVirtualRegistry(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_registry")
	resourceName := "data.harness_platform_har_registry.test"
	_ = os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testDocAccVirtualRegistry(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testDocAccVirtualRegistry(id string) string {
	return fmt.Sprintf(`
	 resource "harness_platform_har_registry" "test" {
	   identifier   = "%[1]s"
	   package_type = "DOCKER"	
       type = "VIRTUAL"
       virtual { }
	 }

	data "harness_platform_har_registry" "test" {
        identifier = harness_platform_har_registry.test.identifier
	}
`, id)
}
