resource "harness_platform_monitored_service" "example" {
  org_id = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name = "name"
    type = "Application"
    description = "description"
    service_ref = "service_ref"
    environment_ref = "environment_ref"
    tags = ["foo:bar", "bar:foo"]
    health_sources {
      name = "name"
      identifier = "identifier"
      type = "ElasticSearch"
      spec = jsonencode({
        connectorRef = "connectorRef"
        feature = "feature"
        queries = [
          {
            name   = "name"
            query = "query"
            index = "index"
            serviceInstanceIdentifier = "serviceInstanceIdentifier"
            timeStampIdentifier = "timeStampIdentifier"
            timeStampFormat = "timeStampFormat"
            messageIdentifier = "messageIdentifier"
          },
          {
            name   = "name2"
            query = "query2"
            index = "index2"
            serviceInstanceIdentifier = "serviceInstanceIdentifier2"
            timeStampIdentifier = "timeStampIdentifier2"
            timeStampFormat = "timeStampFormat2"
            messageIdentifier = "messageIdentifier2"
          }
        ]})
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
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref"
      enabled = true
    }
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref1"
      enabled = false
    }
    template_ref = "template_ref"
    version_label = "version_label"
    enabled = true
  }
}