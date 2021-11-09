package cloudprovider_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceTanzuCloudProviderConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_cloudprovider_tanzu.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTanzuCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckTanzuCloudProviderExists(t, resourceName, name),
				),
			},
			{
				Config: testAccResourceTanzuCloudProvider(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTanzuCloudProviderExists(t, resourceName, name),
				),
			},
		},
	})
}

func TestAccResourceTanzuCloudProviderConnector_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_tanzu.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTanzuCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckTanzuCloudProviderExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*api.Client)
					cp, err := c.CDClient.CloudProviderClient.GetPcfCloudProviderByName(name)
					require.NoError(t, err)
					require.NotNil(t, cp)

					err = c.CDClient.CloudProviderClient.DeleteCloudProvider(cp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceTanzuCloudProvider(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceTanzuCloudProvider(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_cloudprovider_tanzu" "test" {
			name = "%[1]s"
			endpoint = "https://endpoint.com"
			skip_validation = true
			username = "username"
			password_secret_name = harness_encrypted_text.test.name
		}
`, name)
}

func testAccCheckTanzuCloudProviderExists(t *testing.T, resourceName, cloudProviderName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cp := &cac.PcfCloudProvider{}
		err := testAccGetCloudProvider(resourceName, state, cp)
		if err != nil {
			return err
		}
		return nil
	}
}
