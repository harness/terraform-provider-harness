package ng

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnector() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving a Harness connector"),

		ReadContext: dataSourceConnectorRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"branch": {
				Description: "The specified branch of the connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_id": {
				Description: "Unique identifier of the repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"types": {
				Description: fmt.Sprintf("The type of the connector. Available values are %s", strings.Join(nextgen.ConnectorTypesSlice, ", ")),
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"type": {
				Description: "The type of the selected connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "The tags of the connector.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"search_term": {
				Description: "The search term used to find the connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"include_all_connectors_available_at_scope": {
				Description: "Whether to include all connectors available at scope.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"get_default_from_other_repo": {
				Description: "Whether to get default from other repo.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"get_distinct_from_branches": {
				Description: "Whether to get distinct from branches.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"first_result": {
				Description: "When set to true if the query returns more than one result the first item will be selected. When set to false (default) this will return an error.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"connectivity_statuses": {
				Description: fmt.Sprintf("The connectivity status of the connector. Available options are %s", strings.Join(nextgen.ConnectorStatusSlice, ", ")),
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"inheriting_credentials_from_delegate": {
				Description: "Whether the connector inherits credentials from the delegate.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"ccm_connector_filter": {
				Description: "The ccm connector filter.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"features_enabled": {
							Description: fmt.Sprintf("The CCM features that are enabled. Valid options are %s.", strings.Join(nextgen.CCMFeaturesSlice, ", ")),
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
						},
						"aws_account_id": {
							Description: "The AWS account identifier.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"azure_subscription_id": {
							Description: "The Azure subscription identifier.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"azure_tenant_id": {
							Description: "The Azure tenant identifier.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"gcp_project_id": {
							Description: "The GCP project identifier.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"k8s_connector_ref": {
							Description: "The Kubernetes connector reference.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	filterProperties := nextgen.ConnectorFilterProperties{
		FilterType: "Connector",
	}

	searchOptions := &nextgen.ConnectorsApiGetConnectorListV2Opts{
		PageIndex: optional.NewInt32(0),
		PageSize:  optional.NewInt32(2),
		// Body:              optional.NewInterface(filterProperties),
	}

	if attr := d.Get("search_term").(string); attr != "" {
		searchOptions.SearchTerm = optional.NewString(attr)
	}

	if attr := d.Get("org_id").(string); attr != "" {
		searchOptions.OrgIdentifier = optional.NewString(attr)
	}

	if attr := d.Get("project_id").(string); attr != "" {
		searchOptions.ProjectIdentifier = optional.NewString(attr)
	}

	if attr, ok := d.GetOk("include_all_connectors_available_at_scope"); ok {
		searchOptions.IncludeAllConnectorsAvailableAtScope = optional.NewBool(attr.(bool))
	}

	if attr := d.Get("branch").(string); attr != "" {
		searchOptions.Branch = optional.NewString(attr)
	}

	if attr := d.Get("repo_id").(string); attr != "" {
		searchOptions.RepoIdentifier = optional.NewString(attr)
	}

	if attr, ok := d.GetOk("get_default_from_other_repo"); ok {
		searchOptions.GetDefaultFromOtherRepo = optional.NewBool(attr.(bool))
	}

	if attr, ok := d.GetOk("get_distinct_from_branches"); ok {
		searchOptions.GetDistinctFromBranches = optional.NewBool(attr.(bool))
	}

	if attr := d.Get("name").(string); attr != "" {
		filterProperties.ConnectorNames = []string{attr}
	}

	if attr := d.Get("identifier").(string); len(attr) > 0 {
		filterProperties.ConnectorIdentifiers = []string{attr}
	}

	if attr := d.Get("description").(string); len(attr) > 0 {
		filterProperties.ConnectorIdentifiers = []string{attr}
	}

	if attr := d.Get("types").([]interface{}); len(attr) > 0 {
		filterProperties.Types = utils.InterfaceSliceToStringSlice(attr)
	}

	if attr := d.Get("connectivity_statuses").([]interface{}); len(attr) > 0 {
		filterProperties.ConnectivityStatuses = utils.InterfaceSliceToStringSlice(attr)
	}

	if attr, ok := d.GetOk("inheriting_credentials_from_delegate"); ok {
		filterProperties.InheritingCredentialsFromDelegate = attr.(bool)
	}

	if attr := d.Get("tags").([]interface{}); len(attr) > 0 {
		filterProperties.Tags = utils.ExpandTags(attr)
	}

	if attr := d.Get("ccm_connector_filter").([]interface{}); len(attr) > 0 {
		filterProperties.CcmConnectorFilter = expandCcmConnectorFilter(attr)
	}

	resp, _, err := c.NGClient.ConnectorsApi.GetConnectorListV2(ctx, filterProperties, c.AccountId, searchOptions)

	if err != nil {
		return diag.FromErr(err)
	}

	if resp.Data.TotalItems == 0 {
		return diag.Errorf("no connectors found")
	}

	if resp.Data.TotalItems > 1 && !d.Get("first_result").(bool) {
		return diag.Errorf("more than one connector was found that matches the search criteria")
	}

	return readConnectorData(d, resp.Data.Content[0].Connector)

}

func readConnectorData(d *schema.ResourceData, connector *nextgen.ConnectorInfo) diag.Diagnostics {
	d.SetId(connector.Identifier)
	d.Set("name", connector.Name)
	d.Set("identifier", connector.Identifier)
	d.Set("description", connector.Description)
	d.Set("org_id", connector.OrgIdentifier)
	d.Set("project_id", connector.ProjectIdentifier)
	d.Set("tags", utils.FlattenTags(connector.Tags))
	d.Set("type", connector.Type_)

	return nil
}

func expandCcmConnectorFilter(list []interface{}) *nextgen.CcmConnectorFilter {
	if len(list) == 0 {
		return nil
	}

	m := list[0].(map[string]interface{})

	filter := &nextgen.CcmConnectorFilter{}

	if attr := m["features_enabled"].([]interface{}); len(attr) > 0 {
		filter.FeaturesEnabled = utils.InterfaceSliceToStringSlice(attr)
	}

	if attr := m["aws_account_id"].(string); attr != "" {
		filter.AwsAccountId = attr
	}

	if attr := m["azure_subscription_id"].(string); attr != "" {
		filter.AzureSubscriptionId = attr
	}

	if attr := m["azure_tenant_id"].(string); attr != "" {
		filter.AzureTenantId = attr
	}

	if attr := m["gcp_project_id"].(string); attr != "" {
		filter.GcpProjectId = attr
	}

	if attr := m["k8s_connector_ref"].(string); attr != "" {
		filter.K8sConnectorRef = attr
	}

	return filter
}
