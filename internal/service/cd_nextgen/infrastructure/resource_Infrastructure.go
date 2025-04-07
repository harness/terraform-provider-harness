package infrastructure

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

func ResourceInfrastructure() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Infrastructure.",

		ReadContext:   resourceInfrastructureRead,
		UpdateContext: resourceInfrastructureCreateOrUpdate,
		DeleteContext: resourceInfrastructureDelete,
		CreateContext: resourceInfrastructureCreateOrUpdate,
		Importer:      helpers.EnvRelatedResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the Infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"env_id": {
				Description: "Environment Identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: fmt.Sprintf("Type of Infrastructure. Valid values are %s.", strings.Join(nextgen.InfrastructureTypeValues, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
			},
			"yaml": {
				Description:      "Infrastructure YAML." + helpers.Descriptions.YamlText.String(),
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"deployment_type": {
				Description: fmt.Sprintf("Infrastructure deployment type. Valid values are %s.", strings.Join(nextgen.InfrastructureDeploymentypeValues, ", ")),
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"force_delete": {
				Description: "Enable this flag for force deletion of infrastructure",
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
							Description: "Last object identifier (for Github). To be provided only when updating infrastructure.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_commit_id": {
							Description: "Last commit identifier (for Git Repositories other than Github). To be provided only when updating infrastructure.",
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
							Description: "force import infrastructure from remote even if same file path already exist",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"import_from_git": {
							Description: "import infrastructure from git",
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

	// overwrite schema for tags since these are read from the yaml
	if s, ok := resource.Schema["tags"]; ok {
		s.Computed = true
	}

	return resource
}

func resourceInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)
	infraParams := getInfraParams(d)
	resp, httpResp, err := c.InfrastructuresApi.GetInfrastructure(ctx, d.Id(), c.AccountId, env_id, infraParams)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readInfrastructure(d, resp.Data)

	return nil
}

func resourceInfrastructureCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoInfrastructureResponse
	var importResp nextgen.ResponseInfrastructureImportResponse
	var httpResp *http.Response
	id := d.Id()
	infra := buildInfrastructure(d)

	if id == "" {
		if d.Get("git_details.0.import_from_git").(bool) {
			env_id := d.Get("env_id").(string)
			infraParams := infraImportParam(d)
			importResp, httpResp, err = c.InfrastructuresApi.ImportInfrastructure(ctx, c.AccountId, env_id, &infraParams)
		} else {
			infraParams := infraCreateParam(infra, d)
			resp, httpResp, err = c.InfrastructuresApi.CreateInfrastructure(ctx, c.AccountId, &infraParams)
		}
	} else {
		// Check if git details have changed using `d.HasChange` to compare the old and new values.
		connector_ref_changed := d.HasChange("git_details.0.connector_ref")
		filepath_changed := d.HasChange("git_details.0.file_path")
		reponame_changed := d.HasChange("git_details.0.repo_name")

		// If any of the Git-related fields have changed, we set the flag.
		shouldUpdateGitDetails := connector_ref_changed || filepath_changed || reponame_changed

		infraParams := infraUpdateParam(infra, d)
		resp, httpResp, err = c.InfrastructuresApi.UpdateInfrastructure(ctx, c.AccountId, &infraParams)

		if shouldUpdateGitDetails {
			resourceInfrastructureEditGitDetials(ctx, c, d)
		}
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if d.Get("git_details.0.import_from_git").(bool) {
		readImportRes(d, importResp.Data.Identifier)
	} else {
		readInfrastructure(d, resp.Data)
	}

	return nil
}

func resourceInfrastructureEditGitDetials(ctx context.Context, c *nextgen.APIClient, d *schema.ResourceData) diag.Diagnostics {
	id := d.Id()
	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	env_id := d.Get("env_id").(string)
	gitDetails := &nextgen.InfrastructuresApiEditGitDetailsMetadataOpts{
		ConnectorRef: helpers.BuildField(d, "git_details.0.connector_ref"),
		RepoName:     helpers.BuildField(d, "git_details.0.repo_name"),
		FilePath:     helpers.BuildField(d, "git_details.0.file_path"),
	}
	resp, httpResp, err := c.InfrastructuresApi.EditGitDetailsForInfrastructure(ctx, c.AccountId, org_id, project_id, env_id, id, gitDetails)

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

func resourceInfrastructureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)

	_, httpResp, err := c.InfrastructuresApi.DeleteInfrastructure(ctx, d.Id(), c.AccountId, env_id, &nextgen.InfrastructuresApiDeleteInfrastructureOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		ForceDelete:       helpers.BuildFieldBool(d, "force_delete"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildInfrastructure(d *schema.ResourceData) *nextgen.InfrastructureRequest {
	return &nextgen.InfrastructureRequest{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		EnvironmentRef:    d.Get("env_id").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Type_:             d.Get("type").(string),
		Yaml:              d.Get("yaml").(string),
	}
}

func readInfrastructure(d *schema.ResourceData, infra *nextgen.InfrastructureResponse) {
	d.SetId(infra.Infrastructure.Identifier)
	d.Set("org_id", infra.Infrastructure.OrgIdentifier)
	d.Set("project_id", infra.Infrastructure.ProjectIdentifier)
	d.Set("env_id", infra.Infrastructure.EnvironmentRef)
	d.Set("name", infra.Infrastructure.Name)
	d.Set("description", infra.Infrastructure.Description)
	d.Set("tags", helpers.FlattenTags(infra.Infrastructure.Tags))
	d.Set("type", infra.Infrastructure.Type_)
	d.Set("deployment_type", infra.Infrastructure.DeploymentType)
	d.Set("yaml", infra.Infrastructure.Yaml)

	var store_type = helpers.BuildField(d, "git_details.0.store_type")
	var base_branch = helpers.BuildField(d, "git_details.0.base_branch")
	var commit_message = helpers.BuildField(d, "git_details.0.commit_message")
	var connector_ref = helpers.BuildField(d, "git_details.0.connector_ref")

	if infra.Infrastructure.EntityGitDetails != nil {
		d.Set("git_details", []interface{}{readGitDetails(infra, store_type, base_branch, commit_message, connector_ref)})
	}
}

func readGitDetails(infra *nextgen.InfrastructureResponse, store_type optional.String, base_branch optional.String, commit_message optional.String, connector_ref optional.String) map[string]interface{} {
	git_details := map[string]interface{}{
		"branch":         infra.Infrastructure.EntityGitDetails.Branch,
		"file_path":      infra.Infrastructure.EntityGitDetails.FilePath,
		"repo_name":      infra.Infrastructure.EntityGitDetails.RepoName,
		"last_commit_id": infra.Infrastructure.EntityGitDetails.CommitId,
		"last_object_id": infra.Infrastructure.EntityGitDetails.ObjectId,
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

func getInfraParams(d *schema.ResourceData) *nextgen.InfrastructuresApiGetInfrastructureOpts {
	return &nextgen.InfrastructuresApiGetInfrastructureOpts{
		OrgIdentifier:          helpers.BuildField(d, "org_id"),
		ProjectIdentifier:      helpers.BuildField(d, "project_id"),
		Deleted:                helpers.BuildFieldBool(d, "deleted"),
		Branch:                 helpers.BuildField(d, "git_details.0.branch"),
		RepoName:               helpers.BuildField(d, "git_details.0.repo_name"),
		LoadFromCache:          helpers.BuildField(d, "git_details.0.load_from_cache"),
		LoadFromFallbackBranch: helpers.BuildFieldBool(d, "git_details.0.load_from_fallback_branch"),
	}
}

func infraCreateParam(infra *nextgen.InfrastructureRequest, d *schema.ResourceData) nextgen.InfrastructuresApiCreateInfrastructureOpts {
	return nextgen.InfrastructuresApiCreateInfrastructureOpts{
		Body:              optional.NewInterface(infra),
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

func infraUpdateParam(infra *nextgen.InfrastructureRequest, d *schema.ResourceData) nextgen.InfrastructuresApiUpdateInfrastructureOpts {
	return nextgen.InfrastructuresApiUpdateInfrastructureOpts{
		Body:              optional.NewInterface(infra),
		Branch:            helpers.BuildField(d, "git_details.0.branch"),
		FilePath:          helpers.BuildField(d, "git_details.0.file_path"),
		CommitMsg:         helpers.BuildField(d, "git_details.0.commit_message"),
		IsNewBranch:       helpers.BuildFieldBool(d, "git_details.0.is_new_branch"),
		BaseBranch:        helpers.BuildField(d, "git_details.0.base_branch"),
		ConnectorRef:      helpers.BuildField(d, "git_details.0.connector_ref"),
		StoreType:         helpers.BuildField(d, "git_details.0.store_type"),
		LastObjectId:      helpers.BuildField(d, "git_details.0.last_object_id"),
		LastCommitId:      helpers.BuildField(d, "git_details.0.last_commit_id"),
		IsHarnessCodeRepo: helpers.BuildFieldBool(d, "git_details.0.is_harnesscode_repo"),
	}
}

func infraImportParam(d *schema.ResourceData) nextgen.InfrastructuresApiImportInfrastructureOpts {
	return nextgen.InfrastructuresApiImportInfrastructureOpts{
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
