package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceSpotCloudProviderConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
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
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceSpotCloudProviderConnector_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_spot.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSpotCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckSpotCloudProviderExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					testAccConfigureProvider()
					c := testAccProvider.Meta().(*api.Client)
					cp, err := c.CloudProviders().GetSpotInstCloudProviderByName(name)
					require.NoError(t, err)
					require.NotNil(t, cp)

					err = c.CloudProviders().DeleteCloudProvider(cp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceSpotCloudProvider(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
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

		r := testAccGetResource(resourceName, state)
		name := r.Primary.Attributes["name"]

		require.NoError(t, err)
		require.Equal(t, cac.ObjectTypes.SpotInstCloudProvider, cp.Type)
		require.Equal(t, cloudProviderName, cp.Name)
		require.Equal(t, helpers.TestEnvVars.SpotInstAccountId.Get(), cp.AccountId)
		require.Equal(t, name, cp.Token.Name)

		return nil
	}
}
