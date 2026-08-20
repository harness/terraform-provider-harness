package idp_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceIdpScorecard(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_idp_scorecard.test"
	updatedName := id + "_updated"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccScorecardDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIdpScorecard(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "published", "true"),
					resource.TestCheckResourceAttr(resourceName, "filter.0.kind", "component"),
				),
			},
			{
				Config: testAccResourceIdpScorecard(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
			},
		},
	})
}

func testAccScorecardDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		scorecard, _ := testAccGetScorecard(resourceName, state)
		if scorecard != nil {
			return fmt.Errorf("found scorecard: %s", scorecard.Identifier)
		}
		return nil
	}
}

func testAccGetScorecard(resourceName string, state *terraform.State) (*idp.Scorecard, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetIDPClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.ScorecardsApi.GetScorecard(ctx, id, &idp.ScorecardsApiGetScorecardOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return nil, err
	}
	return resp.Scorecard, nil
}

func testAccResourceIdpScorecard(id, name string) string {
	return fmt.Sprintf(`
	%[3]s

	resource "harness_platform_idp_scorecard" "test" {
		identifier          = "%[1]s"
		name                = "%[2]s"
		description         = "Gold standard scorecard"
		published           = true
		weightage_strategy  = "EQUAL_WEIGHTS"

		filter {
			kind = "component"
			type = "service"
		}

		checks {
			identifier = harness_platform_idp_scorecard_check.test.identifier
			custom     = true
		}
	}
	`, id, name, testAccResourceIdpScorecardCheck(id+"_check", id+"_check"))
}
