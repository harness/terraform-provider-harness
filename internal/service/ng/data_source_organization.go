package ng

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving a Harness organization"),

		ReadContext: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"search_term": {
				Description: "Search term used to find the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// "sort_orders": {
			// 	Description: "Sort orders used to order the organization list. Sort orders are listed as `field:order` pairs. For example `name:desc`.",
			// 	Type:        schema.TypeList,
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// 	Optional:    true,
			// },
			"tags": {
				Description: "Tags associated with the project.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"first_result": {
				Description: "When set to true if the query returns more than one result the first item will be selected. When set to false (default) this will return an error.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func dataSourceOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	searchOptions := &nextgen.OrganizationApiGetOrganizationListOpts{
		PageIndex: optional.NewInt32(0),
		PageSize:  optional.NewInt32(2),
	}

	if attr := d.Get("search_term").(string); attr != "" {
		searchOptions.SearchTerm = optional.NewString(attr)
	}

	if attr := d.Get("identifier").(string); attr != "" {
		searchOptions.Identifiers = optional.NewInterface([]string{attr})
	}

	resp, _, err := c.NGClient.OrganizationApi.GetOrganizationList(ctx, c.AccountId, searchOptions)

	if err != nil {
		return diag.FromErr(err)
	}

	if resp.Data.TotalItems == 0 {
		return diag.Errorf("organization not found")
	}

	if resp.Data.TotalItems > 1 && !d.Get("first_result").(bool) {
		return diag.Errorf("more than one organization was found that matches the search criteria")
	}

	readOrganization(d, resp.Data.Content[0].Organization)

	return nil
}
