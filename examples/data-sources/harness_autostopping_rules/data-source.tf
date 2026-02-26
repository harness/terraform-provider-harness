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

# Returns only rules of kind "instance"
data "harness_autostopping_rules" "by_instance_kind" {
  kind = "instance"
}

output "by_instance_kind" {
  value = data.harness_autostopping_rules.by_instance_kind.rules
}

# Returns only rules of kind "k8s"
data "harness_autostopping_rules" "by_k8s_kind" {
  kind = "k8s"
}

output "k8s_rules" {
  value = data.harness_autostopping_rules.by_k8s_kind.rules
}

# Returns rules whose name starts with "myname-" followed by any characters.
# Regex: "myname-.*" matches e.g. "myname-prod", "myname-01", "myname-anything".
data "harness_autostopping_rules" "by_name_prefix" {
  name = "myname-.*"
}

output "rules_by_name_prefix" {
  value = data.harness_autostopping_rules.by_name_prefix.rules
}

# Returns rules whose name starts with "app" or "svc" followed by any characters.
# Regex: "^(app|svc).*" matches e.g. "app-prod", "svc-backend", "appserver" but NOT "myapp".
data "harness_autostopping_rules" "by_name_regex" {
  name = "^(app|svc).*"
}

output "rules_by_name_regex" {
  value = data.harness_autostopping_rules.by_name_regex.rules
}
