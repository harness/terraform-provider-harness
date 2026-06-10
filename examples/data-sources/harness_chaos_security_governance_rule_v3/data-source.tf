# Look up a security governance rule (V3) by identity
data "harness_chaos_security_governance_rule_v3" "by_identity" {
  org_id     = var.org_id
  project_id = var.project_id
  identity   = "<rule_id>"
}

# Look up a security governance rule (V3) by name
data "harness_chaos_security_governance_rule_v3" "by_name" {
  org_id     = var.org_id
  project_id = var.project_id
  name       = "k8s-security-rule"
}

output "security_governance_rule_v3_by_identity" {
  value = data.harness_chaos_security_governance_rule_v3.by_identity
}

output "security_governance_rule_v3_by_name" {
  value = data.harness_chaos_security_governance_rule_v3.by_name
}
