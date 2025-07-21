package setting

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ServiceDiscoverySettingImporter is a basic importer that just sets up the ID structure
var ServiceDiscoverySettingImporter = &schema.ResourceImporter{
	StateContext: schema.ImportStatePassthroughContext,
}

func ResourceSetting() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing service discovery settings.",
		CreateContext: resourceSettingCreateOrUpdate,
		ReadContext:   resourceSettingRead,
		UpdateContext: resourceSettingCreateOrUpdate,
		DeleteContext: resourceSettingDelete,
		Importer:      ServiceDiscoverySettingImporter,

		Schema: map[string]*schema.Schema{
			"org_identifier": {
				Description: "The organization identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_identifier": {
				Description: "The project identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"correlation_id": {
				Description: "Correlation ID for the request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// Image Registry Settings
			"image_registry": {
				Description: "Image registry configuration.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Description: "The account name for the image registry.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"server": {
							Description: "The server URL for the image registry.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"secrets": {
							Description: "List of secrets for the image registry.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			// System fields (read-only)
			"created_at": {
				Description: "Timestamp when the setting was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Timestamp when the setting was last updated.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceSettingCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("failed to get service discovery client")
	}

	// Build the request
	req := svcdiscovery.ServiceSaveSettingRequest{}

	// Set image registry settings if provided
	if v, ok := d.GetOk("image_registry"); ok && len(v.([]interface{})) > 0 {
		imageRegCfg := v.([]interface{})[0].(map[string]interface{})
		req.ImageRegistry = &svcdiscovery.ServiceImageRegistrySetting{
			Account: imageRegCfg["account"].(string),
			Server:  imageRegCfg["server"].(string),
		}

		// Set secrets if provided
		if secrets, ok := imageRegCfg["secrets"].([]interface{}); ok && len(secrets) > 0 {
			req.ImageRegistry.Secrets = make([]string, len(secrets))
			for i, secret := range secrets {
				req.ImageRegistry.Secrets[i] = secret.(string)
			}
		}
	}

	// Get account, org, and project IDs
	accountID := c.AccountId
	orgID := d.Get("org_identifier").(string)
	projectID := d.Get("project_identifier").(string)

	// Set up optional parameters
	opts := &svcdiscovery.SettingApiSaveSettingOpts{}
	if orgID != "" {
		opts.OrganizationIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}
	if correlationID, ok := d.GetOk("correlation_id"); ok {
		opts.CorrelationID = optional.NewString(correlationID.(string))
	}

	// Make the API call
	_, httpResp, err := c.SettingApi.SaveSetting(ctx, req, accountID, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the ID and read back the resource to ensure the state is consistent
	d.SetId(fmt.Sprintf("%s/%s/%s", accountID, orgID, projectID))
	return resourceSettingRead(ctx, d, meta)
}

func resourceSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("failed to get service discovery client")
	}

	// Parse the ID
	accountID, orgID, projectID, err := parseSettingID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Set up optional parameters
	opts := &svcdiscovery.SettingApiGetSettingOpts{}
	if orgID != "" {
		opts.OrganizationIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}
	if correlationID, ok := d.GetOk("correlation_id"); ok {
		opts.CorrelationID = optional.NewString(correlationID.(string))
	}

	// Make the API call
	resp, httpResp, err := c.SettingApi.GetSetting(ctx, accountID, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Set the ID and settings
	d.Set("org_identifier", orgID)
	d.Set("project_identifier", projectID)
	d.Set("created_at", resp.CreatedAt)
	d.Set("updated_at", resp.UpdatedAt)

	// Set image registry settings if available
	if resp.ImageRegistry != nil {
		imageRegistry := map[string]interface{}{
			"account": resp.ImageRegistry.Account,
			"server":  resp.ImageRegistry.Server,
			"secrets": resp.ImageRegistry.Secrets,
		}
		d.Set("image_registry", []interface{}{imageRegistry})
	}

	return nil
}

func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Reset the settings to their default values
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("failed to get service discovery client")
	}

	// Parse the ID
	accountID, orgID, projectID, err := parseSettingID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Set up optional parameters
	opts := &svcdiscovery.SettingApiResetImageRegistrySettingOpts{}
	if orgID != "" {
		opts.OrganizationIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}
	if correlationID, ok := d.GetOk("correlation_id"); ok {
		opts.CorrelationID = optional.NewString(correlationID.(string))
	}

	// Make the API call to reset settings
	_, httpResp, err := c.SettingApi.ResetImageRegistrySetting(ctx, svcdiscovery.ServiceEmpty{}, accountID, opts)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Resource is deleted, remove from state
	d.SetId("")
	return nil
}

// parseSettingID parses the resource ID into account, org, and project IDs
// The ID format can be one of:
// - account_id (account level)
// - account_id/org_identifier (org level)
// - account_id/org_identifier/project_identifier (project level)
func parseSettingID(id string) (string, string, string, error) {
	parts := strings.Split(id, "/")
	switch len(parts) {
	case 1:
		// Account level: account_id
		return parts[0], "", "", nil
	case 2:
		// Org level: account_id/org_identifier
		return parts[0], parts[1], "", nil
	case 3:
		// Project level: account_id/org_identifier/project_identifier
		return parts[0], parts[1], parts[2], nil
	default:
		return "", "", "", fmt.Errorf("invalid ID format, expected account_id[/org_identifier[/project_identifier]]")
	}
}
