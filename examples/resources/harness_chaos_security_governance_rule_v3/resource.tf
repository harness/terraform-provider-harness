# Example of a Security Governance Rule (V3)
resource "harness_chaos_security_governance_rule_v3" "example" {
  org_id         = var.org_id
  project_id     = var.project_id
  name           = "k8s-security-rule"
  description    = "Security governance rule for Kubernetes chaos experiments"
  is_enabled     = true
  condition_ids  = [harness_chaos_security_governance_condition_v3.k8s_condition.id]
  user_group_ids = ["_project_all_users"]
  tags           = ["env:prod", "team:security"]

  time_windows {
    time_zone  = "UTC"
    start_time = 1711238400000
    # Provide either duration or end_time. end_time must be within one year of
    # start_time; using duration avoids that constraint.
    duration = "24h"

    recurrence {
      type  = "Daily"
      until = -1
    }
  }
}

output "security_governance_rule_v3_id" {
  value = harness_chaos_security_governance_rule_v3.example.id
}
