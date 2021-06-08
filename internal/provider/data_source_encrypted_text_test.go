package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
)

func TestAccDataSourceEncryptedTextById(t *testing.T) {

	var (
		expectedName            = "somesecret"
		expectedId              = "MPuZFELfRO-q6rqTcLwFLg"
		expectedSecretManagerId = os.Getenv(envvar.HarnessAccountId)
	)

	resourceName := "data.harness_encrypted_text.secret_by_id"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedTextById(expectedId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile(expectedId)),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile(expectedName)),
					resource.TestMatchResourceAttr(resourceName, "secret_manager_id", regexp.MustCompile(expectedSecretManagerId)),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.application_filter_type", client.ApplicationFilterTypes.All),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.application_id", ""),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.environment_filter_type", client.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.environment_id", ""),
				),
			},
		},
	})
}

func TestAccDataSourceEncryptedTextByName(t *testing.T) {

	var (
		expectedName            = "somesecret"
		expectedId              = "MPuZFELfRO-q6rqTcLwFLg"
		expectedSecretManagerId = os.Getenv(envvar.HarnessAccountId)
	)

	resourceName := "data.harness_encrypted_text.secret_by_name"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedTextByName(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile(expectedId)),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile(expectedName)),
					resource.TestMatchResourceAttr(resourceName, "secret_manager_id", regexp.MustCompile(expectedSecretManagerId)),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.application_filter_type", client.ApplicationFilterTypes.All),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.application_id", ""),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.environment_filter_type", client.EnvironmentFilterTypes.NonProduction),
					resource.TestCheckResourceAttr(resourceName, "usage_scopes.0.environment_id", ""),
				),
			},
		},
	})
}

func testAccDataSourceEncryptedTextById(secretId string) string {
	return fmt.Sprintf(`
		data "harness_encrypted_text" "secret_by_id" {
			id = "%s"
		}
	`, secretId)
}

func testAccDataSourceEncryptedTextByName(name string) string {
	return fmt.Sprintf(`
		data "harness_encrypted_text" "secret_by_name" {
			name = "%s"
		}
	`, name)
}
