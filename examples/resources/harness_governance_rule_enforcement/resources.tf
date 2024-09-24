resource "harness_governance_rule_enforcement" "example" {
  identifier         = "identifier"
  name               = "name"
  cloud_provider     = "AWS/AZURE/GCP"
  rule_ids           = ["rule_id1"]
  rule_set_ids       = ["rule_set_id1"]
  execution_schedule = "0 0 * * * *"
  execution_timezone = "UTC"
  is_enabled         = true
  target_accounts    = ["awsAccountId/azureSubscriptionId/gcpProjectId"]
  target_regions     = ["us-east-1/eastus"]
  is_dry_run         = false
  description        = "description"
}