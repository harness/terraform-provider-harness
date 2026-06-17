# Look up a security governance condition (V3) by identity
data "harness_chaos_security_governance_condition_v3" "by_identity" {
  org_id     = var.org_id
  project_id = var.project_id
  identity   = "<condition_id>"
}

# Look up a security governance condition (V3) by name
data "harness_chaos_security_governance_condition_v3" "by_name" {
  org_id     = var.org_id
  project_id = var.project_id
  name       = "k8s-security-condition"
}

output "security_governance_condition_v3_by_identity" {
  value = data.harness_chaos_security_governance_condition_v3.by_identity
}

output "security_governance_condition_v3_by_name" {
  value = data.harness_chaos_security_governance_condition_v3.by_name
}
