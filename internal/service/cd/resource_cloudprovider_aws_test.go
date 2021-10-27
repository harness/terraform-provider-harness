package cd_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceAwsCloudProvider(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(7))
		resourceName = "harness_cloudprovider_aws.test"
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAwsCloudProviderExists(t, resourceName, name),
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

func TestAccResourceAwsCloudProvider_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsCloudProvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAwsCloudProviderExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*api.Client)
					cp, err := c.CloudProviders().GetAwsCloudProviderByName(name)
					require.NoError(t, err)
					require.NotNil(t, cp)

					err = c.CloudProviders().DeleteCloudProvider(cp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceAwsCloudProvider(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceAwsCloudProvider(name string) string {
	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "aws_access_key" {
		name = "%[1]s_access_key"
		value = "%[2]s"
		secret_manager_id = data.harness_secret_manager.default.id
	}

	resource "harness_encrypted_text" "aws_secret_key" {
		name = "%[1]s_secret_key"
		value = "%[3]s"
		secret_manager_id = data.harness_secret_manager.default.id
	}
	
	resource "harness_cloudprovider_aws" "test" {
		name = "%[1]s"
		access_key_id_secret_name = harness_encrypted_text.aws_access_key.name
		secret_access_key_secret_name = harness_encrypted_text.aws_secret_key.name
	}	
`, name, helpers.TestEnvVars.AwsAccessKeyId.Get(), helpers.TestEnvVars.AwsSecretAccessKey.Get())
}

func testAccCheckAwsCloudProviderExists(t *testing.T, resourceName, cloudProviderName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cp := &cac.AwsCloudProvider{}
		err := testAccGetCloudProvider(resourceName, state, cp)
		if err != nil {
			return err
		}
		return nil
	}
}
