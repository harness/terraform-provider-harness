package idp

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func scorecardSchema(dataSource bool) map[string]*schema.Schema {
	computed := dataSource
	return map[string]*schema.Schema{
		"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
		"name": func() *schema.Schema {
			if dataSource {
				return helpers.GetNameSchema(helpers.SchemaFlagTypes.Computed)
			}
			return helpers.GetNameSchema(helpers.SchemaFlagTypes.Required)
		}(),
		"description": helpers.GetDescriptionSchema(func() helpers.SchemaFlagType {
			if dataSource {
				return helpers.SchemaFlagTypes.Computed
			}
			return helpers.SchemaFlagTypes.Optional
		}()),
		"published": {
			Type:        schema.TypeBool,
			Required:    !computed,
			Computed:    computed,
			Description: "Whether the scorecard is published.",
		},
		"on_demand": {
			Type:        schema.TypeBool,
			Optional:    !computed,
			Computed:    true,
			Description: "Whether the scorecard is evaluated on demand.",
		},
		"tier_group_identifier": {
			Type:        schema.TypeString,
			Optional:    !computed,
			Computed:    true,
			Description: "Identifier of the tier group used to classify scores.",
		},
		"tier_analytics": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Component distribution across scorecard tiers.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"tier_name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Name of the tier.",
					},
					"min_score": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "Minimum score for the tier.",
					},
					"max_score": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "Maximum score for the tier.",
					},
					"tier_colour": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Colour of the tier.",
					},
					"component_count": {
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "Number of components in the tier.",
					},
					"percentage": {
						Type:        schema.TypeFloat,
						Computed:    true,
						Description: "Percentage of components in the tier.",
					},
				},
			},
		},
		"weightage_strategy": func() *schema.Schema {
			s := &schema.Schema{
				Type:        schema.TypeString,
				Optional:    !computed,
				Computed:    true,
				Description: "Weightage strategy for checks. Valid values are EQUAL_WEIGHTS and CUSTOM.",
			}
			if !computed {
				s.ValidateFunc = validation.StringInSlice([]string{"EQUAL_WEIGHTS", "CUSTOM"}, false)
			}
			return s
		}(),
		"checks_missing": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Identifiers of checks referenced by the scorecard that are missing.",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"components": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of components evaluated by the scorecard.",
		},
		"percentage": {
			Type:        schema.TypeFloat,
			Computed:    true,
			Description: "Overall scorecard percentage.",
		},
		"filter": func() *schema.Schema {
			s := &schema.Schema{
				Type:        schema.TypeList,
				Required:    !computed,
				Computed:    computed,
				Description: "Filters that select catalog entities evaluated by the scorecard.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:        schema.TypeString,
							Required:    !computed,
							Computed:    computed,
							Description: "Catalog entity kind to evaluate.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    !computed,
							Computed:    computed,
							Description: "Catalog entity type to evaluate.",
						},
						"owners": {
							Type:        schema.TypeList,
							Optional:    !computed,
							Computed:    true,
							Description: "Entity owners to include.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tags": {
							Type:        schema.TypeList,
							Optional:    !computed,
							Computed:    true,
							Description: "Entity tags to include.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"lifecycle": {
							Type:        schema.TypeList,
							Optional:    !computed,
							Computed:    true,
							Description: "Entity lifecycle stages to include.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"scopes": {
							Type:        schema.TypeList,
							Optional:    !computed,
							Computed:    true,
							Description: "Evaluation scopes (for example ACCOUNT, ORGANIZATION, or PROJECT).",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			}
			if !computed {
				s.MaxItems = 1
			}
			return s
		}(),
		"checks": {
			Type:        schema.TypeSet,
			Required:    !computed,
			Computed:    computed,
			Description: "Checks included in the scorecard.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identifier": {
						Type:        schema.TypeString,
						Required:    !computed,
						Computed:    computed,
						Description: "Identifier of the check.",
					},
					"weightage": {
						Type:        schema.TypeFloat,
						Optional:    !computed,
						Computed:    true,
						Description: "Weightage of the check when using CUSTOM.",
					},
					"custom": {
						Type:        schema.TypeBool,
						Required:    !computed,
						Computed:    computed,
						Description: "Whether the referenced check is custom.",
					},
					"name": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Name of the check.",
					},
					"description": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Description of the check.",
					},
				},
			},
		},
	}
}

func ResourceScorecard() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating IDP scorecards.",
		ReadContext:   resourceScorecardRead,
		CreateContext: resourceScorecardCreateOrUpdate,
		UpdateContext: resourceScorecardCreateOrUpdate,
		DeleteContext: resourceScorecardDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: scorecardSchema(false),
	}
}

func resourceScorecardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, httpResp, err := c.ScorecardsApi.GetScorecard(ctx, id, &idp.ScorecardsApiGetScorecardOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		if isNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if err := readScorecard(d, resp); err != nil {
		return diag.Errorf("failed to read IDP scorecard: %v", err)
	}
	return nil
}

func resourceScorecardCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	req := expandScorecardRequest(d)

	var httpResp *http.Response
	var err error
	if d.Id() == "" {
		_, httpResp, err = c.ScorecardsApi.CreateScorecard(ctx, req, &idp.ScorecardsApiCreateScorecardOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	} else {
		_, httpResp, err = c.ScorecardsApi.UpdateScorecard(ctx, req, d.Id(), &idp.ScorecardsApiUpdateScorecardOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	}
	if err != nil {
		return handleIDPWriteApiError("harness_platform_idp_scorecard", err, d, httpResp)
	}

	d.SetId(req.Scorecard.Identifier)
	return resourceScorecardRead(ctx, d, meta)
}

func resourceScorecardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	httpResp, err := c.ScorecardsApi.DeleteScorecard(ctx, d.Id(), &idp.ScorecardsApiDeleteScorecardOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil
		}
		if isNotFoundError(err) {
			return nil
		}
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func expandScorecardRequest(d *schema.ResourceData) idp.ScorecardRequest {
	req := idp.ScorecardRequest{
		Scorecard: idp.Scorecard{
			Identifier: d.Get("identifier").(string),
			Name:       d.Get("name").(string),
		},
		Checks: expandScorecardChecks(d.Get("checks").(*schema.Set).List()),
	}
	req.Scorecard.Published = d.Get("published").(bool)
	if v, ok := d.GetOk("description"); ok {
		req.Scorecard.Description = v.(string)
	}
	if v, ok := d.GetOk("on_demand"); ok {
		req.Scorecard.OnDemand = v.(bool)
	}
	if v, ok := d.GetOk("tier_group_identifier"); ok {
		req.Scorecard.TierGroupIdentifier = v.(string)
	}
	if v, ok := d.GetOk("weightage_strategy"); ok {
		req.Scorecard.WeightageStrategy = v.(string)
	}
	req.Scorecard.Filter = expandScorecardFilter(d.Get("filter").([]interface{}))
	return req
}

func expandScorecardFilter(items []interface{}) *idp.ScorecardFilter {
	if len(items) == 0 || items[0] == nil {
		return nil
	}
	m := items[0].(map[string]interface{})
	filter := &idp.ScorecardFilter{
		Kind: m["kind"].(string),
	}
	if v, ok := m["type"].(string); ok {
		filter.Type_ = v
	}
	if v, ok := m["owners"].([]interface{}); ok {
		filter.Owners = expandStringList(v)
	}
	if v, ok := m["tags"].([]interface{}); ok {
		filter.Tags = expandStringList(v)
	}
	if v, ok := m["lifecycle"].([]interface{}); ok {
		filter.Lifecycle = expandStringList(v)
	}
	if v, ok := m["scopes"].([]interface{}); ok {
		filter.Scopes = expandStringList(v)
	}
	return filter
}

func expandScorecardChecks(items []interface{}) []idp.ScorecardCheck {
	checks := make([]idp.ScorecardCheck, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		check := idp.ScorecardCheck{
			Identifier: m["identifier"].(string),
			Custom:     m["custom"].(bool),
			Weightage:  m["weightage"].(float64),
		}
		checks = append(checks, check)
	}
	return checks
}

func expandStringList(items []interface{}) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		if s, ok := item.(string); ok && s != "" {
			out = append(out, s)
		}
	}
	return out
}

func readScorecard(d *schema.ResourceData, resp idp.ScorecardDetailsResponse) error {
	scorecard := resp.Scorecard
	if scorecard == nil {
		return nil
	}
	d.SetId(scorecard.Identifier)
	d.Set("identifier", scorecard.Identifier)
	d.Set("name", scorecard.Name)
	d.Set("description", scorecard.Description)
	d.Set("published", scorecard.Published)
	d.Set("on_demand", scorecard.OnDemand)
	d.Set("tier_group_identifier", scorecard.TierGroupIdentifier)
	d.Set("weightage_strategy", scorecard.WeightageStrategy)
	d.Set("checks_missing", scorecard.ChecksMissing)
	d.Set("components", scorecard.Components)
	d.Set("percentage", scorecard.Percentage)

	if scorecard.Filter != nil {
		if err := d.Set("filter", []map[string]interface{}{{
			"kind":      scorecard.Filter.Kind,
			"type":      scorecard.Filter.Type_,
			"owners":    scorecard.Filter.Owners,
			"tags":      scorecard.Filter.Tags,
			"lifecycle": scorecard.Filter.Lifecycle,
			"scopes":    scorecard.Filter.Scopes,
		}}); err != nil {
			return err
		}
	}

	tierAnalytics := make([]map[string]interface{}, 0, len(scorecard.TierAnalytics))
	for _, tier := range scorecard.TierAnalytics {
		tierAnalytics = append(tierAnalytics, map[string]interface{}{
			"tier_name":       tier.TierName,
			"min_score":       tier.MinScore,
			"max_score":       tier.MaxScore,
			"tier_colour":     tier.TierColour,
			"component_count": tier.ComponentCount,
			"percentage":      tier.Percentage,
		})
	}
	if err := d.Set("tier_analytics", tierAnalytics); err != nil {
		return err
	}

	checks := make([]map[string]interface{}, 0, len(resp.Checks))
	for _, check := range resp.Checks {
		checks = append(checks, map[string]interface{}{
			"identifier":  check.Identifier,
			"weightage":   check.Weightage,
			"custom":      check.Custom,
			"name":        check.Name,
			"description": check.Description,
		})
	}
	return d.Set("checks", checks)
}
