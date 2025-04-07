package filters

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceGitOpsFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating Harness GitOps Filters.",

		ReadContext:   resourceGitOpsFiltersRead,
		UpdateContext: resourceGitOpsFiltersCreateOrUpdate,
		DeleteContext: resourceGitOpsFiltersDelete,
		CreateContext: resourceGitOpsFiltersCreateOrUpdate,
		Importer:      helpers.GitOpsFilterImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the GitOps filters.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description:  "Type of GitOps filters.",
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
				Required:    true,
				ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
					jsonStr := v.(string)
					var props map[string]interface{}
					
					if err := json.Unmarshal([]byte(jsonStr), &props); err != nil {
						errs = append(errs, fmt.Errorf("invalid JSON for filter_properties: %s", err))
						return
					}
					
					// Validate that all values are arrays
					for key, val := range props {
						if _, ok := val.([]interface{}); !ok {
							errs = append(errs, fmt.Errorf("property %s must be an array", key))
						}
					}
					
					return
				},
			},
			"filter_visibility": {
				Description:  "This indicates visibility of filters, by default it is Everyone.",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EveryOne", "OnlyCreator"}, false),
			},
		},
	}

	return resource
}

func resourceGitOpsFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	type_ := d.Get("type").(string)

	resp, httpResp, err := c.GitOpsFiltersApi.FilterServiceGet(ctx, id, &nextgen.FiltersApiFilterServiceGetOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		FilterType:        optional.NewString(type_),
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

func resourceGitOpsFiltersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	v1Filter := buildV1Filter(d)

	if id == "" {
		// Create
		resp, httpResp, err := c.GitOpsFiltersApi.FilterServiceCreate(ctx, *v1Filter, &nextgen.FiltersApiFilterServiceCreateOpts{
			AccountIdentifier: optional.NewString(c.AccountId),
		})

		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		d.SetId(resp.Identifier)
		// Read the created resource to get all properties
		return resourceGitOpsFiltersRead(ctx, d, meta)
	}

	// Update
	_, httpResp, err := c.GitOpsFiltersApi.FilterServiceUpdate(ctx, *v1Filter, &nextgen.FiltersApiFilterServiceUpdateOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Read the updated resource to get all properties
	return resourceGitOpsFiltersRead(ctx, d, meta)
}

func resourceGitOpsFiltersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	type_ := d.Get("type").(string)

	_, httpResp, err := c.GitOpsFiltersApi.FilterServiceDelete(ctx, id, &nextgen.FiltersApiFilterServiceDeleteOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		FilterType:        optional.NewString(type_),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildV1Filter(d *schema.ResourceData) *nextgen.V1Filter {
	filter := &nextgen.V1Filter{}

	if attr, ok := d.GetOk("org_id"); ok {
		filter.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		filter.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		filter.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		filter.Identifier = attr.(string)
	}

	// Set filter type
	if attr, ok := d.GetOk("type"); ok {
		filterType := nextgen.V1FilterType(attr.(string))
		filter.FilterType = &filterType
	}

	// Handle filter properties
	if attr, ok := d.GetOk("filter_properties"); ok {
		jsonString := attr.(string)
		var filterProps map[string]interface{}
		
		if err := json.Unmarshal([]byte(jsonString), &filterProps); err == nil {
			// Ensure all values are arrays of strings
			for k, v := range filterProps {
				if arrayVal, isArray := v.([]interface{}); isArray {
					// Convert interface array to string array
					strArray := make([]string, 0, len(arrayVal))
					for _, item := range arrayVal {
						if str, ok := item.(string); ok {
							strArray = append(strArray, str)
						}
					}
					filterProps[k] = strArray
				} else {
					// If not an array, convert to empty array
					filterProps[k] = []string{}
				}
			}
			
			// Set filter properties as interface
			filterPropsInterface := interface{}(filterProps)
			filter.FilterProperties = &filterPropsInterface
		}
	}

	// Handle filter visibility
	if attr, ok := d.GetOk("filter_visibility"); ok {
		visibilityStr := attr.(string)
		visibility := nextgen.V1FilterVisibility(visibilityStr)
		filter.FilterVisibility = &visibility
	}

	return filter
}
