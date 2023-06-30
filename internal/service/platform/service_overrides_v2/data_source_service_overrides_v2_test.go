package service_overrides_v2_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceServiceOverrides(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_service_overrides_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceServiceOverrides(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceServiceOverrides(id string, name string) string {
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
					org_id = harness_platform_organization.test.id
					project_id = harness_platform_project.test.id
					env_id = harness_platform_environment.test.id
					service_id = harness_platform_service.test.id
		            type = "ENV_SERVICE_OVERRIDE"
                    yaml = <<-EOT
              {
                "variables": [
                  {
                    "name": "v1",
                    "type": "String",
                    "value": "val1"
                  }
                ],
                "manifests": [
                  {
                    "manifest": {
                      "identifier": "manifest1",
                      "type": "K8sManifest",
                      "spec": {
                        "store": {
                          "type": "Github",
                          "spec": {
                            "connectorRef": "<+input>",
                            "gitFetchType": "Branch",
                            "paths": [
                              "files1"
                            ],
                            "repoName": "<+input>",
                            "branch": "master"
                          }
                        },
                        "skipResourceVersioning": false
                      }
                    }
                  }
                ],
                "configFiles": [
                  {
                    "configFile": {
                      "identifier": "configFile1",
                      "spec": {
                        "store": {
                          "type": "Harness",
                          "spec": {
                            "files": [
                              "<+org.description>"
                            ]
                          }
                        }
                      }
                    }
                  }
                ]
              }
              EOT
                }

				data "harness_platform_service_overrides_v2" "test" {
                    identifier = harness_platform_service_overrides_v2.test.id
                    org_id = harness_platform_service_overrides_v2.test.org_id
                    project_id = harness_platform_service_overrides_v2.test.project_id
				}
		`, id, name)

}
