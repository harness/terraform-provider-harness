package repo

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
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
			"id": {
				Description: "ID of the repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"version": {
				Description: "Version of the repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"parent_id": {
				Description: "ID of the parent repository.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"uid": {
				Description: "UID of the repository.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"path": {
				Description: "Path of the repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_public": {
				Description: "Whether the repository is public.",
				Type:        schema.TypeBool,
				Required:    true,
			},
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
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	_, _ = c, id

	return nil
}

func resourceRepoCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	_ = c

	return nil
}

func resourceRepoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	_ = c

	return nil
}

func buildRepo(d *schema.ResourceData) *nextgen.Project {
	return &nextgen.Project{
		Identifier:    d.Get("identifier").(string),
		OrgIdentifier: d.Get("org_id").(string),
		Name:          d.Get("name").(string),
		Color:         d.Get("color").(string),
		Description:   d.Get("description").(string),
		Modules:       utils.InterfaceSliceToStringSlice(d.Get("modules").(*schema.Set).List()),
		Tags:          helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
	}
}

// func readRepo(d *schema.ResourceData, Repo *nextgen.Repo) {
// 	d.SetId(Repo.Identifier)
// 	d.Set("identifier", Repo.Identifier)
// 	d.Set("org_id", Repo.OrgIdentifier)
// 	d.Set("name", Repo.Name)
// 	d.Set("color", Repo.Color)
// 	d.Set("description", Repo.Description)
// 	d.Set("modules", Repo.Modules)
// 	d.Set("tags", helpers.FlattenTags(Repo.Tags))
// }
