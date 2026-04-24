package split

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceFMEEnvironment returns the data source for a Split environment in the FME workspace
// for a Harness organization and project. The workspace is resolved via Workspaces.ResolveWorkspaceID;
// the environment is loaded with Environments.FindByName (no workspace_id in schema).
func DataSourceFMEEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Look up a Harness FME (Split) environment by name within the workspace for a Harness organization and project. " +
			"The provider resolves the workspace ID from `org_id` and `project_id` on each read.",

		ReadContext: dataSourceFMEEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Harness organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Harness project identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Split environment in the workspace (e.g. `Production`).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment_id": {
				Description: "The Split environment ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"production": {
				Description: "Whether this is a production environment in Split.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMEEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return diags
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)

	env, err := EnvironmentByOrganizationProjectAndName(ctx, sessionFromMeta(meta), client, orgID, projectID, name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_id", env.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", env.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("production", env.Production); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(env.ID)
	return nil
}
