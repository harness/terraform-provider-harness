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
				Optional:    true,
				ConflictsWith: []string{
					"name", "status", "type", "hostname",
				},
			},
			"name": {
				Description: "The name of the delegate to query for.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": {
				Description: "The hostname of the delegate.",
				Type:        schema.TypeString,
				Optional:    true,
				ConflictsWith: []string{
					"name", "status", "type", "id",
				},
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

	id := d.Get("id").(string)
	name := d.Get("name").(string)
	status := d.Get("status").(string)
	delegateType := d.Get("type").(string)
	hostname := d.Get("hostname").(string)

	var foundDelegate *graphql.Delegate
	var err error

	if id != "" {
		foundDelegate, err = c.CDClient.DelegateClient.GetDelegateById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if hostname != "" {
		foundDelegate, err = c.CDClient.DelegateClient.GetDelegateByHostName(hostname)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		var delegates []*graphql.Delegate
		delegates, _, err = c.CDClient.DelegateClient.ListDelegatesWithFilters(1, 0, name, graphql.DelegateStatus(status), graphql.DelegateType(delegateType))
		if err != nil {
			return diag.FromErr(err)
		}

		if len(delegates) == 0 {
			return diag.Errorf("no delegate found with name %s, status %s, type %s", name, status, delegateType)
		}

		foundDelegate = delegates[0]
	}

	d.SetId(foundDelegate.UUID)
	d.Set("name", foundDelegate.DelegateName)
	d.Set("status", foundDelegate.Status)
	d.Set("type", foundDelegate.DelegateType)
	d.Set("account_id", foundDelegate.AccountId)
	d.Set("profile_id", foundDelegate.DelegateProfileId)
	d.Set("description", foundDelegate.Description)
	d.Set("hostname", foundDelegate.HostName)
	d.Set("ip", foundDelegate.Ip)
	d.Set("last_heartbeat", foundDelegate.LastHeartBeat)
	d.Set("version", foundDelegate.Version)

	return nil
}
