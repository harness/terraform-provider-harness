package registries_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestResourceRegistries(t *testing.T) {
	id := fmt.Sprintf("test_tf2")
	resourceName := "harness_platform_har_registry.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceRegistries(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccResourceRegistries(id string) string {
	return fmt.Sprintf(`
	resource "harness_platform_har_registry" "test" {
		identifier   = "%[1]s"
		space_ref    = "vpCkHKsDSxK9_KYfjCTMKA/QE_Team/DoNotDelete_Ritek_Migration"
		package_type = "DOCKER"
	
		config {
			type = "VIRTUAL"
		}
	
		parent_ref = "vpCkHKsDSxK9_KYfjCTMKA/QE_Team/DoNotDelete_Ritek_Migration"
	}
`, id)
}
