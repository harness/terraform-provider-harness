package delegate

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDelegate() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness delegate. If more than one delegate matches the query the first one will be returned.",

		ReadContext: dataSourceDelegateRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the delegate",
				Type:        schema.TypeString,
				Computed:    true,
			},
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
			"account_id": {
				Description: "The account id the delegate belongs to.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"profile_id": {
				Description: "The id of the profile assigned to the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"host_name": {
				Description: "The host name of the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ip": {
				Description: "The ip address of the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"last_heartbeat": {
				Description: "The last time the delegate was heard from.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"polling_mode_enabled": {
				Description: "Whether the delegate is in polling mode.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"version": {
				Description: "The version of the delegate.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceDelegateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*api.Client)

	name := d.Get("name").(string)
	status := d.Get("status").(string)
	delegateType := d.Get("type").(string)

	delegates, _, err := c.CDClient.DelegateClient.GetDelegateWithFilters(1, 0, name, graphql.DelegateStatus(status), graphql.DelegateType(delegateType))

	if err != nil {
		return diag.FromErr(err)
	}

	if len(delegates) == 0 {
		return diag.Errorf("no delegate found with name %s, status %s, type %s", name, status, delegateType)
	}

	delegate := delegates[0]

	d.SetId(delegate.UUID)
	d.Set("name", delegate.DelegateName)
	d.Set("status", delegate.Status)
	d.Set("type", delegate.DelegateType)
	d.Set("account_id", delegate.AccountId)
	d.Set("profile_id", delegate.DelegateProfileId)
	d.Set("description", delegate.Description)
	d.Set("host_name", delegate.HostName)
	d.Set("ip", delegate.Ip)
	d.Set("last_heartbeat", delegate.LastHeartBeat)
	d.Set("version", delegate.Version)

	return nil
}
