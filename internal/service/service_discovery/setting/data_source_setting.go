package setting

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving service discovery settings.",
		ReadContext: dataSourceSettingRead,
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
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Description: "The account name for the image registry.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"server": {
							Description: "The server URL for the image registry.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"secrets": {
							Description: "List of secrets for the image registry.",
							Type:        schema.TypeList,
							Computed:    true,
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

func dataSourceSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).SDClient
	if c == nil {
		return diag.Errorf("failed to get service discovery client")
	}

	accountID := c.AccountId
	orgID := d.Get("org_identifier").(string)
	projectID := d.Get("project_identifier").(string)

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
	d.SetId(fmt.Sprintf("%s/%s/%s", accountID, orgID, projectID))

	// Set basic fields
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
