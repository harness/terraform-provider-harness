# Lookup by ID
data "harness_autostopping_rule_rds" "by_id" {
  identifier = "12345"
}

# Lookup by name (regex)
data "harness_autostopping_rule_rds" "by_name" {
  name = "^my-rds-rule$"
}

# Lookup by name pattern
data "harness_autostopping_rule_rds" "by_pattern" {
  name = "my-rds-.*-prod"
}
