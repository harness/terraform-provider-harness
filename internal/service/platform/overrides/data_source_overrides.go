package overrides

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceOverrides() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for Harness Overrides V2.",

		ReadContext: dataSourceOverridesRead,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Description: "The service ID to which the override entity is associated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"env_id": {
				Description: "The environment ID to which the override entity is associated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infra_id": {
				Description: "The infrastructure ID to which the override entity is associated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_id": {
				Description: "The cluster ID to which the override entity is associated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of the override entity.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"yaml": {
				Description: "The yaml of the override entity.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "The identifier of the override entity.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"git_details": {
				Description: "Contains parameters related to Git Experience for remote overrides",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"load_from_cache": {
							Description: "Load service yaml from fallback branch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"load_from_fallback_branch": {
							Description: "Load service yaml from fallback branch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"repo_name": {
							Description: "Repo name of remote service override",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	SetScopeDataResourceSchemaForServiceOverride(resource.Schema)

	return resource
}

func dataSourceOverridesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	identifier := d.Get("identifier").(string)

	resp, httpResp, err := c.ServiceOverridesApi.GetServiceOverrides(ctx, identifier, c.AccountId,
		&nextgen.ServiceOverridesApiGetServiceOverridesV2Opts{
			OrgIdentifier:          helpers.BuildField(d, "org_id"),
			ProjectIdentifier:      helpers.BuildField(d, "project_id"),
			RepoName:               helpers.BuildField(d, "repo_name"),
			Branch:                 helpers.BuildField(d, "branch"),
			LoadFromFallbackBranch: helpers.BuildFieldBool(d, "load_from_fallback_branch"),
			LoadFromCache:          helpers.BuildField(d, "load_from_cache"),
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

	readOverrides(d, resp.Data)

	return nil
}

func SetScopeDataResourceSchemaForServiceOverride(s map[string]*schema.Schema) {
	s["project_id"] = helpers.GetProjectIdSchema(helpers.SchemaFlagTypes.Optional)
	s["org_id"] = helpers.GetOrgIdSchema(helpers.SchemaFlagTypes.Optional)
}
