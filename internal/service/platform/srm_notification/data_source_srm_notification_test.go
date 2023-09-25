package srm_notification_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSrmNotification(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6)) //Add with muliptle logs and metrics
	name := id
	resourceName := "data.harness_platform_srm_notification.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSrmNotification(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testSrmNotification(id string, name string) string {
	return fmt.Sprintf(
		`
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

	resource "harness_platform_srm_notification" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name            = "%[2]s"
			type            = "ServiceLevelObjective"
			conditions {
			  type       = "ErrorBudgetRemainingPercentage"
			  spec = jsonencode({
				threshold = 100
			  })
			}
			notification_method {
			  type       = "Slack"
			  spec = jsonencode({
				userGroups : ["userGroups1", "userGroups2"]
				webhookUrl : "https://expamle.slack.com/"
			  })
			}
		}
	}

	data "harness_platform_srm_notification" "test" {
		identifier = harness_platform_srm_notification.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}
