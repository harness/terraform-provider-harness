package filters

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitOpsFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness GitOps Filter.",

		ReadContext: dataSourceGitOpsFiltersRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Filter.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description:  "Type of filter.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"APPLICATION"}, false),
			},
			"org_id": {
				Description: "Organization Identifier for the Entity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter_properties": {
				Description: "Properties of the filters entity defined in Harness as a JSON string. All values should be arrays of strings. Example: jsonencode({\"healthStatus\": [\"Healthy\", \"Degraded\"], \"syncStatus\": [\"Synced\"]})",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"filter_visibility": {
				Description: "This indicates visibility of filters, by default it is Everyone.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func dataSourceGitOpsFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)
	filterType := d.Get("type").(string)

	if id == "" {
		return diag.FromErr(errors.New("identifier must be specified"))
	}

	if filterType == "" {
		return diag.FromErr(errors.New("filter type must be specified"))
	}

	resp, httpResp, err := c.GitOpsFiltersApi.FilterServiceGet(ctx, id, &nextgen.FiltersApiFilterServiceGetOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		FilterType:        optional.NewString(filterType),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	d.SetId(resp.Identifier)
	d.Set("identifier", resp.Identifier)
	d.Set("name", resp.Name)
	d.Set("org_id", resp.OrgIdentifier)
	d.Set("project_id", resp.ProjectIdentifier)

	// Handle filter visibility
	if resp.FilterVisibility != nil {
		d.Set("filter_visibility", string(*resp.FilterVisibility))
	}

	// Handle filter properties
	if resp.FilterProperties != nil {
		if props, ok := (*resp.FilterProperties).(map[string]interface{}); ok {
			// Convert all values to proper arrays
			for k, v := range props {
				if strVal, isString := v.(string); isString {
					// Convert single string to array
					props[k] = []interface{}{strVal}
				} else if _, isArray := v.([]interface{}); !isArray {
					// If not a string or array, convert to empty array
					props[k] = []interface{}{}
				}
			}

			// Convert to JSON string
			jsonBytes, err := json.Marshal(props)
			if err == nil {
				d.Set("filter_properties", string(jsonBytes))
			}
		}
	}

	return nil
}
