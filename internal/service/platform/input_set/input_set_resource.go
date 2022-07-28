package input_set

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInputSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Resource Group",

		ReadContext:   resourceInputSetRead,
		UpdateContext: resourceInputSetCreateOrUpdate,
		CreateContext: resourceInputSetCreateOrUpdate,
		DeleteContext: resourceInputSetDelete,
		Importer:      helpers.PipelineResourceImporter,

		Schema: map[string]*schema.Schema{
			"pipeline_id": {
				Description: "Identifier of the pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipeline_branch": {
				Description: "Github branch of the Pipeline for which the Input Set is to be created",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pipeline_repo_id": {
				Description: "Github Repo identifier of the Pipeline for which the Input Set is to be created",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"branch": {
				Description: "Name of the branch",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_identifier": {
				Description: "Git Sync Config Id",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"root_folder": {
				Description: "Path to the root folder of the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"file_path": {
				Description: "File Path of the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"commit_msg": {
				Description: "Commit Message",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// "is_new_branch": {
			// 	Description: "Checks the new branch",
			// 	Type:        schema.TypeBool,
			// 	Default:     false,
			// 	Optional:    true,
			// },
			"base_branch": {
				Description: "Name of the default branch",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector_ref": {
				Description: "Identifier of Connector needed for CRUD operations on the respective Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"store_type": {
				Description: "Tells whether the Entity is to be saved on Git or not",
				Type:        schema.TypeString,
				Optional:    true,
				/* Enum: "INLINE" "REMOTE" */
			},
			"repo_name": {
				Description: "Name of the repository",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceInputSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	orgId := d.Get("org_id").(string)

	projectId := d.Get("project_id").(string)

	pipelineId := d.Get("pipeline_id").(string)

	resp, _, err := c.InputSetsApi.GetInputSet(ctx, id, c.AccountId, orgId, projectId, pipelineId, &nextgen.InputSetsApiGetInputSetOpts{
		Branch:         helpers.BuildField(d, "branch"),
		RepoIdentifier: helpers.BuildField(d, "repo_identifier"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	if resp.Data == nil {
		return nil
	}

	readInputSet(d, resp.Data)

	return nil

}

func resourceInputSetCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoInputSetResponse

	id := d.Id()
	inputSet := buildInputSet(d)
	orgIdentifier := d.Get("org_id").(string)
	projectIdentifier := d.Get("project_id").(string)
	pipelineIdentifier := d.Get("pipeline_id").(string)

	if id == "" {
		resp, _, err = c.InputSetsApi.PostInputSet(ctx, inputSet.InputSetYaml, c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier, &nextgen.InputSetsApiPostInputSetOpts{
			PipelineBranch: helpers.BuildField(d, "pipeline_branch"),
			PipelineRepoID: helpers.BuildField(d, "pipeline_repo_id"),
			Branch:         helpers.BuildField(d, "branch"),
			RepoIdentifier: helpers.BuildField(d, "repo_id"),
			RootFolder:     helpers.BuildField(d, "root_folder"),
			FilePath:       helpers.BuildField(d, "file_path"),
			CommitMsg:      helpers.BuildField(d, "commit_msg"),
			// IsNewBranch: helpers.BuildField(d, "is_new_branch"), TODO
			BaseBranch: helpers.BuildField(d, "base_branch"),
		})
	} else {
		resp, _, err = c.InputSetsApi.PutInputSet(ctx, inputSet.InputSetYaml, c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier, d.Id(), &nextgen.InputSetsApiPutInputSetOpts{
			PipelineBranch: helpers.BuildField(d, "pipeline_branch"),
			PipelineRepoID: helpers.BuildField(d, "pipeline_repo_id"),
			Branch:         helpers.BuildField(d, "branch"),
			RepoIdentifier: helpers.BuildField(d, "repo_id"),
			RootFolder:     helpers.BuildField(d, "root_folder"),
			FilePath:       helpers.BuildField(d, "file_path"),
			CommitMsg:      helpers.BuildField(d, "commit_msg"),
			// IsNewBranch: helpers.BuildField(d, "is_new_branch"), TODO
			BaseBranch: helpers.BuildField(d, "base_branch"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readInputSet(d, resp.Data)

	return nil
}

func resourceInputSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	orgIdentifier := helpers.BuildField(d, "org_id").Value()
	projectIdentifier := helpers.BuildField(d, "project_id").Value()
	pipelineIdentifier := helpers.BuildField(d, "pipeline_id").Value()

	_, _, err := c.InputSetsApi.DeleteInputSet(ctx, d.Id(), c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier, &nextgen.InputSetsApiDeleteInputSetOpts{})

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildInputSet(d *schema.ResourceData) *nextgen.InputSetResponse {
	inputSet := &nextgen.InputSetResponse{}

	if attr, ok := d.GetOk("account_id"); ok {
		inputSet.AccountId = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		inputSet.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		inputSet.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("pipeline_id"); ok {
		inputSet.PipelineIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("pipeline_branch"); ok {
		inputSet.GitDetails.Branch = attr.(string)
	}

	if attr, ok := d.GetOk("pipeline_repo_id"); ok {
		inputSet.GitDetails.RepoIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("branch"); ok {
		inputSet.GitDetails.Branch = attr.(string)
	}

	if attr, ok := d.GetOk("repo_identifier"); ok {
		inputSet.GitDetails.RepoIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("root_folder"); ok {
		inputSet.GitDetails.RootFolder = attr.(string)
	}

	if attr, ok := d.GetOk("file_path"); ok {
		inputSet.GitDetails.FilePath = attr.(string)
	}

	// if attr, ok := d.GetOk("commit_msg"); ok {
	// 	inputSet.CommitMsg = attr.(string)
	// }

	// if attr, ok := d.GetOk("is_new_branch"); ok {
	// 	inputSet.IsNewBranch = attr.(string)
	// }

	// if attr, ok := d.GetOk("base_branch"); ok {
	// 	inputSet.BaseBranch = attr.(string)
	// }

	// if attr, ok := d.GetOk("connector_ref"); ok {
	// 	inputSet.ConnectorRef = attr.(string)
	// }

	// if attr, ok := d.GetOk("store_type"); ok {
	// 	inputSet.StoreType = attr.(string)
	// }

	if attr, ok := d.GetOk("repo_name"); ok {
		inputSet.GitDetails.RepoName = attr.(string)
	}

	if attr, ok := d.GetOk("yaml"); ok {
		inputSet.InputSetYaml = attr.(string)
	}

	return inputSet
}

func readInputSet(d *schema.ResourceData, inputSet *nextgen.InputSetResponse) {
	d.SetId(inputSet.Identifier)
	d.Set("identifier", inputSet.Identifier)
	d.Set("name", inputSet.Name)
	d.Set("org_id", inputSet.OrgIdentifier)
	d.Set("project_id", inputSet.ProjectIdentifier)
	d.Set("pipeline_id", inputSet.PipelineIdentifier)
	if inputSet.GitDetails != nil {
		d.Set("pipeline_branch", inputSet.GitDetails.Branch)
		d.Set("pipeline_repo_id", inputSet.GitDetails.RepoIdentifier)
		d.Set("branch", inputSet.GitDetails.Branch)
		d.Set("repo_identifier", inputSet.GitDetails.RepoIdentifier)
		d.Set("root_folder", inputSet.GitDetails.RootFolder)
		d.Set("file_path", inputSet.GitDetails.FilePath)
		d.Set("repo_name", inputSet.GitDetails.RepoName)
	}

	// d.Set("commit_msg", inputSet.CommitMsg)
	// d.Set("is_new_branch", inputSet.IsNewBranch)
	// d.Set("base_branch", inputSet.BaseBranch)
	// d.Set("connector_ref", inputSet.ConnectorRef)
	// d.Set("store_type", inputSet.StoreType)
	d.Set("yaml", inputSet.InputSetYaml)
}
