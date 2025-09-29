resource "harness_platform_pipeline_central_notification_rule" "example" {
  identifier                = "identifier"
  name                      = "name"
  status                    = "ENABLED"
  notification_channel_refs = ["account.notification_channel_ref"]
  org = "org_id"
  project = "project_id"

  notification_conditions {
    condition_name = "condition_name"

    notification_event_configs {
      notification_entity = "PIPELINE"
      notification_event  = "PIPELINE_START"

      notification_event_data {
        type              = "PIPELINE"
        scope_identifiers = ["scope_identifier"]
      }

      entity_identifiers = []
    }
  }
  custom_notification_template_ref {
    template_ref = "org.orgTemplate"
    version_label = "1"
    variables {
      name = "variableName"
      value = "1"
      type = "string"
    }
  }

}