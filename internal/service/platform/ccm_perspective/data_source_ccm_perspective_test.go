package ccm_perspective_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCCMPerspective(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(6))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_ccm_perspective.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataCCMPerspective(id, id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", id),
				),
			},
		},
	})
}

func testAccDataCCMPerspective(id string, name string, accountId string) string {
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

	data "harness_platform_ccm_perspective" "test" {
			identifier = harness_platform_ccm_perspective.test.identifier
			account_id = harness_platform_ccm_perspective.test.account_id
			name = harness_platform_ccm_perspective.test.name
			view_version = harness_platform_ccm_perspective.test.view_version
			clone = harness_platform_ccm_perspective.test.clone
		}
`, id, name, accountId)
}
