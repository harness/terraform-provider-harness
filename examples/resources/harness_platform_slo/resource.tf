resource "harness_platform_slo" "example" {
  account_id = "account_id"
  org_id     = "default"
  project_id = "default_project"
  identifier = "TerraformSLO"
  request {
    name              = "TSLO"
    description       = "description"
    tags              = ["foo:bar", "bar:foo"]
    user_journey_refs = ["one", "two"]
    slo_target {
      type                  = "Rolling"
      slo_target_percentage = 10.0
      spec = jsonencode({
        periodLength = "28d"
      })
    }
    type = "Simple"
    spec = jsonencode({
      monitoredServiceRef       = "monitoredServiceRef"
      healthSourceRef           = "healthSourceRef"
      serviceLevelIndicatorType = "serviceLevelIndicatorType"
    })
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref"
      enabled               = true
    }
  }
}