package security_governance_v3

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosSecurityGovernanceRuleV3() *schema.Resource {
	dsSchema := computedifySchemaMap(resourceChaosSecurityGovernanceRuleV3Schema())

	dsSchema["org_id"] = &schema.Schema{
		Description: "The organization ID of the security governance rule",
		Type:        schema.TypeString,
		Required:    true,
	}
	dsSchema["project_id"] = &schema.Schema{
		Description: "The project ID of the security governance rule",
		Type:        schema.TypeString,
		Required:    true,
	}
	dsSchema["identity"] = &schema.Schema{
		Description:  "The identifier of the security governance rule. Exactly one of `identity` or `name` must be provided.",
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"identity", "name"},
	}
	dsSchema["name"] = &schema.Schema{
		Description:  "The name of the security governance rule. Exactly one of `identity` or `name` must be provided.",
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"identity", "name"},
	}

	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Security Governance Rule (V3 / REST API).",
		ReadContext: dataSourceChaosSecurityGovernanceRuleV3Read,
		Schema:      dsSchema,
	}
}

func dataSourceChaosSecurityGovernanceRuleV3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	ruleID, diags := resolveRuleV3ID(ctx, c, d, orgID, projectID)
	if diags.HasError() {
		return diags
	}

	resp, httpResp, err := c.GetruleApi.GetRule(ctx, c.AccountId, ruleID, &chaos.GetruleApiGetRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId(ruleID)

	return setRuleV3Data(d, &resp, orgID, projectID)
}

// resolveRuleV3ID resolves the rule identifier either directly from the
// `identity` input or by searching for an exact `name` match via the list API.
func resolveRuleV3ID(ctx context.Context, c *chaos.APIClient, d *schema.ResourceData, orgID, projectID string) (string, diag.Diagnostics) {
	if identity, ok := d.GetOk("identity"); ok {
		return identity.(string), nil
	}

	name, ok := d.GetOk("name")
	if !ok {
		return "", diag.Errorf("either 'identity' or 'name' must be provided")
	}
	ruleName := name.(string)

	listResp, httpResp, err := c.ListruleApi.ListRule(ctx, c.AccountId, &chaos.ListruleApiListRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
		Search:                 optional.NewString(ruleName),
		Limit:                  optional.NewInt32(100),
	})
	if err != nil {
		return "", helpers.HandleChaosApiError(err, d, httpResp)
	}

	for _, item := range listResp.Rules {
		if item.Name == ruleName {
			return item.RuleId, nil
		}
	}

	return "", diag.Errorf("no security governance rule found with name '%s'", ruleName)
}
