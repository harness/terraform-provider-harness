# Apply once. The first apply registers the session and returns install_command.
# Install that manifest in the cluster, then apply again. The second apply
# confirms discovery, onboards the services, and confirms the service step.
resource "harness_chaos_service_onboarding" "example" {
  org_id     = "default"
  project_id = "demo"

  environment_id    = "prod"
  infrastructure_id = "payments"
}

output "install_command" {
  value = harness_chaos_service_onboarding.example.install_command
}
