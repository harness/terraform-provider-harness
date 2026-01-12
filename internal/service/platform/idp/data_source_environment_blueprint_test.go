package idp_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIdpEnvironmentBlueprint(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resourceName := "data.harness_platform_idp_environment_blueprint.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIdpEnvironmentBlueprint(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "version", "1"),
				),
			},
		},
	})
}

func testAccDataSourceIdpEnvironmentBlueprint(id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_idp_environment_blueprint" "test" {
		identifier = "%[1]s"
		version = "1"
		stable = true
		deprecated = false
		description = "description!"
		yaml = <<-EOT
		apiVersion: harness.io/v1
		kind: EnvironmentBlueprint
		type: long-lived
		identifier: %[1]s
		name: %[1]s
		owner: group:account/_account_all_users
		spec:
		  entities:
		  - identifier: git
		    backend:
		      type: HarnessCD
		      steps:
		        apply:
		          pipeline: gittest
		          branch: main
		        destroy:
		          pipeline: gittest
		          branch: not-main
		  ownedBy:
		  - group:account/_account_all_users
		EOT
	}


    data "harness_platform_idp_environment_blueprint" "test" {
        identifier = harness_platform_idp_environment_blueprint.test.identifier
        version = harness_platform_idp_environment_blueprint.test.version
    }
	`, id)
}
