package ng

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving a Harness project."),

		ReadContext: dataSourceProjectRead,

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
				Computed:    true,
			},
			"color": {
				Description: "Color of the project.",
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed:    true,
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

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)

	resp, _, err := c.NGClient.ProjectApi.GetProject(ctx, id, c.AccountId, &nextgen.ProjectApiGetProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	if err != nil {
		return diag.FromErr(err)
	}

	readProject(d, resp.Data.Project)

	return nil
}
