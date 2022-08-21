package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorHCVAgent(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_hcvagent.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorHCVAgent(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://fake.vault.io:8200"),
					resource.TestCheckResourceAttr(resourceName, "base_secret_path", "/testseretpath"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "test_namespace"),
					resource.TestCheckResourceAttr(resourceName, "sink_path", "/testsinkpath"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorHCVAgent(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://fake.vault.io:8200"),
					resource.TestCheckResourceAttr(resourceName, "base_secret_path", "/testseretpath"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "test_namespace"),
					resource.TestCheckResourceAttr(resourceName, "sink_path", "/testsinkpath"),
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

func testAccResourceConnectorHCVAgent(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_hcvagent" "test" {
			identifier         = "%[1]s"
			name               = "%[2]s"
			description        = "test"
			tags               = ["foo:bar"]
			url                = "https://fake.vault.io:8200"
			base_secret_path   = "/testseretpath"
			namespace          = "test_namespace"
			sink_path          = "/testsinkpath"
			delegate_selectors = ["harness-delegate"]
		  }
`, id, name)
}
