package role_assignments

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRoleAssignments() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving role assignment.",

		ReadContext: resourceRoleAssignmentsRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier for role assignment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_group_identifier": {
				Description: "Resource group identifier.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"role_identifier": {
				Description: "Role identifier.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"principal": {
				Description: "Principal.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope_level": {
							Description: "Scope level.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identifier": {
							Description: "Identifier.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Type.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"disabled": {
				Description: "Disabled or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"managed": {
				Description: "Managed or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project Identifier",
			},
			"org_id": {
				Description: "Org identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	return resource
}
