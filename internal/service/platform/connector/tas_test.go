package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/require"
)

func TestAccResourceConnectorTas_ManualDetails_Secret(t *testing.T) {

	id := fmt.Sprintf("ConnectorTas_ManualDetails_Secret"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_tas.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorTas_manualDetails_secret(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
				),
			},
			{
				Config: testAccResourceConnectorTas_manualDetails_secret(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
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

func TestAccResourceConnectorTas_ManualDetails_UserNameRef(t *testing.T) {

	id := fmt.Sprintf("ConnectorTas_ManualDetails_UserNameRef"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_tas.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorTas_manualDetails_username_ref(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username_ref", "account."+id+"_username"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
				),
			},
			{
				Config: testAccResourceConnectorTas_manualDetails_username_ref(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username_ref", "account."+id+"_username"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
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

func TestAccResourceConnectorTas_DeleteUnderlyingResource(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_tas.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorTas_manualDetails_username_ref(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username_ref", "account."+id+"_username"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.ConnectorsApi.DeleteConnector(ctx, c.AccountId, id, &nextgen.ConnectorsApiDeleteConnectorOpts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceConnectorTas_manualDetails_username_ref(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceConnectorTas_manualDetails_secret(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_tas" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "ManualConfig"
			tas_manual_details {
				endpoint_url = "https://tas.example.com"
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		delegate_selectors = ["harness-delegate"]
	}
`, id, name)
}

func testAccResourceConnectorTas_manualDetails_username_ref(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}
	resource "harness_platform_secret_text" "username_ref" {
		identifier = "%[1]s_username"
		name = "%[2]s_username"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_tas" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "ManualConfig"
			tas_manual_details {
				endpoint_url = "https://tas.example.com"
				username_ref = "account.${harness_platform_secret_text.username_ref.id}"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		delegate_selectors = ["harness-delegate"]
	}
`, id, name)
}
