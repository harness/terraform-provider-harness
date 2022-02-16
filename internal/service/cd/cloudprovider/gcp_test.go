package cloudprovider_test

import (
	"fmt"
	"testing"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceGcpCloudProviderConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_cloudprovider_gcp.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
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
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceGcpCloudProviderConnector_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_gcp.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGcpCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckGcpCloudProviderExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)
					cp, err := c.CDClient.CloudProviderClient.GetGcpCloudProviderByName(name)
					require.NoError(t, err)
					require.NotNil(t, cp)

					err = c.CDClient.CloudProviderClient.DeleteCloudProvider(cp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceGcpCloudProvider(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
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

			usage_scope {
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}
			
			usage_scope {
				environment_filter_type = "PRODUCTION_ENVIRONMENTS"
			}
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
