package pipeline

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipeline() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness pipeline.",

		ReadContext: dataSourcePipelineRead,

		Schema: map[string]*schema.Schema{
			"yaml": {
				Description: "YAML of the pipeline.",
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
					},
				},
			},
			"template_applied": {
				Description: "If true, returns Pipeline YAML with Templates applied on it.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"template_applied_pipeline_yaml": {
				Description: "Pipeline YAML after resolving Templates (returned as a String).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	pipeline_id := d.Get("identifier").(string)
	template_applied := d.Get("template_applied").(bool)
	store_type := helpers.BuildField(d, "git_details.0.store_type")
	base_branch := helpers.BuildField(d, "git_details.0.base_branch")
	commit_message := helpers.BuildField(d, "git_details.0.commit_message")
	connector_ref := helpers.BuildField(d, "git_details.0.connector_ref")

	resp, httpResp, err := c.PipelinesApi.GetPipeline(ctx,
		org_id,
		project_id,
		pipeline_id,
		&nextgen.PipelinesApiGetPipelineOpts{HarnessAccount: optional.NewString(c.AccountId)},
	)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPipeline(d, resp, project_id, org_id, template_applied, store_type, base_branch, commit_message, connector_ref)

	return nil
}
