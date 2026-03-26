package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceFMELargeSegment creates a large segment then looks it up by name.
func TestAccDataSourceFMELargeSegment(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME large segment data source acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	lsName := fmt.Sprintf("tfls_%s", utils.RandStringBytes(8))
	ds := "data.harness_fme_large_segment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMELargeSegment(id, lsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ds, "id"),
					resource.TestCheckResourceAttrSet(ds, "large_segment_id"),
					resource.TestCheckResourceAttr(ds, "name", lsName),
					resource.TestCheckResourceAttrSet(ds, "traffic_type_id"),
				),
			},
		},
	})
}

func testAccDataSourceFMELargeSegment(id, lsName string) string {
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

	resource "harness_fme_large_segment" "created" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[2]s"
		description     = "acc large segment for data source"
	}

	data "harness_fme_large_segment" "test" {
		depends_on = [harness_fme_large_segment.created]
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = harness_fme_large_segment.created.name
	}
	`, id, lsName)
}
