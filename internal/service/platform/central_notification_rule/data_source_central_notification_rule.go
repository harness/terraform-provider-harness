package central_notification_rule

import (
	"context"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"net/http"
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

		ReadContext: dataCentralNotificationRuleRead,

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
										Elem:     &schema.Schema{Type: schema.TypeString},
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

func dataCentralNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org").(string)
	projectID := d.Get("project").(string)
	identifier := d.Get("identifier").(string)
	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	if orgID != "" && projectID != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRule(ctx, orgID, projectID,
			identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	} else if orgID != "" {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleOrg(ctx, orgID, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	} else {
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readCentralNotificationRule(accountID, d, resp)
	return nil
}
