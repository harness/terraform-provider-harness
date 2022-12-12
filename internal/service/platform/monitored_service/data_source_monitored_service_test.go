package monitored_service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMonitoredService(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	orgId := "default"
	projectId := "default_project"
	name := id
	resourceName := "data.harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMonitoredService(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "project_id", projectId),
				),
			},
		},
	})
}

func testAccDataSourceMonitoredService(id string, name string, accountId string) string {
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

	data "harness_platform_monitored_service" "test" {
		account_id = harness_platform_monitored_service.test.account_id
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_monitored_service.test.org_id
		project_id = harness_platform_monitored_service.test.project_id
	}
`, id, name, accountId)
}
