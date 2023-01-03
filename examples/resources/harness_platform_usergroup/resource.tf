resource "harness_platform_usergroup" "sso_type_saml" {
  identifier         = "identifier"
  name               = "name"
  org_id             = "org_id"
  project_id         = "project_id"
  linked_sso_id      = "linked_sso_id"
  externally_managed = false
  users              = ["user_id"]
  notification_configs {
    type              = "SLACK"
    slack_webhook_url = "https://google.com"
  }
  notification_configs {
    type                    = "EMAIL"
    group_email             = "email@email.com"
    send_email_to_all_users = true
  }
  notification_configs {
    type                        = "MSTEAMS"
    microsoft_teams_webhook_url = "https://google.com"
  }
  notification_configs {
    type           = "PAGERDUTY"
    pager_duty_key = "pagerDutyKey"
  }
  linked_sso_display_name = "linked_sso_display_name"
  sso_group_id            = "sso_group_name" // When sso linked type is saml sso_group_id is same as sso_group_name
  sso_group_name          = "sso_group_name"
  linked_sso_type         = "SAML"
  sso_linked              = true
}

resource "harness_platform_usergroup" "sso_type_ldap" {
  identifier         = "identifier"
  name               = "name"
  org_id             = "org_id"
  project_id         = "project_id"
  linked_sso_id      = "linked_sso_id"
  externally_managed = false
  users              = ["user_id"]
  notification_configs {
    type              = "SLACK"
    slack_webhook_url = "https://google.com"
  }
  notification_configs {
    type                    = "EMAIL"
    group_email             = "email@email.com"
    send_email_to_all_users = true
  }
  notification_configs {
    type                        = "MSTEAMS"
    microsoft_teams_webhook_url = "https://google.com"
  }
  notification_configs {
    type           = "PAGERDUTY"
    pager_duty_key = "pagerDutyKey"
  }
  linked_sso_display_name = "linked_sso_display_name"
  sso_group_id            = "sso_group_id"
  sso_group_name          = "sso_group_name"
  linked_sso_type         = "LDAP"
  sso_linked              = true
}

# Create user group by adding user emails
resource "harness_platform_usergroup" "example" {
  identifier         = "identifier"
  name               = "name"
  org_id             = "org_id"
  project_id         = "project_id"
  linked_sso_id      = "linked_sso_id"
  externally_managed = false
  user_emails        = ["user@email.com"]
  notification_configs {
    type              = "SLACK"
    slack_webhook_url = "https://google.com"
  }
  notification_configs {
    type                    = "EMAIL"
    group_email             = "email@email.com"
    send_email_to_all_users = true
  }
  notification_configs {
    type                        = "MSTEAMS"
    microsoft_teams_webhook_url = "https://google.com"
  }
  notification_configs {
    type           = "PAGERDUTY"
    pager_duty_key = "pagerDutyKey"
  }
  linked_sso_display_name = "linked_sso_display_name"
  sso_group_id            = "sso_group_name" // When sso linked type is saml sso_group_id is same as sso_group_name
  sso_group_name          = "sso_group_name"
  linked_sso_type         = "SAML"
  sso_linked              = true
}
