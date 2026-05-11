package ccm_filters_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceCCMFilters_CCM32336_OutOfBandDeleteRecreates verifies that when
// a CCM filter is deleted out-of-band (UI / direct API), the next terraform
// refresh treats the GET as "not found" and re-plans a create instead of
// erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336 (the bug class is: a NextGen GET endpoint
// returning HTTP 500 for a deleted entity causes terraform plan to fail).
// The CCM filter resource Read() routes through helpers.HandleReadApiError
// which clears state on 404 + ENTITY_NOT_FOUND.
func TestAccResourceCCMFilters_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_ccm_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCCMFiltersOrgLevelDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCCMFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "CCMRecommendation"),
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, _, err := c.FilterApi.CcmdeleteFilter(
						ctx, c.AccountId, id, "CCMRecommendation",
						&nextgen.FilterApiCcmdeleteFilterOpts{
							OrgIdentifier: optional.NewString(id),
						},
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config:             testAccResourceCCMFiltersOrgLevel(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceCCMFiltersOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}
