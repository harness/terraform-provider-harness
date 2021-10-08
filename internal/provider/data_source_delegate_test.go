package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDelegate(t *testing.T) {

	var (
		resourceName = "data.harness_delegate.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDelegate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "harness-delegate"),
					resource.TestCheckResourceAttr(resourceName, "ip", "10.12.1.8"),
					resource.TestCheckResourceAttr(resourceName, "host_name", "harness-delegate-ukhyts-1"),
					resource.TestCheckResourceAttr(resourceName, "id", "ZAgU6QnGSAa1vbsZLgdqcQ"),
					resource.TestCheckResourceAttr(resourceName, "account_id", "UKh5Yts7THSMAbccG3HrLA"),
					resource.TestCheckResourceAttr(resourceName, "profile_id", "wVhjS1xATpqGrO-uuRrYEw"),
					resource.TestCheckResourceAttr(resourceName, "type", "KUBERNETES"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttrSet(resourceName, "last_heartbeat"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "version", ""),
				),
			},
		},
	})
}

func testAccDataSourceDelegate() string {
	return `
		data "harness_delegate" "test" {
			name = "harness-delegate"
		}
	`
}
