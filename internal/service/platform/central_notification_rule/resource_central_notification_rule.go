package central_notification_rule

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Optional: true,
			},
			"org": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
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
												"delegate_group_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"frequency": {
													Type:     schema.TypeList,
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
}

func resourceCentralNotificationRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	accountID := c.AccountId
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	scope := getScope(d)
	rule := buildCentralNotificationRule(d, accountID)
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
	return readCentralNotificationRule(accountID, d, resp)
}

func resourceCentralNotificationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	readCentralNotificationRule(accountID, d, resp)

	return nil
}

func resourceCentralNotificationRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountID := c.AccountId
	identifier := d.Get("identifier").(string)
	scope := getScope(d)

	rule := buildCentralNotificationRule(d, accountID)

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
	return readCentralNotificationRule(accountID, d, resp)
}

func resourceCentralNotificationRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return readCentralNotificationRule(accountID, d, resp)
}

func expandStringList(raw interface{}) []string {
	if raw == nil {
		return []string{}
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
			}

			// If notification_event_data provided, marshal raw and also set typed DTO
			if eventDataList, ok := ecMap["notification_event_data"].([]interface{}); ok && len(eventDataList) > 0 {
				if eventData, ok := eventDataList[0].(map[string]interface{}); ok {
					// Marshal raw for debugging/round-trip
					if jsonBytes, err := json.Marshal(eventData); err == nil {
						eventConfigs[j].NotificationEventData = jsonBytes
					}

					// After setting eventConfigs[j].NotificationEventData = jsonBytes
					if t, ok := eventData["type"].(string); ok {
						switch nextgen.ResourceTypeEnum(t) {
						case nextgen.PIPELINE_ResourceTypeEnum:
							rt := nextgen.PIPELINE_ResourceTypeEnum
							eventConfigs[j].PipelineEventNotificationParamsDto = &nextgen.PipelineEventNotificationParamsDto{
								Type_:            &rt,
								ScopeIdentifiers: expandStringList(eventData["scope_identifiers"]),
							}

						case nextgen.DELEGATE_ResourceTypeEnum:
							rt := nextgen.DELEGATE_ResourceTypeEnum
							dto := &nextgen.DelegateEventNotificationParamsDto{
								Type_:            &rt,
								DelegateGroupIds: expandStringList(eventData["delegate_group_ids"]),
							}
							// frequency { type,value } -> FrequencyDto { Key,Value }
							if freqRaw, ok := eventData["frequency"].([]interface{}); ok && len(freqRaw) > 0 {
								if f, ok := freqRaw[0].(map[string]interface{}); ok {
									key, _ := f["type"].(string)
									val, _ := f["value"].(string)
									dto.Frequency = &nextgen.FrequencyDto{Key: key, Value: val}
								}
							}
							eventConfigs[j].DelegateEventNotificationParamsDto = dto

						case nextgen.CHAOS_EXPERIMENT_ResourceTypeEnum:
							rt := nextgen.CHAOS_EXPERIMENT_ResourceTypeEnum
							eventConfigs[j].ChaosExperimentEventNotificationParamsDto = &nextgen.ChaosExperimentEventNotificationParamsDto{
								Type_:              &rt,
								ChaosExperimentIds: expandStringList(eventData["chaos_experiment_ids"]),
							}

						case nextgen.SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum:
							rt := nextgen.SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum
							dto := &nextgen.SloEventNotificationParamsDto{
								Type_:                               &rt,
								ErrorBudgetRemainingPercentage:      0,
								ErrorBudgetRemainingMinutes:         0,
								ErrorBudgetBurnRatePercentage:       0,
								ErrorBudgetBurnRateLookbackDuration: 0,
							}
							if v, ok := eventData["error_budget_remaining_percentage"].(float64); ok {
								dto.ErrorBudgetRemainingPercentage = float32(v)
							}
							if v, ok := eventData["error_budget_remaining_minutes"].(float64); ok {
								dto.ErrorBudgetRemainingMinutes = int32(v)
							}
							if v, ok := eventData["error_budget_burn_rate_percentage"].(float64); ok {
								dto.ErrorBudgetBurnRatePercentage = float32(v)
							}
							if v, ok := eventData["error_budget_burn_rate_lookback_duration"].(float64); ok {
								dto.ErrorBudgetBurnRateLookbackDuration = int32(v)
							}
							eventConfigs[j].SloEventNotificationParamsDto = dto

						case nextgen.STO_EXEMPTION_ResourceTypeEnum:
							rt := nextgen.STO_EXEMPTION_ResourceTypeEnum
							eventConfigs[j].StoExemptionEventNotificationParamsDto = &nextgen.StoExemptionEventNotificationParamsDto{
								Type_:            &rt,
								ScopeIdentifiers: expandStringList(eventData["scope_identifiers"]),
							}

						default:
							// leave typed DTOs nil; serializer will fallback to null
						}
					}
				}
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
		Org:        d.Get("org").(string),
		Project:    d.Get("project").(string),
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

func readCentralNotificationRule(accountIdentifier string, d *schema.ResourceData, notificationRuleDto nextgen.NotificationRuleDto) diag.Diagnostics {
	// Implement read logic as needed
	d.SetId(notificationRuleDto.Identifier)
	d.Set("org", notificationRuleDto.Org)
	d.Set("account", accountIdentifier)
	d.Set("project", notificationRuleDto.Project)
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
			eventConfig := map[string]interface{}{
				"notification_entity": cfg.NotificationEntity,
				"notification_event":  cfg.NotificationEvent,
			}

			// Always set entity_identifiers, even if empty array
			eventConfig["entity_identifiers"] = cfg.EntityIdentifiers

			// Handle notification event data if present
			if len(cfg.NotificationEventData) > 0 {
				var eventData map[string]interface{}
				if err := json.Unmarshal(cfg.NotificationEventData, &eventData); err == nil && eventData != nil {
					// Ensure scope_identifiers is always present as an array, even if empty
					if _, exists := eventData["scope_identifiers"]; !exists {
						eventData["scope_identifiers"] = []interface{}{}
					}
					eventConfig["notification_event_data"] = []interface{}{eventData}
				} else {
					// If unmarshaling failed or eventData is nil, set empty structure
					eventConfig["notification_event_data"] = []interface{}{}
				}
			} else {
				// Set empty notification_event_data structure if not present
				eventConfig["notification_event_data"] = []interface{}{}
			}
			// Append the eventConfig to eventConfigs slice
			eventConfigs = append(eventConfigs, eventConfig)
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

type Scope struct {
	org     string
	project string
	scope   ScopeLevel
}

type ScopeLevel string

const (
	Account ScopeLevel = "account"
	Org     ScopeLevel = "org"
	Project ScopeLevel = "project"
)

func getScope(d *schema.ResourceData) *Scope {
	org := ""
	project := ""

	if attr, ok := d.GetOk("org"); ok {
		org = (attr.(string))
	}

	if attr, ok := d.GetOk("project"); ok {
		project = (attr.(string))
	}

	var scope ScopeLevel
	if org == "" {
		scope = Account
	} else if project == "" {
		scope = Org
	} else {
		scope = Project
	}

	return &Scope{
		org:     org,
		project: project,
		scope:   scope,
	}
}
