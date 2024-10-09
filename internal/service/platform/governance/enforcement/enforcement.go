package governance_enforcement

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceRuleEnforcement() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating, updating, and managing rule enforcement.",
		ReadContext:   resourceRuleEnforcementRead,
		CreateContext: resourceRuleEnforcementCreateOrUpdate,
		UpdateContext: resourceRuleEnforcementCreateOrUpdate,
		DeleteContext: resourceRuleEnforcementDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the rule enforcement.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description for rule enforcement.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cloud_provider": {
				Description:  "The cloud provider for the rule enforcement. It should be either AWS, AZURE or GCP.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS", "GCP", "AZURE"}, false),
			},
			"rule_ids": {
				Description: "List of rule IDs. Either rule_ids or rule_set_ids should be provided.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rule_set_ids": {
				Description: "List of rule set IDs. Either rule_ids or rule_set_ids should be provided.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"execution_schedule": {
				Description: "Execution schedule in cron format.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"execution_timezone": {
				Description: "Timezone for the execution schedule.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"is_enabled": {
				Description: "Indicates if the rule enforcement is enabled. This by default is set to true.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"target_accounts": {
				Description: "List of target account which can be either AWS Account Ids or Azure Subscription Ids or Gcp Project Ids.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_regions": {
				Description: "List of target regions. For GCP it should be left empty but is required in case of AWS or Azure.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_dry_run": {
				Description: "Indicates if the rule enforcement is a dry run. This by default is set to false.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"enforcement_id": {
				Description: "Id of the rule enforcement.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},

		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, v interface{}) error {
			// Ensure at least one of rule_ids or rule_set_ids is provided
			ruleIDs, ruleSetIDs := d.Get("rule_ids").([]interface{}), d.Get("rule_set_ids").([]interface{})
			if len(ruleIDs) == 0 && len(ruleSetIDs) == 0 {
				return fmt.Errorf("either 'rule_ids' or 'rule_set_ids' must be provided")
			}

			// Conditionally require target_regions when cloud_provider is AWS or AZURE
			cloudProvider := d.Get("cloud_provider").(string)
			targetRegions := d.Get("target_regions").([]interface{})
			if (cloudProvider == "AWS" || cloudProvider == "AZURE") && len(targetRegions) == 0 {
				return fmt.Errorf("'target_regions' is required when cloud_provider is 'AWS' or 'AZURE'")
			}

			return nil
		},
	}

	return resource
}

func resourceRuleEnforcementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.RuleEnforcementApi.EnforcementDetails(ctx, c.AccountId, readRuleEnforcementRequest(id))

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		err := readRuleEnforcementResponse(d, resp.Data)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}
	}

	return nil
}

func resourceRuleEnforcementCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoRuleEnforcement
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.RuleEnforcementApi.AddRuleEnforcement(ctx, buildRuleEnforcement(d, false), c.AccountId)
	} else {
		resp, httpResp, err = c.RuleEnforcementApi.UpdateEnforcement(ctx, buildRuleEnforcement(d, true), c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		createOrUpdateRuleEnforcementResponse(d, resp.Data)
	}

	return nil
}

func resourceRuleEnforcementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	_, httpResp, err := c.RuleEnforcementApi.DeleteRuleEnforcement(ctx, c.AccountId, id)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildRuleEnforcement(d *schema.ResourceData, update bool) nextgen.CreateRuleEnforcementDto {
	ruleEnforcement := &nextgen.RuleEnforcement{
		Name:              d.Get("name").(string),
		CloudProvider:     d.Get("cloud_provider").(string),
		RuleIds:           expandStringList(d.Get("rule_ids").([]interface{})),
		RuleSetIDs:        expandStringList(d.Get("rule_set_ids").([]interface{})),
		ExecutionSchedule: d.Get("execution_schedule").(string),
		ExecutionTimezone: d.Get("execution_timezone").(string),
		IsEnabled:         d.Get("is_enabled").(bool),
		TargetAccounts:    expandStringList(d.Get("target_accounts").([]interface{})),
		TargetRegions:     expandStringList(d.Get("target_regions").([]interface{})),
		IsDryRun:          d.Get("is_dry_run").(bool),
		Description:       d.Get("description").(string),
	}

	if update {
		ruleEnforcement.Uuid = d.Id()
	}

	return nextgen.CreateRuleEnforcementDto{
		RuleEnforcement: ruleEnforcement,
	}
}

func createOrUpdateRuleEnforcementResponse(d *schema.ResourceData, ruleEnforcement *nextgen.RuleEnforcement) error {
	d.SetId(ruleEnforcement.Uuid)
	d.Set("name", ruleEnforcement.Name)
	d.Set("cloud_provider", ruleEnforcement.CloudProvider)
	d.Set("rule_ids", ruleEnforcement.RuleIds)
	d.Set("rule_set_ids", ruleEnforcement.RuleSetIDs)
	d.Set("execution_schedule", ruleEnforcement.ExecutionSchedule)
	d.Set("execution_timezone", ruleEnforcement.ExecutionTimezone)
	d.Set("is_enabled", ruleEnforcement.IsEnabled)
	d.Set("target_accounts", ruleEnforcement.TargetAccounts)
	d.Set("target_regions", ruleEnforcement.TargetRegions)
	d.Set("is_dry_run", ruleEnforcement.IsDryRun)
	d.Set("enforcement_id", ruleEnforcement.Uuid)
	d.Set("description", ruleEnforcement.Description)

	return nil
}

func readRuleEnforcementResponse(d *schema.ResourceData, ruleEnforcement *nextgen.EnforcementDetails) error {
	d.Set("name", ruleEnforcement.EnforcementName)
	d.Set("cloud_provider", ruleEnforcement.CloudProvider)
	d.Set("rule_ids", getMapKeys(ruleEnforcement.RuleIds))
	d.Set("rule_set_ids", getMapKeys(ruleEnforcement.RuleSetIds))
	d.Set("execution_schedule", ruleEnforcement.Schedule)
	d.Set("execution_timezone", ruleEnforcement.ExecutionTimezone)
	d.Set("is_enabled", ruleEnforcement.IsEnabled)
	d.Set("target_accounts", ruleEnforcement.Accounts)
	d.Set("target_regions", ruleEnforcement.Regions)
	d.Set("is_dry_run", ruleEnforcement.IsDryRun)
	d.Set("description", ruleEnforcement.Description)

	return nil
}

func readRuleEnforcementRequest(id string) *nextgen.RuleEnforcementApiEnforcementDetailsOpts {
	return &nextgen.RuleEnforcementApiEnforcementDetailsOpts{
		EnforcementId: optional.NewString(id),
	}
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

func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
