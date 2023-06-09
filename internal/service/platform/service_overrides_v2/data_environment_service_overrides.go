package service_overrides_v2

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceServiceOverrides() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for Harness service overrides V2.",

		ReadContext: dataSourceServiceOverridesRead,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Description: "The service ID to which the overrides applies.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"env_id": {
				Description: "The env ID to which the overrides are associated.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"infra_id": {
				Description: "The infrastructure ID to which the overrides are associated",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_id": {
				Description: "The cluster ID to which the overrides are associated",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "The type of the overrides",
				Type:        schema.TypeString,
				Required:    true,
			},
			"spec": {
				Description: "The overrides specification for the service.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	SetScopeDataResourceSchemaForServiceOverride(resource.Schema)

	return resource
}

func dataSourceServiceOverridesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	identifier := d.State().ID

	resp, httpResp, err := c.ServiceOverridesApi.GetServiceOverridesV2(ctx, identifier, c.AccountId,
		&nextgen.ServiceOverridesApiGetServiceOverridesV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil || resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readServiceOverridesV2(d, resp.Data)

	return nil
}

func SetScopeDataResourceSchemaForServiceOverride(s map[string]*schema.Schema) {
	s["project_id"] = helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional)
	s["org_id"] = helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional)
	s["identifier"] = helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required)
}
