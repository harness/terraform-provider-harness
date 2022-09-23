resource "harness_platform_environment_group" "test" {
  identifier = "identifier"
  org_id     = "orgIdentifer"
  project_id = "projectIdentifier"
  color      = "#0063F7"
  yaml       = <<-EOT
  environmentGroup:
    name: "name"
    identifier: "identifier"
    description: "temp"
    orgIdentifier: "orgIdentifer"
    projectIdentifier: "projectIdentifier"
    envIdentifiers: []
    EOT
}
