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

func TestAccResourceIdpScorecardCheck(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_idp_scorecard_check.test"
	updatedName := id + "_updated"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccScorecardCheckDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIdpScorecardCheck(id, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "rule_strategy", "ALL_OF"),
					resource.TestCheckResourceAttr(resourceName, "custom", "true"),
				),
			},
			{
				Config: testAccResourceIdpScorecardCheck(id, updatedName),
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

func testAccScorecardCheckDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		check, _ := testAccGetScorecardCheck(resourceName, state)
		if check != nil {
			return fmt.Errorf("found scorecard check: %s", check.Identifier)
		}
		return nil
	}
}

func testAccGetScorecardCheck(resourceName string, state *terraform.State) (*idp.CheckDetails, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetIDPClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.ChecksApi.GetCheck(ctx, id, &idp.ChecksApiGetCheckOpts{
		HarnessAccount: optional.NewString(c.AccountId),
		Custom:         optional.NewBool(true),
	})
	if err != nil {
		return nil, err
	}
	return resp.CheckDetails, nil
}

func testAccResourceIdpScorecardCheck(id, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_idp_scorecard_check" "test" {
		identifier        = "%[1]s"
		name              = "%[2]s"
		description       = "Ensures a README exists"
		rule_strategy     = "ALL_OF"
		default_behaviour = "FAIL"
		rules {
			data_source_identifier = "github"
			data_point_identifier  = "isFileExists"
			operator               = "=="
			value                  = "true"
			rule_description       = "Repository has a README"

			input_values {
				key   = "filePath"
				value = "README.md"
			}
		}
	}
	`, id, name)
}
