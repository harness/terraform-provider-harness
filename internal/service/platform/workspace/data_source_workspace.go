package workspace

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceWorkspace() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving workspaces.",

		ReadContext: resourceWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Workspace",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Workspace",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the Workspace",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type defines the provisioning tool to use.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"provisioner_version": {
				Description: "Provisioner Version defines the tool version to use",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"provider_connector": {
				Description: "Provider Connector is the reference to the connector for the infrastructure provider",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository": {
				Description: "Repository is the name of the repository to use",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_branch": {
				Description: "Repository Branch in which the code should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"repository_commit": {
				Description: "Repository Commit/Tag in which the code should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"repository_connector": {
				Description: "Repository Connector is the reference to the connector to use for this code",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository_path": {
				Description: "Repository Path is the path in which the infra code resides",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cost_estimation_enabled": {
				Description: "If enabled cost estimation operations will be performed in this workspace",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"terraform_variable": {
				Description: "Terraform variables configured on the workspace",
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable`",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "value is the value of the variable",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable, text or secret",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"environment_variable": {
				Description: "Environment variables configured on the workspace",
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable`",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "value is the value of the variable",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable, text or secret",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"terraform_variable_file": {
				Description: "Terraform variables files configured on the workspace",
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Description: "Repository is the name of the repository to fetch the code from.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_branch": {
							Description: "Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit is set",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_commit": {
							Description: "Repository commit is commit or tag to fetch the variables from. This cannot be set if repository branch is set.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_connector": {
							Description: "Repository connector is the reference to the connector used to fetch the variables.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_path": {
							Description: "Repository path is the path in which the variables reside.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"default_pipelines": {
				Description: "Default pipelines associated with this workspace",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
	return resource
}
