package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorRancher_BearerToken(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_rancher.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorRancher_BearerToken(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bearer_token.0.rancher_url", "https://rancher.cluster.example"),
					resource.TestCheckResourceAttr(resourceName, "bearer_token.0.password_ref", fmt.Sprintf("account.%s", id)),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorRancher_BearerToken(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bearer_token.0.rancher_url", "https://rancher.cluster.example"),
					resource.TestCheckResourceAttr(resourceName, "bearer_token.0.password_ref", fmt.Sprintf("account.%s", id)),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
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

func testAccResourceConnectorRancher_BearerToken(id string, name string) string {
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

		resource "harness_platform_connector_rancher" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			delegate_selectors = ["harness-delegate"]

			bearer_token {
				rancher_url = "https://rancher.cluster.example"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}

			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
	`, id, name)
}
