package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMESegment_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME segment acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	segName := "tfseg_" + testAccFMEAlphanum(8)
	res := "harness_fme_segment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESegment(id, segName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", segName),
					resource.TestCheckResourceAttrSet(res, "traffic_type_id"),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: fmeImportStateIDOrgProjectThird(res, "name"),
			},
		},
	})
}

func testAccResourceFMESegment(id, segName string) string {
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

	resource "harness_fme_segment" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[2]s"
		description     = "acc classic segment"
	}
	`, id, segName)
}
