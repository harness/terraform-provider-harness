package platform

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceOrganization() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness organization",

		ReadContext: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{
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
			"first_result": {
				Description: "When set to true if the query returns more than one result the first item will be selected. When set to false (default) this will return an error.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}

	helpers.SetCommonDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

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

	resp, _, err := c.OrganizationApi.GetOrganizationList(ctx, c.AccountId, searchOptions)

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
