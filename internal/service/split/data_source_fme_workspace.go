package split

import (
	"context"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceFMEWorkspace returns the data source for resolving a Harness FME (Feature Management and Experimentation)
// workspace by Split workspace name or by Harness org_id and project_id (Workspaces.FindByName / FindByOrganizationAndProject).
func DataSourceFMEWorkspace() *schema.Resource {
	return &schema.Resource{
		Description: "Look up a Harness FME (Split) workspace by exact workspace name, or by Harness organization and project identifiers. " +
			"Specify either `name` or both `org_id` and `project_id` (not both).",

		ReadContext: dataSourceFMEWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Exact Split workspace name. Conflicts with `org_id` and `project_id`. Also populated from the API after read.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ConflictsWith: []string{
					"org_id",
					"project_id",
				},
				AtLeastOneOf: []string{
					"name",
					"org_id",
					"project_id",
				},
			},
			"org_id": {
				Description: "Harness organization identifier. Must be set together with `project_id` when not using `name`. Also populated from the API after read.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ConflictsWith: []string{
					"name",
				},
				RequiredWith: []string{
					"project_id",
				},
				AtLeastOneOf: []string{
					"name",
					"org_id",
					"project_id",
				},
			},
			"project_id": {
				Description: "Harness project identifier. Must be set together with `org_id` when not using `name`. Also populated from the API after read.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ConflictsWith: []string{
					"name",
				},
				RequiredWith: []string{
					"org_id",
				},
				AtLeastOneOf: []string{
					"name",
					"org_id",
					"project_id",
				},
			},
			"workspace_id": {
				Description: "The FME (Split) workspace ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Workspace type from the Split API.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"requires_title_and_comments": {
				Description: "Whether the workspace requires title and comments (Split API).",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

type rawConfigAt func(path cty.Path) (cty.Value, diag.Diagnostics)

func fmeWorkspaceRawString(getRaw rawConfigAt, attr string) (value string, fullyKnown bool) {
	v, diags := getRaw(cty.GetAttrPath(attr))
	if diags.HasError() || !v.IsKnown() {
		return "", false
	}
	if v.IsNull() {
		return "", true
	}
	if v.Type() != cty.String {
		return "", true
	}
	return v.AsString(), true
}

func dataSourceFMEWorkspaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return diags
	}

	getRaw := rawConfigAt(d.GetRawConfigAt)
	nameStr, nk := fmeWorkspaceRawString(getRaw, "name")
	orgStr, ok := fmeWorkspaceRawString(getRaw, "org_id")
	projStr, pk := fmeWorkspaceRawString(getRaw, "project_id")

	if !nk || !ok || !pk {
		return diag.Errorf("FME workspace lookup: arguments must be fully known before read")
	}

	var w splitsdk.Workspace
	var err error
	switch {
	case nameStr != "":
		w, err = WorkspaceByName(ctx, sessionFromMeta(meta), client, nameStr)
	case orgStr != "" && projStr != "":
		w, err = WorkspaceByOrganizationAndProject(ctx, sessionFromMeta(meta), client, orgStr, projStr)
	default:
		return diag.Errorf("set either `name` or both `org_id` and `project_id`")
	}
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("workspace_id", w.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", w.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", w.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("org_id", w.OrganizationIdentifier); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("project_id", w.ProjectIdentifier); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("requires_title_and_comments", w.RequiresTitleAndComments); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(w.ID)
	return nil
}
