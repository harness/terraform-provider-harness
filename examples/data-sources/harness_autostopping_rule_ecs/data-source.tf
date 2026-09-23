# Lookup by ID
data "harness_autostopping_rule_ecs" "by_id" {
  identifier = "12345"
}

# Lookup by name (regex)
data "harness_autostopping_rule_ecs" "by_name" {
  name = "^my-ecs-rule$"
}

# Lookup by name pattern
data "harness_autostopping_rule_ecs" "by_pattern" {
  name = "my-ecs-.*-prod"
}
