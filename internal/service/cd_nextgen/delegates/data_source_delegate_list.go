package delegates

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDelegateList() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a list of Harness delegates.",

		ReadContext: dataSourceDelegateListRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fetch_all": {
				Description: "Whether to fetch all delegates.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"status": {
				Description: "Filter delegates by status. Valid values: CONNECTED, DISCONNECTED, ENABLED, DISABLED, WAITING_FOR_APPROVAL, DELETED.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegate_name": {
				Description: "Filter delegates by name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegate_group_identifier": {
				Description: "Filter delegates by group identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegate_tags": {
				Description: "Filter delegates by tags.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"delegate_instance_filter": {
				Description: "Filter delegate instances. Valid values: AVAILABLE, EXPIRED.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auto_upgrade": {
				Description: "Filter delegates by auto upgrade setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_status": {
				Description: "Filter delegates by version status. Valid values: ACTIVE, EXPIRED, EXPIRING, UNSUPPORTED.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filter_type": {
				Description: "Filter type for delegates.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegates": {
				Description: "List of delegates.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Delegate type (e.g., HELM_DELEGATE).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Delegate name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Delegate description.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tags": {
							Description: "Delegate tags.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"connected": {
							Description: "Whether the delegate is connected.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"auto_upgrade": {
							Description: "Auto upgrade setting.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"legacy": {
							Description: "Whether this is a legacy delegate.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDelegateListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	accountId := d.Get("account_id").(string)

	opts := &nextgen.DelegateSetupResourceApiListDelegatesOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		FetchAll:          optional.NewBool(d.Get("fetch_all").(bool)),
	}

	// Create request body with filter values
	body := nextgen.DelegateSetupListDetails{}

	if status, ok := d.GetOk("status"); ok {
		body.Status = nextgen.DelegateStatus(status.(string))
	}
	if delegateName, ok := d.GetOk("delegate_name"); ok {
		body.DelegateName = delegateName.(string)
	}
	if delegateGroupId, ok := d.GetOk("delegate_group_identifier"); ok {
		body.DelegateGroupIdentifier = delegateGroupId.(string)
	}
	if delegateTags, ok := d.GetOk("delegate_tags"); ok {
		tags := make([]string, 0)
		for _, tag := range delegateTags.([]interface{}) {
			tags = append(tags, tag.(string))
		}
		body.DelegateTags = tags
	}
	if instanceFilter, ok := d.GetOk("delegate_instance_filter"); ok {
		body.DelegateInstanceFilter = nextgen.DelegateInstanceFilter(instanceFilter.(string))
	}
	if autoUpgrade, ok := d.GetOk("auto_upgrade"); ok {
		body.AutoUpgrade = autoUpgrade.(string)
	}
	if versionStatus, ok := d.GetOk("version_status"); ok {
		body.VersionStatus = nextgen.DelegateVersionStatus(versionStatus.(string))
	}
	if filterType, ok := d.GetOk("filter_type"); ok {
		body.FilterType = nextgen.DelegateFilterType(filterType.(string))
	}

	resp, httpResp, err := c.DelegateSetupResourceApi.ListDelegates(ctx, body, accountId, opts)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	// Process delegates data using structured response
	delegates := make([]map[string]interface{}, 0, len(resp.Resource))

	for _, delegate := range resp.Resource {
		delegateMap := map[string]interface{}{
			"type":         delegate.Type,
			"name":         delegate.Name,
			"description":  delegate.Description,
			"tags":         delegate.Tags,
			"connected":    delegate.Connected,
			"auto_upgrade": delegate.AutoUpgrade,
			"legacy":       delegate.Legacy,
		}

		delegates = append(delegates, delegateMap)
	}

	d.Set("delegates", delegates)
	d.SetId(accountId)

	return nil
}
