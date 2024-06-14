package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorNexus_Anonymous(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_nexus.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorNexus_anonymous(id, name, nextgen.NexusVersions.V2X.String()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "version", nextgen.NexusVersions.V2X.String()),
					resource.TestCheckResourceAttr(resourceName, "url", "https://nexus.example.com"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorNexus_anonymous(id, updatedName, nextgen.NexusVersions.V3X.String()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "version", nextgen.NexusVersions.V3X.String()),
					resource.TestCheckResourceAttr(resourceName, "url", "https://nexus.example.com"),
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

func TestAccResourceConnectorNexus_UsernamePassword(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_nexus.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorNexus_usernamepassword(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://nexus.example.com"),
					resource.TestCheckResourceAttr(resourceName, "version", "3.x"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "admin"),
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

func testAccResourceConnectorNexus_usernamepassword(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_nexus" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://nexus.example.com"
			delegate_selectors = ["harness-delegate"]
			version = "3.x"
			credentials {
				username = "admin"
				password_ref = "account.doNotDeleteHSM"
			}
		}
`, id, name)
}

func testAccResourceConnectorNexus_anonymous(id string, name string, version string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_nexus" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://nexus.example.com"
			version = "%[3]s"
			delegate_selectors = ["harness-delegate"]
		}
`, id, name, version)
}
