package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceEncryptedText(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"
	updatedValue := value + "-updated"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText(name, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
					// resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
				),
			},
			{
				Config: testAccResourceEncryptedText(name, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
					// resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
				),
			},
		},
	})
}

// func TestAccResourceEncryptedText_secretmanagerid_immutable(t *testing.T) {

// 	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
// 	resourceName := "harness_encrypted_text.test"
// 	value := "someval"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceEncryptedText(name, value),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "value", value),
// 					resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
// 				),
// 			},
// 			{
// 				Config:      testAccResourceEncryptedText_update_secretmanagerid(name, value, "bad_value"),
// 				ExpectError: regexp.MustCompile("secret_manager_id is immutable and cannot be changed once set"),
// 			},
// 		},
// 	})
// }

// func TestAccResourceEncryptedText_UsageScopes(t *testing.T) {

// 	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
// 	updatedName := fmt.Sprintf("%s-updated", name)
// 	resourceName := "harness_encrypted_text.usage_scope_test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceEncryptedText_UsageScopes(name),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.application_filter_type", graphql.ApplicationFilterTypes.All),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.environment_filter_type", graphql.EnvironmentFilterTypes.Production),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.NonProduction),
// 					func(state *terraform.State) error {
// 						et, err := testAccGetEncryptedText(resourceName, state)
// 						require.NoError(t, err)
// 						require.NotNil(t, et)
// 						require.Equal(t, name, et.Name)
// 						require.NotEmpty(t, et.SecretManagerId)
// 						return nil
// 					},
// 				),
// 			},
// 			{
// 				Config: testAccResourceEncryptedText_usageScopes_update(updatedName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All),
// 					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", graphql.EnvironmentFilterTypes.Production),
// 					resource.TestCheckNoResourceAttr(resourceName, "usage_scope.1"),
// 					func(state *terraform.State) error {
// 						et, err := testAccGetEncryptedText(resourceName, state)
// 						require.NoError(t, err)
// 						require.NotNil(t, et)
// 						require.Equal(t, updatedName, et.Name)
// 						require.NotEmpty(t, et.SecretManagerId)
// 						return nil
// 					},
// 				),
// 			},
// 		},
// 	})
// }

func testAccResourceEncryptedText(name string, value string) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%s"
			value = "%s"

			lifecycle {
				ignore_changes = [secret_manager_id]
			}
		}

		output "somvar" {
			value = harness_encrypted_text.test.secret_manager_id
		}
`, name, value)
}

func testAccResourceEncryptedText_UsageScopes(name string) string {
	// nonprod :=
	return fmt.Sprintf(`
	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"

		lifecycle {
			ignore_changes = [secret_manager_id]
		}

		usage_scope {
			application_filter_type = "ALL"
			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
		}

		usage_scope {
			application_filter_type = "ALL"
			environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
		}
	}
	`, name)
}

func testAccResourceEncryptedText_usageScopes_update(name string) string {
	// nonprod :=
	return fmt.Sprintf(`
	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"

		lifecycle {
			ignore_changes = [secret_manager_id]
		}

		usage_scope {
			application_filter_type = "ALL"
			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
		}
	}
	`, name)
}

func testAccResourceEncryptedText_update_secretmanagerid(name string, value string, secretManagerId string) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%s"
			value = "%s"
			secret_manager_id = "%s"

			lifecycle {
				ignore_changes = [secret_manager_id]
			}
		}
`, name, value, secretManagerId)
}

func testAccGetEncryptedText(resourceName string, state *terraform.State) (*graphql.EncryptedText, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Secrets().GetEncryptedTextById(id)
}
