# First apply returns install_command. Apply again after installing it.
resource "harness_chaos_service_onboarding" "example" {
  org_id     = "default"
  project_id = "demo"

  environment_id    = "prod"
  infrastructure_id = "payments"
  service_name      = "bank-of-anthos-config"

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
