package central_notification_rule_test

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"net/http"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCentralNotificationRule(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourceCentralNotificationRule"
	resourceName := "harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckCentralNotificationRuleDestroy(resourceName), // optionally implement a destroy check
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCentralNotificationRuleConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("harness_platform_central_notification_rule.test", "identifier", id),
					resource.TestCheckResourceAttr("harness_platform_central_notification_rule.test", "name", rName),
					resource.TestCheckResourceAttr("harness_platform_central_notification_rule.test", "status", "ENABLED"),
					resource.TestCheckResourceAttr("harness_platform_central_notification_rule.test", "notification_channel_refs.#", "1"),
					resource.TestCheckResourceAttr("harness_platform_central_notification_rule.test", "notification_conditions.#", "1"),
				),
			},
		},
	})
}

func testAccResourceCentralNotificationRuleConfig(name, id string) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
  identifier = "%[1]s"
  name       = "%[2]s"
}

resource "harness_platform_project" "test" {
  identifier = "%[1]s"
  name       = "%[2]s"
  org_id     = harness_platform_organization.test.id
  color      = "#472848"
}

resource "harness_platform_central_notification_channel" "test" {
  identifier                = "%[1]s_channel"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  name                      = "%[2]s Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_central_notification_rule" "test" {
  identifier                = "`+id+`"
  name                      = "`+name+`"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "test-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"
      notification_event_data = {
        type = "PIPELINE"
      }
    }
  }
}
`, id, name)
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccGetCentralNotificationRule(resourceName string, state *terraform.State) (*nextgen.NotificationRuleDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	id := r.Primary.ID
	org := buildField(r, "org").Value()
	project := buildField(r, "project").Value()

	var (
		resp     nextgen.NotificationRuleDto
		httpResp *http.Response
		err      error
	)

	if org != "" && project != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRule(ctx, org, project, id,
			&nextgen.NotificationRulesApiGetNotificationRuleOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	} else if org != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleOrg(ctx, org, id,
			&nextgen.NotificationRulesApiGetNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	} else {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleAccount(ctx, id,
			&nextgen.NotificationRulesApiGetNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(c.AccountId),
			})
	}

	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil, nil // expected
		}
		return nil, err
	}

	return &resp, nil
}

func testAccCheckCentralNotificationRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, err := testAccGetCentralNotificationRule(resourceName, state)
		if err != nil {
			return err
		}
		if rule != nil {
			return fmt.Errorf("found notification rule: %s", rule.Identifier)
		}
		return nil
	}
}
