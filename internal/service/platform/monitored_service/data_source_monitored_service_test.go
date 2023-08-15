package monitored_service_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMonitoredService(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_monitored_service.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceMonitoredService(id string, name string) string {
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
		org_id = harness_platform_organization.test.id
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
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
`, id, name)
}
