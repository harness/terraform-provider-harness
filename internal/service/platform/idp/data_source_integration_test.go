package idp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIdpIntegration(t *testing.T) {
	id := fmt.Sprintf("tf_acc_%s", strings.ToLower(utils.RandStringBytes(8)))
	const dataSourceName = "data.harness_platform_idp_integration.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIdpIntegration(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", id),
					resource.TestCheckResourceAttr(dataSourceName, "name", id),
					resource.TestCheckResourceAttr(dataSourceName, "type", "Custom"),
					resource.TestCheckResourceAttr(dataSourceName, "custom.0.source_label", "orders"),
					resource.TestCheckResourceAttr(dataSourceName, "kind.0.kind", "custom"),
				),
			},
		},
	})
}

func TestAccDataSourceIdpIntegrationKindMapping(t *testing.T) {
	id := fmt.Sprintf("tf_acc_%s", strings.ToLower(utils.RandStringBytes(8)))
	const dataSourceName = "data.harness_platform_idp_integration_kind_mapping.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIdpIntegrationKindMapping(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "id", id+"/custom"),
					resource.TestCheckResourceAttr(dataSourceName, "ingestion.0.identifier", ".service.id"),
					resource.TestCheckResourceAttr(dataSourceName, "ingestion.0.properties.team", ".service.team"),
				),
			},
		},
	})
}

func testAccResourceIdpIntegrationCustom(id string) string {
	return fmt.Sprintf(`
resource "harness_platform_idp_integration" "test" {
  identifier = "%[1]s"
  name       = "%[1]s"
  type       = "Custom"

  custom {
    mode         = "api"
    source_label = "orders"
  }

  kind {
    kind = "custom"

    mapping {
      ingestion {
        identifier = ".service.id"
        name       = ".service.name"
        properties = {
          team = ".service.team"
        }
      }
    }
  }
}
`, id)
}

func testAccDataSourceIdpIntegration(id string) string {
	return fmt.Sprintf(`
%s

data "harness_platform_idp_integration" "test" {
  identifier = harness_platform_idp_integration.test.identifier
}
`, testAccResourceIdpIntegrationCustom(id))
}

func testAccDataSourceIdpIntegrationKindMapping(id string) string {
	return fmt.Sprintf(`
%s

data "harness_platform_idp_integration_kind_mapping" "test" {
  integration_id = harness_platform_idp_integration.test.identifier
  kind           = "custom"
}
`, testAccResourceIdpIntegrationCustom(id))
}
