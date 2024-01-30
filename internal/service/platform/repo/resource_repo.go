package repo

import (
	"context"
	"net/http"

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
		UpdateContext: resourceRepoCreateOrUpdate,
		DeleteContext: resourceRepoDelete,
		CreateContext: resourceRepoCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"created_by": {
				Description: "ID of the user who created the repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"created": {
				Description: "Timestamp when the repository was created.",
				Type:        schema.TypeInt,
				Optional:    true,
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
				Optional:    true,
			},
			"git_url": {
				Description: "Git URL of the repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": {
				Description: "ID of the repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"importing": {
				Description: "Whether the repository is being imported.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_public": {
				Description: "Whether the repository is public.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"num_closed_pulls": {
				Description: "Number of closed pull requests.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_forks": {
				Description: "Number of forks.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_merged_pulls": {
				Description: "Number of merged pull requests.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_open_pulls": {
				Description: "Number of open pull requests.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"num_pulls": {
				Description: "Total number of pull requests.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"parent_id": {
				Description: "ID of the parent repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"path": {
				Description: "Path of the repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"size": {
				Description: "Size of the repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size_updated": {
				Description: "Timestamp when the repository size was last updated.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"uid": {
				Description: "UID of the repository.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"updated": {
				Description: "Timestamp when the repository was last updated.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func resourceRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	id := d.Id()

	repo, resp, err := c.RepositoryApi.FindRepository(ctx, id)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRepo(d, &repo)

	return nil
}

func resourceRepoCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	var err error
	var repo code.TypesRepository
	var resp *http.Response
	id := d.Id()

	if id == "" {
		repo, resp, err = c.RepositoryApi.CreateRepository(ctx, &code.RepositoryApiCreateRepositoryOpts{})
		if err != nil {
			return helpers.HandleApiError(err, d, resp)
		}
	} else {
		repo, resp, err = c.RepositoryApi.UpdateRepository(ctx, d.Get("uid").(string), &code.RepositoryApiUpdateRepositoryOpts{})
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

func resourceRepoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	resp, err := c.RepositoryApi.DeleteRepository(ctx, d.Get("uid").(string))
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	return nil
}

func buildRepo(d *schema.ResourceData) *code.TypesRepository {
	return &code.TypesRepository{
		Created:        d.Get("created").(int32),
		CreatedBy:      d.Get("created_by").(int32),
		DefaultBranch:  d.Get("default_branch").(string),
		Description:    d.Get("description").(string),
		ForkId:         d.Get("fork_id").(int32),
		GitUrl:         d.Get("git_url").(string),
		Id:             d.Get("id").(int32),
		Importing:      d.Get("importing").(bool),
		IsPublic:       d.Get("is_public").(bool),
		NumClosedPulls: d.Get("num_closed_pulls").(int32),
		NumForks:       d.Get("num_forks").(int32),
		NumMergedPulls: d.Get("num_merged_pulls").(int32),
		NumOpenPulls:   d.Get("num_open_pulls").(int32),
		NumPulls:       d.Get("num_pulls").(int32),
		ParentId:       d.Get("parent_id").(int32),
		Path:           d.Get("path").(string),
		Size:           d.Get("size").(int32),
		SizeUpdated:    d.Get("size_updated").(int32),
		Uid:            d.Get("uid").(string),
		Updated:        d.Get("updated").(int32),
	}
}

func readRepo(d *schema.ResourceData, resp *code.TypesRepository) {
	d.Set("created", resp.Created)
	d.Set("created_by", resp.CreatedBy)
	d.Set("default_branch", resp.DefaultBranch)
	d.Set("description", resp.Description)
	d.Set("fork_id", resp.ForkId)
	d.Set("git_url", resp.GitUrl)
	d.Set("id", resp.Id)
	d.Set("importing", resp.Importing)
	d.Set("is_public", resp.IsPublic)
	d.Set("num_closed_pulls", resp.NumClosedPulls)
	d.Set("num_forks", resp.NumForks)
	d.Set("num_merged_pulls", resp.NumMergedPulls)
	d.Set("num_open_pulls", resp.NumOpenPulls)
	d.Set("num_pulls", resp.NumPulls)
	d.Set("parent_id", resp.ParentId)
	d.Set("path", resp.Path)
	d.Set("size", resp.Size)
	d.Set("size_updated", resp.SizeUpdated)
	d.Set("uid", resp.Uid)
	d.Set("updated", resp.Updated)
}
