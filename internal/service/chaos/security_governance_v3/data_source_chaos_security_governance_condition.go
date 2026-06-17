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

// computedifySchema converts a resource schema element into a fully computed
// data-source schema element, recursing into nested resources.
func computedifySchema(s *schema.Schema) *schema.Schema {
	out := &schema.Schema{
		Type:        s.Type,
		Description: s.Description,
		Computed:    true,
	}
	switch e := s.Elem.(type) {
	case *schema.Schema:
		out.Elem = &schema.Schema{Type: e.Type}
	case *schema.Resource:
		out.Elem = &schema.Resource{Schema: computedifySchemaMap(e.Schema)}
	}
	return out
}

func computedifySchemaMap(in map[string]*schema.Schema) map[string]*schema.Schema {
	out := make(map[string]*schema.Schema, len(in))
	for k, v := range in {
		out[k] = computedifySchema(v)
	}
	return out
}

func DataSourceChaosSecurityGovernanceConditionV3() *schema.Resource {
	dsSchema := computedifySchemaMap(resourceChaosSecurityGovernanceConditionV3Schema())

	// Inputs: scope + the condition identifier to look up.
	dsSchema["org_id"] = &schema.Schema{
		Description: "The organization ID of the security governance condition",
		Type:        schema.TypeString,
		Required:    true,
	}
	dsSchema["project_id"] = &schema.Schema{
		Description: "The project ID of the security governance condition",
		Type:        schema.TypeString,
		Required:    true,
	}
	dsSchema["identity"] = &schema.Schema{
		Description:  "The identifier of the security governance condition. Exactly one of `identity` or `name` must be provided.",
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"identity", "name"},
	}
	dsSchema["name"] = &schema.Schema{
		Description:  "The name of the security governance condition. Exactly one of `identity` or `name` must be provided.",
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"identity", "name"},
	}

	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Security Governance Condition (V3 / REST API).",
		ReadContext: dataSourceChaosSecurityGovernanceConditionV3Read,
		Schema:      dsSchema,
	}
}

func dataSourceChaosSecurityGovernanceConditionV3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	conditionID, diags := resolveConditionV3ID(ctx, c, d, orgID, projectID)
	if diags.HasError() {
		return diags
	}

	resp, httpResp, err := c.GetconditionApi.GetCondition(ctx, c.AccountId, conditionID, &chaos.GetconditionApiGetConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId(conditionID)

	return setConditionV3Data(d, &resp, orgID, projectID)
}

// resolveConditionV3ID resolves the condition identifier either directly from
// the `identity` input or by searching for an exact `name` match via the list API.
func resolveConditionV3ID(ctx context.Context, c *chaos.APIClient, d *schema.ResourceData, orgID, projectID string) (string, diag.Diagnostics) {
	if identity, ok := d.GetOk("identity"); ok {
		return identity.(string), nil
	}

	name, ok := d.GetOk("name")
	if !ok {
		return "", diag.Errorf("either 'identity' or 'name' must be provided")
	}
	conditionName := name.(string)

	listResp, httpResp, err := c.ListconditionApi.ListCondition(ctx, c.AccountId, &chaos.ListconditionApiListConditionOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
		Search:                 optional.NewString(conditionName),
		Limit:                  optional.NewInt32(100),
	})
	if err != nil {
		return "", helpers.HandleChaosApiError(err, d, httpResp)
	}

	for _, item := range listResp.Conditions {
		if item.Name == conditionName {
			return item.ConditionId, nil
		}
	}

	return "", diag.Errorf("no security governance condition found with name '%s'", conditionName)
}
