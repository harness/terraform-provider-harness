---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_usergroup Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness User Group. Linking SSO providers with User Groups:
  
  	The following fields need to be populated for LDAP SSO Providers:
  	
  	- linked_sso_id
  	
  	- linked_sso_display_name
  	
  	- sso_group_id
  	
  	- sso_group_name
  	
  	- linked_sso_type
  	
  	- sso_linked
  	
  	The following fields need to be populated for SAML SSO Providers:
  	
  	- linked_sso_id
  	
  	- linked_sso_display_name
  	
  	- sso_group_name
  	
  	- sso_group_id // same as sso_group_name
  	
  	- linked_sso_type
  	
  	- sso_linked
---

# harness_platform_usergroup (Resource)

Resource for creating a Harness User Group. Linking SSO providers with User Groups:

		The following fields need to be populated for LDAP SSO Providers:
		
		- linked_sso_id
		
		- linked_sso_display_name
		
		- sso_group_id
		
		- sso_group_name
		
		- linked_sso_type
		
		- sso_linked
		
		The following fields need to be populated for SAML SSO Providers:
		
		- linked_sso_id
		
		- linked_sso_display_name
		
		- sso_group_name
		
		- sso_group_id // same as sso_group_name
		
		- linked_sso_type
		
		- sso_linked

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `description` (String) Description of the resource.
- `externally_managed` (Boolean) Whether the user group is externally managed.
- `linked_sso_display_name` (String) Name of the linked SSO.
- `linked_sso_id` (String) The SSO account ID that the user group is linked to.
- `linked_sso_type` (String) Type of linked SSO.
- `notification_configs` (Block List) List of notification settings. (see [below for nested schema](#nestedblock--notification_configs))
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `sso_group_id` (String) Identifier of the userGroup in SSO.
- `sso_group_name` (String) Name of the SSO userGroup.
- `sso_linked` (Boolean) Whether sso is linked or not.
- `tags` (Set of String) Tags to associate with the resource.
- `user_emails` (List of String) List of user emails in the UserGroup. Either provide list of users or list of user emails.
- `users` (List of String) List of users in the UserGroup. Either provide list of users or list of user emails. (Should be null for SSO managed)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--notification_configs"></a>
### Nested Schema for `notification_configs`

Optional:

- `group_email` (String) Group email.
- `microsoft_teams_webhook_url` (String) Url of Microsoft teams webhook.
- `pager_duty_key` (String) Pager duty key.
- `send_email_to_all_users` (Boolean) Send email to all the group members.
- `slack_webhook_url` (String) Url of slack webhook.
- `type` (String) Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level user group
terraform import harness_platform_usergroup.example <usergroup_id>

# Import org level user group
terraform import harness_platform_usergroup.example <ord_id>/<usergroup_id>

# Import project level user group
terraform import harness_platform_usergroup.example <org_id>/<project_id>/<usergroup_id>
```
