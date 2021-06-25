package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceApplication(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_application.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccApplicationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
					testAccApplicationCreation(t, resourceName, expectedName),
				),
			},
			{
				Config: testAccResourceApplication(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "my description"),
					testAccApplicationCreation(t, resourceName, updatedName),
				),
			},
		},
	})
}

func testAccApplicationCreation(t *testing.T, resourceName string, appName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, err := testAccGetApplication(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, appName, app.Name)

		return nil
	}
}

func testAccGetApplication(resourceName string, state *terraform.State) (*graphql.Application, error) {
	r := testAccGetResource(resourceName, state)
	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID

	return c.Applications().GetApplicationById(id)
}

func testAccApplicationDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, _ := testAccGetApplication(resourceName, state)
		if app != nil {
			return fmt.Errorf("Found application: %s", app.Id)
		}

		return nil
	}
}

func testAccResourceApplication(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%s"
			description = "my description"
		}
`, name)
}
