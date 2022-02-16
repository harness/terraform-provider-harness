package environment_test

import (
	"fmt"
	"testing"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceEnvironment(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "harness_environment.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironment(name, cac.EnvironmentTypes.NonProd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckEnvironmentExists(t, resourceName, name, cac.EnvironmentTypes.NonProd),
				),
			},
			{
				Config: testAccResourceEnvironment(name, cac.EnvironmentTypes.Prod),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckEnvironmentExists(t, resourceName, name, cac.EnvironmentTypes.Prod),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.ID
					app_id := primary.Attributes["app_id"]
					return fmt.Sprintf("%s/%s", app_id, id), nil
				},
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceEnvironment_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "harness_environment.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironment(name, cac.EnvironmentTypes.NonProd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckEnvironmentExists(t, resourceName, name, cac.EnvironmentTypes.NonProd),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)
					app, err := c.CDClient.ApplicationClient.GetApplicationByName(name)
					require.NoError(t, err)
					require.NotNil(t, app)

					env, err := c.CDClient.ConfigAsCodeClient.GetEnvironmentByName(app.Id, name)
					require.NoError(t, err)
					require.NotNil(t, env)

					err = c.CDClient.ConfigAsCodeClient.DeleteEnvironment(app.Name, env.Name)
					require.NoError(t, err)
				},
				Config:             testAccResourceEnvironment(name, cac.EnvironmentTypes.NonProd),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetEnvironment(resourceName string, state *terraform.State) (*cac.Environment, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	svcId := r.Primary.ID
	appId := r.Primary.Attributes["app_id"]

	return c.CDClient.ConfigAsCodeClient.GetEnvironmentById(appId, svcId)
}

func testAccCheckEnvironmentExists(t *testing.T, resourceName string, name string, envType cac.EnvironmentType) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, err := testAccGetEnvironment(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, env)
		require.Equal(t, name, env.Name)
		require.Equal(t, envType, env.EnvironmentType)
		// require.Len(t, env.)

		return nil
	}
}

func testAccEnvironmentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetEnvironment(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment: %s", env.Id)
		}

		return nil
	}
}

func testAccResourceEnvironment(name string, envType cac.EnvironmentType) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_service_kubernetes" "test" {
			app_id = harness_application.test.id
			name = "%[1]s"
			helm_version = "V2"
			description = "description"
			
			variable {
				name = "test"
				value = "test_value"
				type = "TEXT"
			}

			variable {
				name = "test2"
				value = "test_value2"
				type = "TEXT"
			}
		}

		resource "harness_environment" "test" {
			app_id = harness_application.test.id
			name = "%[1]s"
			type = "%[2]s"

			variable_override {
				service_name = harness_service_kubernetes.test.name
				name = "test"
				value = "override"
				type = "TEXT"
			}

			variable_override {
				name = "test2"
				value = "override2"
				type = "TEXT"
			}
		}
`, name, envType)
}
