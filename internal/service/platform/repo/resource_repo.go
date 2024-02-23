package repo

import (
	"context"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepo() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Repo.",

		ReadContext:   resourceRepoRead,
		CreateContext: resourceRepoCreateOrUpdate,
		UpdateContext: resourceRepoCreateOrUpdate,
		DeleteContext: resourceRepoDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func resourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	repo, resp, err := c.RepositoryApi.FindRepository(
		ctx,
		c.AccountId,
		id,
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRepo(d, &repo, orgID.Value(), projectID.Value())

	return nil
}

func resourceRepoCreateOrUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	var err error
	var repo code.TypesRepository
	var resp *http.Response

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")

	// check
	var sourceRepo string
	source := getSource(d)
	if source != nil {
		sourceRepo = source["repo"].(string)
	}

	// determine what type of write the change requires
	switch {
	case id != "":
		// update repo
		body := buildRepoBody(d)
		body.DefaultBranch = ""
		repo, resp, err = c.RepositoryApi.UpdateRepository(
			ctx,
			c.AccountId,
			id,
			&code.RepositoryApiUpdateRepositoryOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	case sourceRepo != "":
		// import repo
		repo, resp, err = c.RepositoryApi.ImportRepository(
			ctx,
			c.AccountId,
			&code.RepositoryApiImportRepositoryOpts{
				Body:              optional.NewInterface(buildRepoImportBody(d)),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	default:
		// create repo
		repo, resp, err = c.RepositoryApi.CreateRepository(
			ctx,
			c.AccountId,
			&code.RepositoryApiCreateRepositoryOpts{
				Body:              optional.NewInterface(buildRepoBody(d)),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	}
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepo(d, &repo, orgID.Value(), projectID.Value())
	return nil
}

func resourceRepoDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	id := d.Id()

	resp, err := c.RepositoryApi.DeleteRepository(
		ctx,
		c.AccountId,
		id,
		&code.RepositoryApiDeleteRepositoryOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	return nil
}

func buildRepoBody(d *schema.ResourceData) *code.OpenapiCreateRepositoryRequest {
	return &code.OpenapiCreateRepositoryRequest{
		DefaultBranch: d.Get("default_branch").(string),
		Description:   d.Get("description").(string),
		GitIgnore:     d.Get("git_ignore").(string),
		Identifier:    d.Get("identifier").(string),
		IsPublic:      false, // d.Get("is_public").(bool),
		License:       d.Get("license").(string),
		Readme:        d.Get("readme").(bool),
	}
}

func buildRepoImportBody(d *schema.ResourceData) *code.ReposImportBody {
	importBody := &code.ReposImportBody{
		Description: d.Get("description").(string),
		Identifier:  d.Get("identifier").(string),
	}

	source := getSource(d)
	if source != nil {
		providerType := code.ImporterProviderType(strings.ToLower(source["type"].(string)))
		importBody.Provider = &code.ImporterProvider{
			Host:     source["host"].(string),
			Password: source["password"].(string),
			Type_:    &providerType,
			Username: source["username"].(string),
		}
		importBody.ProviderRepo = source["repo"].(string)
	}

	return importBody
}

func getSource(d *schema.ResourceData) map[string]interface{} {
	srcSet := d.Get("source").(*schema.Set)
	srcList := srcSet.List()
	if len(srcList) == 0 {
		return nil
	}
	srcElem := srcList[0]
	return srcElem.(map[string]interface{})
}

func readRepo(d *schema.ResourceData, repo *code.TypesRepository, orgId string, projectId string) {
	d.SetId(repo.Identifier)
	d.Set("org_id", orgId)
	d.Set("project_id", projectId)
	d.Set("name", repo.Identifier)
	d.Set("identifier", repo.Identifier)
	d.Set("created", repo.Created)
	d.Set("created_by", repo.CreatedBy)
	d.Set("default_branch", repo.DefaultBranch)
	d.Set("description", repo.Description)
	d.Set("git_url", repo.GitUrl)
	d.Set("importing", repo.Importing)
	// d.Set("is_public", repo.IsPublic)
	d.Set("path", repo.Path)
	d.Set("updated", repo.Updated)
}

func createOnlyFieldDiffFunc(k, oldValue, newValue string, d *schema.ResourceData) bool { return true }

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identifier": {
			Description: "Identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the resource.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_by": {
			Description: "ID of the user who created the repository.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created": {
			Description: "Timestamp when the repository was created.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"default_branch": {
			Description:      "Default branch of the repository.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: createOnlyFieldDiffFunc, // handle default branch update
		},
		"description": {
			Description: "Description of the repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"git_url": {
			Description: "Git URL of the repository.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"id": {
			Description: "ID of the repository.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"importing": {
			Description: "Whether the repository is being imported.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		// "is_public": {
		// 	Description: "Whether the repository is public.",
		// 	Type:        schema.TypeBool,
		// 	Optional:    true,
		// },
		"path": {
			Description: "Path of the repository.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated": {
			Description: "Timestamp when the repository was last updated.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"git_ignore": {
			Description:      "Repository should be created with specified predefined gitignore file.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: createOnlyFieldDiffFunc,
		},
		"license": {
			Description:      "Repository should be created with specified predefined license file.",
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: createOnlyFieldDiffFunc,
		},
		"readme": {
			Description:      "Repository should be created with readme file.",
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: createOnlyFieldDiffFunc,
		},
		"source": {
			Description:      "Provider related configurations.",
			Type:             schema.TypeSet,
			Optional:         true,
			DiffSuppressFunc: createOnlyFieldDiffFunc,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"username": {
						Description: "The username for authentication when importing.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"password": {
						Description: "The password for authentication when importing.",
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
					},
					"type": {
						Description: "The type of SCM provider (github, gitlab, bitbucket, stash, gitea, gogs) when importing.",
						Type:        schema.TypeString, //type enum here
						Optional:    true,
					},
					"host": {
						Description: "The host URL for the import source.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"repo": {
						Description: "The provider repository to import from.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
}
