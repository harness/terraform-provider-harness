# Example of fetching a notification channel by identifier
data "harness_central_notification_channel" "example" {
  identifier = "email_channel"
  org_id     = "my_org"
  project_id = "my_project"
}
