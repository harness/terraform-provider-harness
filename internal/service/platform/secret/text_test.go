package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/require"
)

func TestAccSecretText_inline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := secretValue + "updated"
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_text_inline(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				Config: testAccResourceSecret_text_inline(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func TestAccResourceSecretText_reference(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := secretValue + "updated"
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_text_reference(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "azureSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", secretValue),
				),
			},
			{
				Config: testAccResourceSecret_text_reference(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "azureSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}
func TestProjectSecretText_inline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := secretValue + "updated"
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceSecretText_inline(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
					resource.TestCheckResourceAttr(resourceName, "value", secretValue),
				),
			},
			{
				Config: testProjectResourceSecretText_inline(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func TestOrgSecretText_inline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := secretValue + "updated"
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceSecretText_inline(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				Config: testOrgResourceSecretText_inline(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Inline"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func TestOrgResourceSecretText_reference(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedValue := secretValue + "updated"
	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecretDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceSecretText_reference(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", id),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", secretValue),
				),
			},
			{
				Config: testOrgResourceSecretText_reference(id, updatedName, updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", id),
					resource.TestCheckResourceAttr(resourceName, "value_type", "Reference"),
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"value"},
			},
		},
	})
}

func TestAccSecretText_DeleteUnderLyingResource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	secretValue := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	resourceName := "harness_platform_secret_text.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecret_text_inline(id, name, secretValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.SecretsApi.DeleteSecretV2(ctx, id, c.AccountId, &nextgen.SecretsApiDeleteSecretV2Opts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceSecret_text_inline(id, name, secretValue),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceSecret_text_inline(id string, name string, secretValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "%[3]s"
		}
`, id, name, secretValue)
}

func testAccResourceSecret_text_reference(id string, name string, secretValue string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "azureSecretManager"
			value_type = "Reference"
			value = "%[3]s"
		}
`, id, name, secretValue)
}

func testProjectResourceSecretText_inline(id string, name string, secretValue string) string {
	return fmt.Sprintf(`
  resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  resource "harness_platform_project" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    org_id = harness_platform_organization.test.id
    color = "#472848"
	}
	
    resource "harness_platform_secret_text" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      description = "test"
      tags = ["foo:bar"]
      secret_manager_identifier = "harnessSecretManager"
      value_type = "Inline"
      value = "%[3]s"
      org_id = harness_platform_organization.test.id
      project_id = harness_platform_project.test.id
    	depends_on = [time_sleep.wait_3_seconds]
	}
	resource "time_sleep" "wait_3_seconds" {
		create_duration = "3s"
		depends_on = [harness_platform_project.test]
		
	}
`, id, name, secretValue)
}

func testOrgResourceSecretText_inline(id string, name string, secretValue string) string {
	return fmt.Sprintf(`
  resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }
  
    resource "harness_platform_secret_text" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      description = "test"
      tags = ["foo:bar"]
      secret_manager_identifier = "harnessSecretManager"
      value_type = "Inline"
      value = "%[3]s"
      org_id = harness_platform_organization.test.id
      depends_on = [time_sleep.wait_4_seconds]
    }
    resource "time_sleep" "wait_4_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
		}
`, id, name, secretValue)
}

func testOrgResourceSecretText_reference(id string, name string, secretValue string) string {
	return fmt.Sprintf(`
  resource "harness_platform_organization" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
  }

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		org_id = harness_platform_organization.test.id

		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.azuretest"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"
		is_default = false

		azure_environment_type = "AZURE"
	  depends_on = [time_sleep.wait_3_seconds]
    }
    resource "time_sleep" "wait_3_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_organization.test]
			
		}

    resource "harness_platform_secret_text" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      description = "test"
      tags = ["foo:bar"]

      secret_manager_identifier = "%[1]s"
      value_type = "Reference"
      value = "%[3]s"
      org_id = harness_platform_organization.test.id
      depends_on = [time_sleep.wait_4_seconds]
    }
    resource "time_sleep" "wait_4_seconds" {
			create_duration = "3s"
			depends_on = [harness_platform_connector_azure_key_vault.test]
			
		}
`, id, name, secretValue)
}
