---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_central_notification_channel Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a central notification channel in Harness.
---

# harness_platform_central_notification_channel (Data Source)

Data source for retrieving a central notification channel in Harness.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the notification channel.

### Optional

- `account` (String) Account identifier associated with this notification channel.
- `channel` (Block List) Configuration details of the notification channel. (see [below for nested schema](#nestedblock--channel))
- `created` (Number) Timestamp when the notification channel was created.
- `last_modified` (Number) Timestamp when the notification channel was last modified.
- `name` (String) Name of the notification channel.
- `notification_channel_type` (String) Type of notification channel. One of: EMAIL, SLACK, PAGERDUTY, MSTeams, WEBHOOK, DATADOG.
- `org` (String) Identifier of the organization the notification channel is scoped to.
- `project` (String) Identifier of the project the notification channel is scoped to.
- `status` (String) Status of the notification channel. Possible values are ENABLED or DISABLED.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--channel"></a>
### Nested Schema for `channel`

Optional:

- `api_key` (String) API key for the webhook or integration.
- `datadog_urls` (List of String) List of Datadog webhook URLs.
- `delegate_selectors` (List of String) List of delegate selectors to use for sending notifications.
- `email_ids` (List of String) List of email addresses to notify.
- `execute_on_delegate` (Boolean) Whether to execute the notification logic on delegate.
- `headers` (Block List) Custom HTTP headers to include in webhook requests. (see [below for nested schema](#nestedblock--channel--headers))
- `ms_team_keys` (List of String) List of Microsoft Teams integration keys.
- `pager_duty_integration_keys` (List of String) List of PagerDuty integration keys.
- `slack_webhook_urls` (List of String) List of Slack webhook URLs to send notifications to.
- `user_groups` (Block List) List of user groups to notify. (see [below for nested schema](#nestedblock--channel--user_groups))
- `webhook_urls` (List of String) List of generic webhook URLs.

<a id="nestedblock--channel--headers"></a>
### Nested Schema for `channel.headers`

Required:

- `key` (String) Header key name.
- `value` (String) Header value.


<a id="nestedblock--channel--user_groups"></a>
### Nested Schema for `channel.user_groups`

Optional:

- `identifier` (String) Identifier of the user group.
