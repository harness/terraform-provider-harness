package split

import (
	"context"
	"sort"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMEEnvironment manages a Split environment in the FME workspace for a Harness org and project.
func ResourceFMEEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Create, update, and delete a Harness FME (Split) environment. Import id format: `org_id/project_id/environment_id`.",

		CreateContext: resourceFMEEnvironmentCreate,
		ReadContext:   resourceFMEEnvironmentRead,
		UpdateContext: resourceFMEEnvironmentUpdate,
		DeleteContext: resourceFMEEnvironmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEEnvironmentImport,
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Harness organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Harness project identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Environment name in Split (max 20 characters per Split API).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"production": {
				Description: "Whether this is a production environment.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"environment_id": {
				Description: "The Split environment ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"bootstrap_api_token_ids": {
				Description: "IDs of API keys auto-created by Split when the environment is created. " +
					"Only populated from the create response; the Split API does not return these on read. " +
					"Stored in Terraform state so the provider can delete them before destroying the environment. " +
					"Empty after `terraform import` unless you set this attribute manually (not recommended).",
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"change_permissions": {
				Description: "Change permission and approval settings for this environment. " +
					"Controls whether kills are allowed, whether approvals are required for changes, " +
					"and who can approve or skip approvals. " +
					"Note: the Split API does not return these on read; values are preserved from create/update responses.",
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_kills": {
							Description: "Whether kill operations are allowed in this environment.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"are_approvals_required": {
							Description: "Whether approvals are required before changes take effect.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"are_approvers_restricted": {
							Description: "Whether only specific users/groups/API keys can approve changes.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"approvers":             permissionEntityListSchema("Users, groups, or API keys that can approve changes."),
						"approval_skippable_by": permissionEntityListSchema("Users, groups, or API keys that can skip the approval requirement."),
					},
				},
			},
		},
	}
}

func resourceFMEEnvironmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, envID, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	d.SetId(envID)
	return []*schema.ResourceData{d}, nil
}

func resourceFMEEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req := splitsdk.CreateEnvironmentRequest{
		Name:              d.Get("name").(string),
		Production:        d.Get("production").(bool),
		ChangePermissions: expandChangePermissions(d),
	}
	env, err := client.Environments.Create(wsID, req)
	if err != nil {
		return diag.FromErr(err)
	}
	var tokenStrs []string
	for _, tok := range env.ApiTokens {
		if tok.ID != "" {
			tokenStrs = append(tokenStrs, tok.ID)
		}
	}
	sort.Strings(tokenStrs)
	bootstrapIDs := make([]interface{}, len(tokenStrs))
	for i, id := range tokenStrs {
		bootstrapIDs[i] = id
	}
	if err := d.Set("bootstrap_api_token_ids", bootstrapIDs); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(env.ID)
	return resourceFMEEnvironmentRead(ctx, d, meta)
}

func resourceFMEEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	env, err := client.Environments.FindByID(wsID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if env == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", env.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("production", env.Production); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", env.ID); err != nil {
		return diag.FromErr(err)
	}
	// Split list/get responses omit apiTokens; keep IDs captured at create (or import) in state.
	preserved := d.Get("bootstrap_api_token_ids")
	if preserved == nil {
		preserved = []interface{}{}
	}
	sorted := sortInterfaceStringSlice(preserved)
	if err := d.Set("bootstrap_api_token_ids", sorted); err != nil {
		return diag.FromErr(err)
	}
	// Split Get does not return change_permissions; preserve whatever is already in state.
	if prev := d.Get("change_permissions"); prev != nil {
		if err := d.Set("change_permissions", prev); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceFMEEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	name := d.Get("name").(string)
	prod := d.Get("production").(bool)
	req := splitsdk.UpdateEnvironmentRequest{
		Name:              &name,
		Production:        &prod,
		ChangePermissions: expandChangePermissions(d),
	}
	if _, err := client.Environments.Update(wsID, d.Id(), req); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMEEnvironmentRead(ctx, d, meta)
}

func resourceFMEEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	seen := map[string]struct{}{}
	for _, raw := range fmeEnvironmentBootstrapTokenIDs(d) {
		seen[raw] = struct{}{}
		_ = client.ApiKeys.Delete(raw)
	}
	env, err := client.Environments.FindByID(wsID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if env != nil {
		for _, tok := range env.ApiTokens {
			if tok.ID == "" {
				continue
			}
			if _, ok := seen[tok.ID]; ok {
				continue
			}
			_ = client.ApiKeys.Delete(tok.ID)
		}
	}
	if err := client.Environments.Delete(wsID, d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func sortInterfaceStringSlice(raw interface{}) []interface{} {
	list, ok := raw.([]interface{})
	if !ok || len(list) == 0 {
		if list == nil {
			return []interface{}{}
		}
		return list
	}
	strs := make([]string, 0, len(list))
	for _, item := range list {
		s, ok := item.(string)
		if !ok || s == "" {
			continue
		}
		strs = append(strs, s)
	}
	sort.Strings(strs)
	out := make([]interface{}, len(strs))
	for i, s := range strs {
		out[i] = s
	}
	return out
}

func fmeEnvironmentBootstrapTokenIDs(d *schema.ResourceData) []string {
	raw := d.Get("bootstrap_api_token_ids")
	if raw == nil {
		return nil
	}
	list, ok := raw.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	out := make([]string, 0, len(list))
	for _, item := range list {
		s, ok := item.(string)
		if !ok || s == "" {
			continue
		}
		out = append(out, s)
	}
	return out
}
