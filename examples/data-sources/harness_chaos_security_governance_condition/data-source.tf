# Example of looking up a security governance condition by name
data "harness_chaos_security_governance_condition" "by_name" {
  org_id     = var.org_id
  project_id = var.project_id
  name       = "k8s-security-condition"
}

# Example of looking up a security governance condition by ID
data "harness_chaos_security_governance_condition" "by_id" {
  org_id     = var.org_id
  project_id = var.project_id
  id         = "<condition_id>"
}

# Output the retrieved conditions
output "security_governance_condition_by_name" {
  value = data.harness_chaos_security_governance_condition.by_name
}

output "security_governance_condition_by_id" {
  value = data.harness_chaos_security_governance_condition.by_id
}
