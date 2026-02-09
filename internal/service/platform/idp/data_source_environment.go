package idp

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironment() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving IDP environments.",
		ReadContext: dataSourceEnvironmentRead,
		Schema: map[string]*schema.Schema{
			"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
			"org_id":     helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Required),
			"project_id": helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Required),
			"name":       helpers.GetNameSchema(helpers.SchemaFlagTypes.Optional),
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner of the environment",
			},
			"blueprint_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Blueprint to base the environment on",
			},
			"blueprint_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the blueprint to base the environment on",
			},
			"based_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Based on environment reference. This should be passed as <orgIdentifier>.<projectIdentifier>/<environmentIdentifier>",
			},
			"target_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "target state of the environment. If different from the current, a pipeline will be triggered to update the environment",
			},
			"overrides": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Overrides for environment blueprint inputs in YAML format",
			},
			"inputs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Additional inputs for controlling the environment. This should be passed as a map of key-value pairs in YAML format",
			},
		},
	}

	return resource
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	info := getEnvironmentInfo(d)

	resp, httpResp, err := c.EntitiesApi.GetEntity(ctx, info.Scope, environmentKind, info.Identifier, &idp.EntitiesApiGetEntityOpts{
		HarnessAccount:    optional.NewString(c.AccountId),
		OrgIdentifier:     info.OrgID,
		ProjectIdentifier: info.ProjectID,
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if err := readEnvironment(d, resp); err != nil {
		return diag.Errorf("failed to read environment from datasource. Err: %v", err)
	}

	return nil
}
