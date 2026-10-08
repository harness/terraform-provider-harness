# Lookup by ID
data "harness_autostopping_rule_scale_group" "by_id" {
  identifier = "12345"
}

# Lookup by name (regex)
data "harness_autostopping_rule_scale_group" "by_name" {
  name = "^my-scale-group-rule$"
}

# Lookup by name pattern
data "harness_autostopping_rule_scale_group" "by_pattern" {
  name = "my-asg-.*-prod"
}
