package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMESegmentEnvironmentAssociation_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME segment environment association acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	segName := "tfseg_" + testAccFMEAlphanum(8)
	res := "harness_fme_segment_environment_association.test"
	var envID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESegmentEnvironmentAssociation(id, envName, segName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "segment_name", segName),
					resource.TestCheckResourceAttrSet(res, "environment_id"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: fmeImportStatePrimaryID(res),
				Check:             testAccFMECaptureAttr(res, "environment_id", &envID),
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check: resource.ComposeTestCheckFunc(
					testAccFMEVerifySegmentGone(id, id, segName),
					testAccFMEVerifySegmentEnvAssociationInactive(id, id, envID, segName),
					testAccFMEVerifyEnvironmentGone(id, id, envID),
				),
			},
		},
	})
}

func testAccResourceFMESegmentEnvironmentAssociation(id, envName, segName string) string {
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

	resource "harness_fme_segment" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
	}

	resource "harness_fme_segment_environment_association" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		environment_id = harness_fme_environment.test.environment_id
		segment_name   = harness_fme_segment.test.name
		depends_on     = [harness_fme_segment.test, harness_fme_environment.test]
	}
	`, id, envName, segName)
}
