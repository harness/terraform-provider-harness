package input_set_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceInputSet(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccInputSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSet(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				Config: testAccResourceInputSet(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.PipelineResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccInputSetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		inputSet, _ := testAccGetInputSet(resourceName, state)
		if inputSet != nil {
			return fmt.Errorf("Found input set: %s", inputSet.Identifier)
		}
		return nil
	}
}

func TestAccResourceInputSetRemote(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccInputSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSetRemote(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				Config: testAccResourceInputSetRemote(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.PipelineResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type"},
			},
		},
	})

}

func TestAccResourceInputSetInline(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccInputSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSetInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				Config: testAccResourceInputSetInline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.PipelineResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func TestAccResourceInputSetImportFromGit(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccInputSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSetImportFromGit(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "inputset"),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", "DoNotDeletePipeline"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.PipelineResourceImportStateIdFunc(resourceName),
                ImportStateVerifyIgnore: []string{"git_import_info.0.branch_name", "git_import_info.0.connector_ref", "git_import_info.0.file_path","git_import_info.0.repo_name", "import_from_git", "pipeline_import_request.0.pipeline_description", "pipeline_import_request.0.pipeline_name", "git_import_info.#", "git_import_info.0.%", "pipeline_import_request.#", "pipeline_import_request.0.%", "git_details.0.connector_ref", "git_details.0.connector_ref", "git_details.0.store_type", "git_details.0.store_type", "git_import_info.0.is_force_import","input_set_import_request.#", "input_set_import_request.0.%", "input_set_import_request.0.input_set_description", "input_set_import_request.0.input_set_name"},
			},
		},
	})

}


func testAccGetInputSet(resourceName string, state *terraform.State) (*nextgen.InputSetResponseBody, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetClientWithContext()
	id := r.Primary.ID
	orgIdentifier := buildField(r, "org_id").Value()
	projectIdentifier := buildField(r, "project_id").Value()
	pipelineIdentifier := buildField(r, "pipeline_id").Value()
	var branch_name optional.String

	branch_name = buildField(r, "git_details.0.branch_name")
	parent_entity_connector_ref := buildField(r, "git_details.0.parent_entity_connector_ref")
	parent_entity_repo_name := buildField(r, "git_details.0.parent_entity_repo_name")

	resp, _, err := c.InputSetsApi.GetInputSet(ctx, orgIdentifier, projectIdentifier, id, pipelineIdentifier, &nextgen.InputSetsApiGetInputSetOpts{
		HarnessAccount:           optional.NewString(c.AccountId),
		BranchName:               branch_name,
		ParentEntityConnectorRef: parent_entity_connector_ref,
		ParentEntityRepoName:     parent_entity_repo_name,
	})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func TestAccResourceInputSet_DeleteUnderlyingResource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	project_id := id
	org_id := id
	pipeline_id := id
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSet(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetClientWithContext()
					_, err := c.InputSetsApi.DeleteInputSet(ctx, org_id, project_id, id, pipeline_id, &nextgen.InputSetsApiDeleteInputSetOpts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceInputSet(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccResourceInputSetInline(id string, name string) string {
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
    resource "harness_platform_pipeline" "test" {
        identifier = "%[1]s"
        org_id = harness_platform_project.test.org_id
        project_id = harness_platform_project.test.id
        name = "%[2]s"
yaml = <<-EOT
pipeline:
    name: %[2]s
    identifier: %[1]s
    allowStageExecutions: false
    projectIdentifier: ${harness_platform_project.test.id}
    orgIdentifier: ${harness_platform_project.test.org_id}
    tags: {}
    stages:
        - stage:
            name: dep
            identifier: dep
            description: ""
            type: Deployment
            spec:
                serviceConfig:
                    serviceRef: service
                    serviceDefinition:
                        type: Kubernetes
                        spec:
                            variables: []
                infrastructure:
                    environmentRef: testenv
                    infrastructureDefinition:
                        type: KubernetesDirect
                        spec:
                            connectorRef: testconf
                            namespace: test
                            releaseName: release-<+INFRA_KEY>
                    allowSimultaneousDeployments: false
                execution:
                    steps:
                        - stepGroup:
                                name: Canary Deployment
                                identifier: canaryDepoyment
                                steps:
                                    - step:
                                        name: Canary Deployment
                                        identifier: canaryDeployment
                                        type: K8sCanaryDeploy
                                        timeout: 10m
                                        spec:
                                            instanceSelection:
                                                type: Count
                                                spec:
                                                    count: 1
                                            skipDryRun: false
                                    - step:
                                        name: Canary Delete
                                        identifier: canaryDelete
                                        type: K8sCanaryDelete
                                        timeout: 10m
                                        spec: {}
                                rollbackSteps:
                                    - step:
                                        name: Canary Delete
                                        identifier: rollbackCanaryDelete
                                        type: K8sCanaryDelete
                                        timeout: 10m
                                        spec: {}
                        - stepGroup:
                                name: Primary Deployment
                                identifier: primaryDepoyment
                                steps:
                                    - step:
                                        name: Rolling Deployment
                                        identifier: rollingDeployment
                                        type: K8sRollingDeploy
                                        timeout: 10m
                                        spec:
                                            skipDryRun: false
                                rollbackSteps:
                                    - step:
                                        name: Rolling Rollback
                                        identifier: rollingRollback
                                        type: K8sRollingRollback
                                        timeout: 10m
                                        spec: {}
                    rollbackSteps: []
            tags: {}
            failureStrategies:
                - onFailure:
                        errors:
                            - AllErrors
                        action:
                            type: StageRollback
    variables:
        - name: key
          type: String
          default: value
          value: <+input>.allowedValues(value)
                EOT
}

resource "harness_platform_input_set" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    tags = [
        "foo:bar",
    ]
    org_id = harness_platform_organization.test.id
    project_id = harness_platform_project.test.id
    pipeline_id = harness_platform_pipeline.test.id
    depends_on = [time_sleep.wait_5_seconds]

    yaml = <<-EOT
inputSet:
  identifier: "%[1]s"
  name: "%[2]s"
  tags:
    foo: "bar"
  orgIdentifier: "${harness_platform_organization.test.id}"
  projectIdentifier: "${harness_platform_project.test.id}"
  pipeline:
    identifier: "${harness_platform_pipeline.test.id}"
    variables:
    - name: "key"
      type: "String"
      value: "value"
EOT
}

resource "time_sleep" "wait_5_seconds" {
    depends_on = [harness_platform_pipeline.test]
    create_duration = "5s"
}
    `, id, name)
}

func testAccResourceInputSetRemote(id string, name string) string {
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
    resource "harness_platform_pipeline" "test" {
        identifier = "%[1]s"
        org_id = harness_platform_project.test.org_id
        project_id = harness_platform_project.test.id
        name = "%[2]s"
        git_details {
            branch_name = "main"
            commit_message = "Commit"
            file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
            connector_ref = "account.Jajoo"
            store_type = "REMOTE"
            repo_name = "jajoo_git"
        }
yaml = <<-EOT
pipeline:
    name: %[2]s
    identifier: %[1]s
    allowStageExecutions: false
    projectIdentifier: ${harness_platform_project.test.id}
    orgIdentifier: ${harness_platform_project.test.org_id}
    tags: {}
    stages:
        - stage:
            name: dep
            identifier: dep
            description: ""
            type: Deployment
            spec:
                serviceConfig:
                    serviceRef: service
                    serviceDefinition:
                        type: Kubernetes
                        spec:
                            variables: []
                infrastructure:
                    environmentRef: testenv
                    infrastructureDefinition:
                        type: KubernetesDirect
                        spec:
                            connectorRef: testconf
                            namespace: test
                            releaseName: release-<+INFRA_KEY>
                    allowSimultaneousDeployments: false
                execution:
                    steps:
                        - stepGroup:
                                name: Canary Deployment
                                identifier: canaryDepoyment
                                steps:
                                    - step:
                                        name: Canary Deployment
                                        identifier: canaryDeployment
                                        type: K8sCanaryDeploy
                                        timeout: 10m
                                        spec:
                                            instanceSelection:
                                                type: Count
                                                spec:
                                                    count: 1
                                            skipDryRun: false
                                    - step:
                                        name: Canary Delete
                                        identifier: canaryDelete
                                        type: K8sCanaryDelete
                                        timeout: 10m
                                        spec: {}
                                rollbackSteps:
                                    - step:
                                        name: Canary Delete
                                        identifier: rollbackCanaryDelete
                                        type: K8sCanaryDelete
                                        timeout: 10m
                                        spec: {}
                        - stepGroup:
                                name: Primary Deployment
                                identifier: primaryDepoyment
                                steps:
                                    - step:
                                        name: Rolling Deployment
                                        identifier: rollingDeployment
                                        type: K8sRollingDeploy
                                        timeout: 10m
                                        spec:
                                            skipDryRun: false
                                rollbackSteps:
                                    - step:
                                        name: Rolling Rollback
                                        identifier: rollingRollback
                                        type: K8sRollingRollback
                                        timeout: 10m
                                        spec: {}
                    rollbackSteps: []
            tags: {}
            failureStrategies:
                - onFailure:
                        errors:
                            - AllErrors
                        action:
                            type: StageRollback
    variables:
        - name: key
          type: String
          default: value
          value: <+input>.allowedValues(value)
                EOT
}

resource "harness_platform_input_set" "test" {
    identifier = "%[1]s"
    name = "%[2]s"
    tags = [
        "foo:bar",
    ]
    org_id = harness_platform_organization.test.id
    project_id = harness_platform_project.test.id
    pipeline_id = harness_platform_pipeline.test.id
    git_details {
        branch_name = "main"
        commit_message = "Commit"
        file_path = ".harness/GitEnabledInputSet%[1]s.yaml"
        connector_ref = "account.Jajoo"
        store_type = "REMOTE"
        repo_name = "jajoo_git"
    }

    depends_on = [time_sleep.wait_5_seconds]

    yaml = <<-EOT
inputSet:
  identifier: "%[1]s"
  name: "%[2]s"
  tags:
    foo: "bar"
  orgIdentifier: "${harness_platform_organization.test.id}"
  projectIdentifier: "${harness_platform_project.test.id}"
  pipeline:
    identifier: "${harness_platform_pipeline.test.id}"
    variables:
    - name: "key"
      type: "String"
      value: "value"
EOT
}

resource "time_sleep" "wait_5_seconds" {
    depends_on = [harness_platform_pipeline.test]
    create_duration = "5s"
}
    `, id, name)
}

func testAccResourceInputSet(id string, name string) string {
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

        resource "harness_platform_pipeline" "test" {
						identifier = "%[1]s"
						org_id = harness_platform_project.test.org_id
						project_id = harness_platform_project.test.id
						name = "%[2]s"
            yaml = <<-EOT
                pipeline:
                    identifier: %[1]s
                    name: %[2]s
                    allowStageExecutions: false
                    projectIdentifier: ${harness_platform_project.test.id}
                    orgIdentifier: ${harness_platform_project.test.org_id}
                    tags: {}
                    stages:
                        - stage:
                            name: dep
                            identifier: dep
                            description: ""
                            type: Deployment
                            spec:
                                serviceConfig:
                                    serviceRef: service
                                    serviceDefinition:
                                        type: Kubernetes
                                        spec:
                                            variables: []
                                infrastructure:
                                    environmentRef: testenv
                                    infrastructureDefinition:
                                        type: KubernetesDirect
                                        spec:
                                            connectorRef: testconf
                                            namespace: test
                                            releaseName: release-<+INFRA_KEY>
                                    allowSimultaneousDeployments: false
                                execution:
                                    steps:
                                        - stepGroup:
                                                name: Canary Deployment
                                                identifier: canaryDepoyment
                                                steps:
                                                    - step:
                                                        name: Canary Deployment
                                                        identifier: canaryDeployment
                                                        type: K8sCanaryDeploy
                                                        timeout: 10m
                                                        spec:
                                                            instanceSelection:
                                                                type: Count
                                                                spec:
                                                                    count: 1
                                                            skipDryRun: false
                                                    - step:
                                                        name: Canary Delete
                                                        identifier: canaryDelete
                                                        type: K8sCanaryDelete
                                                        timeout: 10m
                                                        spec: {}
                                                rollbackSteps:
                                                    - step:
                                                        name: Canary Delete
                                                        identifier: rollbackCanaryDelete
                                                        type: K8sCanaryDelete
                                                        timeout: 10m
                                                        spec: {}
                                        - stepGroup:
                                                name: Primary Deployment
                                                identifier: primaryDepoyment
                                                steps:
                                                    - step:
                                                        name: Rolling Deployment
                                                        identifier: rollingDeployment
                                                        type: K8sRollingDeploy
                                                        timeout: 10m
                                                        spec:
                                                            skipDryRun: false
                                                rollbackSteps:
                                                    - step:
                                                        name: Rolling Rollback
                                                        identifier: rollingRollback
                                                        type: K8sRollingRollback
                                                        timeout: 10m
                                                        spec: {}
                                    rollbackSteps: []
                            tags: {}
                            failureStrategies:
                                - onFailure:
                                        errors:
                                            - AllErrors
                                        action:
                                          type: StageRollback
                    variables:
                        - name: key
                          type: String
                          default: value
                          value: <+input>.allowedValues(value)
            EOT
        }

				resource "harness_platform_input_set" "test" {
						identifier = "%[1]s"
						name = "%[2]s"
						tags = [
							"foo:bar",
                        ]
						org_id = harness_platform_organization.test.id
						project_id = harness_platform_project.test.id
						pipeline_id = harness_platform_pipeline.test.id
                        depends_on = [time_sleep.wait_5_seconds]

						yaml = <<-EOT
                inputSet:
                  identifier: "%[1]s"
                  name: "%[2]s"
                  tags:
                    foo: "bar"
                  orgIdentifier: "${harness_platform_organization.test.id}"
                  projectIdentifier: "${harness_platform_project.test.id}"
                  pipeline:
                    identifier: "${harness_platform_pipeline.test.id}"
                    variables:
                    - name: "key"
                      type: "String"
                      value: "value"
            EOT
				}

                resource "time_sleep" "wait_5_seconds" {
                    depends_on = [harness_platform_pipeline.test]
                    create_duration = "5s"
                }
        `, id, name)
}

func testAccResourceInputSetImportFromGit(id string, name string) string {
	return fmt.Sprintf(`
        resource "harness_platform_organization" "test" {
					identifier = "%[1]s"
					name = "%[2]s"
				}
        resource "harness_platform_input_set" "test" {
                        identifier = "inputset"
                        org_id = "default"
						project_id = "DoNotDelete_Amit"
                        name = "inputset"
                        pipeline_id = "DoNotDeletePipeline"
                        import_from_git = true
                        git_import_info {
                            branch_name = "main"
                            file_path = ".harness/inputset.yaml"
                            connector_ref = "account.DoNotDeleteGithub"
                            repo_name = "open-repo"
                        }
                        input_set_import_request {
                            input_set_name = "inputset"
                            input_set_description = ""
                        }
                }
        `, id, name)
}
