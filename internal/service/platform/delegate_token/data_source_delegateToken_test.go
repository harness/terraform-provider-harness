package delegatetoken_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDelegateToken(t *testing.T) {
	name := "tempName223421311q"
	account_id := os.Getenv("HARNESS_ACCOUNT_ID")

	resourceName := "data.harness_platform_delegatetoken.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tesAccDataSourceDelegateToken1(name, account_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func tesAccDataSourceDelegateToken(name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_delegatetoken" "test" {
			name = "%[1]s"
			account_id = "%[2]s"
		}

		data "harness_platform_delegatetoken" "test" {
			name = harness_platform_delegatetoken.test.name
			account_id = harness_platform_delegatetoken.test.account_id
		}
	`, name, accountId)
}

func tesAccDataSourceDelegateToken1(name string, accountId string) string {
	return fmt.Sprintf(`
		data "harness_platform_delegatetoken" "test" {
			name = "%[1]s"
			account_id = "%[2]s"
		}
	`, name, accountId)
}
