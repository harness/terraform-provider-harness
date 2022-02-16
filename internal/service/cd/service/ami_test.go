package service_test

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

func TestAccResourceAMIService(t *testing.T) {

	var (
		name               = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		description        = "some description"
		updatedDescription = "updated description"
		resourceName       = "harness_service_ami.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAMIService(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAMIServiceExists(t, resourceName, name, description),
				),
			},
			{
				Config: testAccResourceAMIService(name, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAMIServiceExists(t, resourceName, name, updatedDescription),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: serviceImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceAMIService_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		description  = "some description"
		resourceName = "harness_service_ami.test"
		serviceId    = ""
		appId        = ""
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAMIService(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAMIServiceExists(t, resourceName, name, description),
					func(state *terraform.State) error {
						svc, _ := testAccGetService(resourceName, state)
						serviceId = svc.Id
						appId = svc.ApplicationId
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)
					svc, err := c.CDClient.ConfigAsCodeClient.GetServiceById(appId, serviceId)
					require.NoError(t, err)
					require.NotNil(t, svc)

					err = c.CDClient.ConfigAsCodeClient.DeleteService(svc.ApplicationId, svc.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceAMIService(name, description),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckAMIServiceExists(t *testing.T, resourceName string, name string, description string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		svc, err := testAccGetService(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, svc)
		require.Equal(t, name, svc.Name)
		require.Equal(t, cac.ArtifactTypes.AMI, svc.ArtifactType)
		require.Equal(t, cac.DeploymentTypes.AMI, svc.DeploymentType)
		require.Equal(t, description, svc.Description)

		return nil
	}
}

func testAccResourceAMIService(name string, description string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_service_ami" "test" {
			app_id = harness_application.test.id
			name = "%[1]s"
			description = "%[2]s"
		}

`, name, description)
}
