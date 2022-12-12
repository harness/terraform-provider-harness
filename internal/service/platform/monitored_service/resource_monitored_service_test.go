package monitored_service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceMonitoredService(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	org := "default"
	project := "default_project"
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName, org, project),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredService(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceMonitoredService(id, updatedName, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceMonitoredService_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	org := "default"
	project := "default_project"
	resourceName := "harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccMonitoredServiceDestroy(resourceName, org, project),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMonitoredService(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					resp, _, err := c.MonitoredServiceApi.DeleteMonitoredService(ctx, id, c.AccountId, org, project)
					require.NoError(t, err)
					require.True(t, resp.Resource)
				},
				Config:             testAccResourceMonitoredService(id, name, accountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetMonitoredService(resourceName string, org string, project string, state *terraform.State) (*nextgen.MonitoredServiceDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.MonitoredServiceApi.GetMonitoredService(ctx, id, c.AccountId, org, project)
	if err != nil {
		return nil, err
	}

	return resp.Data.MonitoredService, nil
}

func testAccMonitoredServiceDestroy(resourceName string, org string, project string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		monitoredService, _ := testAccGetMonitoredService(resourceName, org, project, state)
		if monitoredService != nil {
			return fmt.Errorf("Found monitored service: %s", monitoredService.Identifier)
		}

		return nil
	}
}

func testAccResourceMonitoredService(id string, name string, accountId string) string {
	return fmt.Sprintf(`
		resource "harness_platform_monitored_service" "test" {
			account_id = "%[3]s"
			org_id     = "default"
			project_id = "default_project"
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
				dependencies {
					monitored_service_identifier = "monitored_service_identifier"
					type = "KUBERNETES"
					dependency_metadata = jsonencode({
						namespace = "namespace"
						workload = "workload"
						type = "KUBERNETES"
					})
				}
				notification_rule_refs {
					notification_rule_ref = "notification_rule_ref"
					enabled = true
				}
				notification_rule_refs {
					notification_rule_ref = "notification_rule_ref1"
					enabled = false
				}
				template_ref = "template_ref"
				version_label = "version_label"
				enabled = true
			}
		}
`, id, name, accountId)
}
