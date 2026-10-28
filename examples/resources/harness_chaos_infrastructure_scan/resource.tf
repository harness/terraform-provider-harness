# Scan an existing chaos infrastructure. infrastructure_id and environment_id
# are the same values as harness_chaos_infrastructure_v2. Terraform reads that
# infrastructure and sends its discovery agent.
# Apply this after install_command has been applied and the agent has discovered
# resources. A scan before that can finish ERRORED or NO_RESOURCES_FOUND.
resource "harness_chaos_infrastructure_scan" "example" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_infra_scan"
  name       = "Payments infrastructure scan"

  detection_method = "ai"

  infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id
  environment_id    = harness_chaos_infrastructure_v2.payments.environment_id

  # Refreshes this same scan when that infrastructure is created or updated.
  # updated_at is read from the API on every refresh. If the server changes it
  # on its own, the next apply runs the scan again.
  retry_token = harness_chaos_infrastructure_v2.payments.updated_at
}

# Imports and callers that already know the discovery agent use this form.
resource "harness_chaos_infrastructure_scan" "by_agent" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_agent_scan"
  name       = "Payments agent scan"

  agent_identity       = "payments-agent"
  environment_identity = "prod"
}
