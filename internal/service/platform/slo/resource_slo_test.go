package slo_test

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSlo(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccSloDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSloForTimeWindow(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testAccResourceSloForRequest(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceSlo_DeleteUnderlyingResource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSloForTimeWindow(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testAccResourceSloForRequest(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccProvider.Meta().(*internal.Session).GetPlatformClient()
					resp, _, err := c.SloApi.DeleteSLODataNg(ctx, c.AccountId, id, id, id)
					require.NoError(t, err)
					require.NotNil(t, resp)
					require.Equal(t, resp.Resource, true)
				},
				Config:             testAccResourceSloForTimeWindow(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetSlo(resourceName string, state *terraform.State) (*nextgen.ServiceLevelObjectiveV2Dto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.SloApi.GetServiceLevelObjectiveNg(ctx, id, c.AccountId, buildField(r, "org_id").Value(), buildField(r, "project_id").Value())
	if err != nil {
		return nil, err
	}

	return resp.Resource.ServiceLevelObjectiveV2, nil
}

func testAccSloDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		slo, _ := testAccGetSlo(resourceName, state)
		if slo != nil {
			return fmt.Errorf("Found SLO: %s", slo.Identifier)
		}

		return nil
	}
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceSloForTimeWindow(id string, name string) string {
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
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			yaml = <<-EOT
			service:
			  name: %[2]s
			  identifier: %[1]s
			EOT
		}

		resource "harness_platform_monitored_service" "test" {
			depends_on = [
				time_sleep.wait_4_seconds
			]
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = harness_platform_service.test.id
				environment_ref = harness_platform_environment.test.id
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "%[1]s"
					type = "ElasticSearch"
					spec = jsonencode({
						connectorRef = "connectorRef"
						feature = "feature"
						queries = [
							{
								name   = "name"
								query = "query"
								index = "index"
								serviceInstanceIdentifier = "serviceInstanceIdentifier"
								timeStampIdentifier = "timeStampIdentifier"
								timeStampFormat = "timeStampFormat"
								messageIdentifier = "messageIdentifier"
							},
							{
								name   = "name2"
								query = "query2"
								index = "index2"
								serviceInstanceIdentifier = "serviceInstanceIdentifier2"
								timeStampIdentifier = "timeStampIdentifier2"
								timeStampFormat = "timeStampFormat2"
								messageIdentifier = "messageIdentifier2"
							}
						]
					})
				}
				change_sources {
					name = "csName1"
					identifier = "harness_cd_next_gen"
					type = "HarnessCDNextGen"
					enabled = true
					spec = jsonencode({
					})
					category = "Deployment"
				}
				template_ref = "template_ref"
				version_label = "version_label"
			}
		}

		resource "harness_platform_slo" "test" {
			depends_on = [
				harness_platform_monitored_service.test,
			]
			org_id = harness_platform_monitored_service.test.org_id
			project_id = harness_platform_monitored_service.test.project_id
			identifier = "%[1]s"
			request {
				  name = "%[2]s"
				  description = "description"
				  tags = ["foo:bar", "bar:foo"]
				  user_journey_refs = ["one", "two"]
				  slo_target {
						type = "Calender"
						slo_target_percentage = 10
						spec = jsonencode({
							type = "Monthly"
							spec = {
								dayOfMonth = 5
							}
						})
				  }
				  type = "Simple"
				  spec = jsonencode({
						monitoredServiceRef = harness_platform_monitored_service.test.id
						serviceLevelIndicatorType = "Availability"
						serviceLevelIndicators = [
							{
								name = "name"
								identifier = "%[1]s"
								type = "Window"
								spec = {
									type = "Threshold"
									spec = {
										metric1 = "metric1"
										thresholdValue = 10
										thresholdType = ">"
									}
                                sliMissingDataType = "Good"
								}
							}
						]
				  })

			}
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_environment.test, harness_platform_service.test]
			destroy_duration = "4s"
		}
`, id, name)
}

func testAccResourceSloForRequest(id string, name string) string {
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
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			yaml = <<-EOT
			service:
			  name: %[2]s
			  identifier: %[1]s
			EOT
		}

		resource "harness_platform_monitored_service" "test" {
			depends_on = [
				time_sleep.wait_4_seconds
			]
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = harness_platform_service.test.id
				environment_ref = harness_platform_environment.test.id
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "%[1]s"
					type = "ElasticSearch"
					spec = jsonencode({
						connectorRef = "connectorRef"
						feature = "feature"
						queries = [
							{
								name   = "name"
								query = "query"
								index = "index"
								serviceInstanceIdentifier = "serviceInstanceIdentifier"
								timeStampIdentifier = "timeStampIdentifier"
								timeStampFormat = "timeStampFormat"
								messageIdentifier = "messageIdentifier"
							},
							{
								name   = "name2"
								query = "query2"
								index = "index2"
								serviceInstanceIdentifier = "serviceInstanceIdentifier2"
								timeStampIdentifier = "timeStampIdentifier2"
								timeStampFormat = "timeStampFormat2"
								messageIdentifier = "messageIdentifier2"
							}
						]
					})
				}
				change_sources {
					name = "csName1"
					identifier = "harness_cd_next_gen"
					type = "HarnessCDNextGen"
					enabled = true
					spec = jsonencode({
					})
					category = "Deployment"
				}
				template_ref = "template_ref"
				version_label = "version_label"
			}
		}

         resource "harness_platform_slo" "test" {
			depends_on = [
				harness_platform_monitored_service.test,
			]
			org_id = harness_platform_monitored_service.test.org_id
			project_id = harness_platform_monitored_service.test.project_id
			identifier = "%[1]s"
			request {
				  name = "%[2]s"
				  description = "description"
				  tags = ["foo:bar", "bar:foo"]
				  user_journey_refs = ["one", "two"]
				  slo_target {
						type = "Calender"
						slo_target_percentage = 10
						spec = jsonencode({
							type = "Monthly"
							spec = {
								dayOfMonth = 5
							}
						})
				  }
				  type = "Simple"
				  spec = jsonencode({
						monitoredServiceRef = harness_platform_monitored_service.test.id
						serviceLevelIndicatorType = "Availability"
						serviceLevelIndicators = [
							{
								name = "name"
								identifier = "%[1]s"
								type = "Request"
								spec = {
                                    eventType = "Good"
                                    metric1 = "metric123"
                                    metric2 = "metric456"
								}
							}
						]
				  })
			}
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_environment.test, harness_platform_service.test]
			destroy_duration = "4s"
		}
`, id, name)
}
