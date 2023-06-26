package ccm_perspective_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceCCMPerspective(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_ccm_perspective.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCCMPerspective(id, id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccCCMPerspective(resourceName string, state *terraform.State) (*nextgen.CeView, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.CloudCostPerspectivesApi.GetPerspective(ctx, c.AccountId, id)

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func testAccCCMPerspectiveDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		pers, _ := testAccCCMPerspective(resourceName, state)
		if pers != nil {
			return fmt.Errorf("Found ccm perspective: %s", pers.Uuid)
		}

		return nil
	}
}

func testAccResourceCCMPerspective(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_ccm_perspective" "test" {
			identifier = "%[1]s"
			account_id = "%[3]s"
			name = "%[2]s"
			clone = false
			view_version = "v1"
			view_state = "DRAFT"
			view_type = "CUSTOMER"
			lifecycle {
				ignore_changes = [
					identifier,
				]
			}
  	}
`, id, name, accountId)
}
