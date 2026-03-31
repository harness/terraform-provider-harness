package fault_template

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
				Description: "Template revision to query (defaults to v1 if not specified)",
				Type:        schema.TypeString,
				Optional:    true,
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

		// Retry logic to handle eventual consistency issues
		var resp chaos.ChaosfaulttemplateGetFaultTemplateResponse
		var httpResp interface{}
		var apiErr error

		// Get revision from user input or default to "v1"
		revision := "v1"
		if v, ok := d.GetOk("revision"); ok && v.(string) != "" {
			revision = v.(string)
		}
		log.Printf("[DEBUG] Using revision: %s", revision)

		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			resp, httpResp, apiErr = c.FaulttemplateApi.GetFaultTemplate(ctx, accountID, orgID, projectID, hubIdentity, revision, identityStr, nil)

			// If successful, break out of retry loop
			if apiErr == nil && resp.Data != nil {
				break
			}

			// If this is the last retry, return the error
			if i == maxRetries-1 {
				if apiErr != nil {
					return helpers.HandleChaosReadApiError(apiErr, d, httpResp.(*http.Response))
				}
				if resp.Data == nil {
					return diag.Errorf("fault template not found: %s", identityStr)
				}
			}

			// Wait before retrying (exponential backoff: 1s, 2s, 4s, 8s, 16s)
			waitTime := (1 << i) // 2^i seconds
			log.Printf("[DEBUG] Fault template not found yet, retrying in %d seconds (attempt %d/%d)", waitTime, i+1, maxRetries)
			select {
			case <-ctx.Done():
				return diag.Errorf("context cancelled while waiting for fault template")
			case <-time.After(time.Duration(waitTime) * time.Second):
				// Continue to next retry
			}
		}

		if resp.Data == nil {
			return diag.Errorf("fault template not found: %s", identityStr)
		}

		// Set the ID
		d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, resp.Data.Identity))

		// Set data source fields (simplified compared to resource)
		setDataSourceFaultTemplateData(d, resp.Data)
		return nil

	} else if name, ok := d.GetOk("name"); ok {
		// If name is provided, list and filter
		nameStr := name.(string)
		log.Printf("[DEBUG] Fetching fault template by name: %s", nameStr)

		// Retry logic to handle eventual consistency issues with ListFaultTemplate
		var foundTemplate *chaos.ChaosfaulttemplateChaosFaultTemplate
		var apiErr error

		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			resp, httpResp, err := c.FaulttemplateApi.ListFaultTemplate(ctx, accountID, nil)
			apiErr = err

			if err == nil && resp.Data != nil && len(resp.Data) > 0 {
				// Find exact match
				for j, t := range resp.Data {
					if t.Name == nameStr {
						foundTemplate = &resp.Data[j]
						break
					}
				}

				// If found, break out of retry loop
				if foundTemplate != nil {
					break
				}
			}

			// If this is the last retry, return the error
			if i == maxRetries-1 {
				if apiErr != nil {
					return helpers.HandleChaosReadApiError(apiErr, d, httpResp)
				}
				if foundTemplate == nil {
					totalFound := 0
					if resp.Data != nil {
						totalFound = len(resp.Data)
					}
					return diag.Errorf("fault template not found with name: %s (found %d templates total)", nameStr, totalFound)
				}
			}

			// Wait before retrying (exponential backoff: 1s, 2s, 4s, 8s, 16s)
			waitTime := (1 << i)
			log.Printf("[DEBUG] Fault template with name %s not found yet, retrying in %d seconds (attempt %d/%d)", nameStr, waitTime, i+1, maxRetries)
			select {
			case <-ctx.Done():
				return diag.Errorf("context cancelled while waiting for fault template")
			case <-time.After(time.Duration(waitTime) * time.Second):
				// Continue to next retry
			}
		}

		if foundTemplate == nil {
			return diag.Errorf("fault template not found with name: %s after %d retries", nameStr, maxRetries)
		}

		// Set the ID
		d.SetId(fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, hubIdentity, foundTemplate.Identity))

		// Set data source fields (simplified compared to resource)
		setDataSourceFaultTemplateData(d, foundTemplate)
		return nil

	} else {
		return diag.Errorf("either 'identity' or 'name' must be specified")
	}
}

// setDataSourceFaultTemplateData sets only the fields that exist in the data source schema
// This is different from the resource which has a full spec block
func setDataSourceFaultTemplateData(d *schema.ResourceData, template *chaos.ChaosfaulttemplateChaosFaultTemplate) {
	// Basic identifiers
	d.Set("account_id", template.AccountID)
	d.Set("hub_identity", template.HubRef)
	d.Set("identity", template.Identity)
	d.Set("name", template.Name)
	d.Set("description", template.Description)

	// Metadata
	d.Set("infrastructure_type", template.InfraType)
	d.Set("infrastructures", template.Infras)
	d.Set("tags", template.Tags)
	d.Set("category", template.Category)
	d.Set("type", template.Type_)
	d.Set("is_default", template.IsDefault)
	d.Set("is_enterprise", template.IsEnterprise)
	d.Set("is_removed", template.IsRemoved)
	d.Set("permissions_required", template.PermissionsRequired)

	// Computed fields
	d.Set("revision", template.Revision)
	d.Set("hub_ref", template.HubRef)
	d.Set("created_at", template.CreatedAt)
	d.Set("created_by", template.CreatedBy)
	d.Set("updated_at", template.UpdatedAt)
	d.Set("updated_by", template.UpdatedBy)

	// Variables
	if len(template.Variables) > 0 {
		variables := make([]map[string]interface{}, len(template.Variables))
		for i, variable := range template.Variables {
			varMap := map[string]interface{}{
				"name":     variable.Name,
				"required": variable.Required,
			}

			if variable.Value != nil {
				if val, ok := (*variable.Value).(string); ok {
					varMap["value"] = val
				}
			}

			if variable.Description != "" {
				varMap["description"] = variable.Description
			}

			if variable.Type_ != nil {
				varMap["type"] = strings.ToLower(string(*variable.Type_))
			}

			variables[i] = varMap
		}
		d.Set("variables", variables)
	}

	// Simplified fault_name field (data source doesn't have full spec)
	// Extract fault name from template YAML if available
	if template.Template != "" {
		// Simple extraction - look for faultName in YAML
		// This is a simplified version for the data source
		if strings.Contains(template.Template, "faultName:") {
			lines := strings.Split(template.Template, "\n")
			for _, line := range lines {
				if strings.Contains(line, "faultName:") {
					parts := strings.Split(line, ":")
					if len(parts) >= 2 {
						faultName := strings.TrimSpace(parts[1])
						faultName = strings.Trim(faultName, "\"'")
						d.Set("fault_name", faultName)
						break
					}
				}
			}
		}
	}
}
