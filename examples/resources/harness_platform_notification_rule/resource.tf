# Sample resource for SLO
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
        //        webhook_url = "http://myslackwebhookurl.com" // used for Slack
        //        integrationKey = "test-pd-integration-key" // used for PagerDuty
        //        msTeamKeys = ["ms-teams-key1", "ms-teams-key2"] // used for MsTeams
      })
    }
    type = "ServiceLevelObjective"
    conditions {
      type = "ErrorBudgetBurnRate"
      spec = jsonencode({
        threshold = 1
      })
    }
    conditions {
      type = "ErrorBudgetRemainingPercentage"
      spec = jsonencode({
        threshold = 30
      })
    }
    conditions {
      type = "ErrorBudgetRemainingMinutes"
      spec = jsonencode({
        threshold = 300
      })
    }
  }
}

# Sample resource for Monitored Service
resource "harness_platform_notification_rule" "example1" {
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
        //        webhook_url = "http://myslackwebhookurl.com" // used for Slack
        //        integrationKey = "test-pd-integration-key" // used for PagerDuty
        //        msTeamKeys = ["ms-teams-key1", "ms-teams-key2"] // used for MsTeams
      })
    }
    type = "MonitoredService"
    conditions {
      type = "ChangeImpact"
      spec = jsonencode({
        threshold        = 33
        period           = "30m"
        changeCategories = ["Deployment", "Infrastructure"]
      })
    }
    conditions {
      type = "HealthScore"
      spec = jsonencode({
        threshold = 33
        period    = "30m"
      })
    }
    conditions {
      type = "ChangeObserved"
      spec = jsonencode({
        changeCategories = ["Deployment", "Alert", "ChaosExperiment"]
      })
    }
    conditions {
      type = "DeploymentImpactReport"
      spec = jsonencode({})
    }
  }
}
