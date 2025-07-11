package central_notification_channel

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCentralNotificationChannelService() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a central notification channel in Harness.",

		ReadContext: resourceCentralNotificationChannelRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the notification channel.",
			},
			"org": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identifier of the organization the notification channel is scoped to.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identifier of the project the notification channel is scoped to.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the notification channel.",
			},
			"notification_channel_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of notification channel. One of: EMAIL, SLACK, PAGERDUTY, MSTeams, WEBHOOK, DATADOG.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Status of the notification channel. Possible values are ENABLED or DISABLED.",
			},
			"account": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account identifier associated with this notification channel.",
			},
			"created": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Timestamp when the notification channel was created.",
			},
			"last_modified": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Timestamp when the notification channel was last modified.",
			},
			"channel": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Configuration details of the notification channel.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of email addresses to notify.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"slack_webhook_urls": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Slack webhook URLs to send notifications to.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"webhook_urls": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of generic webhook URLs.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"pager_duty_integration_keys": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of PagerDuty integration keys.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"ms_team_keys": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Microsoft Teams integration keys.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"user_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of user groups to notify.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identifier": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Identifier of the user group.",
									},
								},
							},
						},
						"headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Custom HTTP headers to include in webhook requests.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Header key name.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Header value.",
									},
								},
							},
						},
						"datadog_urls": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Datadog webhook URLs.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"api_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "API key for the webhook or integration.",
						},
						"delegate_selectors": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of delegate selectors to use for sending notifications.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"execute_on_delegate": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to execute the notification logic on delegate.",
						},
					},
				},
			},
		},
	}
}
