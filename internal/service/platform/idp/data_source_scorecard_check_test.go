package idp_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIdpScorecardCheck(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_platform_idp_scorecard_check.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIdpScorecardCheck(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
				),
			},
		},
	})
}

func testAccDataSourceIdpScorecardCheck(id string) string {
	return fmt.Sprintf(`
	%s

	data "harness_platform_idp_scorecard_check" "test" {
		identifier = harness_platform_idp_scorecard_check.test.identifier
	}
	`, testAccResourceIdpScorecardCheck(id, id))
}
