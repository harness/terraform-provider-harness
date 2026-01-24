package fault_template

import (
	"context"
	"fmt"
	"log"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFaultTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Fault Template. " +
			"Supports lookup by identity (recommended) or name.",

		ReadContext: dataSourceFaultTemplateRead,

		Schema: map[string]*schema.Schema{
			"hub_identity": {
				Description: "Identity of the chaos hub",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identity": {
				Description:   "Unique identifier of the fault template (recommended)",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Description:   "Name of the fault template (may have timing issues with newly created templates)",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"identity"},
			},
			"org_id": {
				Description: "Organization identifier",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier",
				Type:        schema.TypeString,
				Optional:    true,
			},

			// Computed fields - Phase 1 fields
			"account_id": {
				Description: "Account identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the fault template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infrastructure_type": {
				Description: "Infrastructure type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infrastructures": {
				Description: "List of supported infrastructures",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Description: "Tags associated with the fault template",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"category": {
				Description: "Fault categories",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Description: "Fault type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_default": {
				Description: "Whether this is a default template",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_enterprise": {
				Description: "Whether this is an enterprise-only template",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_removed": {
				Description: "Soft delete flag",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"permissions_required": {
				Description: "Required permissions for the fault",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"revision": {
				Description: "Template revision",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"hub_ref": {
				Description: "Hub reference",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Creation timestamp",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"created_by": {
				Description: "Creator user ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Update timestamp",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"updated_by": {
				Description: "Updater user ID",
				Type:        schema.TypeString,
				Computed:    true,
			},

			// Variables
			"variables": {
				Description: "Template variables",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Variable name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Variable value",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Variable type",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"required": {
							Description: "Whether the variable is required",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"description": {
							Description: "Variable description",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},

			// Simplified spec fields for data source
			"fault_name": {
				Description: "Name of the fault",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFaultTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)

	// If identity is provided, fetch directly (RECOMMENDED)
	if identity, ok := d.GetOk("identity"); ok {
		identityStr := identity.(string)
		log.Printf("[DEBUG] Fetching fault template by identity: %s", identityStr)

		resp, httpResp, apiErr := c.FaulttemplateApi.GetFaultTemplate(ctx, accountID, orgID, projectID, hubIdentity, "latest", identityStr, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil {
			return diag.Errorf("fault template not found: %s", identityStr)
		}

		// Set the ID
		d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, resp.Data.Identity))

		// Set all the data
		if err := setFaultTemplateData(d, resp.Data); err != nil {
			return diag.FromErr(err)
		}
		return nil

	} else if name, ok := d.GetOk("name"); ok {
		// If name is provided, list and filter
		nameStr := name.(string)
		log.Printf("[DEBUG] Fetching fault template by name: %s (Note: may have timing issues with newly created templates)", nameStr)

		resp, httpResp, apiErr := c.FaulttemplateApi.ListFaultTemplate(ctx, accountID, nil)
		if apiErr != nil {
			return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
		}

		if resp.Data == nil || len(resp.Data) == 0 {
			return diag.Errorf("fault template not found with name: %s", nameStr)
		}

		// Find exact match
		var foundTemplate *chaos.ChaosfaulttemplateChaosFaultTemplate
		for i, t := range resp.Data {
			if t.Name == nameStr {
				foundTemplate = &resp.Data[i]
				break
			}
		}

		if foundTemplate == nil {
			return diag.Errorf("fault template not found with name: %s (found %d templates total)", nameStr, len(resp.Data))
		}

		// Set the ID
		d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, foundTemplate.Identity))

		// Set all the data
		if err := setFaultTemplateData(d, foundTemplate); err != nil {
			return diag.FromErr(err)
		}
		return nil

	} else {
		return diag.Errorf("either 'identity' or 'name' must be specified")
	}
}

// setFaultTemplateData is reused from resource_fault_template.go
// Note: The function is already defined in resource_fault_template.go
// and will set all Phase 1 fields including variables
