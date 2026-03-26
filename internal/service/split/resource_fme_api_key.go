package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceFMEApiKey creates and deletes a Split API key (server-side or client-side only).
func ResourceFMEApiKey() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) API key. Only `server_side` and `client_side` keys are supported. The raw key value is only available immediately after create. Split may omit `id` on create and only return `key`; the provider then uses that value as `id` and for delete. Import id format: `org_id/project_id/<id_or_key_from_Split>`.",

		CreateContext: resourceFMEApiKeyCreate,
		ReadContext:   resourceFMEApiKeyRead,
		DeleteContext: resourceFMEApiKeyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEApiKeyImport,
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
				Description: "API key display name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"api_key_type": {
				Description: "Split API key type. Must be `server_side` or `client_side`.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"server_side",
					"client_side",
				}, false),
			},
			"environment_id": {
				Description: "Split environment ID the key is scoped to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"api_key": {
				Description: "The secret API key value (only set on initial create).",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"key_id": {
				Description: "Identifier used with the Split delete API (same as `id`). When the create response includes `id`, that is used; otherwise the returned `key` value.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMEApiKeyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, keyID, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	d.SetId(keyID)
	return []*schema.ResourceData{d}, nil
}

func resourceFMEApiKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req := splitsdk.KeyRequest{
		Name:    d.Get("name").(string),
		KeyType: d.Get("api_key_type").(string),
		Workspace: &splitsdk.KeyWorkspaceRequest{
			Type: "workspace",
			Id:   wsID,
		},
		Environments: []splitsdk.KeyEnvironmentRequest{
			{Type: "environment", Id: d.Get("environment_id").(string)},
		},
	}
	resp, err := client.ApiKeys.Create(req)
	if err != nil {
		return diag.FromErr(err)
	}
	splitKeyID := splitAPIKeyResourceID(resp)
	if splitKeyID == "" {
		return diag.Errorf("api key create returned success but empty id and key (debug: go run ./examples/split_debug_api_key -org-id=... -project-id=... -environment-id=... -name=unique_name)")
	}
	d.SetId(splitKeyID)
	if err := d.Set("key_id", splitKeyID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("api_key", resp.Key); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMEApiKeyRead(ctx, d, meta)
}

func resourceFMEApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	if err := d.Set("key_id", d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// splitAPIKeyResourceID picks the Terraform resource id and Split DELETE path segment.
// Production create responses often omit JSON "id" but include "key" (harness-go-sdk KeyResponse still fills Key).
func splitAPIKeyResourceID(resp *splitsdk.KeyResponse) string {
	if resp == nil {
		return ""
	}
	if resp.Id != "" {
		return resp.Id
	}
	return resp.Key
}

func resourceFMEApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	if err := client.ApiKeys.Delete(d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("api key delete: %w", err))
	}
	return nil
}
