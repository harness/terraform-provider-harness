package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGcpCloudProviderConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_cloudprovider_gcp.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGcpCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckGcpCloudProviderExists(t, resourceName, name),
				),
			},
			{
				Config: testAccResourceGcpCloudProvider(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGcpCloudProviderExists(t, resourceName, name),
				),
				ExpectError: regexp.MustCompile("name is immutable"),
			},
		},
	})
}

func testAccResourceGcpCloudProvider(name string) string {
	return fmt.Sprintf(`
		resource "harness_cloudprovider_gcp" "test" {
			name = "%[1]s"
			skip_validation = true
			delegate_selectors = ["testing"]
		}
`, name)
}

func testAccCheckGcpCloudProviderExists(t *testing.T, resourceName, cloudProviderName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cp := &cac.GcpCloudProvider{}
		err := testAccGetCloudProvider(resourceName, state, cp)
		if err != nil {
			return err
		}
		return nil
	}
}
