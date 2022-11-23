package trigger

import (
	"context"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTrigger() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness trigger.",

		ReadContext: dataSourceTriggerRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the trigger.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the trigger.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"app_id": {
				Description: "The id of the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The trigger description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"condition": {
				Description: "The condition that will execute the Trigger: On new artifact, On pipeline completion, On Cron schedule, On webhook, On New Manifest.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trigger_condition_type": {
							Description: "Trigger condition.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"on_webhook": {
							Description: "On webhook.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"webhook_details": {
										Description: "Webhook details.",
										Type:        schema.TypeSet,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"header": {
													Description: "Header.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"method": {
													Description: "Method.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"payload": {
													Description: "Payload.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"webhook_url": {
													Description: "Webhook URL.",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"webhook_token": {
													Description: "Webhook token.",
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTriggerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}

	var trigger *graphql.Trigger
	var err error

	if id, ok := d.GetOk("id"); ok {
		trigger, err = c.TriggerClient.GetTriggerById(id.(string))
	} else if name, ok := d.GetOk("name"); ok {
		trigger, err = c.TriggerClient.GetTriggerByName(name.(string), d.Get("app_id").(string))
	} else {
		return diag.Errorf("Must specify either `id` or `name`.")
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(trigger.Id)
	d.Set("name", trigger.Name)
	d.Set("description", trigger.Description)
	d.Set("condition", []map[string]interface{}{
		{
			"trigger_condition_type": trigger.Condition.TriggerConditionType,
			"on_webhook": []map[string]interface{}{
				{
					"webhook_details": []map[string]interface{}{
						{
							"payload":       trigger.Condition.WebhookDetails.Payload,
							"header":        trigger.Condition.WebhookDetails.Header,
							"method":        trigger.Condition.WebhookDetails.Method,
							"webhook_url":   trigger.Condition.WebhookDetails.WebhookUrl,
							"webhook_token": getWebhookToken(trigger.Condition.WebhookDetails.WebhookUrl),
						},
					},
				},
			},
		},
	})

	return nil
}

func getWebhookToken(webhookurl string) string {
	first := strings.LastIndex(webhookurl, "webhooks/")
	last := strings.LastIndex(webhookurl, "?accountId")
	result := ""
	for c, ch := range webhookurl {
		if c > first+8 && c < last {
			result = result + string(ch)
		}
	}
	return result
}
