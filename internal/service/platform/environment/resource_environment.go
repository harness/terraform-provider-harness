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
						"load_from_fallbackBranch": {
							Description: "If the Entity is to be fetched from fallbackBranch",
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

	resp, httpResp, err := c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Id(), c.AccountId, &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		ParentEntityConnectorRef: helpers.BuildField(d, "git_details.parent_entity_connector_ref"),
		ParentEntityRepoName: helpers.BuildField(d, "git_details.parent_entity_repo_name"),
		LoadFromFallbackBranch: helpers.BuildFieldForBoolean(d, "git_details.load_from_fallbackBranch"),
	})

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
	var httpResp *http.Response
	id := d.Id()
	env := buildEnvironment(d)

	if id == "" {
		resp, httpResp, err = c.EnvironmentsApi.CreateEnvironmentV2(ctx, c.AccountId, &nextgen.EnvironmentsApiCreateEnvironmentV2Opts{
			Body: optional.NewInterface(env),
			StoreType: helpers.BuildField(d, "git_details.store_type"),
			ConnectorRef: helpers.BuildField(d, "git_details.connector_ref"),
			CommitMsg: helpers.BuildField(d, "git_details.commit_message"),
			IsHarnessCodeRepo: helpers.BuildFieldForBoolean(d, "git_details.is_harnesscode_repo"),
			IsNewBranch: helpers.BuildFieldForBoolean(d, "git_details.is_new_branch"),
			Branch: helpers.BuildField(d, "git_details.branch_name"),
			RepoName:  helpers.BuildField(d, "git_details.repo_name"),
			FilePath: helpers.BuildField(d, "git_details.file_path"),
			BaseBranch: helpers.BuildField(d, "git_details.base_branch"),
		})
	} else {
		resp, httpResp, err = c.EnvironmentsApi.UpdateEnvironmentV2(ctx, c.AccountId, &nextgen.EnvironmentsApiUpdateEnvironmentV2Opts{
			Body: optional.NewInterface(env),
			StoreType: helpers.BuildField(d, "git_details.store_type"),
			ConnectorRef: helpers.BuildField(d, "git_details.connector_ref"),
			CommitMsg: helpers.BuildField(d, "git_details.commit_message"),
			IsHarnessCodeRepo: helpers.BuildFieldForBoolean(d, "git_details.is_harnesscode_repo"),
			IsNewBranch: helpers.BuildFieldForBoolean(d, "git_details.is_new_branch"),
			Branch: helpers.BuildField(d, "git_details.branch_name"),
			LastCommitId: helpers.BuildField(d, "git_details.lat_commit_id"),
			RepoIdentifier: helpers.BuildField(d, "git_details.repo_name"),
			FilePath: helpers.BuildField(d, "git_details.file_path"),
			LastObjectId: helpers.BuildField(d, "git_details.last_object_id"),
			BaseBranch: helpers.BuildField(d, "git_details.base_branch"),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readEnvironment(d, resp.Data.Environment)

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
