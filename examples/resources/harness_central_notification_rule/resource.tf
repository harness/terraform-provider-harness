resource "harness_central_notification_rule" "high_severity_rule" {
  name        = "High Severity Alerts Rule"
  identifier  = "high_severity_alerts_rule"
  description = "Rule for high severity alerts notification"
  org_id      = "my_org"
  project_id  = "my_project"

  # Rule conditions
  conditions {
    type  = "PIPELINE"
    event = "PIPELINE_START"

    # Filter for high severity events
    filter {
      type  = "severity"
      value = "HIGH"
    }
  }

  # Notification channels to notify
  notification_channels = [
    "email_notification_channel"
  ]

  # Execution settings
  enabled = true

  tags = {
    environment = "production"
    team        = "devops"
  }
}
