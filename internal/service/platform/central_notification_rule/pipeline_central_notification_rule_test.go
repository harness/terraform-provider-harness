package central_notification_rule_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcePipelineCentralNotificationRule_basic(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.0.type", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.0.scope_identifiers.#", "0"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_accountLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_Account"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"
	expectedRuleIdentifier := id + "_channel_account"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleAccountConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", expectedRuleIdentifier),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_START"),
					resource.TestCheckNoResourceAttr(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_projectLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_Project"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleProjectConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_FAILED"),
					resource.TestCheckResourceAttrSet(resourceName, "org"),
					resource.TestCheckResourceAttrSet(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_orgLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_Org"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"
	expectedRuleIdentifier := id + "_channel_org"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleOrgConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", expectedRuleIdentifier),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_SUCCESS"),
					resource.TestCheckResourceAttrSet(resourceName, "org"),
					resource.TestCheckNoResourceAttr(resourceName, "project"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_multipleEvents(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_MultipleEvents"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleMultipleEventsConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.condition_name", "pipeline-start-condition"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_START"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.1.condition_name", "pipeline-success-condition"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.1.notification_event_configs.0.notification_event", "PIPELINE_SUCCESS"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.2.condition_name", "pipeline-failed-condition"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.2.notification_event_configs.0.notification_event", "PIPELINE_FAILED"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_multipleChannels(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_MultipleChannels"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleMultipleChannelsConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "2"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_disabled(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_Disabled"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleDisabledConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_withNotificationEventData(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_EventData"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleScopeIdentifiersConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.0.type", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event_data.0.scope_identifiers.#", "0"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceName]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceName)
						}
						fmt.Printf("DEBUG: Full state for %s:\n", resourceName)
						for k, v := range rs.Primary.Attributes {
							if strings.Contains(k, "notification_event_data") || strings.Contains(k, "scope_identifier") {
								fmt.Printf("  %s = %s\n", k, v)
							}
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_allPipelineEvents(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_AllEvents"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleAllEventsConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.#", "6"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_START"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.1.notification_event_configs.0.notification_event", "PIPELINE_SUCCESS"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.2.notification_event_configs.0.notification_event", "PIPELINE_FAILED"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.3.notification_event_configs.0.notification_event", "STAGE_START"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.4.notification_event_configs.0.notification_event", "STAGE_SUCCESS"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.5.notification_event_configs.0.notification_event", "STAGE_FAILED"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_pipelineStartEvent(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_PipelineStart"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleSingleEventConfig(rName, id, "PIPELINE_START"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_START"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_pipelineSuccessEvent(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_PipelineSuccess"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleSingleEventConfig(rName, id, "PIPELINE_SUCCESS"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_SUCCESS"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_stageStartEvent(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_StageStart"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleSingleEventConfig(rName, id, "STAGE_START"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "STAGE_START"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_stageSuccessEvent(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_StageSuccess"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleSingleEventConfig(rName, id, "STAGE_SUCCESS"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "STAGE_SUCCESS"),
				),
			},
		},
	})
}

func TestAccResourcePipelineCentralNotificationRule_stageFailedEvent(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	rName := "TestAccResourcePipelineCentralNotificationRule_StageFailed"
	resourceName := "harness_platform_pipeline_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineCentralNotificationRuleSingleEventConfig(rName, id, "STAGE_FAILED"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "STAGE_FAILED"),
				),
			},
		},
	})
}

func testAccResourcePipelineCentralNotificationRuleConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "test-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleAccountConfig(name, id string) string {
	channelIdentifier := id + "_channel"
	ruleIdentifier := channelIdentifier + "_account"
	scopeChannelIdentifier := "account." + channelIdentifier

	config := fmt.Sprintf(`
resource "harness_platform_central_notification_channel" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  depends_on                = [harness_platform_central_notification_channel.test]
  identifier                = "%[3]s"
  name                      = "%[2]s"
  status                    = "ENABLED"
  notification_channel_refs = ["%[4]s"]

  notification_conditions {
    condition_name = "account-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, channelIdentifier, name, ruleIdentifier, scopeChannelIdentifier)
	fmt.Printf("DEBUG AccountLevel Config: name=%s, id=%s, channelId=%s, ruleId=%s, scopeChannelId=%s\nConfig:\n%s\n", name, id, channelIdentifier, ruleIdentifier, scopeChannelIdentifier, config)
	return config
}

func testAccResourcePipelineCentralNotificationRuleOrgConfig(name, id string) string {
	channelIdentifier := id + "_channel"
	ruleIdentifier := channelIdentifier + "_org"
	scopeChannelIdentifier := "org." + channelIdentifier

	config := fmt.Sprintf(`
resource "harness_platform_organization" "test" {
  identifier = "%[1]s"
  name       = "%[2]s"
}

resource "harness_platform_central_notification_channel" "test" {
  identifier                = "%[3]s"
  org                       = harness_platform_organization.test.id
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  depends_on                = [harness_platform_central_notification_channel.test]
  identifier                = "%[4]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  status                    = "ENABLED"
  notification_channel_refs = ["%[5]s"]

  notification_conditions {
    condition_name = "org-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_SUCCESS"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name, channelIdentifier, ruleIdentifier, scopeChannelIdentifier)
	fmt.Printf("DEBUG OrgLevel Config: name=%s, id=%s, channelId=%s, ruleId=%s, scopeChannelId=%s\nConfig:\n%s\n", name, id, channelIdentifier, ruleIdentifier, scopeChannelIdentifier, config)
	return config
}

func testAccResourcePipelineCentralNotificationRuleMultipleEventsConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "pipeline-start-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "pipeline-success-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_SUCCESS"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "pipeline-failed-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleMultipleChannelsConfig(name, id string) string {
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

resource "harness_platform_central_notification_channel" "email_test" {
  identifier                = "%[1]s_email_channel"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  name                      = "%[2]s_Email_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_central_notification_channel" "email2_test" {
  identifier                = "%[1]s_email2_channel"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  name                      = "%[2]s_Email2_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify2@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [
    harness_platform_central_notification_channel.email_test.identifier,
    harness_platform_central_notification_channel.email2_test.identifier
  ]

  notification_conditions {
    condition_name = "multiple-channels-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleDisabledConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "DISABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "disabled-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleScopeIdentifiersConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "scope-identifiers-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleAllEventsConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "pipeline-start-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "pipeline-success-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_SUCCESS"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "pipeline-failed-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "stage-start-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "STAGE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "stage-success-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "STAGE_SUCCESS"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }

  notification_conditions {
    condition_name = "stage-failed-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "STAGE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}

func testAccResourcePipelineCentralNotificationRuleSingleEventConfig(name, id, event string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "single-event-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "%[3]s"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name, event)
}

func buildPipelineField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccGetPipelineCentralNotificationRule(resourceName string, state *terraform.State) (*nextgen.NotificationRuleDto, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	if r == nil || r.Primary == nil {
		return nil, nil
	}
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	id := r.Primary.ID
	org := buildPipelineField(r, "org").Value()
	project := buildPipelineField(r, "project").Value()

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
		// Any 404 error or resource not found error means the resource is destroyed
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		// Check for various error patterns that indicate the resource doesn't exist
		errMsg := err.Error()
		if strings.Contains(errMsg, "does not exist") ||
			strings.Contains(errMsg, "not found") ||
			strings.Contains(errMsg, "RESOURCE_NOT_FOUND") ||
			strings.Contains(errMsg, "Project with identifier") ||
			strings.Contains(errMsg, "Organization with identifier") ||
			strings.Contains(errMsg, "404") {
			return nil, nil
		}
		return nil, err
	}

	return &resp, nil
}

func testAccCheckPipelineCentralNotificationRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, err := testAccGetPipelineCentralNotificationRule(resourceName, state)
		if err != nil {
			return err
		}
		if rule != nil {
			return fmt.Errorf("found pipeline notification rule: %s", rule.Identifier)
		}
		return nil
	}
}

func testAccResourcePipelineCentralNotificationRuleProjectConfig(name, id string) string {
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
  name                      = "%[2]s_Channel"
  notification_channel_type = "EMAIL"
  status                    = "ENABLED"

  channel {
    email_ids = ["notify@harness.io"]
  }
}

resource "harness_platform_pipeline_central_notification_rule" "test" {
  identifier                = "%[1]s"
  name                      = "%[2]s"
  org                       = harness_platform_organization.test.id
  project                   = harness_platform_project.test.id
  status                    = "ENABLED"
  notification_channel_refs = [harness_platform_central_notification_channel.test.identifier]

  notification_conditions {
    condition_name = "project-condition"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_FAILED"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
    }
  }
}
`, id, name)
}
