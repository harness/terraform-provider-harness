package monitored_service_test

import (
	"fmt"
	"github.com/antihax/optional"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceMonitoredService(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testAccResourceMonitoredService(id, updatedName),
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

func testAccGetMonitoredService(resourceName string, state *terraform.State) (*nextgen.MonitoredService, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	resp, _, err := c.MonitoredServiceApi.GetMonitoredService(ctx, id, c.AccountId, buildField(r, "org_id").Value(), buildField(r, "project_id").Value())
	if err != nil {
		return nil, err
	}

	return resp.Data.MonitoredService, nil
}

func testAccMonitoredServiceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		monitoredService, _ := testAccGetMonitoredService(resourceName, state)
		if monitoredService != nil {
			return fmt.Errorf("Found monitored service: %s", monitoredService.Identifier)
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

func testAccResourceMonitoredService(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "DatadogLog"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "feature"
					queries = [
						{
							name   = "name"
							query = "query"
							indexes = ["index"]
							serviceInstanceIdentifier = "serviceInstanceIdentifier"
						}
					]})
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

				enabled = true
			}
		}
		resource "harness_platform_monitored_service" "test1" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "service_ref1_environment_ref"
			request {
				name = "service_ref1_environment_ref"
				type = "Application"
				description = "description"
				service_ref = "service_ref1"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "DatadogLog"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "feature"
					queries = [
						{
							name   = "name"
							query = "query"
							indexes = ["index"]
							serviceInstanceIdentifier = "serviceInstanceIdentifier"
						}
					]})
				}
				dependencies {
					monitored_service_identifier = "%[1]s"
				}
			}
		}
`, id, name)
}

func TestMonitoredServiceWithoutChangeSource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testMonitoredServiceWithoutChangeSource(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testMonitoredServiceWithoutChangeSource(id, updatedName),
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

func testMonitoredServiceWithoutChangeSource(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "DatadogLog"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "feature"
					queries = [
						{
							name   = "name"
							query = "query"
							indexes = ["index"]
							serviceInstanceIdentifier = "serviceInstanceIdentifier"
						}
					]})
				}

				enabled = true
			}
		}
`, id, name)
}

func TestMonitoredServiceWithoutEnabled(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testMonitoredServiceWithoutEnabled(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testMonitoredServiceWithoutEnabled(id, updatedName),
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

func testMonitoredServiceWithoutEnabled(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "DatadogLog"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "feature"
					queries = [
						{
							name   = "name"
							query = "query"
							indexes = ["index"]
							serviceInstanceIdentifier = "serviceInstanceIdentifier"
						}
					]})
				}
			}
		}
`, id, name)
}

func TestAccResourceMonitoredServiceWithAppD(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredServiceWithAppD(id, name),
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

func TestAccResourceMonitoredServiceWithGCPLogs(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredServiceWithGCPLogs(id, name),
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

func TestAccResourceMonitoredServiceWithSplunkLogs(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredServiceWithSplunkLogs(id, name),
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

func TestAccResourceMonitoredServiceWithDynatrace(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredServiceWithDynatrace(id, name),
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

func TestAccResourceMonitoredServiceWithNewRelic(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredServiceWithNewRelic(id, name),
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

func testAccResourceMonitoredServiceWithAppD(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "AppDynamics"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "Application Monitoring"
					metricPacks = [ {
						identifier= "Errors" 
					} ]
					applicationName = "cv-app"
					tierName = "docker-tier"
					})
				}
			}
		}
`, id, name)
}

func testAccResourceMonitoredServiceWithGCPLogs(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "StackdriverLog"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "Cloud Logs"
					queries = [
							{
								name   = "GCO Logs Query"
								query = "error"
								messageIdentifier = "['jsonPayload'].['message']"
								serviceInstanceIdentifier = "['resource'].['labels'].['pod_name']"
							}
						]
					})
				}
			}
		}
`, id, name)
}

func testAccResourceMonitoredServiceWithSplunkLogs(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "Splunk"
					spec = jsonencode({
					connectorRef = "connectorRef"
					feature = "Splunk Cloud Logs"
					queries = [
							{
								identifier = "SPLUNK_Logs_Query"
								name   = "SPLUNK Logs Query"
								query = "index=_internal \" error \" NOT debug source=*splunkd.log*"
								serviceInstanceIdentifier = "['host']"
							}
						]
					})
				}
			}
		}
`, id, name)
}

func testAccResourceMonitoredServiceWithDynatrace(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "Dynatrace"
					spec = jsonencode({
					connectorRef = "account.dynatraceconnectorforautomation"
					feature = "dynatrace_apm"
					metricPacks = [ {
						identifier= "Performance"
					},
					{
						identifier= "Infrastructure"
					}]
					serviceId = "SERVICE-D739201C4CBBA618"
					serviceMethodIds = [
						"SERVICE_METHOD-F3988BEE84FF7388"
					]
					serviceName = ":4444"
					})
				}
			}
		}
`, id, name)
}

func testAccResourceMonitoredServiceWithNewRelic(id string, name string) string {
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

		resource "harness_platform_monitored_service" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				name = "%[2]s"
				type = "Application"
				description = "description"
				service_ref = "service_ref"
				environment_ref = "environment_ref"
				tags = ["foo:bar", "bar:foo"]
				health_sources {
					name = "name"
					identifier = "identifier"
					type = "NewRelic"
					spec = jsonencode({
					connectorRef = "account.Newrelicautomation_do_not_delete"
					feature = "apm"
					applicationId = "107019083"
					applicationName = "My Application"
					metricPacks = [ {
							identifier = "Performance"
						} ]
					})
				}
			}
		}
`, id, name)
}
