package central_notification_rule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePipelineCentralNotificationRule() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating a Harness Notification Rule for Pipeline",
		CreateContext: resourcePipelineCentralNotificationRuleCreate,
		ReadContext:   resourcePipelineCentralNotificationRuleRead,
		UpdateContext: resourcePipelineCentralNotificationRuleUpdate,
		DeleteContext: resourcePipelineCentralNotificationRuleDelete,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account identifier associated with this notification channel.",
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
										Type:     schema.TypeList,
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
}

func resourcePipelineCentralNotificationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	accountID := c.AccountId
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	scope := getScope(d)
	rule := buildPipelineCentralNotificationRule(d)
	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationRulesApi.CreateNotificationRule(ctx, scope.org, scope.project, &nextgen.NotificationRulesApiCreateNotificationRuleOpts{
			Body:           optional.NewInterface(rule),
			HarnessAccount: optional.NewString(accountID),
		})

	case Org:
		resp, httpResp, err = c.NotificationRulesApi.CreateNotificationRuleOrg(ctx, scope.org, &nextgen.NotificationRulesApiCreateNotificationRuleOrgOpts{
			Body:           optional.NewInterface(rule),
			HarnessAccount: optional.NewString(accountID),
		})
	default:
		resp, httpResp, err = c.NotificationRulesApi.CreateNotificationRuleAccount(ctx, &nextgen.NotificationRulesApiCreateNotificationRuleAccountOpts{
			Body:           optional.NewInterface(rule),
			HarnessAccount: optional.NewString(accountID),
		})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(resp.Identifier)
	return readPipelineCentralNotificationRule(accountID, d, resp)
}

func resourcePipelineCentralNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)
	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRule(ctx, scope.org, scope.project,
			identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleOrg(ctx, scope.org, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		resp, httpResp, err = c.NotificationRulesApi.GetNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiGetNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readPipelineCentralNotificationRule(accountID, d, resp)

	return nil
}

func resourcePipelineCentralNotificationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)

	rule := buildPipelineCentralNotificationRule(d)

	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRule(ctx, scope.org, scope.project,
			identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	case Org:
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRuleOrg(ctx, scope.org, identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleOrgOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		resp, httpResp, err = c.NotificationRulesApi.UpdateNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiUpdateNotificationRuleAccountOpts{
				Body:           optional.NewInterface(rule),
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return readPipelineCentralNotificationRule(accountID, d, resp)
}

func resourcePipelineCentralNotificationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)

	var resp nextgen.NotificationRuleDto
	var httpResp *http.Response
	var err error
	switch scope.scope {
	case Project:
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRule(ctx, scope.org, scope.project, identifier, &nextgen.NotificationRulesApiDeleteNotificationRuleOpts{
			HarnessAccount: optional.NewString(accountID),
		})
	case Org:
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRuleOrg(ctx, scope.org, identifier,
			&nextgen.NotificationRulesApiDeleteNotificationRuleOrgOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	default:
		httpResp, err = c.NotificationRulesApi.DeleteNotificationRuleAccount(ctx, identifier,
			&nextgen.NotificationRulesApiDeleteNotificationRuleAccountOpts{
				HarnessAccount: optional.NewString(accountID),
			})
	}
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return readPipelineCentralNotificationRule(accountID, d, resp)
}

func expandPipelineStringList(raw interface{}) []string {
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

func expandPipelineNotificationTemplateVariables(raw []interface{}) []nextgen.NotificationTemplateInputsDto {
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

func expandPipelineNotificationConditions(raw []interface{}) []nextgen.NotificationConditionDto {
	result := make([]nextgen.NotificationConditionDto, len(raw))

	for i, cond := range raw {
		condition := cond.(map[string]interface{})
		eventConfigsRaw := condition["notification_event_configs"].([]interface{})
		eventConfigs := make([]nextgen.NotificationEventConfigDto, len(eventConfigsRaw))

		for j, ec := range eventConfigsRaw {
			ecMap := ec.(map[string]interface{})
			evtCfg := nextgen.NotificationEventConfigDto{
				NotificationEntity: ecMap["notification_entity"].(string),
				NotificationEvent:  ecMap["notification_event"].(string),
				EntityIdentifiers:  expandPipelineStringList(ecMap["entity_identifiers"]),
			}
			if rawEventData, ok := ecMap["notification_event_data"]; ok && rawEventData != nil {
				if eventDataList, ok := rawEventData.([]interface{}); ok && len(eventDataList) > 0 {
					eventDataMap, _ := eventDataList[0].(map[string]interface{})
					if typeStr, ok := eventDataMap["type"].(string); ok && typeStr != "" {
						t := nextgen.ResourceTypeEnum(typeStr)
						switch t {
						case nextgen.PIPELINE_ResourceTypeEnum:
							evtCfg.PipelineEventNotificationParamsDto = &nextgen.PipelineEventNotificationParamsDto{
								Type_:            &t,
								ScopeIdentifiers: expandPipelineStringList(eventDataMap["scope_identifiers"]),
							}
						default:
							panic(fmt.Sprintf("unsupported resource type: %s", t))
						}
					}
				}
			}
			eventConfigs[j] = evtCfg
		}

		result[i] = nextgen.NotificationConditionDto{
			ConditionName:            condition["condition_name"].(string),
			NotificationEventConfigs: eventConfigs,
		}
	}

	return result
}

func buildPipelineCentralNotificationRule(d *schema.ResourceData) nextgen.NotificationRuleDto {
	rule := nextgen.NotificationRuleDto{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		Org:        d.Get("org").(string),
		Project:    d.Get("project").(string),
		Status: func() *nextgen.Status {
			s := nextgen.Status(d.Get("status").(string))
			return &s
		}(),
		NotificationChannelRefs: expandPipelineStringList(d.Get("notification_channel_refs")),
		NotificationConditions:  expandPipelineNotificationConditions(d.Get("notification_conditions").([]interface{})),
	}

	if v, ok := d.GetOk("custom_notification_template_ref"); ok {
		ref := v.([]interface{})[0].(map[string]interface{})
		templateRef := nextgen.CustomNotificationTemplateDto{
			TemplateRef:  ref["template_ref"].(string),
			VersionLabel: ref["version_label"].(string),
			Variables:    expandPipelineNotificationTemplateVariables(ref["variables"].([]interface{})),
		}
		rule.CustomNotificationTemplateRef = &templateRef
	}

	return rule
}

func readPipelineCentralNotificationRule(accountIdentifier string, d *schema.ResourceData, notificationRuleDto nextgen.NotificationRuleDto) diag.Diagnostics {
	d.SetId(notificationRuleDto.Identifier)
	if notificationRuleDto.Org != "" {
		d.Set("org", notificationRuleDto.Org)
	}
	if notificationRuleDto.Project != "" {
		d.Set("project", notificationRuleDto.Project)
	}
	d.Set("account", accountIdentifier)
	d.Set("identifier", notificationRuleDto.Identifier)
	d.Set("name", notificationRuleDto.Name)
	d.Set("status", notificationRuleDto.Status)
	d.Set("notification_channel_refs", notificationRuleDto.NotificationChannelRefs)
	d.Set("created", notificationRuleDto.Created)
	d.Set("last_modified", notificationRuleDto.LastModified)
	// Convert notification_conditions
	var conditions []map[string]interface{}
	for _, cond := range notificationRuleDto.NotificationConditions {
		var eventConfigs []map[string]interface{}
		for _, cfg := range cond.NotificationEventConfigs {
			eventDataList := []interface{}{}
			// Safely read cfg.NotificationEventData.Type_
			if cfg.PipelineEventNotificationParamsDto != nil && cfg.PipelineEventNotificationParamsDto.Type_ != nil {
				eventData := make(map[string]interface{})

				eventData["type"] = string(*cfg.PipelineEventNotificationParamsDto.Type_)
				eventData["scope_identifiers"] = cfg.PipelineEventNotificationParamsDto.ScopeIdentifiers
				eventDataList = []interface{}{eventData}
			} else {
				panic(fmt.Sprintf("unsupported notification event data in read: %+v", cfg))
			}

			eventConfigs = append(eventConfigs, map[string]interface{}{
				"notification_entity":     cfg.NotificationEntity,
				"notification_event":      cfg.NotificationEvent,
				"entity_identifiers":      cfg.EntityIdentifiers,
				"notification_event_data": eventDataList,
			})
		}

		conditions = append(conditions, map[string]interface{}{
			"condition_name":             cond.ConditionName,
			"notification_event_configs": eventConfigs,
		})
	}
	d.Set("notification_conditions", conditions)

	if notificationRuleDto.CustomNotificationTemplateRef != nil {
		custom := map[string]interface{}{
			"template_ref":  notificationRuleDto.CustomNotificationTemplateRef.TemplateRef,
			"version_label": notificationRuleDto.CustomNotificationTemplateRef.VersionLabel,
		}

		if len(notificationRuleDto.CustomNotificationTemplateRef.Variables) > 0 {
			var vars []map[string]interface{}
			for _, v := range notificationRuleDto.CustomNotificationTemplateRef.Variables {
				vars = append(vars, map[string]interface{}{
					"name":  v.Name,
					"value": v.Value,
					"type":  v.Type_,
				})
			}
			custom["variables"] = vars
		}

		d.Set("custom_notification_template_ref", []interface{}{custom})
	}

	return nil
}
