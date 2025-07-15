package workspace

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Description: "Repository Tag in which the code should be accessed",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"repository_sha": {
				Description: "Repository Commit SHA in which the code should be accessed",
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
							Description: "Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit or sha is set",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_commit": {
							Description: "Repository commit is tag to fetch the variables from. This cannot be set if repository branch or sha is set.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_sha": {
							Description: "Repository commit is SHA to fetch the variables from. This cannot be set if repository branch or commit is set.",
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
			"variable_sets": {
				Description: "Variable sets to use.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connector": {
				Description: "Provider connectors configured on the Workspace. Only one connector of a type is supported",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector_ref": {
							Description: "Connector Ref is the reference to the connector",
							Type:        schema.TypeString,
							Required:    true,
						},
						"type": {
							Description:  "Type is the connector type of the connector. Supported types: aws, azure, gcp",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"aws", "azure", "gcp"}, false),
						},
					},
				},
			},
		},
	}
	resource.Schema["tags"] = helpers.GetTagsSchema(helpers.SchemaFlagTypes.Optional)
	return resource
}
