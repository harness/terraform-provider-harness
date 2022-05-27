package application_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceApplication(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_application.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication(expectedName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					testAccApplicationCreation(t, resourceName, expectedName),
				),
			},
			{
				Config: testAccResourceApplication(updatedName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
					testAccApplicationCreation(t, resourceName, updatedName),
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

func TestAccResourceApplication_DeleteUnderlyingResource(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	resourceName := "harness_application.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication(expectedName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					testAccApplicationCreation(t, resourceName, expectedName),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*internal.Session).CDClient
					app, err := c.ApplicationClient.GetApplicationByName(expectedName)
					require.NoError(t, err)
					require.NotNil(t, app)

					err = c.ApplicationClient.DeleteApplication(app.Id)
					require.NoError(t, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Config:             testAccResourceApplication(expectedName, true),
			},
		},
	})
}

func testAccApplicationCreation(t *testing.T, resourceName string, appName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, err := acctest.TestAccGetApplication(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, appName, app.Name)

		return nil
	}
}

func testAccApplicationDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, _ := acctest.TestAccGetApplication(resourceName, state)
		if app != nil {
			return fmt.Errorf("Found application: %s", app.Id)
		}

		return nil
	}
}

func testAccResourceApplication(name string, withTags bool) string {

	tags := ""
	if withTags {
		tags = `tags = ["test:val", "foo:bar"]`
	}

	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%s"
			description = "my description"

			%s
		}
`, name, tags)
}
