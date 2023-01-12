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
