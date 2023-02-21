package environment_service_overrides

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironmentServiceOverrides() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for creating a Harness environment service overrides.",

		ReadContext: dataSourceEnvironmentServiceOverridesRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the service overrides.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_id": {
				Description: "The service ID to which the overrides applies.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"env_id": {
				Description: "The env ID to which the overrides associated.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Environment Service Overrides YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	SetProjectLevelDataResourceSchemaForServiceOverride(resource.Schema)

	return resource
}

func dataSourceEnvironmentServiceOverridesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgId := d.Get("org_id").(string)
	projId := d.Get("project_id").(string)
	envId := d.Get("env_id").(string)

	resp, httpResp, err := c.EnvironmentsApi.GetServiceOverridesList(ctx, c.AccountId, orgId, projId, envId,
		&nextgen.EnvironmentsApiGetServiceOverridesListOpts{
			ServiceIdentifier: helpers.BuildField(d, "service_id"),
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

	readEnvironmentServiceOverridesList(d, resp.Data)

	return nil
}

func SetProjectLevelDataResourceSchemaForServiceOverride(s map[string]*schema.Schema) {
	s["project_id"] = helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Required)
	s["org_id"] = helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Required)
}
