# Data sources to verify the rules
data "harness_chaos_security_governance_rule" "example" {
  id         = "<rule_id>"
  org_id     = "<org_id>"
  project_id = "<project_id>"
}

data "harness_chaos_security_governance_rule" "example_linux" {
  id         = "<rule_id>"
  org_id     = "<org_id>"
  project_id = "<project_id>"
}

data "harness_chaos_security_governance_rule" "example_windows" {
  id         = "<rule_id>"
  org_id     = "<org_id>"
  project_id = "<project_id>"
}
