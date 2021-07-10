package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("harness_cloudprovider_aws", &resource.Sweeper{
		Name: "harness_cloudprovider_aws",
		F:    testSweepCloudProviders,
	})
}

func TestAccResourceAwsCloudProvider(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_aws.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsCloudProvider(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAwsCloudProviderExists(t, resourceName, name),
				),
			},
		},
	})
}

func testAccResourceAwsCloudProvider(name string, useAccessKeys bool) string {

	credentialsConfig := ""

	if useAccessKeys {
		credentialsConfig = fmt.Sprintf(`
			access_keys {
				access_key_id = "%[1]s"
				encrypted_secret_access_key_secret_name = harness_encrypted_text.test.name
			}
		`, helpers.TestEnvVars.AwsAccessKeyId.Get())
	}

	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "test" {
		name = "%[1]s"
		value = "%[3]s"
		secret_manager_id = data.harness_secret_manager.default.id
	}
	
	resource "harness_cloudprovider_aws" "test" {
		name = "%[1]s"

		credentials {
			%[2]s
		}
	}	
`, name, credentialsConfig, helpers.TestEnvVars.AwsSecretAccessKey.Get())
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
