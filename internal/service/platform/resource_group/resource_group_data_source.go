package resource_group

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceResourceGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "DataSource for looking up resource group in harness.",
		ReadContext: resourceResourceGroupRead,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "Account Identifier of the account",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"allowed_scope_levels": {
				Description: "The scope levels at which this resource group can be used",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"included_scopes": {
				Description: "Included scopes",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Description: "Account Identifier of the account",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"org_id": {
							Description: "Organization Identifier",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"project_id": {
							Description: "Project Identifier",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"filter": {
							Description: "Can be one of these 2 EXCLUDING_CHILD_SCOPES or INCLUDING_CHILD_SCOPES",
							Type:        schema.TypeString,
							Computed:    true,
						},
					}},
			},
			"resource_filter": {
				Description: "Contains resource filter for a resource group",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_all_resources": {
							Description: "Include all resource or not",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"resources": {
							Description: "Resources for a resource group",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_type": {
										Description: "Type of the resource",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"identifiers": {
										Description: "List of the identifiers",
										Type:        schema.TypeSet,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"attribute_filter": {
										Description: "Used to filter resources on their attributes",
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_name": {
													Description: "Name of the attribute",
													Type:        schema.TypeString,
													Computed:    true,
												},
												"attribute_values": {
													Description: "Value of the attributes",
													Type:        schema.TypeSet,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
