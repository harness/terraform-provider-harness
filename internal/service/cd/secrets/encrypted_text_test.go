package secrets_test

import (
	"fmt"
	"testing"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

const (
	awsSecretManager = "LoyX8CUFQWm3pk_kCLPn2w"
)

func TestAccResourceEncryptedText(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"
	updatedValue := value + "-updated"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText(name, value, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
				),
			},
			{
				Config: testAccResourceEncryptedText(name, updatedValue, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func TestAccResourceEncryptedText_Reference(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"
	updatedValue := value + "-updated"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText_Reference(name, value, awsSecretManager, "test/secret/micah"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "secret_reference", "test/secret/micah"),
				),
			},
			{
				Config: testAccResourceEncryptedText_Reference(name, updatedValue, awsSecretManager, "test/secret/micah2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "secret_reference", "test/secret/micah2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret_reference"},
			},
		},
	})
}

func TestAccResourceEncryptedText_secretmanagerid_immutable(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText(name, value, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
				),
			},
			{
				Config:             testAccResourceEncryptedText(name, value, "foo"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceEncryptedText_secretmanagerid_DeleteUnderlyingResource(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText(name, value, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)
					secret, err := c.CDClient.SecretClient.GetEncryptedTextByName(name)
					require.NoError(t, err)
					require.NotNil(t, secret)

					err = c.CDClient.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
					require.NoError(t, err)
				},
				Config:             testAccResourceEncryptedText(name, value, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceEncryptedText_UsageScopes(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s-updated", name)
	resourceName := "harness_encrypted_text.usage_scope_test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText_UsageScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.environment_filter_type", graphql.EnvironmentFilterTypes.Production.String()),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction.String()),
					func(state *terraform.State) error {
						et, err := testAccGetEncryptedText(resourceName, state)
						require.NoError(t, err)
						require.NotNil(t, et)
						require.Equal(t, name, et.Name)
						require.NotEmpty(t, et.SecretManagerId)
						return nil
					},
				),
			},
			{
				Config: testAccResourceEncryptedText_usageScopes_update(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.Production.String()),
					resource.TestCheckNoResourceAttr(resourceName, "usage_scope.1"),
					func(state *terraform.State) error {
						et, err := testAccGetEncryptedText(resourceName, state)
						require.NoError(t, err)
						require.NotNil(t, et)
						require.Equal(t, updatedName, et.Name)
						require.NotEmpty(t, et.SecretManagerId)
						return nil
					},
				),
			},
		},
	})
}

func testAccResourceEncryptedText(name string, value string, secretMangerId string) string {

	if secretMangerId == "" {
		secretMangerId = "data.harness_secret_manager.default.id"
	} else {
		secretMangerId = fmt.Sprintf("\"%s\"", secretMangerId)
	}

	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

	
		resource "harness_encrypted_text" "test" {
			name = "%s"
			value = "%s"
			secret_manager_id = %[3]s
		}
`, name, value, secretMangerId)
}

func testAccResourceEncryptedText_Reference(name string, value string, secretMangerId string, ref string) string {

	if secretMangerId == "" {
		secretMangerId = "data.harness_secret_manager.default.id"
	} else {
		secretMangerId = fmt.Sprintf("\"%s\"", secretMangerId)
	}

	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

	
		resource "harness_encrypted_text" "test" {
			name = "%s"
			secret_manager_id = %[3]s
			secret_reference = "%[4]s"
		}
`, name, value, secretMangerId, ref)
}

func testAccResourceEncryptedText_UsageScopes(name string) string {
	// nonprod :=
	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"
		secret_manager_id = data.harness_secret_manager.default.id

		usage_scope {
			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
		}

		usage_scope {
			environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
		}
	}
	`, name)
}

func testAccResourceEncryptedText_usageScopes_update(name string) string {
	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"
		secret_manager_id = data.harness_secret_manager.default.id
		usage_scope {
			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
		}
	}
	`, name)
}

func testAccGetEncryptedText(resourceName string, state *terraform.State) (*graphql.EncryptedText, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.CDClient.SecretClient.GetEncryptedTextById(id)
}
