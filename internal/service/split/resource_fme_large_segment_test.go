package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMELargeSegment_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME large segment resource acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	lsName := "tfls_" + testAccFMEAlphanum(8)
	res := "harness_fme_large_segment.test"
	resAssoc := "harness_fme_large_segment_environment_association.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMELargeSegment(id, envName, lsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", lsName),
					resource.TestCheckResourceAttrSet(res, "large_segment_id"),
					resource.TestCheckResourceAttr(resAssoc, "segment_name", lsName),
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

func testAccResourceFMELargeSegment(id, envName, lsName string) string {
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

	resource "harness_fme_environment" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[2]s"
	}

	data "harness_fme_traffic_type" "user" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "user"
	}

	resource "harness_fme_large_segment" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
		description     = "acc large segment"
		depends_on      = [harness_fme_environment.test]
	}

	resource "harness_fme_large_segment_environment_association" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		environment_id = harness_fme_environment.test.environment_id
		segment_name   = harness_fme_large_segment.test.name
		depends_on     = [harness_fme_large_segment.test, harness_fme_environment.test]
	}
	`, id, envName, lsName)
}
