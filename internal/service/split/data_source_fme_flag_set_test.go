package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceFMEFlagSet creates a flag set in the test workspace, then looks it up with the data source.
func TestAccDataSourceFMEFlagSet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME flag set acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	fsName := fmt.Sprintf("tf_fs_ds_%s", testAccRandomFlagSetSuffix(8))
	res := "harness_fme_flag_set.created"
	ds := "data.harness_fme_flag_set.test"
	var flagSetID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMEFlagSet(id, fsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ds, "id"),
					resource.TestCheckResourceAttrSet(ds, "flag_set_id"),
					resource.TestCheckResourceAttrPair(ds, "flag_set_id", res, "flag_set_id"),
					resource.TestCheckResourceAttrPair(ds, "id", res, "id"),
					resource.TestCheckResourceAttr(ds, "name", fsName),
					testAccFMECaptureAttr(res, "flag_set_id", &flagSetID),
				),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyFlagSetGone(id, id, flagSetID),
			},
		},
	})
}

func testAccDataSourceFMEFlagSet(id, flagSetName string) string {
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

	resource "harness_fme_flag_set" "created" {
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		name        = "%[2]s"
		description = "acc data source flag set"
	}

	data "harness_fme_flag_set" "test" {
		depends_on = [harness_fme_flag_set.created]
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = harness_fme_flag_set.created.name
	}
	`, id, flagSetName)
}
