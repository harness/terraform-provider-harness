package ng

import (
	"context"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Resource for creating a Harness project."),

		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		CreateContext: resourceProjectCreate,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// <org_id>/<project_id>
				parts := strings.Split(d.Id(), "/")
				d.Set("org_id", parts[0])
				d.SetId(parts[1])

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"color": {
				Description: "Color of the project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"modules": {
				Description: "Modules in the project.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Description: "Description of the project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags associated with the project.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	orgId := d.Get("org_id").(string)

	resp, _, err := c.NGClient.ProjectApi.GetProject(ctx, id, c.AccountId, &nextgen.ProjectApiGetProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readProject(d, resp.Data.Project)

	return nil
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	project := buildProject(d)

	resp, _, err := c.NGClient.ProjectApi.PostProject(ctx, nextgen.ProjectRequest{Project: project}, c.AccountId, &nextgen.ProjectApiPostProjectOpts{OrgIdentifier: optional.NewString(d.Get("org_id").(string))})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readProject(d, resp.Data.Project)

	return nil
}

func buildProject(d *schema.ResourceData) *nextgen.Project {
	return &nextgen.Project{
		Identifier:    d.Get("identifier").(string),
		OrgIdentifier: d.Get("org_id").(string),
		Name:          d.Get("name").(string),
		Color:         d.Get("color").(string),
		Description:   d.Get("description").(string),
		Modules:       utils.InterfaceSliceToStringSlice(d.Get("modules").(*schema.Set).List()),
		Tags:          utils.ExpandTags(d.Get("tags").(*schema.Set).List()),
	}
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	project := buildProject(d)

	resp, _, err := c.NGClient.ProjectApi.PutProject(ctx, nextgen.ProjectRequest{Project: project}, c.AccountId, project.Identifier, &nextgen.ProjectApiPutProjectOpts{OrgIdentifier: optional.NewString(d.Get("org_id").(string))})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readProject(d, resp.Data.Project)

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	_, _, err := c.NGClient.ProjectApi.DeleteProject(ctx, d.Id(), c.AccountId, &nextgen.ProjectApiDeleteProjectOpts{OrgIdentifier: optional.NewString(d.Get("org_id").(string))})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func readProject(d *schema.ResourceData, project *nextgen.Project) {
	d.SetId(project.Identifier)
	d.Set("identifier", project.Identifier)
	d.Set("org_id", project.OrgIdentifier)
	d.Set("name", project.Name)
	d.Set("color", project.Color)
	d.Set("description", project.Description)
	d.Set("modules", project.Modules)
	d.Set("tags", utils.FlattenTags(project.Tags))
}
