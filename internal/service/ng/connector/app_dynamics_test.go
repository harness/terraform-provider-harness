package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnector_appd_usernamepassword(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_appdynamics_usernamepassword(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.url", "https://appdynamics.com/"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.account_name", "myaccount"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.username_password.0.password_ref", "account.acctest_appd_password"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_appdynamics_usernamepassword(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.url", "https://appdynamics.com/"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.account_name", "myaccount"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.username_password.0.password_ref", "account.acctest_appd_password"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.delegate_selectors.#", "1"),
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

func TestAccResourceConnector_appd_token(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_appdynamics_token(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.url", "https://appdynamics.com/"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.account_name", "myaccount"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.api_token.0.client_id", "admin"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.api_token.0.client_secret_ref", "account.acctest_appd_password"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_appdynamics_token(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.url", "https://appdynamics.com/"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.account_name", "myaccount"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.api_token.0.client_id", "admin"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.api_token.0.client_secret_ref", "account.acctest_appd_password"),
					resource.TestCheckResourceAttr(resourceName, "app_dynamics.0.delegate_selectors.#", "1"),
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

func testAccResourceConnector_appdynamics_usernamepassword(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			app_dynamics {
				url = "https://appdynamics.com/"
				account_name = "myaccount"
				delegate_selectors = ["harness-delegate"]
				username_password {
					username = "admin"
					password_ref = "account.acctest_appd_password"
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_appdynamics_token(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			app_dynamics {
				url = "https://appdynamics.com/"
				account_name = "myaccount"
				delegate_selectors = ["harness-delegate"]
				api_token {
					client_id = "admin"
					client_secret_ref = "account.acctest_appd_password"
				}
			}
		}
`, id, name)
}
