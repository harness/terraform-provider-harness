package split

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceFMEApiKey creates and deletes a Split API key (server-side or client-side only).
func ResourceFMEApiKey() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) API key. Only `server_side` and `client_side` keys are supported. The raw key value is only available immediately after create. Split may omit `id` on create and only return `key`; the provider then uses that value as `id` and for delete. Import id format: `org_id/project_id/<id_or_key_from_Split>` when the Admin API returns key metadata on GET, or `org_id/project_id/environment_id/api_key_type/name/key_id` (six segments; `name` must not contain `/`).",

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
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	parts := strings.Split(d.Id(), "/")
	var orgID, projectID, keyID, envID, apiKeyType, name string
	switch len(parts) {
	case 3:
		var err error
		orgID, projectID, keyID, err = ParseImportID3(d.Id())
		if err != nil {
			return nil, err
		}
		var ok bool
		var getErr error
		name, apiKeyType, envID, ok, getErr = tryGetAPIKeyMetadata(client, keyID)
		if getErr != nil {
			return nil, getErr
		}
		if !ok {
			return nil, fmt.Errorf(
				"cannot import harness_fme_api_key with id %q: could not read key metadata from Split (use org_id/project_id/environment_id/api_key_type/name/key_id with six segments; api_key_type must be server_side or client_side; name must not contain '/')",
				d.Id(),
			)
		}
	case 6:
		var err error
		orgID, projectID, envID, apiKeyType, name, keyID, err = ParseImportID6(d.Id())
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf(
			"harness_fme_api_key import id must be org_id/project_id/key_id or org_id/project_id/environment_id/api_key_type/name/key_id (got %d segments)",
			len(parts),
		)
	}
	if apiKeyType != "server_side" && apiKeyType != "client_side" {
		return nil, fmt.Errorf("import: api_key_type must be server_side or client_side, got %q", apiKeyType)
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("environment_id", envID); err != nil {
		return nil, err
	}
	if err := d.Set("api_key_type", apiKeyType); err != nil {
		return nil, err
	}
	if err := d.Set("name", name); err != nil {
		return nil, err
	}
	d.SetId(keyID)
	if err := d.Set("key_id", keyID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

const splitAPIKeysPath = "/internal/api/v2/apiKeys"

// tryGetAPIKeyMetadata loads name, api key type, and environment id via GET when the Split Admin API supports it.
func tryGetAPIKeyMetadata(client *splitsdk.APIClient, keyID string) (name, apiKeyType, envID string, ok bool, err error) {
	u := client.BasePath + splitAPIKeysPath + "/" + url.PathEscape(keyID)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return "", "", "", false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", false, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", false, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return parseSplitAPIKeyMetadataJSON(body)
	case http.StatusNotFound, http.StatusMethodNotAllowed:
		return "", "", "", false, nil
	default:
		return "", "", "", false, fmt.Errorf("api key metadata: %d %s", resp.StatusCode, string(body))
	}
}

func parseSplitAPIKeyMetadataJSON(body []byte) (name, apiKeyType, envID string, ok bool, err error) {
	var payload struct {
		Name         string `json:"name"`
		ApiKeyType   string `json:"apiKeyType"`
		Type         string `json:"type"`
		Environments []struct {
			ID string `json:"id"`
		} `json:"environments"`
		EnvironmentIDs []string `json:"environmentIds"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", "", "", false, err
	}
	kt := payload.ApiKeyType
	if kt == "" {
		kt = payload.Type
	}
	env := ""
	switch {
	case len(payload.Environments) > 0:
		env = payload.Environments[0].ID
	case len(payload.EnvironmentIDs) > 0:
		env = payload.EnvironmentIDs[0]
	}
	if payload.Name == "" || kt == "" || env == "" {
		return "", "", "", false, nil
	}
	return payload.Name, kt, env, true, nil
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
