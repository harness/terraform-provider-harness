package repo

import (
	"context"
	"net/http"
	"strings"
	"time"

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

	repo, resp, err := c.RepositoryApi.GetRepository(
		ctx,
		c.AccountId,
		id,
		&code.RepositoryApiGetRepositoryOpts{
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

	// If import is in progress, wait for it to complete
	if repo.Importing {
		if err := waitForImportCompletion(ctx, c.RepositoryApi, repo.Identifier, c.AccountId, orgID, projectID); err != nil {
			return diag.FromErr(err)
		}
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
	var srcElem map[string]interface{}
	if attr, ok := d.GetOk("source"); ok {
		if attr != nil && len(attr.(*schema.Set).List()) > 0 {
			srcElem = attr.(*schema.Set).List()[0].(map[string]interface{})
		}
	}
	return srcElem
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
	// d.Set("is_public", repo.IsPublic)
	d.Set("path", repo.Path)
	d.Set("updated", repo.Updated)
}

func waitForImportCompletion(ctx context.Context, api *code.RepositoryApiService, importID string, accountID string,
	projectID optional.String, orgID optional.String) error {
	for {
		repo, _, err := api.FindRepository(
			ctx,
			accountID,
			importID,
			&code.RepositoryApiFindRepositoryOpts{
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)

		if err != nil {
			return err
		}

		if !repo.Importing {
			return nil
		}

		select {
		case <-time.After(5 * time.Second):
			// Sleep for 5 seconds
		case <-ctx.Done():
			// Context canceled, return with error
			return ctx.Err()
		}
	}
}

func createOnlyFieldDiffFunc(k, oldValue, newValue string, d *schema.ResourceData) bool {
	return d.Id() != ""
}

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identifier": {
			Description: "Identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the repository.",
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
			Description:      "Default branch of the repository (Applicate only for create).",
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
			Description: "Internal ID of the repository.",
			Type:        schema.TypeString,
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
			Description:      "Configuration for importing an existing repository from SCM provider.",
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
						Description: "The full identifier of the repository on the SCM provider's platform.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
}
