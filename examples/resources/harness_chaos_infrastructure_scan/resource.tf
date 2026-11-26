# Scan an existing chaos infrastructure.
resource "harness_chaos_infrastructure_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_infra_scan"
  name       = "Payments infrastructure scan"

  detection_method = "ai"

  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id
  environment_id    = harness_chaos_infrastructure_v2.payments.environment_id

  # Run this scan again when the infrastructure changes.
  retry_token = harness_chaos_infrastructure_v2.payments.updated_at
}

# Scan by discovery agent.
resource "harness_chaos_infrastructure_scan" "by_agent" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_agent_scan"
  name       = "Payments agent scan"

  agent_identity       = "payments-agent"
  environment_identity = "prod"
}
