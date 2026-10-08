# Lookup by ID
data "harness_autostopping_rule_k8s" "by_id" {
  identifier = "12345"
}

# Lookup by name (regex)
data "harness_autostopping_rule_k8s" "by_name" {
  name = "^my-k8s-rule$"
}

# Lookup by name pattern
data "harness_autostopping_rule_k8s" "by_pattern" {
  name = "my-k8s-.*-prod"
}
