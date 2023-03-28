package service_test

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

func TestAccResourceService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceServiceWithoutServiceDefinition(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestProjectResourceService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testProjectResourceServiceWithoutServiceDefinition(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestOrgResourceService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				Config: testOrgResourceServiceWithoutServiceDefinition(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceServiceWithYaml(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	varValue := t.Name()
	updatedVarValue := fmt.Sprintf("%s_updated", varValue)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceServiceWithYaml(id, name, varValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				Config: testAccResourceServiceWithYaml(id, updatedName, updatedVarValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceServiceWithYamlAccountLevel(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	varValue := t.Name()
	updatedVarValue := fmt.Sprintf("%s_updated", varValue)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceServiceWithYamlAccountLevel(id, name, varValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceServiceWithYamlAccountLevel(id, updatedName, updatedVarValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceServiceWithYamlOrgLevel(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	varValue := t.Name()
	updatedVarValue := fmt.Sprintf("%s_updated", varValue)
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceServiceWithYamlOrgLevel(id, name, varValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				Config: testAccResourceServiceWithYamlOrgLevel(id, updatedName, updatedVarValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceService_DeleteUnderlyingResource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.ServicesApi.DeleteServiceV2(ctx, id, c.AccountId, &nextgen.ServicesApiDeleteServiceV2Opts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:             testProjectResourceService(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
func TestForceDeleteService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	varValue := t.Name()
	resourceName := "harness_platform_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceServiceForForceDeletion(id, name, varValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetService(resourceName string, state *terraform.State) (*nextgen.ServiceResponseDetails, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.ServicesApi.GetServiceV2(ctx, id, c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil || resp.Data.Service == nil {
		return nil, nil
	}

	return resp.Data.Service, nil
}

func testAccServiceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetService(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found service: %s", env.Identifier)
		}

		return nil
	}
}

func testAccResourceService(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      
    }
`, id, name)
}
func testAccResourceServiceWithoutServiceDefinition(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
      EOT
    }
`, id, name)
}

func testProjectResourceService(id string, name string) string {
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

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
    }
`, id, name)
}
func testProjectResourceServiceWithoutServiceDefinition(id string, name string) string {
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

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
      EOT
    }
`, id, name)
}

func testOrgResourceService(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_organization" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
    }

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_organization.test.id
      
    }
`, id, name)
}
func testOrgResourceServiceWithoutServiceDefinition(id string, name string) string {
	return fmt.Sprintf(`
    resource "harness_platform_organization" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
    }

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_organization.test.id
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
      EOT
    }
`, id, name)
}
func testAccResourceServiceWithYaml(id string, name string, varValue string) string {
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

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
              configFiles:
                - configFile:
                    identifier: configFile1
                    spec:
                      store:
                        type: Harness
                        spec:
                          files:
                            - <+org.description>
              variables:
                - name: var1
                  type: String
                  value: %[3]s
                - name: var2
                  type: String
                  value: val2
            type: Kubernetes
          gitOpsEnabled: false
      EOT
    }
`, id, name, varValue)
}
func testAccResourceServiceForForceDeletion(id string, name string, varValue string) string {
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

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_project.test.org_id
      project_id = harness_platform_project.test.id
      force_delete = true
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
              configFiles:
                - configFile:
                    identifier: configFile1
                    spec:
                      store:
                        type: Harness
                        spec:
                          files:
                            - <+org.description>
              variables:
                - name: var1
                  type: String
                  value: %[3]s
                - name: var2
                  type: String
                  value: val2
            type: Kubernetes
          gitOpsEnabled: false
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
                    environmentRef: <+input>
                    deployToAll: false
                    environmentInputs: <+input>
                    serviceOverrideInputs: <+input>
                    infrastructureDefinitions: <+input>
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
`, id, name, varValue)
}

func testAccResourceServiceWithYamlOrgLevel(id string, name string, varValue string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      org_id = harness_platform_organization.test.id
			
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
              configFiles:
                - configFile:
                    identifier: configFile1
                    spec:
                      store:
                        type: Harness
                        spec:
                          files:
                            - <+org.description>
              variables:
                - name: var1
                  type: String
                  value: %[3]s
                - name: var2
                  type: String
                  value: val2
            type: Kubernetes
          gitOpsEnabled: false
      EOT
    }
`, id, name, varValue)
}

func testAccResourceServiceWithYamlAccountLevel(id string, name string, varValue string) string {
	return fmt.Sprintf(`
    resource "harness_platform_service" "test" {
      identifier = "%[1]s"
      name = "%[2]s"
      yaml = <<-EOT
        service:
          name: %[2]s
          identifier: %[1]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: K8sManifest
                    spec:
                      store:
                        type: Github
                        spec:
                          connectorRef: <+input>
                          gitFetchType: Branch
                          paths:
                            - files1
                          repoName: <+input>
                          branch: master
                      skipResourceVersioning: false
              configFiles:
                - configFile:
                    identifier: configFile1
                    spec:
                      store:
                        type: Harness
                        spec:
                          files:
                            - <+org.description>
              variables:
                - name: var1
                  type: String
                  value: %[3]s
                - name: var2
                  type: String
                  value: val2
            type: Kubernetes
          gitOpsEnabled: false
      EOT
    }
`, id, name, varValue)
}
