package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitConnector(t *testing.T) {

	var (
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "data.harness_git_connector.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitConnector(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "branch", "master"),
					resource.TestCheckResourceAttr(resourceName, "url_type", "REPO"),
					resource.TestCheckResourceAttr(resourceName, "username", "someuser"),
				),
			},
		},
	})
}

func testAccDataSourceGitConnector(name string) string {
	return fmt.Sprintf(`
	data "harness_secret_manager" "test" {
		default = true
	}

	resource "harness_encrypted_text" "test" {
		name 							= "%[1]s"
		value 					  = "foo"
		secret_manager_id = data.harness_secret_manager.test.id
	}

	resource "harness_git_connector" "test" {
		name = "%[1]s"
		url = "https://github.com/micahlmartin/harness-demo"
		branch = "master"
		password_secret_id = harness_encrypted_text.test.id
		url_type = "REPO"
		username = "someuser"
	}	

		data "harness_git_connector" "test" {
			id = harness_git_connector.test.id
		}
`, name)
}
