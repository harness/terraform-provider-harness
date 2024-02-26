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
				Type:        schema.TypeString,
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
						"branch_name": {
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
	var importResp nextgen.EnvironmentImportResponseDto
	var httpResp *http.Response
	id := d.Id()
	env := buildEnvironment(d)

	if id == "" {
		if d.Get("import_from_git").(bool) {
			envParams := envImportParam(env, d)
			importResp, httpResp, err = c.EnvironmentsApi.ImportEnvironment(ctx, c.AccountId, &envParams)
		} else {
			envParams := envCreateParam(env, d)
			resp, httpResp, err = c.EnvironmentsApi.CreateEnvironmentV2(ctx, c.AccountId, &envParams)
		}
	} else {
		envParams := envUpdateParam(env, d)
		resp, httpResp, err = c.EnvironmentsApi.UpdateEnvironmentV2(ctx, c.AccountId, &envParams)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if d.Get("import_from_git").(bool) {
		readImportRes(d, importResp.EnvIdentifier)
	} else {
	    readEnvironment(d, resp.Data.Environment)
	}

	return nil
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.EnvironmentsApi.DeleteEnvironmentV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentsApiDeleteEnvironmentV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		ForceDelete:       helpers.BuildFieldForBoolean(d, "force_delete"),
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
}

func getEnvParams(d *schema.ResourceData) *nextgen.EnvironmentsApiGetEnvironmentV2Opts {
	svcParams := &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
		OrgIdentifier:                 helpers.BuildField(d, "org_id"),
		ProjectIdentifier:             helpers.BuildField(d, "project_id"),
		Deleted:                       helpers.BuildFieldBool(d, "deleted"),
		Branch:                        helpers.BuildField(d, "git_details.0.branch"),
		RepoName:                      helpers.BuildField(d, "git_details.0.repo_name"),
		LoadFromCache:                 helpers.BuildField(d, "git_details.0.load_from_cache"),
		LoadFromFallbackBranch:        helpers.BuildFieldBool(d, "git_details.0.load_from_fallback_branch"),
	}
	return svcParams
}

func envCreateParam(svc *nextgen.EnvironmentRequest, d *schema.ResourceData) nextgen.EnvironmentsApiCreateEnvironmentV2Opts {
	return nextgen.EnvironmentsApiCreateEnvironmentV2Opts{
		Body:              optional.NewInterface(svc),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
	}
}

func envUpdateParam(svc *nextgen.EnvironmentRequest, d *schema.ResourceData) nextgen.EnvironmentsApiUpdateEnvironmentV2Opts {
	return nextgen.EnvironmentsApiUpdateEnvironmentV2Opts{
		Body:              optional.NewInterface(svc),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		IfMatch: helpers.BuildField(d, "if_match"),
		LastObjectId: helpers.BuildField(d, "git_details.0.last_object_id"),
		LastCommitId: helpers.BuildField(d, "git_details.0.last_commit_id"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
	}
}

func envImportParam(svc *nextgen.EnvironmentRequest, d *schema.ResourceData) nextgen.EnvironmentsV2ApiImportEnvironmentOpts {
	return nextgen.EnvironmentsV2ApiImportEnvironmentOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier:     helpers.BuildField(d, "project_id"),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harness_code_repo"),
		RepoName:          helpers.BuildField(d, "git_details.0.repo_name"),
		IsForceImport: helpers.BuildFieldBool(d, "git_details.0.is_force_import"),
	}
}

func readImportRes(d *schema.ResourceData, identifier string) {
	d.SetId(identifier)
	d.Set("identifier", identifier)
}