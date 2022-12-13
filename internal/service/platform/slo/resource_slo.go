package slo

import (
	"context"
	"encoding/json"
	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func ResourceSloService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a SLO.",

		CreateContext: resourceSloCreate,
		ReadContext:   resourceSloRead,
		UpdateContext: resourceSloUpdate,
		DeleteContext: resourceSloDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account Identifier for / of the SLO.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier for / of the SLO.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier of the SLO.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier for / of the SLO.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"request": {
				Description: "Request for creating / updating SLO.",
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name of the SLO.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"description": {
							Description: "Description for the SLO.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tags": {
							Description: "Tags for the SLO.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_journey_refs": {
							Description: "User journey reference list for the SLO.",
							Type:        schema.TypeList,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"slo_target": {
							Description: "SLO Target specification.",
							Type:        schema.TypeSet,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of the SLO target.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"slo_target_percentage": {
										Description: "Target percentage for the SLO.",
										Type:        schema.TypeFloat,
										Required:    true,
									},
									"spec": {
										Description: "Specification of the SLO Target.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"type": {
							Description: "Type of the SLO.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"spec": {
							Description: "Specification of the SLO.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"notification_rule_refs": {
							Description: "Notification rule references for the SLO.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification_rule_ref": {
										Description: "Notification rule reference for the SLO.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"enabled": {
										Description: "Enable / Disable notification rule reference for the SLO.",
										Type:        schema.TypeBool,
										Required:    true,
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

func resourceSloCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	createSloRequest := buildSloRequest(d, identifier)
	respCreate, httpRespCreate, errCreate := c.SloApi.SaveSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier,
		&nextgen.SloApiSaveSLODataNgOpts{
			Body: optional.NewInterface(createSloRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readSlo(d, &respCreate.Resource, accountIdentifier)
	return nil
}

func resourceSloRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	identifier := d.Get("identifier").(string)
	if attr, ok := d.GetOk("account_id"); ok {
		accountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	resp, httpResp, err := c.SloApi.GetServiceLevelObjectiveNg(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readSlo(d, &resp.Resource, accountIdentifier)
	return nil
}

func resourceSloUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier, identifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	updateSloRequest := buildSloRequest(d, identifier)
	respCreate, httpRespCreate, errCreate := c.SloApi.UpdateSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier,
		&nextgen.SloApiUpdateSLODataNgOpts{
			Body: optional.NewInterface(updateSloRequest),
		})

	if errCreate != nil {
		return helpers.HandleApiError(errCreate, d, httpRespCreate)
	}

	readSlo(d, &respCreate.Resource, accountIdentifier)
	return nil
}

func resourceSloDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	identifier := d.Get("identifier").(string)
	if attr, ok := d.GetOk("account_id"); ok {
		accountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	_, httpResp, err := c.SloApi.DeleteSLODataNg(ctx, accountIdentifier, orgIdentifier, projectIdentifier, identifier)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildSloRequest(d *schema.ResourceData, identifier string) *nextgen.ServiceLevelObjectiveV2Dto {
	serviceLevelObjectiveV2Dto := &nextgen.ServiceLevelObjectiveV2Dto{}

	if attr, ok := d.GetOk("org_id"); ok {
		serviceLevelObjectiveV2Dto.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		serviceLevelObjectiveV2Dto.ProjectIdentifier = attr.(string)
	}

	serviceLevelObjectiveV2Dto.Identifier = identifier

	if attr, ok := d.GetOk("request"); ok {
		request := attr.([]interface{})[0].(map[string]interface{})

		serviceLevelObjectiveV2Dto.Name = request["name"].(string)
		serviceLevelObjectiveV2Dto.Description = request["description"].(string)

		tags := map[string]string{}
		for _, t := range request["tags"].(*schema.Set).List() {
			tagStr := t.(string)
			parts := strings.Split(tagStr, ":")
			tags[parts[0]] = parts[1]
		}
		serviceLevelObjectiveV2Dto.Tags = tags

		userJourneyRefsReq := request["user_journey_refs"].([]interface{})
		userJourneyRefsList := make([]string, len(userJourneyRefsReq))
		for i, userJourneyRef := range userJourneyRefsReq {
			userJourneyRefsList[i] = userJourneyRef.(string)
		}
		serviceLevelObjectiveV2Dto.UserJourneyRefs = userJourneyRefsList

		sloTarget := request["slo_target"].(*schema.Set).List()[0].(map[string]interface{})
		sloTargetType := sloTarget["type"].(string)
		sloTargetPercentage := sloTarget["slo_target_percentage"].(float64)
		sloTargetSpec := sloTarget["spec"].(string)
		if sloTargetType == "Rolling" {
			data := nextgen.RollingSloTargetSpec{}
			json.Unmarshal([]byte(sloTargetSpec), &data)

			sloTargetDto := nextgen.SloTargetDto{
				Type_: nextgen.SLOTargetType(sloTargetType),
				SloTargetPercentage: sloTargetPercentage,
				Rolling: &data,
			}

			serviceLevelObjectiveV2Dto.SloTarget = &sloTargetDto
		}
		if sloTargetType == "Calender" {
			data := nextgen.CalenderSloTargetSpec{}
			json.Unmarshal([]byte(sloTargetSpec), &data)

			sloTargetDto := nextgen.SloTargetDto{
				Type_: nextgen.SLOTargetType(sloTargetType),
				SloTargetPercentage: sloTargetPercentage,
				Calender: &data,
			}

			serviceLevelObjectiveV2Dto.SloTarget = &sloTargetDto
		}

		sloType := request["type"].(string)
		sloSpec := request["spec"].(string)
		if sloType == "Simple" {
			data := nextgen.SimpleServiceLevelObjectiveSpec{}
			json.Unmarshal([]byte(sloSpec), &data)

			serviceLevelObjectiveV2Dto.Type_ = nextgen.SLOType(sloType)
			serviceLevelObjectiveV2Dto.Simple = &data
		}
		if sloType == "Composite" {
			data := nextgen.CompositeServiceLevelObjectiveSpec{}
			json.Unmarshal([]byte(sloSpec), &data)

			serviceLevelObjectiveV2Dto.Type_ = nextgen.SLOType(sloType)
			serviceLevelObjectiveV2Dto.Composite = &data
		}

		notificationRuleRefsReq := request["notification_rule_refs"].([]interface{})
		notificationRuleRefs := make([]nextgen.NotificationRuleRefDto, len(notificationRuleRefsReq))
		for i, notificationRuleRef := range notificationRuleRefsReq {
			test := notificationRuleRef.(map[string]interface{})
			notificationRuleRefDto := &nextgen.NotificationRuleRefDto{
				NotificationRuleRef: test["notification_rule_ref"].(string),
				Enabled: test["enabled"].(bool),
			}
			notificationRuleRefs[i] = *notificationRuleRefDto
		}
		serviceLevelObjectiveV2Dto.NotificationRuleRefs = notificationRuleRefs
	}

	return serviceLevelObjectiveV2Dto
}

func readSlo(d *schema.ResourceData, serviceLevelObjectiveV2Response **nextgen.ServiceLevelObjectiveV2Response, accountIdentifier string) {
	serviceLevelObjectiveV2Dto := &(*serviceLevelObjectiveV2Response).ServiceLevelObjectiveV2

	d.SetId((*serviceLevelObjectiveV2Dto).Identifier)

	d.Set("account_id", accountIdentifier)
	d.Set("org_id", (*serviceLevelObjectiveV2Dto).OrgIdentifier)
	d.Set("project_id", (*serviceLevelObjectiveV2Dto).ProjectIdentifier)
	d.Set("identifier", (*serviceLevelObjectiveV2Dto).Identifier)
	d.Set("request", []map[string]interface{}{
		{
			"name": (*serviceLevelObjectiveV2Dto).Name,
			"description": (*serviceLevelObjectiveV2Dto).Description,
			"tags": helpers.FlattenTags((*serviceLevelObjectiveV2Dto).Tags),
			"user_journey_refs": (*serviceLevelObjectiveV2Dto).UserJourneyRefs,
			"type": (*serviceLevelObjectiveV2Dto).Type_,
		},
	})
}
