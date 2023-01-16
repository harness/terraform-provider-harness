package input_set

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInputSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness input set.",

		ReadContext: dataSourceInputSetRead,

		Schema: map[string]*schema.Schema{
			"pipeline_id": {
				Description: "Identifier of the pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains parameters related to creating an Entity for Git Experience.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch_name": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"file_path": {
							Description: "File path of the Entity in the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"commit_message": {
							Description: "Commit message used for the merge commit.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"base_branch": {
							Description: "Name of the default branch (this checks out a new branch titled by branch_name).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connector_ref": {
							Description: "Identifier of the Harness Connector used for CRUD operations on the Entity.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"store_type": {
							Description: "Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_object_id": {
							Description: "Last object identifier (for Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_commit_id": {
							Description: "Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"parent_entity_connector_ref": {
							Description: "Connector reference for Parent Entity (Pipeline).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"parent_entity_repo_name": {
							Description: "Repository name for Parent Entity (Pipeline).",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceInputSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	pipelineId := d.Get("pipeline_id").(string)

	branch_name := helpers.BuildField(d, "git_details.0.branch_name")
	parent_entity_connector_ref := helpers.BuildField(d, "git_details.0.parent_entity_connector_ref")
	parent_entity_repo_name := helpers.BuildField(d, "git_details.0.parent_entity_repo_name")

	store_type := helpers.BuildField(d, "git_details.0.store_type")
	base_branch := helpers.BuildField(d, "git_details.0.base_branch")
	commit_message := helpers.BuildField(d, "git_details.0.commit_message")
	connector_ref := helpers.BuildField(d, "git_details.0.connector_ref")

	resp, httpResp, err := c.InputSetsApi.GetInputSet(ctx,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("identifier").(string),
		d.Get("pipeline_id").(string),
		&nextgen.InputSetsApiGetInputSetOpts{
			HarnessAccount:           optional.NewString(c.AccountId),
			BranchName:               branch_name,
			ParentEntityConnectorRef: parent_entity_connector_ref,
			ParentEntityRepoName:     parent_entity_repo_name,
		},
	)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readInputSet(d, &resp, pipelineId, store_type, base_branch, commit_message, connector_ref)

	return nil
}
