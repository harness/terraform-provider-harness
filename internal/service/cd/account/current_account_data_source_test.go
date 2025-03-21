package account_test

import (
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCurrentAccount(t *testing.T) {

	var (
		resourceName = "data.harness_current_account.test"
		accountId    = os.Getenv("HARNESS_ACCOUNT_ID")
		endpoint     = os.Getenv("HARNESS_ENDPOINT")
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCurrentAccount(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "endpoint", endpoint),
				),
			},
		},
	})
}

func testAccDataSourceCurrentAccount() string {
	return `data "harness_current_account" "test" {}`
}
