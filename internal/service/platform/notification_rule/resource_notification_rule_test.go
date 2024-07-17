package notification_rule_test

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

func TestAccResourceNotificationRule(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_notification_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccNotificationRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNotificationRule(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			{
				Config: testAccResourceNotificationRule(id, updatedName),
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

func testAccGetNotificationRule(resourceName string, state *terraform.State) (*nextgen.NotificationRule, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	resp, _, err := c.SrmNotificationApiService.GetSrmNotification(ctx, id, c.AccountId, buildField(r, "org_id").Value(), buildField(r, "project_id").Value())
	if err != nil {
		return nil, err
	}
	if resp.Resource == nil {
		return nil, fmt.Errorf("empty resource received in response")
	}

	return resp.Resource.NotificationRule, nil
}

func testAccNotificationRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		notificationRule, _ := testAccGetNotificationRule(resourceName, state)
		if notificationRule != nil {
			return fmt.Errorf("Found notification rule: %s", notificationRule.Identifier)
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

func testAccResourceNotificationRule(id string, name string) string {
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

		resource "harness_platform_notification_rule" "test" {
			depends_on = [
				harness_platform_organization.test,
				harness_platform_project.test,
			]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				  name = "%[2]s"

				  type = "ServiceLevelObjective"
				  conditions {
					type = "ErrorBudgetRemainingPercentage"
					spec = jsonencode({
					  threshold = 30
					})
				  }
				  notification_method {
					spec = jsonencode({
					  webhook_url = "http://myslackwebhookurl.com"
					  user_groups = ["account.dsd"]
					})
					type = "Slack"
				  }
			}
		}
`, id, name)
}
