resource "harness_platform_slo" "example" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name              = "name"
    description       = "description"
    tags              = ["foo:bar", "bar:foo"]
    user_journey_refs = ["one", "two"]
    slo_target {
      type                  = "Calender"
      slo_target_percentage = 10
      spec = jsonencode({
        type = "Monthly"
        spec = {
          dayOfMonth = 5
        }
      })
    }
    type = "Simple"
    spec = jsonencode({
      monitoredServiceRef       = "monitoredServiceRef"
      serviceLevelIndicatorType = "Availability"
      healthSourceRef           = "healthSourceRef" 
      serviceLevelIndicators = [
        {
          name       = "name"
          identifier = "identifier"
          type       = "Window"
          spec = {
            type = "Threshold"
            spec = {
              metric1        = "metric1"
              thresholdValue = 10
              thresholdType  = ">"
            }
            sliMissingDataType = "Good"
          }
        }
      ]
    })
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref"
      enabled               = true
    }
  }
}
