package environment_service_overrides_test

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

func TestAccEnvServiceOverrides(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	name := id
	expectedid := id + "_" + id
	resourceName := "harness_environment_service_overrides.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvServiceOverridesDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccEnvServiceOverrides(expectedid, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", expectedid),
					resource.TestCheckResourceAttr(resourceName, "project_id", expectedid),
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

func testAccGetPlatformEnvServiceOverrides(resourceName string, state *terraform.State) (*nextgen.PageResponseServiceOverrideResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	envId := r.Primary.Attributes["env_id"]
	serviceId := r.Primary.Attributes["service_id"]

	resp, _, err := c.EnvironmentsApi.GetServiceOverridesList(ctx, c.AccountId, orgId, projId, envId,
		&nextgen.EnvironmentsApiGetServiceOverridesListOpts{
			ServiceIdentifier: optional.NewString(serviceId),
		})

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func testAccEnvServiceOverridesDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformEnvServiceOverrides(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment service override")
		}

		return nil
	}
}

func testAccEnvServiceOverrides(id string, name string) string {
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
            type: Kubernetes
          gitOpsEnabled: false
		  EOT
		}

		resource "harness_environment_service_overrides" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			env_id = harness_platform_environment.test.id
			service_id = harness_platform_service.test.id
			yaml = <<-EOT
        serviceOverrides:
          environmentRef: ${harness_platform_environment.test.id}
          serviceRef: ${harness_platform_service.test.id}
          variables:
           - name: asda
             type: String
             value: asddad
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
                 identifier: configFileEnv
                 spec:
                   store:
                     type: Harness
                     spec:
                       files:
                         - account:/Add-ons/svcOverrideTest
                       secretFiles: []
		  EOT
		}
`, id, name)
}
