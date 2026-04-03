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
	var trafficTypeID string
	var attributeAPIID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMETrafficTypeAttribute(id, attrID, "ACC TT Attribute", false, []string{"acc_a", "acc_b"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "identifier", attrID),
					resource.TestCheckResourceAttr(res, "display_name", "ACC TT Attribute"),
					resource.TestCheckResourceAttr(res, "data_type", "string"),
					resource.TestCheckResourceAttr(res, "is_searchable", "false"),
					resource.TestCheckResourceAttr(res, "suggested_values.#", "2"),
					resource.TestCheckResourceAttrSet(res, "attribute_id"),
					testAccFMECaptureAttr(res, "traffic_type_id", &trafficTypeID),
					testAccFMECaptureAttr(res, "attribute_id", &attributeAPIID),
				),
			},
			{
				Config: testAccResourceFMETrafficTypeAttribute(id, attrID, "ACC TT Attribute Updated", true, []string{"acc_a", "acc_b", "acc_c"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "display_name", "ACC TT Attribute Updated"),
					resource.TestCheckResourceAttr(res, "is_searchable", "true"),
					resource.TestCheckResourceAttr(res, "suggested_values.#", "3"),
					testAccFMECaptureAttr(res, "traffic_type_id", &trafficTypeID),
					testAccFMECaptureAttr(res, "attribute_id", &attributeAPIID),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: fmeImportStateIDOrgProjectTTFourth(res),
				Check: resource.ComposeTestCheckFunc(
					testAccFMECaptureAttr(res, "traffic_type_id", &trafficTypeID),
					testAccFMECaptureAttr(res, "attribute_id", &attributeAPIID),
				),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyTrafficTypeAttributeGone(id, id, trafficTypeID, attributeAPIID),
			},
		},
	})
}

func testAccResourceFMETrafficTypeAttribute(id, attrIdentifier, displayName string, searchable bool, suggested []string) string {
	sugHCL := "[]"
	if len(suggested) > 0 {
		sugHCL = testAccHCLStringList(suggested)
	}
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
		org_id             = harness_platform_organization.test.id
		project_id         = harness_platform_project.test.id
		traffic_type_id    = data.harness_fme_traffic_type.user.traffic_type_id
		identifier         = "%[2]s"
		display_name       = "%[3]s"
		data_type          = "string"
		is_searchable      = %[4]t
		suggested_values   = %[5]s
	}
	`, id, attrIdentifier, displayName, searchable, sugHCL)
}
