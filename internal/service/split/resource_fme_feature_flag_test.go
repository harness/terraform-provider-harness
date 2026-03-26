package split_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFMEFeatureFlag_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME feature flag acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	flagName := "tfflag_" + testAccFMEAlphanum(8)
	res := "harness_fme_feature_flag.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEFeatureFlag(id, flagName, "acc feature flag v1", []string{"acc_test"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", flagName),
					resource.TestCheckResourceAttr(res, "description", "acc feature flag v1"),
					resource.TestCheckResourceAttr(res, "tags.#", "1"),
					resource.TestCheckResourceAttrSet(res, "flag_id"),
				),
			},
			{
				Config: testAccResourceFMEFeatureFlag(id, flagName, "acc feature flag v2", []string{"acc_test", "acc_post_update"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "description", "acc feature flag v2"),
					resource.TestCheckResourceAttr(res, "tags.#", "2"),
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

func testAccHCLStringList(ss []string) string {
	parts := make([]string, len(ss))
	for i, s := range ss {
		parts[i] = fmt.Sprintf("%q", s)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func testAccResourceFMEFeatureFlag(id, flagName, description string, tags []string) string {
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

	resource "harness_fme_feature_flag" "test" {
		org_id          = harness_platform_organization.test.id
		project_id      = harness_platform_project.test.id
		traffic_type_id = data.harness_fme_traffic_type.user.traffic_type_id
		name            = "%[2]s"
		description     = "%[3]s"
		tags            = %[4]s
	}
	`, id, flagName, description, testAccHCLStringList(tags))
}
