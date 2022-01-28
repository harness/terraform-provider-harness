package delegate

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDelegateIds() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get a list of delegate ID's matching the specified search criteria.",

		ReadContext: dataSourceDelegateIdsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the delegate to query for.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: fmt.Sprintf("The status of the delegate to query for. Valid values are %s", strings.Join(graphql.DelegateStatusSlice, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: fmt.Sprintf("The type of the delegate to query for. Valid values are %s", strings.Join(graphql.DelegateTypesSlice, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegate_ids": {
				Description: "A list of delegate ID's matching the specified search criteria.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceDelegateIdsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*api.Client)

	name := d.Get("name").(string)
	status := d.Get("status").(string)
	delegateType := d.Get("type").(string)

	delegates, _, err := c.CDClient.DelegateClient.ListDelegatesWithFilters(1, 0, name, graphql.DelegateStatus(status), graphql.DelegateType(delegateType))
	if err != nil {
		return diag.FromErr(err)
	}

	if len(delegates) == 0 {
		return diag.Errorf("no delegates found with name %s, status %s, type %s", name, status, delegateType)
	}

	delegate_ids := []string{}
	for _, delegate := range delegates {
		delegate_ids = append(delegate_ids, delegate.UUID)
	}

	d.SetId(fmt.Sprintf("%d", utils.StringHashcode(name+status+delegateType)))
	d.Set("delegate_ids", delegate_ids)

	return nil
}
