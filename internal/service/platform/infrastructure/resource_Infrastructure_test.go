package infrastructure_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceInfrastructure(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfrastructure(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
				),
			},
			{
				Config: testAccResourceInfrastructure(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceInfrastructure_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfrastructure(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					envId := id
					resp, _, err := c.InfrastructuresApi.DeleteInfrastructure(ctx, id, c.AccountId, envId, &nextgen.InfrastructuresApiDeleteInfrastructureOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
					require.True(t, resp.Data)
				},
				Config:             testAccResourceInfrastructure(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetPlatformInfrastructure(resourceName string, state *terraform.State) (*nextgen.InfrastructureResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	envId := r.Primary.Attributes["env_id"]

	resp, _, err := c.InfrastructuresApi.GetInfrastructureList((ctx), c.AccountId, envId, &nextgen.InfrastructuresApiGetInfrastructureListOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return &resp.Data.Content[0], nil
}

func testAccInfrastructureDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		infra, _ := testAccGetPlatformInfrastructure(resourceName, state)
		if infra != nil {
			return fmt.Errorf("Found infrastructure: %s", infra.Infrastructure.Identifier)
		}

		return nil
	}
}

func testAccResourceInfrastructure(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			yaml = <<-EOT
			   infrastructureDefinition:
         name: "%[2]s"
         identifier: "%[1]s"
         description: ""
         tags:
           asda: ""
         orgIdentifier: ${harness_platform_organization.test.id}
         projectIdentifier: ${harness_platform_project.test.id}
         environmentRef: ${harness_platform_environment.test.id}
         deploymentType: Kubernetes
         type: KubernetesDirect
         spec:
          connectorRef: account.gfgf
          namespace: asdasdsa
          releaseName: release-<+INFRA_KEY>
          allowSimultaneousDeployments: false
      EOT
		}

`, id, name)
}
