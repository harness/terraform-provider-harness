resource "harness_platform_slo" "slo" {
  depends_on = [
    harness_platform_monitored_service.service_ref_environment_ref,
  ]
  org_id     = harness_platform_organization.terraform_org.id
  project_id = harness_platform_project.terraform_project.id
  identifier = "slo"
  request {
    name              = "slo"
    description       = "description"
    tags              = ["foo:bar", "bar:foo"]
    user_journey_refs = ["one", "two"]
    slo_target {
      type                  = "Calender"
      slo_target_percentage = 10
      spec                  = jsonencode({
        type = "Monthly"
        spec = {
          dayOfMonth = 5
        }
      })
    }
    type = "Simple"
    spec = jsonencode({
      monitoredServiceRef       = harness_platform_monitored_service.service_ref_environment_ref.id
      healthSourceRef           = "prometheus"
      serviceLevelIndicatorType = "Availability"
      serviceLevelIndicators    = [
        {
          name       = "name"
          identifier = "slo"
          type       = "Availability"
          spec       = {
            type = "Threshold"
            spec = {
              metric1        = "prometheus_metric"
              thresholdValue = 10
              thresholdType  = ">"
            }
          }
          sliMissingDataType = "Good"
        }
      ]
    })
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref"
      enabled               = true
    }
  }
}

