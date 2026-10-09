package setting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Enum values of the NG settings API. The generated SDK models use plain strings for these.
const (
	settingUpdateTypeUpdate  = "UPDATE"
	settingUpdateTypeRestore = "RESTORE"

	settingSourceAccount = "ACCOUNT"
	settingSourceOrg     = "ORG"
	settingSourceProject = "PROJECT"
	settingSourceDefault = "DEFAULT"
)

// settingCategories lists the setting categories, searched in order when importing a setting.
var settingCategories = []string{
	"CD", "CI", "CE", "CV", "CF", "STO", "CORE", "PMS", "TEMPLATESERVICE", "GOVERNANCE", "CHAOS", "SCIM",
	"GIT_EXPERIENCE", "CONNECTORS", "EULA", "NOTIFICATIONS", "SUPPLY_CHAIN_ASSURANCE", "USER",
	"MODULES_VISIBILITY", "DBOPS", "IR", "AR",
}

func ResourceSetting() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing a Harness default setting (Account/Organization/Project Settings) at account, organization or project scope. " +
			"Destroying this resource restores the setting to the value inherited from the parent scope (or the Harness default).",
		ReadContext:   resourceSettingRead,
		CreateContext: resourceSettingCreateOrUpdate,
		UpdateContext: resourceSettingCreateOrUpdate,
		DeleteContext: resourceSettingDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the setting, e.g. `default_repo_for_git_experience`.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization. Omit to manage the setting at account scope.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description:  "Unique identifier of the project. Omit to manage the setting at account or organization scope.",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"org_id"},
			},
			"value": {
				Description: "Value of the setting. Boolean and number settings are passed as strings, e.g. `\"true\"` or `\"30\"`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"allow_overrides": {
				Description: "Whether child scopes can override this setting. Setting this to `false` removes existing overrides in child scopes.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"name": {
				Description: "Display name of the setting.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"category": {
				Description: "Category of the setting, e.g. `GIT_EXPERIENCE`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"group_identifier": {
				Description: "Group identifier of the setting.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value_type": {
				Description: "Type of the setting value: `String`, `Boolean` or `Number`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"allowed_values": {
				Description: "Allowed values of the setting, if restricted.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_modified_at": {
				Description: "Last modification timestamp of the setting.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, err := findSetting(ctx, c, d)
	if err != nil {
		return err
	}

	// The setting is no longer set at this scope (restored outside Terraform, or never set):
	// it now inherits its value, so remove it from state and let Terraform set it again.
	if resp == nil || resp.Setting.SettingSource != settingSource(d) {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readSetting(d, resp.Setting, resp.LastModifiedAt)

	return nil
}

func resourceSettingCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	result, diags := updateSetting(ctx, c, d, settingUpdateTypeUpdate)
	if diags != nil {
		return diags
	}

	d.SetId(d.Get("identifier").(string))
	if result.Setting != nil {
		readSetting(d, result.Setting, result.LastModifiedAt)
	}

	return nil
}

func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, diags := updateSetting(ctx, c, d, settingUpdateTypeRestore)
	return diags
}

// updateSetting sends a single UPDATE or RESTORE request. The settings API reports per-item
// failures in the response body (updateStatus=false) rather than as an HTTP error.
func updateSetting(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData, updateType string) (*nextgen.SettingUpdateResponseDto, diag.Diagnostics) {
	identifier := d.Get("identifier").(string)
	req := []nextgen.SettingRequestDto{{
		Identifier:     identifier,
		Value:          d.Get("value").(string),
		AllowOverrides: d.Get("allow_overrides").(bool),
		UpdateType:     updateType,
	}}

	resp, httpResp, err := c.SettingApi.UpdateSettingValue(ctx, req, c.AccountId, &nextgen.SettingApiUpdateSettingValueOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})
	if err != nil {
		return nil, helpers.HandleApiError(err, d, httpResp)
	}

	for i := range resp.Data {
		result := &resp.Data[i]
		if result.Identifier != identifier {
			continue
		}
		if !result.UpdateStatus {
			return nil, diag.Errorf("failed to %s setting %s: %s", updateType, identifier, result.ErrorMessage)
		}
		return result, nil
	}

	return nil, diag.Errorf("failed to %s setting %s: no result returned", updateType, identifier)
}

// findSetting looks up a setting at the resource scope. The single-setting GET endpoint only
// returns the value, so the list endpoint is used to get allow_overrides and the setting source.
// The category is known after create; on import every category is searched.
func findSetting(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData) (*nextgen.SettingResponseDto, diag.Diagnostics) {
	identifier := d.Get("identifier").(string)
	if identifier == "" {
		identifier = d.Id()
		d.Set("identifier", identifier)
	}

	categories := settingCategories
	opts := &nextgen.SettingApiGetSettingsListOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}
	knownCategory := false
	if category, ok := d.GetOk("category"); ok {
		categories = []string{category.(string)}
		opts.Group = helpers.BuildField(d, "group_identifier")
		knownCategory = true
	}

	var skippedErr error
	for _, category := range categories {
		resp, httpResp, err := c.SettingApi.GetSettingsList(ctx, c.AccountId, category, opts)
		if err != nil {
			// When searching on import, a category the server doesn't support (400 or 404) is
			// skipped. Any other failure (auth, server or network error) is returned, so it
			// isn't reported as "setting not found".
			if !knownCategory && isCategoryUnavailable(httpResp) {
				skippedErr = fmt.Errorf("category %s: %w", category, err)
				continue
			}
			return nil, helpers.HandleReadApiError(err, d, httpResp)
		}
		for i := range resp.Data {
			if resp.Data[i].Setting != nil && resp.Data[i].Setting.Identifier == identifier {
				return &resp.Data[i], nil
			}
		}
	}

	if !knownCategory {
		notFound := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("setting %s not found at the given scope", identifier),
		}
		if skippedErr != nil {
			notFound.Detail = fmt.Sprintf("Some categories were skipped; last error: %v", skippedErr)
		}
		return nil, diag.Diagnostics{notFound}
	}

	return nil, nil
}

// isCategoryUnavailable reports whether a list call failed because the server doesn't support
// the category (for example a module that isn't available on this cluster).
func isCategoryUnavailable(httpResp *http.Response) bool {
	return httpResp != nil && (httpResp.StatusCode == http.StatusBadRequest || httpResp.StatusCode == http.StatusNotFound)
}

// settingSource returns the source a setting has when it is set at the resource scope.
func settingSource(d *schema.ResourceData) string {
	if _, ok := d.GetOk("project_id"); ok {
		return settingSourceProject
	}
	if _, ok := d.GetOk("org_id"); ok {
		return settingSourceOrg
	}
	return settingSourceAccount
}

func readSetting(d *schema.ResourceData, setting *nextgen.SettingDto, lastModifiedAt int64) {
	d.SetId(setting.Identifier)
	d.Set("identifier", setting.Identifier)
	d.Set("value", setting.Value)
	d.Set("allow_overrides", setting.AllowOverrides)
	d.Set("name", setting.Name)
	d.Set("category", setting.Category)
	d.Set("group_identifier", setting.GroupIdentifier)
	d.Set("value_type", setting.ValueType)
	d.Set("allowed_values", setting.AllowedValues)
	d.Set("last_modified_at", lastModifiedAt)
}
