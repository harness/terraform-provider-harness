package infrastructure_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceInfrastructure(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfrastructure(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceInfrastructureAccountLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfrastructureAccountLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceInfrastructureOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfrastructureOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testAccDataSourceInfrastructure(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type = "PreProduction"
		}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			deployment_type = "Kubernetes"
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

		data "harness_platform_infrastructure" "test" {
			identifier = harness_platform_infrastructure.test.id
			name = harness_platform_infrastructure.test.name
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id = harness_platform_environment.test.id
		}
`, id, name)
}

func testAccDataSourceInfrastructureAccountLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			type = "PreProduction"
		}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			deployment_type = "Kubernetes"
			yaml = <<-EOT
			   infrastructureDefinition:
         name: "%[2]s"
         identifier: "%[1]s"
         description: ""
         tags:
           asda: ""
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

		data "harness_platform_infrastructure" "test" {
			identifier = harness_platform_infrastructure.test.id
			env_id = harness_platform_environment.test.id
		}
`, id, name)
}

func testAccDataSourceInfrastructureOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			type = "PreProduction"
		}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			deployment_type = "Kubernetes"
			yaml = <<-EOT
			   infrastructureDefinition:
         name: "%[2]s"
         identifier: "%[1]s"
         description: ""
         tags:
           asda: ""
         orgIdentifier: ${harness_platform_organization.test.id}
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

		data "harness_platform_infrastructure" "test" {
			identifier = harness_platform_infrastructure.test.id
			org_id = harness_platform_organization.test.id
			env_id = harness_platform_environment.test.id
		}
`, id, name)
}
