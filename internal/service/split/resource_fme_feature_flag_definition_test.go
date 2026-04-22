package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMEFeatureFlagDefinition_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME feature flag definition acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	flagName := "tfflag_" + testAccFMEAlphanum(8)
	res := "harness_fme_feature_flag_definition.test"
	var envID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEFeatureFlagDefinition(id, envName, flagName, "off"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "flag_name", flagName),
					resource.TestCheckResourceAttrSet(res, "definition_id"),
					testAccFMECaptureAttr(res, "environment_id", &envID),
				),
			},
			{
				Config: testAccResourceFMEFeatureFlagDefinition(id, envName, flagName, "on"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "flag_name", flagName),
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
					testAccFMEVerifyFeatureFlagGone(id, id, flagName),
					testAccFMEVerifyFeatureFlagDefinitionGone(id, id, envID, flagName),
					testAccFMEVerifyEnvironmentGone(id, id, envID),
				),
			},
		},
	})
}

func testAccResourceFMEFeatureFlagDefinition(id, envName, flagName, defaultTreatment string) string {
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

	resource "harness_fme_feature_flag" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[3]s"
	}

	locals {
		ff_definition = jsonencode({
			treatments = [
				{ name = "on" },
				{ name = "off" },
			]
			defaultTreatment  = "%[4]s"
			defaultRule         = [{ treatment = "%[4]s", size = 100 }]
			trafficAllocation   = 100
		})
	}

	resource "harness_fme_feature_flag_definition" "test" {
		org_id           = harness_platform_organization.test.id
		project_id       = harness_platform_project.test.id
		environment_id   = harness_fme_environment.test.environment_id
		flag_name        = harness_fme_feature_flag.test.name
		definition       = local.ff_definition
		depends_on       = [harness_fme_feature_flag.test, harness_fme_environment.test]
	}
	`, id, envName, flagName, defaultTreatment)
}
