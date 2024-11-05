package environment_group_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceEnvironmentGroup(t *testing.T) {

	name := t.Name()
	color := "#0063F7"
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"
	updatedColor := "#0063F8"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroup(id, name, color),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", "#0063F7"),
				),
			},
			{
				Config: testAccResourceEnvironmentGroup(id, name, updatedColor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", updatedColor),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceEnvironmentGroupOrgLevel(t *testing.T) {

	name := t.Name()
	color := "#0063F7"
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"
	updatedColor := "#0063F8"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroupOrgLevel(id, name, color),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", "#0063F7"),
				),
			},
			{
				Config: testAccResourceEnvironmentGroupOrgLevel(id, name, updatedColor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", updatedColor),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceEnvironmentGroupAccountLevel(t *testing.T) {

	name := t.Name()
	color := "#0063F7"
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"
	updatedColor := "#0063F8"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroupAccountLevel(id, name, color),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "color", "#0063F7"),
				),
			},
			{
				Config: testAccResourceEnvironmentGroupAccountLevel(id, name, updatedColor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "color", updatedColor),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml"},
				ImportStateIdFunc:       acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetPlatformEnvironmentGroup(resourceName string, state *terraform.State) (*nextgen.EnvironmentGroupResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	branch := r.Primary.Attributes["branch"]
	repoIdentifier := r.Primary.Attributes["repoIdentifier"]
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.EnvironmentGroupApi.GetEnvironmentGroup((ctx), id, c.AccountId, &nextgen.EnvironmentGroupApiGetEnvironmentGroupOpts{
		Branch:            optional.NewString(branch),
		RepoIdentifier:    optional.NewString(repoIdentifier),
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil || resp.Data.EnvGroup == nil {
		return nil, nil
	}

	return resp.Data.EnvGroup, nil
}

func testAccEnvironmentGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformEnvironmentGroup(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment group: %s", env.Identifier)
		}

		return nil
	}
}
func TestAccResourceEnvironmentGroupForceDelete(t *testing.T) {

	name := t.Name()
	color := "#0063F7"
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroupForceDelete(id, name, color),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", "#0063F7"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"yaml", "force_delete"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceEnvironmentGroup(id string, name string, color string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}

		resource "harness_platform_environment_group" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			color = "%[3]s"
			yaml = <<-EOT
			   environmentGroup:
			                 name: "%[1]s"
			                 identifier: "%[1]s"
			                 description: "temp"
			                 orgIdentifier: ${harness_platform_project.test.org_id}
			                 projectIdentifier: ${harness_platform_project.test.id}
			                 envIdentifiers: []
		  EOT
		}
`, id, name, color)
}
func testAccResourceEnvironmentGroupOrgLevel(id string, name string, color string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_environment_group" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.identifier
			color = "%[3]s"
			yaml = <<-EOT
			   environmentGroup:
			                 name: "%[1]s"
			                 identifier: "%[1]s"
			                 description: "temp"
			                 orgIdentifier: ${harness_platform_organization.test.identifier}
			                 envIdentifiers: []
		  EOT
		}
`, id, name, color)
}
func testAccResourceEnvironmentGroupAccountLevel(id string, name string, color string) string {
	return fmt.Sprintf(`
		resource "harness_platform_environment_group" "test" {
			identifier = "%[1]s"
			color = "%[3]s"
			yaml = <<-EOT
			   environmentGroup:
			                 name: "%[1]s"
			                 identifier: "%[1]s"
			                 description: "temp"
			                 envIdentifiers: []
		  EOT
		}
`, id, name, color)
}
func testAccResourceEnvironmentGroupForceDelete(id string, name string, color string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}

		resource "harness_platform_environment_group" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			color = "%[3]s"
			force_delete = true
			yaml = <<-EOT
			   environmentGroup:
			                 name: "%[1]s"
			                 identifier: "%[1]s"
			                 description: "temp"
			                 orgIdentifier: ${harness_platform_project.test.org_id}
			                 projectIdentifier: ${harness_platform_project.test.id}
			                 envIdentifiers: []
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
                        name: p2
                        identifier: p2
                        description: ""
                        type: Deployment
                        spec:
                          deploymentType: Kubernetes
                          service:
                            serviceRef: <+input>
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
                          environmentGroup:
                            envGroupRef: "%[1]s"
                            metadata:
                              parallel: true
                            deployToAll: <+input>
                            environments: <+input>
                        tags: {}
                        failureStrategies:
                          - onFailure:
                              errors:
                                - AllErrors
                              action:
                                type: StageRollback

                            EOT
    }
`, id, name, color)
}
