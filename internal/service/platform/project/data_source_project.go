package project

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProject() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness project.",

		ReadContext: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
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
		},
	}

	helpers.SetOrgLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)
	orgId := d.Get("org_id").(string)

	resp, _, err := c.ProjectApi.GetProject(ctx, id, c.AccountId, &nextgen.ProjectApiGetProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	if err != nil {
		return diag.FromErr(err)
	}

	readProject(d, resp.Data.Project)

	return nil
}
