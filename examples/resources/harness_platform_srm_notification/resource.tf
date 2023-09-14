#Sample template for Elastic Search Log Health Source
resource "harness_platform_srm_notification" "example" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "MonitoredService"
    conditions {
      type       = "ERROR_BUDGET_REMAINING_PERCENTAGE"
      spec = jsonencode({
        threshold = 100
      })
    }
    notificationMethod {
      type       = "Slack"
      spec = jsonencode({
        userGroups : ["userGroups1", "userGroups2"]
        webhookUrl : "https://expamle.slack.com/"
      })
    }
  }
}