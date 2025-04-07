package service_overrides_v2_test

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

func TestAccServiceOverrides_ProjectScope(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	resourceName := "harness_platform_service_overrides_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceOverridesProjectScope(id, name),
				Check: resource.ComposeTestCheckFunc(
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

func TestAccServiceOverrides_OrgScope(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	resourceName := "harness_platform_service_overrides_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceOverridesOrgScope(id, name),
				Check: resource.ComposeTestCheckFunc(
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

func TestAccServiceOverrides_AccountScope(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	resourceName := "harness_platform_service_overrides_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceOverridesAccountScope(id, name),
				Check:  resource.ComposeTestCheckFunc(),
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

func TestAccRemoteServiceOverrides(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	resourceName := "harness_platform_service_overrides_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config:             testRemoteServiceOverrides(id, name),
				Destroy:            false,
				ExpectNonEmptyPlan: true,
				Check:              resource.ComposeTestCheckFunc(),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type",
					"git_details.#", "git_details.0.%", "git_details.0.base_branch", "git_details.0.branch", "git_details.0.file_path", "git_details.0.is_harnesscode_repo", "git_details.0.is_new_branch",
					"git_details.0.last_commit_id", "git_details.0.last_object_id", "git_details.0.load_from_cache", "git_details.0.load_from_fallback_branch", "git_details.0.repo_name", "git_details.0.import_from_git", "git_details.0.is_force_import", "git_details.0.parent_entity_connector_ref", "git_details.0.parent_entity_repo_name", "git_details.0.is_harness_code_repo", "yaml"},
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccServiceOverridesEditGitDetails(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	resourceName := "harness_platform_service_overrides_v2.test"
	identifier := id + "_override" // Explicit identifier for easier debugging

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// First create a service override with initial Git details
				Config:             testAccServiceOverridesGitDetails(id, name, "main", "initial-path"),
				ExpectNonEmptyPlan: true,
				Destroy:            false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "git_details.0.file_path", ".harness/automation/overrides/initial-path.yaml"),
				),
			},
			{
				Config:             testAccServiceOverridesGitDetails(id, name, "main", "updated-path"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", identifier),
					resource.TestCheckResourceAttr(resourceName, "git_details.0.file_path", ".harness/automation/overrides/updated-path.yaml"),
				),
			},
		},
	})
}

func testAccGetPlatformServiceOverrides(resourceName string, state *terraform.State) (*nextgen.ServiceOverridesResponseDtov2, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	identifier := r.Primary.ID

	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.ServiceOverridesApi.GetServiceOverridesV2(ctx, identifier, c.AccountId,
		&nextgen.ServiceOverridesApiGetServiceOverridesV2Opts{
			OrgIdentifier:     optional.NewString(orgId),
			ProjectIdentifier: optional.NewString(projId),
		})

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func testAccServiceOverridesDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformServiceOverrides(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found service overrides")
		}

		return nil
	}
}

func testAccServiceOverridesProjectScope(id string, name string) string {
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

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			yaml = <<-EOT
        service:
          name: %[1]s
          identifier: %[2]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: Values
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
            type: Kubernetes
          gitOpsEnabled: false
		  EOT
		}

        resource "harness_platform_service_overrides_v2" "test" {
          org_id     = harness_platform_organization.test.id
          project_id = harness_platform_project.test.id
          env_id     = harness_platform_environment.test.id
          service_id = harness_platform_service.test.id
          type       = "ENV_SERVICE_OVERRIDE"
          yaml = <<-EOT
variables:
  - name: v1
    type: String
    value: val1
    required: false
manifests:
  - manifest:
      identifier: manifest1
      type: Values
      spec:
        store:
          spec:
            connectorRef: <+input>
            gitFetchType: Branch
            branch: master
            commitId: null
            paths:
              - files1
            folderPath: null
            repoName: <+input>
          type: Github
        optionalValuesYaml: null
configFiles:
  - configFile:
      identifier: configFile1
      spec:
        store:
          spec:
            files:
              - <+org.description>
            secretFiles: null
          type: Harness
              EOT
}
`, id, name)
}

func testRemoteServiceOverrides(id string, name string) string {
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
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
            git_details {
				store_type = "REMOTE"
				connector_ref = "account.TF_GitX_connector"  
				repo_name = "pcf_practice"
				file_path = ".harness/automation/service/%[1]s.yaml"
				branch = "main"
            }
			yaml = <<-EOT
        service:
          name: %[1]s
          identifier: %[2]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: Values
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
            type: Kubernetes
          gitOpsEnabled: false
		  EOT
		}

  resource "harness_platform_service_overrides_v2" "test" {
    org_id     = harness_platform_organization.test.id
    project_id = harness_platform_project.test.id
     env_id     = harness_platform_environment.test.id
     service_id = harness_platform_service.test.id
     type       = "ENV_SERVICE_OVERRIDE"
     git_details {
       store_type = "REMOTE"
       connector_ref = "account.TF_GitX_connector"  
       repo_name = "pcf_practice"
       file_path = ".harness/automation/overrides/%[1]s.yaml"
       branch = "main"
       }
       yaml = <<-EOT
       variables:
         - name: v1
           type: String
           value: val1
       manifests:
         - manifest:
             identifier: manifest1
             type: Values
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
                     EOT
}`, id, name)
}

func testAccServiceOverridesOrgScope(id string, name string) string {
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

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			yaml = <<-EOT
        service:
          name: %[1]s
          identifier: %[2]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: Values
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
            type: Kubernetes
          gitOpsEnabled: false
		  EOT
		}

		resource "harness_platform_service_overrides_v2" "test" {
			org_id = harness_platform_organization.test.id
			env_id = "org.${harness_platform_environment.test.id}"
			service_id = "org.${harness_platform_service.test.id}"
            type = "ENV_SERVICE_OVERRIDE"
            yaml = <<-EOT
variables:
  - name: v1
    type: String
    value: val1
    required: false
manifests:
  - manifest:
      identifier: manifest1
      type: Values
      spec:
        store:
          spec:
            connectorRef: <+input>
            gitFetchType: Branch
            branch: master
            commitId: null
            paths:
              - files1
            folderPath: null
            repoName: <+input>
          type: Github
        optionalValuesYaml: null
configFiles:
  - configFile:
      identifier: configFile1
      spec:
        store:
          spec:
            files:
              - <+org.description>
            secretFiles: null
          type: Harness
              EOT
		}
`, id, name)
}

func testAccServiceOverridesAccountScope(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
  	}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			yaml = <<-EOT
        service:
          name: %[1]s
          identifier: %[2]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: Values
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
            type: Kubernetes
          gitOpsEnabled: false
		  EOT
		}

		resource "harness_platform_service_overrides_v2" "test" {
            env_id = "account.${harness_platform_environment.test.id}"
			service_id = "account.${harness_platform_service.test.id}"
            type = "ENV_SERVICE_OVERRIDE"
            yaml = <<-EOT
variables:
  - name: v1
    type: String
    value: val1
    required: false
manifests:
  - manifest:
      identifier: manifest1
      type: Values
      spec:
        store:
          spec:
            connectorRef: <+input>
            gitFetchType: Branch
            branch: master
            commitId: null
            paths:
              - files1
            folderPath: null
            repoName: <+input>
          type: Github
        optionalValuesYaml: null
configFiles:
  - configFile:
      identifier: configFile1
      spec:
        store:
          spec:
            files:
              - <+org.description>
            secretFiles: null
          type: Harness
              EOT
		}
`, id, name)
}

func testAccServiceOverridesGitDetails(id string, name string, branch string, filePath string) string {
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
			project_id = harness_platform_project.test.id
			tags = ["foo:bar", "baz"]
			type = "PreProduction"
		}

		resource "harness_platform_service" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			git_details {
				store_type = "REMOTE"
				connector_ref = "account.TF_GitX_connector"  
				repo_name = "pcf_practice"
				file_path = ".harness/automation/service/%[1]s.yaml"
				branch = "main"
			}
			yaml = <<-EOT
        service:
          name: %[1]s
          identifier: %[2]s
          serviceDefinition:
            spec:
              manifests:
                - manifest:
                    identifier: manifest1
                    type: Values
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
            type: Kubernetes
          gitOpsEnabled: false
		EOT
		}

		resource "harness_platform_service_overrides_v2" "test" {
			identifier = "%[1]s_override"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id     = harness_platform_environment.test.id
			service_id = harness_platform_service.test.id
			type       = "ENV_SERVICE_OVERRIDE"
			git_details {
				store_type = "REMOTE"
				connector_ref = "account.TF_GitX_connector"  
				repo_name = "pcf_practice"
				file_path = ".harness/automation/overrides/%[3]s.yaml"
				branch = "%[4]s"
				commit_message = "Update service override for %[1]s"
			}
			yaml = <<-EOT
        variables:
          - name: v1
            type: String
            value: val1
            required: false
        manifests:
          - manifest:
              identifier: manifest1
              type: Values
              spec:
                store:
                  type: Github
                  spec:
                    connectorRef: <+input>
                    gitFetchType: Branch
                    branch: master
                    commitId: null
                    paths:
                      - files1
                    folderPath: null
                    repoName: <+input>
                  optionalValuesYaml: null
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
                    secretFiles: null
			EOT
		}`, id, name, filePath, branch)
}
