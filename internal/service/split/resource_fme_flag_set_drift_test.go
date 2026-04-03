package split_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	splitpkg "github.com/harness/terraform-provider-harness/internal/service/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/require"
)

// TestAccFMEFlagSet_destroyedOutOfBand deletes the flag set via the Split API then expects a non-empty plan to recreate it.
func TestAccFMEFlagSet_destroyedOutOfBand(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME flag set drift acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	fsName := fmt.Sprintf("tf_drift_%s", testAccRandomFlagSetSuffix(8))
	res := "harness_fme_flag_set.test"
	cfg := testAccFMEFlagSetOnly(id, fsName, "drift test flag set")
	var flagSetID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(res, "flag_set_id"),
					resource.TestCheckResourceAttr(res, "name", fsName),
					testAccFMECaptureAttr(res, "flag_set_id", &flagSetID),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					meta := acctest.TestAccGetApiClientFromProvider()
					ctx := context.Background()
					client, cctx := meta.GetSplitClientWithContext(ctx)
					wsID, err := splitpkg.AccWorkspaceID(cctx, meta, id, id)
					require.NoError(t, err)
					fs, err := client.FlagSets.FindByName(wsID, fsName)
					require.NoError(t, err)
					require.NotNil(t, fs, "flag set should exist before out-of-band delete")
					require.NoError(t, client.FlagSets.Delete(fs.ID))
				},
				Config:             cfg,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccFMEHarnessOrgProjectOnly(id),
				Check:  testAccFMEVerifyFlagSetGone(id, id, flagSetID),
			},
		},
	})
}

func testAccFMEFlagSetOnly(id, fsName, description string) string {
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

	resource "harness_fme_flag_set" "test" {
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		name        = "%[2]s"
		description = "%[3]s"
	}
	`, id, fsName, description)
}
