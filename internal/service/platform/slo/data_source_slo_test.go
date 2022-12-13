package slo_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSlo(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_slo.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSlo(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceSlo(id string, name string) string {
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

	resource "harness_platform_slo" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			  name = "%[2]s"
			  description = "description"
			  tags = ["foo:bar", "bar:foo"]
			  user_journey_refs = ["one", "two"]
			  slo_target {
					type = "Rolling"
					slo_target_percentage = 10.0
					spec = jsonencode({
						periodLength = "28d"
					})
			  }
			  type = "Simple"
			  spec = jsonencode({
					monitoredServiceRef = harness_platform_monitored_service.test.identifier
					healthSourceRef = "healthSourceRef"
					serviceLevelIndicatorType = "serviceLevelIndicatorType"
			  })
			  notification_rule_refs {
					notification_rule_ref = "notification_rule_ref"
					enabled = true
			  }
		}
	}

	data "harness_platform_slo" "test" {
		identifier = harness_platform_slo.test.identifier
		org_id = harness_platform_slo.test.org_id
		project_id = harness_platform_slo.test.project_id
	}
`, id, name)
}
