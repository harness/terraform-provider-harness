terraform {
  required_providers {
    harness = {
      source  = "harness/harness"
      version = "0.40.2"
    }
  }
}

# Returns all autostopping rules without any filtering.
data "harness_autostopping_rules" "all" {}

output "all_rules" {
  value = data.harness_autostopping_rules.all.rules
}

# get the rule ids as a list of integers
output "all_rule_ids" {
  value = [for r in data.harness_autostopping_rules.all.rules : r.id]
}

# Returns only rules of kind "instance"
data "harness_autostopping_rules" "by_instance_kind" {
  kind = "instance"
}

output "by_instance_kind" {
  value = data.harness_autostopping_rules.by_instance_kind.rules
}

# get the rule ids as a list of integers
output "by_instance_kind_ids" {
  value = [for r in data.harness_autostopping_rules.by_instance_kind.rules : r.id]
}

# Returns only rules of kind "k8s"
data "harness_autostopping_rules" "by_k8s_kind" {
  kind = "k8s"
}

output "k8s_rules" {
  value = data.harness_autostopping_rules.by_k8s_kind.rules
}

# get the rule ids as a list of integers
output "k8s_rule_ids" {
  value = [for r in data.harness_autostopping_rules.by_k8s_kind.rules : r.id]
}

# Returns rules whose name starts with "myname-" followed by any characters.
# Regex: "myname-.*" matches e.g. "myname-prod", "myname-01", "myname-anything".
data "harness_autostopping_rules" "by_name_prefix" {
  name = "myname-.*"
}

output "rules_by_name_prefix" {
  value = data.harness_autostopping_rules.by_name_prefix.rules
}

# get the rule ids as a list of integers
output "rules_by_name_prefix_ids" {
  value = [for r in data.harness_autostopping_rules.by_name_prefix.rules : r.id]
}

# Returns rules whose name starts with "app" or "svc" followed by any characters.
# Regex: "^(app|svc).*" matches e.g. "app-prod", "svc-backend", "appserver" but NOT "myapp".
data "harness_autostopping_rules" "by_name_regex" {
  name = "^(app|svc).*"
}

output "rules_by_name_regex" {
  value = data.harness_autostopping_rules.by_name_regex.rules
}

# get the rule ids as a list of integers
output "rules_by_name_regex_ids" {
  value = [for r in data.harness_autostopping_rules.by_name_regex.rules : r.id]
}
