package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
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
					resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
				),
			},
			{
				Config: testAccResourceEncryptedText(name, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
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
				Config: testAccResourceEncryptedText(name, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
				),
			},
			{
				Config:      testAccResourceEncryptedText_update_secretmanagerid(name, value, "bad_value"),
				ExpectError: regexp.MustCompile("secret_manager_id is immutable and cannot be changed once set"),
			},
		},
	})
}

func TestAccResourceEncryptedText_UsageScopes(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_encrypted_text.usage_scope_test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEncryptedText_UsageScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.application_filter_type", client.ApplicationFilterTypes.All),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.1.environment_filter_type", client.EnvironmentFilterTypes.Production),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", client.ApplicationFilterTypes.All),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", client.EnvironmentFilterTypes.NonProduction),
				),
			},
			{
				Config: testAccResourceEncryptedText_usageScopes_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", client.ApplicationFilterTypes.All),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", client.EnvironmentFilterTypes.Production),
					resource.TestCheckNoResourceAttr(resourceName, "usage_scope.1"),
				),
			},
		},
	})
}

func testAccResourceEncryptedText(name string, value string) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name = "%s"
			value = "%s"
		}
`, name, value)
}

func testAccResourceEncryptedText_UsageScopes(name string) string {
	// nonprod :=
	return fmt.Sprintf(`
	resource "harness_encrypted_text" "usage_scope_test" {
		name = "%s"
		value = "someval"

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
		}
`, name, value, secretManagerId)
}
