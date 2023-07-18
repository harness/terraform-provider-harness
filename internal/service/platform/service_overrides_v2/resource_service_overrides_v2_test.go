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
manifests:
  - manifest:
      identifier: manifest1
      type: K8sManifest
      spec:
        store:
          type: Github
          spec:
            connectorRef: "<+input>"
            gitFetchType: Branch
            paths:
              - files1
            repoName: "<+input>"
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
              - "<+org.description>"
              EOT
}
`, id, name)
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
manifests:
  - manifest:
      identifier: manifest1
      type: K8sManifest
      spec:
        store:
          type: Github
          spec:
            connectorRef: "<+input>"
            gitFetchType: Branch
            paths:
              - files1
            repoName: "<+input>"
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
              - "<+org.description>"

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

		resource "harness_platform_service_overrides_v2" "test" {
            env_id = "account.${harness_platform_environment.test.id}"
			service_id = "account.${harness_platform_service.test.id}"
            type = "ENV_SERVICE_OVERRIDE"
            yaml = <<-EOT
variables:
  - name: v1
    type: String
    value: val1
manifests:
  - manifest:
      identifier: manifest1
      type: K8sManifest
      spec:
        store:
          type: Github
          spec:
            connectorRef: "<+input>"
            gitFetchType: Branch
            paths:
              - files1
            repoName: "<+input>"
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
              - "<+org.description>"
              EOT
		}
`, id, name)
}
