resource "harness_platform_pipeline_central_notification_rule" "projExample" {
  identifier                = "identifier"
  name                      = "name"
  status                    = "ENABLED"
  notification_channel_refs = ["account.channel"]
  org = "default"
  project = "proj0"

  notification_conditions {
    condition_name = "pipelineRuleProjectConditionName"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = []
      }
      entity_identifiers = []
    }
  }
}

resource "harness_platform_pipeline_central_notification_rule" "orgExample" {
  identifier                = "identifier"
  name                      = "name"
  status                    = "ENABLED"
  notification_channel_refs = ["channel"]
  org = "default"

  notification_conditions {
    condition_name = "pipelineRuleOrgConditionName"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = ["proj0", "random"]
      }
      entity_identifiers = []
    }
  }
}

resource "harness_platform_pipeline_central_notification_rule" "accountExample" {
  identifier                = "identifier"
  name                      = "name"
  status                    = "DISABLED"
  notification_channel_refs = ["org.channel"]

  notification_conditions {
    condition_name = "pipelineRuleConditionName"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = ["org"]
      }
      entity_identifiers = []
    }
  }
}