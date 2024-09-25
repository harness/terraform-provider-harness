package governance_rule_set

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceRuleSet() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating, updating, and managing rule.",
		ReadContext:   resourceRuleSetRead,
		CreateContext: resourceRuleSetCreateOrUpdate,
		UpdateContext: resourceRuleSetCreateOrUpdate,
		DeleteContext: resourceRuleDelete,
		Importer:      helpers.AccountLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the rule set.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description for rule set.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cloud_provider": {
				Description:  "The cloud provider for the rule set. It should be either AWS, AZURE or GCP.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS", "GCP", "AZURE"}, false),
			},
			"rule_ids": {
				Description: "List of rule IDs",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rule_set_id": {
				Description: "Id of the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceRuleSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.RuleSetsApi.ListRuleSets(ctx, readRuleSetRequest(id), c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		err := readRuleSetResponse(d, resp.Data)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}
	}

	return nil
}

func readRuleSetRequest(id string) nextgen.CreateRuleSetFilterDto {
	return nextgen.CreateRuleSetFilterDto{
		RuleSet: &nextgen.RuleSetRequest{
			RuleSetIds: []string{id},
		},
	}
}

func readRuleSetResponse(d *schema.ResourceData, ruleSetsList *nextgen.RuleSetList) error {
	ruleSet := ruleSetsList.RuleSet[0]

	d.Set("name", ruleSet.Name)
	d.Set("cloud_provider", ruleSet.CloudProvider)
	d.Set("description", ruleSet.Description)
	d.Set("rule_ids", ruleSet.RulesIdentifier)

	return nil
}

func resourceRuleSetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoRuleSet
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.RuleSetsApi.AddRuleSet(ctx, buildRuleSet(d, false), c.AccountId)
	} else {
		resp, httpResp, err = c.RuleSetsApi.UpdateRuleSet(ctx, buildRuleSet(d, true), c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		createOrUpdateRuleSetResponse(d, resp.Data)
	}

	return nil
}

func buildRuleSet(d *schema.ResourceData, update bool) nextgen.CreateRuleSetDto {
	ruleSet := &nextgen.RuleSet{
		Name:            d.Get("name").(string),
		CloudProvider:   d.Get("cloud_provider").(string),
		Description:     d.Get("description").(string),
		RulesIdentifier: expandStringList(d.Get("rule_ids").([]interface{})),
		IsOOTB:          false,
	}

	if update {
		ruleSet.Uuid = d.Id()
	}

	return nextgen.CreateRuleSetDto{
		RuleSet: ruleSet,
	}
}

func createOrUpdateRuleSetResponse(d *schema.ResourceData, ruleSet *nextgen.RuleSet) error {
	d.SetId(ruleSet.Uuid)
	d.Set("rule_set_id", ruleSet.Uuid)
	d.Set("name", ruleSet.Name)
	d.Set("cloud_provider", ruleSet.CloudProvider)
	d.Set("description", ruleSet.Description)
	d.Set("rule_ids", ruleSet.RulesIdentifier)

	return nil
}

func resourceRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	_, httpResp, err := c.RuleSetsApi.DeleteRuleSet(ctx, c.AccountId, id)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func expandStringList(givenStringListInterface []interface{}) []string {
	var expandedStringList []string

	if len(givenStringListInterface) > 0 {
		for _, id := range givenStringListInterface {
			expandedStringList = append(expandedStringList, id.(string))
		}
	}
	return expandedStringList
}
