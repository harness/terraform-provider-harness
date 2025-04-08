package environment

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceEnvironment() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness environment.",

		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentCreateOrUpdate,
		DeleteContext: resourceEnvironmentDelete,
		CreateContext: resourceEnvironmentCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description:  fmt.Sprintf("The type of environment. Valid values are %s", strings.Join(nextgen.EnvironmentTypeValues, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(nextgen.EnvironmentTypeValues, false),
			},
			"yaml": {
				Description:      "Environment YAML." + helpers.Descriptions.YamlText.String(),
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"force_delete": {
				Description: "Enable this flag for force deletion of environments",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains parameters related to creating an Entity for Git Experience.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_path": {
							Description: "File path of the Entity in the repository.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"commit_message": {
							Description: "Commit message used for the merge commit.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_new_branch": {
							Description: "If a new branch creation is requested.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"base_branch": {
							Description: "Name of the default branch (this checks out a new branch titled by branch_name).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"connector_ref": {
							Description: "Identifier of the Harness Connector used for CRUD operations on the Entity." + helpers.Descriptions.ConnectorRefText.String(),
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"parent_entity_connector_ref": {
							Description: "Identifier of the Harness Connector used for CRUD operations on the Parent Entity." + helpers.Descriptions.ConnectorRefText.String(),
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"store_type": {
							Description:  "Specifies whether the Entity is to be stored in Git or not. Possible values: INLINE, REMOTE.",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"INLINE", "REMOTE"}, false),
							Computed:     true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"parent_entity_repo_name": {
							Description: "Name of the repository where parent entity lies.",
							Type:        schema.TypeString,
							Optional:    true,
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
						"is_harnesscode_repo": {
							Description: "If the gitProvider is HarnessCode",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"load_from_cache": {
							Description: "If the Entity is to be fetched from cache",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"load_from_fallback_branch": {
							Description: "If the Entity is to be fetched from fallbackBranch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_force_import": {
							Description: "force import environment from remote even if same file path already exist",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"import_from_git": {
							Description: "import environment from git",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	envParams := getEnvParams(d)
	resp, httpResp, err := c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Id(), c.AccountId, envParams)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readEnvironment(d, resp.Data.Environment)

	return nil
}

func resourceEnvironmentCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoEnvironmentResponse
	var importResp nextgen.ResponseEnvironmentImportResponseDto
	var httpResp *http.Response
	id := d.Id()
	env := buildEnvironment(d)
	shouldUpdateGitDetails := false

	if id == "" {
		if d.Get("git_details.0.import_from_git").(bool) {
			envParams := envImportParam(d)
			importResp, httpResp, err = c.EnvironmentsApi.ImportEnvironment(ctx, c.AccountId, &envParams)
		} else {
			envParams := envCreateParam(env, d)
			resp, httpResp, err = c.EnvironmentsApi.CreateEnvironmentV2(ctx, c.AccountId, &envParams)
		}
	} else {
		// Check if git details have changed using `d.HasChange` to compare the old and new values.
		connector_ref_changed := d.HasChange("git_details.0.connector_ref")
		filepath_changed := d.HasChange("git_details.0.file_path")
		reponame_changed := d.HasChange("git_details.0.repo_name")

		// If any of the Git-related fields have changed, we set the flag.
		shouldUpdateGitDetails = connector_ref_changed || filepath_changed || reponame_changed

		envParams := envUpdateParam(env, d)
		resp, httpResp, err = c.EnvironmentsApi.UpdateEnvironmentV2(ctx, c.AccountId, &envParams)

		if shouldUpdateGitDetails {
			resourceEnviornmentEditGitDetials(ctx, c, d)
		}
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if d.Get("git_details.0.import_from_git").(bool) {
		readImportRes(d, importResp.Data.EnvIdentifier)
	} else {
		if shouldUpdateGitDetails {
			envParams := getEnvParams(d)
			resp, httpResp, err = c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Id(), c.AccountId, envParams)

			if err != nil {
				return helpers.HandleReadApiError(err, d, httpResp)
			}
			readEnvironment(d, resp.Data.Environment)
		}
	}

	return nil
}

func resourceEnviornmentEditGitDetials(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData) diag.Diagnostics {
	id := d.Id()
	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	gitDetails := &nextgen.EnvironmentApiEditGitDetailsMetadataOpts{
		ConnectorRef: helpers.BuildField(d, "git_details.0.connector_ref"),
		RepoName:     helpers.BuildField(d, "git_details.0.repo_name"),
		FilePath:     helpers.BuildField(d, "git_details.0.file_path"),
	}
	resp, httpResp, err := c.EnvironmentsApi.EditGitDetailsForEnviornment(ctx, c.AccountId, org_id, project_id, id, gitDetails)

	if httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(resp.Data.Identifier)
	d.Set("identifier", resp.Data.Identifier)

	return nil
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.EnvironmentsApi.DeleteEnvironmentV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentsApiDeleteEnvironmentV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		ForceDelete:       helpers.BuildFieldBool(d, "force_delete"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildEnvironment(d *schema.ResourceData) *nextgen.EnvironmentRequest {
	return &nextgen.EnvironmentRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Name:              d.Get("name").(string),
		Color:             d.Get("color").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Type_:             d.Get("type").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readEnvironment(d *schema.ResourceData, env *nextgen.EnvironmentResponseDetails) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("name", env.Name)
	d.Set("color", env.Color)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("type", env.Type_.String())
	if d.Get("yaml").(string) != "" {
		d.Set("yaml", env.Yaml)
	}

	var store_type = helpers.BuildField(d, "git_details.0.store_type")
	var base_branch = helpers.BuildField(d, "git_details.0.base_branch")
	var commit_message = helpers.BuildField(d, "git_details.0.commit_message")
	var connector_ref = helpers.BuildField(d, "git_details.0.connector_ref")

	if env.EntityGitDetails != nil {
		d.Set("git_details", []interface{}{readGitDetails(env, store_type, base_branch, commit_message, connector_ref)})
	}
}

func readGitDetails(env *nextgen.EnvironmentResponseDetails, store_type optional.String, base_branch optional.String, commit_message optional.String, connector_ref optional.String) map[string]interface{} {
	git_details := map[string]interface{}{
		"branch":         env.EntityGitDetails.Branch,
		"file_path":      env.EntityGitDetails.FilePath,
		"repo_name":      env.EntityGitDetails.RepoName,
		"last_commit_id": env.EntityGitDetails.CommitId,
		"last_object_id": env.EntityGitDetails.ObjectId,
	}
	if store_type.IsSet() {
		git_details["store_type"] = store_type.Value()
	}
	if base_branch.IsSet() {
		git_details["base_branch"] = base_branch.Value()
	}
	if commit_message.IsSet() {
		git_details["commit_message"] = commit_message.Value()
	}
	if connector_ref.IsSet() {
		git_details["connector_ref"] = connector_ref.Value()
	}
	if connector_ref.Value() == "" {
		git_details["is_harness_code_repo"] = true
	}
	return git_details
}

func getEnvParams(d *schema.ResourceData) *nextgen.EnvironmentsApiGetEnvironmentV2Opts {
	return &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
		OrgIdentifier:          helpers.BuildField(d, "org_id"),
		ProjectIdentifier:      helpers.BuildField(d, "project_id"),
		Deleted:                helpers.BuildFieldBool(d, "deleted"),
		Branch:                 helpers.BuildField(d, "git_details.0.branch"),
		RepoName:               helpers.BuildField(d, "git_details.0.repo_name"),
		LoadFromCache:          helpers.BuildField(d, "git_details.0.load_from_cache"),
		LoadFromFallbackBranch: helpers.BuildFieldBool(d, "git_details.0.load_from_fallback_branch"),
	}
}

func envCreateParam(env *nextgen.EnvironmentRequest, d *schema.ResourceData) nextgen.EnvironmentsApiCreateEnvironmentV2Opts {
	return nextgen.EnvironmentsApiCreateEnvironmentV2Opts{
		Body:              optional.NewInterface(env),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harnesscode_repo"),
	}
}

func envUpdateParam(env *nextgen.EnvironmentRequest, d *schema.ResourceData) nextgen.EnvironmentsApiUpdateEnvironmentV2Opts {
	return nextgen.EnvironmentsApiUpdateEnvironmentV2Opts{
		Body:              optional.NewInterface(env),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		IfMatch:           helpers.BuildField(d, "if_match"),
		LastObjectId:      helpers.BuildField(d, "git_details.0.last_object_id"),
		LastCommitId:      helpers.BuildField(d, "git_details.0.last_commit_id"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harnesscode_repo"),
	}
}

func envImportParam(d *schema.ResourceData) nextgen.EnvironmentsV2ApiImportEnvironmentOpts {
	return nextgen.EnvironmentsV2ApiImportEnvironmentOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harnesscode_repo"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		IsForceImport:     helpers.BuildFieldBool(d, "git_details.0.is_force_import"),
	}
}

func readImportRes(d *schema.ResourceData, identifier string) {
	d.SetId(identifier)
	d.Set("identifier", identifier)
}
