# Lookup by ID
data "harness_autostopping_rule_vm" "by_id" {
  identifier = "12345"
}

# Lookup by name (regex)
data "harness_autostopping_rule_vm" "by_name" {
  name = "^my-vm-rule$"
}

# Lookup by name pattern
data "harness_autostopping_rule_vm" "by_pattern" {
  name = "my-vm-.*-prod"
}
