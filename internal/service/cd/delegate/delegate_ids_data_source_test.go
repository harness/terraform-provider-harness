package delegate_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDelegateIds(t *testing.T) {
	t.Skip("Skipping until we figure out how to get the tests passing properly in CI")

	delegateName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "data.harness_delegate_ids.test"

	defer deleteDelegate(t, delegateName)

	acctest.TestAccPreCheck(t)
	delegate := createDelegateContainer(t, delegateName)

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegateIds(delegateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "delegate_ids.0", delegate.UUID),
					resource.TestCheckResourceAttr(resourceName, "delegate_ids.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceDelegateIds(name string) string {
	return fmt.Sprintf(`
		data "harness_delegate_ids" "test" {
			name = "%[1]s"
		}
	`, name)
}
