package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const defaultFMETrafficTypeName = "user"

func TestAccDataSourceFMETrafficType(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME traffic type acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	ds := "data.harness_fme_traffic_type.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMETrafficType(id, defaultFMETrafficTypeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ds, "id"),
					resource.TestCheckResourceAttrSet(ds, "traffic_type_id"),
					resource.TestCheckResourceAttr(ds, "org_id", id),
					resource.TestCheckResourceAttr(ds, "project_id", id),
					resource.TestCheckResourceAttr(ds, "name", defaultFMETrafficTypeName),
				),
			},
		},
	})
}

// TestAccDataSourceFMETrafficType_matchesResource creates a traffic type and asserts the data source returns the same ids.
func TestAccDataSourceFMETrafficType_matchesResource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME traffic type acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	ttName := "tfttds_" + testAccFMEAlphanum(8)
	res := "harness_fme_traffic_type.created"
	ds := "data.harness_fme_traffic_type.lookup"
	var trafficTypeID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMETrafficTypeWithResource(id, ttName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(ds, "traffic_type_id", res, "traffic_type_id"),
					resource.TestCheckResourceAttrPair(ds, "id", res, "id"),
					resource.TestCheckResourceAttr(ds, "name", ttName),
					testAccFMECaptureAttr(res, "traffic_type_id", &trafficTypeID),
				),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyTrafficTypeGone(id, id, trafficTypeID),
			},
		},
	})
}

func testAccDataSourceFMETrafficTypeWithResource(id, ttName string) string {
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

	resource "harness_fme_traffic_type" "created" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[2]s"
	}

	data "harness_fme_traffic_type" "lookup" {
		depends_on = [harness_fme_traffic_type.created]
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = harness_fme_traffic_type.created.name
	}
	`, id, ttName)
}

func testAccDataSourceFMETrafficType(id, name string) string {
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

	data "harness_fme_traffic_type" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[2]s"
	}
	`, id, name)
}
