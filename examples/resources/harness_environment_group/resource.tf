resource "platform_environment_group" "test" {
  identifier = "identifier"
  org_id     = "orgIdentifer"
  project_id = "projectIdentifier"
  color = "#0063F7"
  yaml       = <<-EOT
   ---
  environmentGroup:
    name: "%[1]s"
    identifier: "%[1]s"
    description: "temp"
    orgIdentifier: ${harness_platform_project.test.org_id}
    projectIdentifier: ${harness_platform_project.test.id}
    envIdentifiers: []
    EOT
}
