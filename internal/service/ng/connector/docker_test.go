package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnector_docker_DockerHub(t *testing.T) {

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
				Config: testAccResourceConnector_docker_DockerHub(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.type", "DockerHub"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.url", "https://hub.docker.com"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.credentials.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.credentials.0.password_ref", "account.TEST_k8s_client_test"),
				),
			},
			{
				Config: testAccResourceConnector_docker_DockerHub(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.type", "DockerHub"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.url", "https://hub.docker.com"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.credentials.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.credentials.0.password_ref", "account.TEST_k8s_client_test"),
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

func TestAccResourceConnector_docker_DockerHub_Anonymous(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_docker_anonymous(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.type", "DockerHub"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.url", "https://hub.docker.com"),
					resource.TestCheckResourceAttr(resourceName, "docker_registry.0.delegate_selectors.#", "1"),
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

func testAccResourceConnector_docker_DockerHub(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			docker_registry {
				type = "DockerHub"
				url = "https://hub.docker.com"
				delegate_selectors = ["harness-delegate"]
				credentials {
					username = "admin"
					password_ref = "account.TEST_k8s_client_test"
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_docker_anonymous(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			docker_registry {
				type = "DockerHub"
				url = "https://hub.docker.com"
				delegate_selectors = ["harness-delegate"]
			}
		}
`, id, name)
}
