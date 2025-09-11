# Example of fetching the default notification template set
data "harness_default_notification_template_set" "example" {
  identifier = "default"
  org_id     = "my_org"
  project_id = "my_project"
}