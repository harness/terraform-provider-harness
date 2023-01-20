resource "harness_platform_monitored_service" "service_ref_environment_ref" {
  org_id     = "terraform_org"
  project_id = "terraform_project"
  identifier = "service_ref_environment_ref"
  request {
    name = "service_ref_environment_ref"
    type = "Application"
    description = "new_description_new"
    service_ref = "service_ref"
    environment_ref = "environment_ref"
    tags = ["foo:bar", "bar:foo"]
    health_sources {
      name = "prometheus"
      identifier = "prometheus"
      type = "Prometheus"
      spec = jsonencode({
        connectorRef = "connectorRef"
        feature = "feature"
        metricDefinitions = [
          {
            identifier   = "prometheus_metric"
            metricName = "Prometheus Metric"
            riskProfile = {
              category =  "Errors"
            }
            sli =  {
              enabled =  true
            },
            query = "sum(abc{identifier=\"slo-ratiobased-unsuccessfulCalls-datapattern\"})",
            groupName = "t2",
            isManualQuery = true
          }
        ]
      })
    }
    change_sources {
      name = "csName1"
      identifier = "harness_cd_next_gen"
      type = "HarnessCDNextGen"
      enabled = true
      spec = jsonencode({
      })
      category = "Deployment"
    }

    template_ref = "template_ref"
    version_label = "version_label"
    enabled = true
  }
}


