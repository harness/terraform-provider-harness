resource "harness_central_notification_channel" "email_channel" {
  name        = "Email Notification Channel"
  identifier  = "email_notification_channel"
  description = "Email channel for critical alerts"
  org_id      = "my_org"
  project_id  = "my_project"

  # Email channel configuration
  email_config {
    recipients         = ["team@example.com", "alerts@example.com"]
    subject           = "[ALERT] Harness Notification"
    send_to_all_admins = false
  }

  # Notification preferences
  notification_method = "EMAIL"
  enabled            = true

  tags = {
    environment = "production"
    team        = "devops"
  }
}
