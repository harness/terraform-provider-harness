package variable_set

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceVariableSet() *schema.Resource {
	resource := &schema.Resource{
		Description: `Data source for retrieving Variable Sets.

The Variable Set is looked up with ` + "`identifier`" + ` at the scope implied by ` + "`org_id`" + ` and ` + "`project_id`" + `:
omit both for an account level Variable Set, set ` + "`org_id`" + ` for an org level Variable Set, and set both for a
project level Variable Set.

The exported ` + "`id`" + ` is the bare identifier, without a scope prefix. When referencing a Variable Set from a
resource in a lower scope, such as ` + "`harness_platform_workspace`" + `, prefix the reference with ` + "`account.`" + ` for an
account level Variable Set or ` + "`org.`" + ` for an org level Variable Set. An unprefixed reference is resolved against
the consuming resource's own project.`,

		ReadContext: resourceVarsetRead,

		// identifier, name, description, tags, org_id and project_id come from
		// SetMultiLevelDatasourceSchemaIdentifierRequired below.
		Schema: map[string]*schema.Schema{
			"connector": {
				Description: "Provider connectors configured on the Variable Set.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector_ref": {
							Description: "Connector Ref is the reference to the connector",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Type is the connector type of the connector. Supported types: aws, azure, gcp",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"environment_variable": {
				Description: "Environment variables configured on the Variable Set",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field contains the value of the variable. For secret value types this contains a reference to a Harness secret.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable, either string or secret.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"terraform_variable": {
				Description: "Terraform variables configured on the Variable Set.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key is the identifier for the variable.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "Value is the value of the variable. For string value types this field contains the value of the variable. For secret value types this contains a reference to a Harness secret.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value_type": {
							Description: "Value type indicates the value type of the variable, either string or secret.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"terraform_variable_file": {
				Description: "Terraform variables files configured on the Variable Set",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Description: "Repository is the name of the repository the variables are fetched from.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_branch": {
							Description: "Repository branch is the name of the branch the variables are fetched from.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_commit": {
							Description: "Repository commit is the tag the variables are fetched from.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repository_sha": {
							Description: "Repository sha is the commit SHA the variables are fetched from.",
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
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	// The lookup only uses identifier plus the org and project scope, so name is really an
	// output. It stays Optional because existing configs may set it, and setting Computed
	// lets the value from the API populate it when they do not.
	resource.Schema["name"].Computed = true
	resource.Schema["name"].Description = "Name of the Variable Set. This is an output; a value set here is ignored by the lookup."

	resource.Schema["identifier"].Description = "Identifier of the Variable Set. Do not include a scope prefix here; use org_id and project_id to select the scope."
	resource.Schema["org_id"].Description = "Organization identifier of the organization the Variable Set resides in. Leave empty to look up an account level Variable Set."
	resource.Schema["project_id"].Description = "Project identifier of the project the Variable Set resides in. Leave empty to look up an account or org level Variable Set."
	resource.Schema["description"].Description = "Description of the Variable Set."

	// Variable Sets have no tags in the API, so this is always empty. It comes from the
	// shared helper and is kept for compatibility with configs that already reference it.
	resource.Schema["tags"].Description = "Tags are not supported on Variable Sets. This attribute is always empty."

	return resource
}
