resource "harness_platform_notification_rule" "example" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name = "name"
    notification_method {
      type = "Slack"
      spec = jsonencode({
        webhook_url = "http://myslackwebhookurl.com"
        user_groups = ["account.test"]
      })
    }
    type = "ServiceLevelObjective"
    conditions {
      type = "ErrorBudgetRemainingPercentage"
      spec = jsonencode({
        threshold = 30
      })
    }
  }
}
