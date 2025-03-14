package variable_set

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceVariableSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving Variable Sets.",

		ReadContext: resourceVarsetRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Variable Set.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the organization the Variable Set resides in.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"project_id": {
				Description: "Project identifier of the project the Variable Set resides in.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the Variable Set.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Variable Set.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connector": {
				Description: "Provider connectors configured on the Variable Set. Only one connector of a type is supported",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
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
			"environment_variable": {
				Description: "Environment variables configured on the Variable Set",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the Variable Set.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field should contain the value of the variable. For secret value types this should contain a reference to a valid harness secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable. Currently we support string and secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"terraform_variable": {
				Description: "Terraform variables configured on the Variable Set. Terraform variable keys must be unique within the Variable Set.",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable. Must be unique within the Variable Set.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field should contain the value of the variable. For secret value types this should contain a reference to a valid harness secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable. Currently we support string and secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"terraform_variable_file": {
				Description: "Terraform variables files configured on the Variable Set",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Description: "Repository is the name of the repository to fetch the code from.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repository_branch": {
							Description: "Repository branch is the name of the branch to fetch the variables from. This cannot be set if repository commit or sha is set",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_commit": {
							Description: "Repository commit is tag to fetch the variables from. This cannot be set if repository branch or sha is set.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_sha": {
							Description: "Repository commit is SHA to fetch the variables from. This cannot be set if repository branch or commit is set.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"repository_connector": {
							Description: "Repository connector is the reference to the connector used to fetch the variables.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repository_path": {
							Description: "Repository path is the path in which the variables reside.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
