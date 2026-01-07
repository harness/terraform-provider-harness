# Example: List GCP projects using a GCP cloud connector
data "harness_platform_gcp_projects" "example" {
  connector_id = "my_gcp_connector"
}

# Example: List GCP projects using a GCP secret manager connector
data "harness_platform_gcp_projects" "example_secret_manager" {
  connector_id = "my_gcp_secret_manager_connector"
}

# Example: List GCP projects with org and project scope
data "harness_platform_gcp_projects" "example_scoped" {
  connector_id = "my_gcp_connector"
  org_id       = "my_org"
  project_id   = "my_project"
}

# Output the projects
output "gcp_projects" {
  value = data.harness_platform_gcp_projects.example.projects
}

