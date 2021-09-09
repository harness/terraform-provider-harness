package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func init() {
	resource.AddTestSweepers("harness_encrypted_text", &resource.Sweeper{
		Name: "harness_encrypted_text",
		F:    testSweepHarnessEncryptedText,
		Dependencies: []string{
			"harness_cloudprovider_aws",
			"harness_git_connector",
			"harness_ssh_credential",
			"harness_winrm_credential",
		},
	})
}

func testSweepHarnessEncryptedText(r string) error {
	c := testAccGetApiClientFromProvider()

	limit := 500
	offset := 0
	hasMore := true

	for hasMore {

		secrets, _, err := c.Secrets().ListEncryptedTextSecrets(limit, offset)

		if err != nil {
			return err
		}

		for _, secret := range secrets {
			if strings.HasPrefix(secret.Name, "Test") {
				if err = c.Secrets().DeleteSecret(secret.UUID, graphql.SecretTypes.EncryptedText); err != nil {
					return err
				}
			}
		}

		hasMore = len(secrets) == limit
	}

	return nil
}

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
		},
	})
}

func TestAccResourceEncryptedText_secretmanagerid_immutable(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.test"
	value := "someval"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
					testAccConfigureProvider()
					c := testAccProvider.Meta().(*api.Client)
					secret, err := c.Secrets().GetEncryptedTextByName(name)
					require.NoError(t, err)
					require.NotNil(t, secret)

					err = c.Secrets().DeleteSecret(secret.Id, secret.SecretType)
					require.NoError(t, err)
				},
				Config:             testAccResourceTanzuCloudProvider(name),
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText_UsageScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.application_filter_type", graphql.ApplicationFilterTypes.All.String()),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.environment_filter_type", graphql.EnvironmentFilterTypes.Production.String()),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All.String()),
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
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", graphql.ApplicationFilterTypes.All.String()),
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
	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"
		secret_manager_id = data.harness_secret_manager.default.id
		usage_scope {
			application_filter_type = "ALL"
			environment_filter_type = "PRODUCTION_ENVIRONMENTS"
		}
	}
	`, name)
}

func testAccGetEncryptedText(resourceName string, state *terraform.State) (*graphql.EncryptedText, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Secrets().GetEncryptedTextById(id)
}
