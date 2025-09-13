package central_notification_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func notificationEntityEnum() []string {
	return []string{
		"PIPELINE",
		"DELEGATE",
		"CONNECTOR",
		"CHAOS_EXPERIMENT",
		"SERVICE_LEVEL_OBJECTIVE",
		"STO_EXEMPTION",
	}
}

func DataSourceCentralNotificationRuleService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Notification Rule.",

		ReadContext: resourceCentralNotificationRuleRead,

		Schema: map[string]*schema.Schema{
			"org": {
				Description: "Identifier of the organization in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "Identifier of the project in which the Notification Rule is configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier of the Notification Rule.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account identifier associated with this notification channel.",
			},
			"created": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification rule was created.",
			},
			"last_modified": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Timestamp when the notification rule was last modified.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"notification_channel_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"notification_conditions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"condition_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"notification_event_configs": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification_entity": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(notificationEntityEnum(), false),
									},
									"notification_event": {
										Type:     schema.TypeString,
										Required: true,
									},
									"notification_event_data": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"scope_identifiers": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"delegate_group_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"frequency": {
													Type:     schema.TypeMap,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"chaos_experiment_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"error_budget_remaining_percentage": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"error_budget_remaining_minutes": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"error_budget_burn_rate_percentage": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"error_budget_burn_rate_lookback_duration": {
													Type:     schema.TypeInt,
													Optional: true,
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
			"custom_notification_template_ref": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_ref": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version_label": {
							Type:     schema.TypeString,
							Required: true,
						},
						"variables": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return resource
}
