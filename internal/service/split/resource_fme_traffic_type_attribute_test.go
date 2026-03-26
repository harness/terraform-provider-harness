package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMETrafficTypeAttribute_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME traffic type attribute acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	attrID := "tfattr_" + testAccFMEAlphanum(8)
	res := "harness_fme_traffic_type_attribute.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMETrafficTypeAttribute(id, attrID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "identifier", attrID),
					resource.TestCheckResourceAttr(res, "display_name", "ACC TT Attribute"),
					resource.TestCheckResourceAttr(res, "data_type", "STRING"),
					resource.TestCheckResourceAttrSet(res, "attribute_id"),
				),
			},
		},
	})
}

func testAccResourceFMETrafficTypeAttribute(id, attrIdentifier string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	data "harness_fme_traffic_type" "user" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "user"
	}

	resource "harness_fme_traffic_type_attribute" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		identifier      = "%[2]s"
		display_name    = "ACC TT Attribute"
		data_type       = "STRING"
		is_searchable   = false
		suggested_values = ["acc_a", "acc_b"]
	}
	`, id, attrIdentifier)
}
