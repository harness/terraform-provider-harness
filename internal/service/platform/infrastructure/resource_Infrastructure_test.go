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
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testAccResourceInfrastructure(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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
func TestAccResourceInfrastructureForceDelete(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfrastructureForceDelete(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
				ImportStateIdFunc:       acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
			},
		},
	})
}
func TestAccResourceMultipleInfrastructure(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	id2 := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	name2 := id2
	updatedName2 := fmt.Sprintf("%s_updated", name2)
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_infrastructure.test"
	resourceName2 := "harness_platform_infrastructure.test2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMultipleInfrastructure(id, name, id2, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName2, "id", id2),
					resource.TestCheckResourceAttr(resourceName2, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testAccResourceMultipleInfrastructure(id, updatedName, id2, updatedName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName2, "id", id2),
					resource.TestCheckResourceAttr(resourceName2, "name", updatedName2),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
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

func TestAccResourceInfrastructureAccountLevel(t *testing.T) {

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
				Config: testAccResourceInfrastructureAccountLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
				),
			},
			{
				Config: testAccResourceInfrastructureAccountLevel(id, updatedName),
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
func TestAccResourceInfrastructureAccountLevelForceDelete(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfrastructureAccountLevelForceDelete(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
				ImportStateIdFunc:       acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceInfrastructureOrgLevel(t *testing.T) {

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
				Config: testAccResourceInfrastructureOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				Config: testAccResourceInfrastructureOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
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

func TestAccResourceInfrastructureOrgLevelForceDelete(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInfrastructureOrgLevelForceDelete(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
				ImportStateIdFunc:       acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceInfrastructure_DeleteUnderlyingResource(t *testing.T) {
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

func TestResourceRemoteInfrastructure(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testResourceRemoteInfrastructure(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type",
					"git_details.#", "git_details.0.%", "git_details.0.base_branch", "git_details.0.branch", "git_details.0.file_path", "git_details.0.is_harnesscode_repo", "git_details.0.is_new_branch",
					"git_details.0.last_commit_id", "git_details.0.last_object_id", "git_details.0.load_from_cache", "git_details.0.load_from_fallback_branch", "git_details.0.repo_name", "git_details.0.import_from_git", "git_details.0.is_force_import", "git_details.0.parent_entity_connector_ref", "git_details.0.parent_entity_repo_name", "yaml"},
			},
		},
	})
}

func TestResourceImportRemoteInfrastructure(t *testing.T) {
	resourceName := "harness_platform_infrastructure.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfrastructureDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testResourceImportRemoteInfrastructure(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "accountInfra"),
					resource.TestCheckResourceAttr(resourceName, "name", "accountInfra"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.EnvRelatedResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type",
					"git_details.#", "git_details.0.%", "git_details.0.base_branch", "git_details.0.branch", "git_details.0.file_path", "git_details.0.is_harnesscode_repo", "git_details.0.is_new_branch",
					"git_details.0.last_commit_id", "git_details.0.last_object_id", "git_details.0.load_from_cache", "git_details.0.load_from_fallback_branch", "git_details.0.repo_name", "git_details.0.import_from_git", "git_details.0.is_force_import", "git_details.0.parent_entity_connector_ref", "git_details.0.parent_entity_repo_name"},
			},
		},
	})
}

func testResourceImportRemoteInfrastructure() string {
	return fmt.Sprintf(`
	resource "harness_platform_infrastructure" "test" {
		identifier  = "accountInfra"
		name        = "accountInfra"
		env_id = "DoNotDeleteTerraformResourceEnv"
		type = "KubernetesDirect"
		git_details { 
			connector_ref = "account.DoNotDeleteRTerraformResource"
			repo_name = "terraform-test"
			file_path = ".harness/accountInfra.yaml"
			branch = "main"
			import_from_git = "true"
			is_force_import = "true"
		}
		yaml = <<-EOT
		infrastructureDefinition:
        name: accountInfra
        identifier: accountInfra
        environmentRef: DoNotDeleteTerraformResourceEnv
        type: KubernetesDirect
        deploymentType: Kubernetes
        spec:
         connectorRef: <+input>
         namespace: <+input>
         releaseName: <+input>
        allowSimultaneousDeployments: false
      EOT
	}
`)
}

func testResourceRemoteInfrastructure(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_environment" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		tags = ["foo:bar", "baz"]
		type = "PreProduction"
	}

	resource "harness_platform_infrastructure" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		env_id = harness_platform_environment.test.id
		type = "KubernetesDirect"
		git_details {
			store_type = "REMOTE"
			connector_ref = "account.DoNotDeleteRTerraformResource"
			repo_name = "terraform-test"
			file_path = ".harness/%[1]s.yaml"
			branch = "main"
		}
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
          connectorRef: "<+input>"
          namespace: "<+input>"
          releaseName: "<+input>"
         allowSimultaneousDeployments: false
      EOT
	}
`, id, name)
}

func testAccGetPlatformInfrastructure(resourceName string, state *terraform.State) (*nextgen.InfrastructureResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	envId := r.Primary.Attributes["env_id"]

	resp, _, err := c.InfrastructuresApi.GetInfrastructure((ctx), id, c.AccountId, envId, &nextgen.InfrastructuresApiGetInfrastructureOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
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

func testAccResourceMultipleInfrastructure(id string, name string, id2 string, name2 string) string {
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

		resource "harness_platform_infrastructure" "test2" {
			identifier = "%[3]s"
			name = "%[4]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			yaml = <<-EOT
			   infrastructureDefinition:
         name: "%[4]s"
         identifier: "%[3]s"
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

`, id, name, id2, name2)
}

func testAccResourceInfrastructureAccountLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
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

`, id, name)
}

func testAccResourceInfrastructureOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
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

func testAccResourceInfrastructureForceDelete(id string, name string) string {
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
			force_delete = true
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
			force_delete = true
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
        resource "harness_platform_pipeline" "test" {
        identifier = "%[1]s"
        org_id = harness_platform_project.test.org_id
        project_id = harness_platform_project.test.id
        name = "%[2]s"
        yaml = <<-EOT
        pipeline:
          name: "%[2]s"
          identifier: "%[1]s"
          projectIdentifier: ${harness_platform_project.test.id}
          orgIdentifier: ${harness_platform_project.test.org_id}
          tags: {}
          stages:
            - stage:
                name: p3
                identifier: p3
                description: ""
                type: Deployment
                spec:
                  deploymentType: Kubernetes
                  service:
                    serviceRef: "%[1]s"
                    serviceInputs:
                      serviceDefinition:
                        type: Kubernetes
                        spec:
                          artifacts:
                            primary:
                              primaryArtifactRef: <+input>
                              sources: <+input>
                  environment:
                    environmentRef: "%[1]s"
                    deployToAll: false
                    environmentInputs: <+input>
                    serviceOverrideInputs: <+input>
                    infrastructureDefinitions:
                     - identifier: "%[1]s"
                  execution:
                    steps:
                      - step:
                          name: Rollout Deployment
                          identifier: rolloutDeployment
                          type: K8sRollingDeploy
                          timeout: 10m
                          spec:
                            skipDryRun: false
                            pruningEnabled: false
                    rollbackSteps:
                      - step:
                          name: Rollback Rollout Deployment
                          identifier: rollbackRolloutDeployment
                          type: K8sRollingRollback
                          timeout: 10m
                          spec:
                            pruningEnabled: false
                tags: {}
                failureStrategies:
                  - onFailure:
                      errors:
                        - AllErrors
                      action:
                        type: StageRollback
        EOT
}


`, id, name)
}
func testAccResourceInfrastructureAccountLevelForceDelete(id string, name string) string {
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
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
			force_delete = true
  	}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			force_delete = true
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
        resource "harness_platform_pipeline" "test" {
        identifier = "%[1]s"
        org_id = harness_platform_project.test.org_id
        project_id = harness_platform_project.test.id
        name = "%[2]s"
        yaml = <<-EOT
        pipeline:
          name: "%[2]s"
          identifier: "%[1]s"
          projectIdentifier: ${harness_platform_project.test.id}
          orgIdentifier: ${harness_platform_project.test.org_id}
          tags: {}
          stages:
            - stage:
                name: p3
                identifier: p3
                description: ""
                type: Deployment
                spec:
                  deploymentType: Kubernetes
                  service:
                    serviceRef: "%[1]s"
                    serviceInputs:
                      serviceDefinition:
                        type: Kubernetes
                        spec:
                          artifacts:
                            primary:
                              primaryArtifactRef: <+input>
                              sources: <+input>
                  environment:
                    environmentRef: "account.%[1]s"
                    deployToAll: false
                    infrastructureDefinitions:
                     - identifier: "%[1]s"
                  execution:
                    steps:
                      - step:
                          name: Rollout Deployment
                          identifier: rolloutDeployment
                          type: K8sRollingDeploy
                          timeout: 10m
                          spec:
                            skipDryRun: false
                            pruningEnabled: false
                    rollbackSteps:
                      - step:
                          name: Rollback Rollout Deployment
                          identifier: rollbackRolloutDeployment
                          type: K8sRollingRollback
                          timeout: 10m
                          spec:
                            pruningEnabled: false
                tags: {}
                failureStrategies:
                  - onFailure:
                      errors:
                        - AllErrors
                      action:
                        type: StageRollback
        EOT
}
`, id, name)
}
func testAccResourceInfrastructureOrgLevelForceDelete(id string, name string) string {
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
			org_id = harness_platform_organization.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
			force_delete = true
  	}

		resource "harness_platform_infrastructure" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			env_id = harness_platform_environment.test.id
			type = "KubernetesDirect"
			force_delete = true
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
        resource "harness_platform_pipeline" "test" {
        identifier = "%[1]s"
        org_id = harness_platform_project.test.org_id
        project_id = harness_platform_project.test.id
        name = "%[2]s"
        yaml = <<-EOT
        pipeline:
          name: "%[2]s"
          identifier: "%[1]s"
          projectIdentifier: ${harness_platform_project.test.id}
          orgIdentifier: ${harness_platform_project.test.org_id}
          tags: {}
          stages:
            - stage:
                name: p3
                identifier: p3
                description: ""
                type: Deployment
                spec:
                  deploymentType: Kubernetes
                  service:
                    serviceRef: "%[1]s"
                    serviceInputs:
                      serviceDefinition:
                        type: Kubernetes
                        spec:
                          artifacts:
                            primary:
                              primaryArtifactRef: <+input>
                              sources: <+input>
                  environment:
                    environmentRef: "org.%[1]s"
                    deployToAll: false
                    infrastructureDefinitions:
                     - identifier: "%[1]s"
                  execution:
                    steps:
                      - step:
                          name: Rollout Deployment
                          identifier: rolloutDeployment
                          type: K8sRollingDeploy
                          timeout: 10m
                          spec:
                            skipDryRun: false
                            pruningEnabled: false
                    rollbackSteps:
                      - step:
                          name: Rollback Rollout Deployment
                          identifier: rollbackRolloutDeployment
                          type: K8sRollingRollback
                          timeout: 10m
                          spec:
                            pruningEnabled: false
                tags: {}
                failureStrategies:
                  - onFailure:
                      errors:
                        - AllErrors
                      action:
                        type: StageRollback
        EOT
}


`, id, name)
}
