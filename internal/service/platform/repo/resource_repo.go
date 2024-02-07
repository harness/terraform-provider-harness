package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RepoBody struct {
	DefaultBranch string `json:"default_branch"`
	Description   string `json:"description"`
	ForkID        int    `json:"fork_id"`
	GitIgnore     string `json:"git_ignore"`
	Identifier    string `json:"identifier"`
	IsPublic      bool   `json:"is_public"`
	License       string `json:"license"`
	ParentRef     string `json:"parent_ref"`
	Readme        bool   `json:"readme"`
}

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

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	accountID := d.Get("account_id").(string)
	path := d.Get("path").(string)

	repo, resp, err := c.RepositoryApi.FindRepository(
		ctx,
		accountID,
		path,
		&code.RepositoryApiFindRepositoryOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		},
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRepo(d, &repo)

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

	accountID := d.Get("account_id").(string)
	path := d.Get("path").(string)
	id := d.Id()

	bodyJson, err := json.Marshal(buildRepoBody(d))
	if err != nil {
		fmt.Println("Error:", err)
		return helpers.HandleApiError(err, d, nil)
	}

	body := optional.NewInterface(bodyJson)
	orgId := d.Get("org_identifier").(string)
	prjId := d.Get("project_identifier").(string)

	if id == "" {
		repo, resp, err = c.RepositoryApi.CreateRepository(
			ctx, accountID,
			&code.RepositoryApiCreateRepositoryOpts{
				Body:              body,
				OrgIdentifier:     optional.NewString(orgId),
				ProjectIdentifier: optional.NewString(prjId),
			},
		)
		if err != nil {
			return helpers.HandleApiError(err, d, resp)
		}
	} else {
		repo, resp, err = c.RepositoryApi.UpdateRepository(
			ctx,
			accountID,
			path,
			&code.RepositoryApiUpdateRepositoryOpts{
				Body:              optional.NewInterface(bodyJson),
				OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
				ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
			},
		)
		if err != nil {
			return helpers.HandleApiError(err, d, resp)
		}
	}

	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepo(d, &repo)

	return nil
}

func resourceRepoDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	createdBy := strconv.Itoa(d.Get("created_by").(int))
	path := d.Get("path").(string)

	resp, err := c.RepositoryApi.DeleteRepository(
		ctx,
		createdBy,
		path, &code.RepositoryApiDeleteRepositoryOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	return nil
}

func buildRepoBody(d *schema.ResourceData) *RepoBody {
	return &RepoBody{
		DefaultBranch: d.Get("default_branch").(string),
		Description:   d.Get("description").(string),
		ForkID:        d.Get("fork_id").(int),
		GitIgnore:     d.Get("git_ignore").(string),
		Identifier:    d.Get("identifier").(string),
		IsPublic:      d.Get("is_public").(bool),
		License:       d.Get("license").(string),
		ParentRef:     d.Get("parent_ref").(string),
		Readme:        d.Get("readme").(bool),
	}
}

func readRepo(d *schema.ResourceData, repo *code.TypesRepository) {
	d.SetId(repo.Identifier)

	d.Set("created", repo.Created)
	d.Set("created_by", repo.CreatedBy)
	d.Set("default_branch", repo.DefaultBranch)
	d.Set("description", repo.Description)
	d.Set("fork_id", repo.ForkId)
	d.Set("git_url", repo.GitUrl)
	d.Set("id", repo.Id)
	d.Set("importing", repo.Importing)
	d.Set("is_public", repo.IsPublic)
	d.Set("num_closed_pulls", repo.NumClosedPulls)
	d.Set("num_forks", repo.NumForks)
	d.Set("num_merged_pulls", repo.NumMergedPulls)
	d.Set("num_open_pulls", repo.NumOpenPulls)
	d.Set("num_pulls", repo.NumPulls)
	d.Set("parent_id", repo.ParentId)
	d.Set("path", repo.Path)
	d.Set("size", repo.Size)
	d.Set("size_updated", repo.SizeUpdated)
	d.Set("updated", repo.Updated)
}

func createSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{
		"account_id": {
			Description: "ID of the account who created the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"org_identifier": {
			Description: "Identifier of the organization.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"project_identifier": {
			Description: "Identifier of the project.",
			Type:        schema.TypeString,
			Optional:    true,
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
			Description: "Default branch of the repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"description": {
			Description: "Description of the repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fork_id": {
			Description: "ID of the forked repository.",
			Type:        schema.TypeInt,
			Computed:    true,
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
		"is_public": {
			Description: "Whether the repository is public.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"num_closed_pulls": {
			Description: "Number of closed pull requests.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"num_forks": {
			Description: "Number of forks.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"num_merged_pulls": {
			Description: "Number of merged pull requests.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"num_open_pulls": {
			Description: "Number of open pull requests.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"num_pulls": {
			Description: "Total number of pull requests.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"parent_id": {
			Description: "ID of the parent repository.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"path": {
			Description: "Path of the repository.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"size": {
			Description: "Size of the repository.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"size_updated": {
			Description: "Timestamp when the repository size was last updated.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated": {
			Description: "Timestamp when the repository was last updated.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"git_ignore": {
			Description: "Gitignore file for the repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"identifier": {
			Description: "Identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"license": {
			Description: "License for the repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"parent_ref": {
			Description: "Reference to the parent repository.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"readme": {
			Description: "Whether the repository has a readme.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}
