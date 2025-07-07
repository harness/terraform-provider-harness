package central_notification_rule

import (
	"context"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func ResourceCentralNotificationRule() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating a Harness Notification Rule",

		CreateContext: resourceCentralNotificationRuleCreate,
		ReadContext:   resourceCentralNotificationRuleRead,
		UpdateContext: resourceCentralNotificationRuleUpdate,
		DeleteContext: resourceCentralNotificationRuleDelete,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"notification_channel_refs": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"notification_conditions": {
				Type:     schema.TypeList,
				Required: true,
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
										Type:     schema.TypeString,
										Required: true,
									},
									"notification_event": {
										Type:     schema.TypeString,
										Required: true,
									},
									"entity_identifiers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
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
}

func resourceCentralNotificationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	rule := buildCentralNotificationRule(d, accountID)
	resp, httpResp, err := c.NotificationRulesApi.CreateNotificationRule(ctx, orgID, projectID, &nextgen.NotificationRulesApiCreateNotificationRuleOpts{
		Body:           optional.NewInterface(rule),
		HarnessAccount: optional.NewString(accountID),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(resp.Identifier)
	return readCentralNotificationRule(d, resp)
}

func resourceCentralNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	identifier := d.Id()
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

	readCentralNotificationRule(d, resp)

	return nil
}

func resourceCentralNotificationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	identifier := d.Get("identifier").(string)

	rule := buildCentralNotificationRule(d, accountID)

	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	if orgID != "" && projectID != "" {
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRule(ctx, orgID, projectID,
			identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	} else if orgID != "" {
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRuleOrg(ctx, orgID, identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleOrgOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	} else {
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleAccountOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return readCentralNotificationRule(d, resp)
}

func resourceCentralNotificationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	identifier := d.Get("identifier").(string)

	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	if orgID != "" && projectID != "" {
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRule(ctx, orgID, projectID, identifier, &nextgen.NotificationRulesApiDeleteNotificationRuleOpts{
			HarnessAccount: optional.NewString(accountID),
		})
	} else if orgID != "" {
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRuleOrg(ctx, orgID, identifier,
			&nextgen.NotificationRulesApiDeleteNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	} else {
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiDeleteNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return readCentralNotificationRule(d, resp)
}

func expandStringList(raw interface{}) []string {
	if raw == nil {
		return nil
	}
	rawList := raw.([]interface{})
	strList := make([]string, len(rawList))
	for i, val := range rawList {
		strList[i] = val.(string)
	}
	return strList
}

func expandNotificationTemplateVariables(raw []interface{}) []nextgen.NotificationTemplateInputsDto {
	result := make([]nextgen.NotificationTemplateInputsDto, len(raw))
	for i, item := range raw {
		data := item.(map[string]interface{})
		result[i] = nextgen.NotificationTemplateInputsDto{
			Name:  data["name"].(string),
			Value: data["value"].(string),
			Type_: data["type"].(string),
		}
	}
	return result
}

func expandNotificationConditions(raw []interface{}) []nextgen.NotificationConditionDto {
	result := make([]nextgen.NotificationConditionDto, len(raw))

	for i, cond := range raw {
		condition := cond.(map[string]interface{})
		eventConfigsRaw := condition["notification_event_configs"].([]interface{})
		eventConfigs := make([]nextgen.NotificationEventConfigDto, len(eventConfigsRaw))

		for j, ec := range eventConfigsRaw {
			ecMap := ec.(map[string]interface{})
			eventConfigs[j] = nextgen.NotificationEventConfigDto{
				NotificationEntity: ecMap["notification_entity"].(string),
				NotificationEvent:  ecMap["notification_event"].(string),
				EntityIdentifiers:  expandStringList(ecMap["entity_identifiers"]),
				NotificationEventData: &nextgen.NotificationEventParamsDto{
					Type_: func() *nextgen.ResourceTypeEnum {
						val := nextgen.ResourceTypeEnum(ecMap["notification_event_data"].(map[string]interface{})["type"].(string))
						return &val
					}(),
				},
			}
		}

		result[i] = nextgen.NotificationConditionDto{
			ConditionName:            condition["condition_name"].(string),
			NotificationEventConfigs: eventConfigs,
		}
	}

	return result
}

func buildCentralNotificationRule(d *schema.ResourceData, accountID string) nextgen.NotificationRuleDto {
	rule := nextgen.NotificationRuleDto{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Org:        d.Get("org_id").(string),
		Project:    d.Get("project_id").(string),
		Status: func() *nextgen.Status {
			s := nextgen.Status(d.Get("status").(string))
			return &s
		}(),
		NotificationChannelRefs: expandStringList(d.Get("notification_channel_refs")),
		NotificationConditions:  expandNotificationConditions(d.Get("notification_conditions").([]interface{})),
	}

	if v, ok := d.GetOk("custom_notification_template_ref"); ok {
		ref := v.([]interface{})[0].(map[string]interface{})
		templateRef := nextgen.CustomNotificationTemplateDto{
			TemplateRef:  ref["template_ref"].(string),
			VersionLabel: ref["version_label"].(string),
			Variables:    expandNotificationTemplateVariables(ref["variables"].([]interface{})),
		}
		rule.CustomNotificationTemplateRef = &templateRef
	}

	return rule
}

func readCentralNotificationRule(d *schema.ResourceData, notificationRuleDto nextgen.NotificationRuleDto) diag.Diagnostics {
	// Implement read logic as needed
	d.SetId(notificationRuleDto.Identifier)
	d.Set("org_id", notificationRuleDto.Org)
	d.Set("project_id", notificationRuleDto.Project)
	d.Set("identifier", notificationRuleDto.Identifier)
	d.Set("name", notificationRuleDto.Name)
	d.Set("status", notificationRuleDto.Status)
	d.Set("notification_channel_refs", notificationRuleDto.NotificationChannelRefs)
	d.Set("custom_notification_template_ref", notificationRuleDto.CustomNotificationTemplateRef)
	return nil
}
