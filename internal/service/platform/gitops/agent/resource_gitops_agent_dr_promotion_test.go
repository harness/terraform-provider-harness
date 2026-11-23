package agent_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Promoting a SECONDARY node to PRIMARY via harness_platform_gitops_agent_dr_promotion must
// be reflected back on the promotion resource's current_primary attribute.
func TestAccResourceGitopsAgentDrPromotion(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	id = strings.ReplaceAll(id, "_", "")
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	primaryResourceName := "harness_platform_gitops_agent.primary"
	secondaryResourceName := "harness_platform_gitops_agent.secondary"
	promotionResourceName := "harness_platform_gitops_agent_dr_promotion.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			if err := testAccResourceGitopsAgentDestroy(primaryResourceName)(s); err != nil {
				return err
			}
			return testAccResourceGitopsAgentDestroy(secondaryResourceName)(s)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgentDrPrimarySecondary(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(primaryResourceName, "disaster_recovery_live_type", "PRIMARY"),
					resource.TestCheckResourceAttr(secondaryResourceName, "disaster_recovery_live_type", "SECONDARY"),
				),
			},
			{
				Config: testAccResourceGitopsAgentDrPromotionConfig(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(promotionResourceName, "current_primary", id),
				),
			},
		},
	})
}

func testAccResourceGitopsAgentDrPromotionConfig(id string, accountId string) string {
	return testAccResourceGitopsAgentDrPrimarySecondary(id, accountId) + fmt.Sprintf(`
		resource "harness_platform_gitops_agent_dr_promotion" "test" {
			agent_id      = harness_platform_gitops_agent.secondary.identifier
			dr_identifier = "%[1]s"

			depends_on = [
				harness_platform_gitops_agent.primary,
				harness_platform_gitops_agent.secondary,
			]
		}
		`, id)
}
