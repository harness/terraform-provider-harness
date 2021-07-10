package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSpotCloudProviderConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_cloudprovider_spot.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSpotCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckSpotCloudProviderExists(t, resourceName, name),
				),
			},
			{
				Config: testAccResourceSpotCloudProvider(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotCloudProviderExists(t, resourceName, name),
				),
				ExpectError: regexp.MustCompile("name is immutable"),
			},
		},
	})
}

func testAccResourceSpotCloudProvider(name string) string {
	return fmt.Sprintf(`
	 	data "harness_secret_manager" "default" {
			default = true
		}

    resource "harness_encrypted_text" "test" {
			name = "%s"
			secret_manager_id = data.harness_secret_manager.default.id
			value = "%[3]s"
		}

		resource "harness_cloudprovider_spot" "test" {
			name = "%[1]s"
			account_id = "%[2]s"
			token_secret_name = harness_encrypted_text.test.name
		}
`, name, helpers.TestEnvVars.SpotInstAccountId.Get(), helpers.TestEnvVars.SpotInstToken.Get())
}

func testAccCheckSpotCloudProviderExists(t *testing.T, resourceName, cloudProviderName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cp := &cac.SpotInstCloudProvider{}
		err := testAccGetCloudProvider(resourceName, state, cp)
		if err != nil {
			return err
		}
		return nil
	}
}
