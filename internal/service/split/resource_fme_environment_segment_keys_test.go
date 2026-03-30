package split_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMEEnvironmentSegmentKeys_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME environment segment keys acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	envName := "tf" + testAccFMEAlphanum(10)
	segName := "tfseg_" + testAccFMEAlphanum(8)
	res := "harness_fme_environment_segment_keys.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironmentSegmentKeys(id, envName, segName, []string{"acc_key_1", "acc_key_2"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "segment_name", segName),
					resource.TestCheckResourceAttr(res, "keys.#", "2"),
				),
			},
			{
				Config: testAccResourceFMEEnvironmentSegmentKeys(id, envName, segName, []string{"acc_key_1", "acc_key_3"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "keys.#", "2"),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: fmeImportStatePrimaryID(res),
			},
		},
	})
}

func testAccResourceFMEEnvironmentSegmentKeys(id, envName, segName string, keys []string) string {
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

	resource "harness_fme_segment_environment_association" "act" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		environment_id = harness_fme_environment.test.environment_id
		segment_name   = harness_fme_segment.test.name
		depends_on     = [harness_fme_segment.test, harness_fme_environment.test]
	}

	resource "harness_fme_environment_segment_keys" "test" {
		org_id         = harness_platform_organization.test.id
		project_id     = harness_platform_project.test.id
		environment_id = harness_fme_environment.test.environment_id
		segment_name   = harness_fme_segment.test.name
		keys           = %[4]s
		depends_on     = [harness_fme_segment_environment_association.act]
	}
	`, id, envName, segName, testAccHCLStringList(keys))
}
