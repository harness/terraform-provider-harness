package cd_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
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

// func TestAccResourceApplication_DeleteUnderlyingResource(t *testing.T) {

// 	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
// 	resourceName := "harness_application.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		CheckDestroy:      testAccApplicationDestroy(resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceApplication(expectedName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
// 					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
// 					testAccApplicationCreation(t, resourceName, expectedName),
// 				),
// 			},
// 			{
// 				PreConfig: func() {
// 					acctest.TestAccConfigureProvider()
// 					c := acctest.TestAccProvider.Meta().(*api.Client)
// 					app, err := c.Applications().GetApplicationByName(expectedName)
// 					require.NoError(t, err)
// 					require.NotNil(t, app)

// 					err = c.Applications().DeleteApplication(app.Id)
// 					require.NoError(t, err)
// 				},
// 				PlanOnly:           true,
// 				ExpectNonEmptyPlan: true,
// 				Config:             testAccResourceApplication(expectedName),
// 			},
// 		},
// 	})
// }

// func testAccApplicationCreation(t *testing.T, resourceName string, appName string) resource.TestCheckFunc {
// 	return func(state *terraform.State) error {
// 		app, err := testAccGetApplication(resourceName, state)
// 		require.NoError(t, err)
// 		require.NotNil(t, app)
// 		require.Equal(t, appName, app.Name)

// 		return nil
// 	}
// }

// func testAccGetApplication(resourceName string, state *terraform.State) (*graphql.Application, error) {
// 	r := acctest.TestAccGetResource(resourceName, state)
// 	c := testAccGetApiClientFromProvider()
// 	id := r.Primary.ID

// 	return c.Applications().GetApplicationById(id)
// }

func testAccApplicationGitSyncDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, _ := testAccGetApplication(resourceName, state)
		if app != nil {
			return fmt.Errorf("Found application: %s", app.Id)
		}

		return nil
	}
}

func testAccResourceApplicationGitSync(name string) string {
	connector := testAccResourceGitConnector_default(name)
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
