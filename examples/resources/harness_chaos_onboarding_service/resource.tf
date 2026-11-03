# Creates one more discovered service on the session from harness_chaos_service_onboarding.
# This resource does not register an onboarding session.
resource "harness_chaos_onboarding_service" "example" {
  org_id     = "default"
  project_id = "demo"

  environment_id        = "prod"
  infrastructure_id     = "payments"
  onboarding_id         = harness_chaos_service_onboarding.example.onboarding_id
  service_name = "bank-of-anthos-accounts-db"

  probes {
    probe_id = "pod-status"
  }
}
