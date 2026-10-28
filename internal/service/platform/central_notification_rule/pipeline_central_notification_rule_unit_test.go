package central_notification_rule

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testPipelineRuleDto(params *nextgen.PipelineEventNotificationParamsDto) nextgen.NotificationRuleDto {
	return nextgen.NotificationRuleDto{
		Identifier: "rule1",
		Name:       "rule1",
		NotificationConditions: []nextgen.NotificationConditionDto{
			{
				ConditionName: "cond",
				NotificationEventConfigs: []nextgen.NotificationEventConfigDto{
					{
						NotificationEntity:                 "PIPELINE",
						NotificationEvent:                  "PIPELINE_FAILED",
						PipelineEventNotificationParamsDto: params,
					},
				},
			},
		},
	}
}

// When the API returns no notification_event_data, read must not fabricate a
// default block in state - doing so causes perpetual plan drift (PL-75312).
func TestReadPipelineCentralNotificationRule_nullEventData(t *testing.T) {
	r := ResourcePipelineCentralNotificationRule()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{})

	diags := readPipelineCentralNotificationRule("acc", d, testPipelineRuleDto(nil))
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	conditions := d.Get("notification_conditions").([]interface{})
	if len(conditions) != 1 {
		t.Fatalf("expected 1 condition, got %d", len(conditions))
	}
	eventConfigs := conditions[0].(map[string]interface{})["notification_event_configs"].([]interface{})
	if len(eventConfigs) != 1 {
		t.Fatalf("expected 1 event config, got %d", len(eventConfigs))
	}
	eventData := eventConfigs[0].(map[string]interface{})["notification_event_data"].([]interface{})
	if len(eventData) != 0 {
		t.Fatalf("expected no notification_event_data block, got %v", eventData)
	}
}

// When the API returns notification_event_data, read must populate it in state.
func TestReadPipelineCentralNotificationRule_withEventData(t *testing.T) {
	r := ResourcePipelineCentralNotificationRule()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{})

	pipelineType := nextgen.PIPELINE_ResourceTypeEnum
	params := &nextgen.PipelineEventNotificationParamsDto{Type_: &pipelineType}

	diags := readPipelineCentralNotificationRule("acc", d, testPipelineRuleDto(params))
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	conditions := d.Get("notification_conditions").([]interface{})
	eventConfigs := conditions[0].(map[string]interface{})["notification_event_configs"].([]interface{})
	eventData := eventConfigs[0].(map[string]interface{})["notification_event_data"].([]interface{})
	if len(eventData) != 1 {
		t.Fatalf("expected 1 notification_event_data block, got %d", len(eventData))
	}
	block := eventData[0].(map[string]interface{})
	if block["type"].(string) != "PIPELINE" {
		t.Fatalf("expected type PIPELINE, got %v", block["type"])
	}
	if scopeIds := block["scope_identifiers"].([]interface{}); len(scopeIds) != 0 {
		t.Fatalf("expected empty scope_identifiers, got %v", scopeIds)
	}
}
