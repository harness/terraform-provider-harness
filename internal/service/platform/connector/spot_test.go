package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorSpot_DelegateNull(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	execute(t, testAccResourceConnectorSpotPermanentTokenWithNullExecuteOnDelegate(id, name), "true", id, name)
}

func TestAccResourceConnectorSpot_DelegateTrue(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	execute(t, testAccResourceConnectorSpotPermanentTokenWithExecuteOnDelegateTrue(id, name), "true", id, name)
}

func TestAccResourceConnectorSpot_DelegateFalse(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	execute(t, testAccResourceConnectorSpotPermanentTokenWithExecuteOnDelegateFalse(id, name), "false", id, name)
}

func execute(t *testing.T, config string, executeOnDelegate string, id string, name string) {
	resourceName := "harness_platform_connector_spot.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permanent_token.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "permanent_token.0.spot_account_id_ref", "account.TEST_spot_account_id"),
					resource.TestCheckResourceAttr(resourceName, "permanent_token.0.execute_on_delegate", executeOnDelegate),
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

func testAccResourceConnectorSpotPermanentTokenWithNullExecuteOnDelegate(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_spot" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			permanent_token {
				spot_account_id_ref = "account.TEST_spot_account_id"
				api_token_ref = "account.TEST_spot_api_token"
				delegate_selectors = ["harness-delegate"]
			}
		}
`, id, name)
}

func testAccResourceConnectorSpotPermanentTokenWithExecuteOnDelegateTrue(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_spot" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			permanent_token {
				spot_account_id_ref = "account.TEST_spot_account_id"
				api_token_ref = "account.TEST_api_token_ref"
				delegate_selectors = ["harness-delegate"]
				execute_on_delegate = true
			}
		}
`, id, name)
}

func testAccResourceConnectorSpotPermanentTokenWithExecuteOnDelegateFalse(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_spot" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			permanent_token {
				spot_account_id_ref = "account.TEST_spot_account_id"
				api_token_ref = "account.TEST_api_token_ref"
				delegate_selectors = ["harness-delegate"]
				execute_on_delegate = false
			}

		}
`, id, name)
}
