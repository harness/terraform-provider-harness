package idp_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIdpEnvironment(t *testing.T) {
	t.Skip("Skipping test as it takes 2+ minutes to run and requires elaborate Harness setup")

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resourceName := "data.harness_platform_idp_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIdpEnvironment(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", "default"),
					resource.TestCheckResourceAttr(resourceName, "project_id", "ssem"),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "blueprint_identifier", "noop"),
					resource.TestCheckResourceAttr(resourceName, "blueprint_version", "v1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "target_state", "inactive"),
				),
			},
		},
	})
}

func testAccDataSourceIdpEnvironment(id string) string {
	str := fmt.Sprintf(`
	resource "harness_platform_idp_environment" "test" {
	    identifier = "%[1]s"
		org_id = "default"
		project_id = "ssem"
		name = "%[1]s"
		owner = "user:account/admin@harness.io"
		blueprint_identifier = "noop"
		blueprint_version = "v1.0.0"
		target_state = "inactive"
		overrides = <<-EOT
        config: {}
        entities: {}
        EOT
	}

	data "harness_platform_idp_environment" "test" {
		identifier = harness_platform_idp_environment.test.identifier
		org_id = harness_platform_idp_environment.test.org_id
		project_id = harness_platform_idp_environment.test.project_id
	}
	`, id)

	return str
}
