# Example of fetching a notification rule by identifier
data "harness_central_notification_rule" "example" {
  identifier = "high_severity_rule"
  org_id     = "my_org"
  project_id = "my_project"
}