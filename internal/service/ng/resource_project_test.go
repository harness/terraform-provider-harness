package ng_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceProject(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_project.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccProjectDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProject(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceProject(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.ID
					org_id := primary.Attributes["org_id"]
					return fmt.Sprintf("%s/%s", org_id, id), nil
				},
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
// 					testAccConfigureProvider()
// 					c := testAccProvider.Meta().(*api.Client)
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

// func TestAccResourceApplication_Import(t *testing.T) {

// 	resourceName := "harness_application.test"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceApplication("test"),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
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
// 	r := testAccGetResource(resourceName, state)
// 	c := testAccGetApiClientFromProvider()
// 	id := r.Primary.ID

// 	return c.Applications().GetApplicationById(id)
// }

func testAccGetProject(resourceName string, state *terraform.State) (*nextgen.Project, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]

	resp, _, err := c.NGClient.ProjectApi.GetProject(context.Background(), id, c.AccountId, &nextgen.ProjectApiGetProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	if err != nil {
		return nil, err
	}

	return resp.Data.Project, nil
}

func testAccProjectDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		project, _ := testAccGetProject(resourceName, state)
		if project != nil {
			return fmt.Errorf("Found project: %s", project.Identifier)
		}

		return nil
	}
}

func testAccResourceProject(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = "test"
			color = "#0063F7"
		}
`, id, name)
}
