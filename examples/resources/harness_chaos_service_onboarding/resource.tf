# Apply once. The first apply registers the session and returns install_command.
# Install that manifest in the cluster, then apply again. The second apply
# confirms discovery and creates one chaos service with these probes.
# Another discovered service on this infrastructure is a harness_chaos_onboarding_service
# that uses this resource's onboarding_id.
resource "harness_chaos_service_onboarding" "example" {
  org_id     = "default"
  project_id = "demo"

  environment_id        = "prod"
  infrastructure_id     = "payments"
  # Name shown on the service card. Use discovered_service_id when two services share this name.
  service_name = "bank-of-anthos-config"

  probes {
    probe_id = "http-check"
    inputs = {
      url = "http://access-control"
    }
  }

  probes {
    probe_id = "pod-status"
  }
}

output "install_command" {
  value = harness_chaos_service_onboarding.example.install_command
}
