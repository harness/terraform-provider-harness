package application_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceApplicationGitSync(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_application_gitsync.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApplicationGitSyncDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationGitSync(expectedName),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccApplicationGitSyncDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, _ := acctest.TestAccGetApplication(resourceName, state)
		if app != nil {
			return fmt.Errorf("Found application: %s", app.Id)
		}

		return nil
	}
}

func testAccResourceApplicationGitSync(name string) string {
	connector := acctest.TestAccResourceGitConnector_default(name)
	return fmt.Sprintf(`
		%[2]s

		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_application_gitsync" "test" {
			app_id = harness_application.test.id
			connector_id = harness_git_connector.test.id
			branch = "main"
			enabled = false
		}
`, name, connector)
}
