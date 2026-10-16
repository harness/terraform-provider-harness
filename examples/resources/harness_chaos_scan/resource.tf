# Scan a pipeline execution whose status is Success and that deployed with a
# Kubernetes CD stage. Apply waits until the scan finishes. An ERRORED or
# ABORTED scan fails the apply. A COMPLETED scan can still have no risks.
resource "harness_chaos_scan" "pipeline" {
  org_id     = "default"
  project_id = "demo"
  identity   = "cart_deploy_scan"
  name       = "Cart deploy scan"

  # UI dropdown: "ai" is AI Scan, "ruleBased" is Rule-based Scan.
  # "AI" and "RiskAnalyzer" are accepted as the same choices.
  # Omit detection_method to let the API use AI Scan.
  detection_method = "ruleBased"

  pipeline_scan {
    pipeline_identity     = "cart_deploy"
    pipeline_execution_id = "exec-123"
  }
}

# Scan a Kubernetes infrastructure. infrastructure_id and environment_id are
# the chaos infrastructure, the same values as harness_chaos_infrastructure_v2.
# Terraform reads the linked discovery agent and sends the same source the UI
# sends for Scan infrastructure.
resource "harness_chaos_scan" "infrastructure" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_infra_scan"
  name       = "Payments infrastructure scan"

  # AI Scan. Use "ruleBased" for a Rule-based Scan.
  detection_method = "ai"

  infrastructure_scan {
    infrastructure_id = harness_chaos_infrastructure_v2.payments.infra_id
    environment_id    = harness_chaos_infrastructure_v2.payments.environment_id
  }

  # Refreshes this same scan when the Terraform infrastructure is created or
  # updated. An apply that leaves updated_at unchanged does not scan again.
  retry_token = harness_chaos_infrastructure_v2.payments.updated_at
}

# The discovery agent can still be set directly. Imports use this form because
# the API stores the agent, not the chaos infrastructure id.
resource "harness_chaos_scan" "infrastructure_by_agent" {
  org_id     = "default"
  project_id = "demo"
  identity   = "payments_agent_scan"
  name       = "Payments agent scan"

  infrastructure_scan {
    agent_identity       = "payments-agent"
    environment_identity = "prod"
  }
}

# Risks are computed. They stay empty until the scan finishes.
output "pipeline_scan_risks" {
  value = harness_chaos_scan.pipeline.risks
}
