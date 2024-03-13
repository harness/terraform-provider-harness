package project

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProjectList() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness project.",

		ReadContext: dataSourceProjectListRead,

		Schema: map[string]*schema.Schema{
			"projects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}

	helpers.SetOrgLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceProjectListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	orgId := d.Get("org_id").(string)

	var err error
	var httpResp *http.Response

	var resp nextgen.ResponseDtoPageResponseProjectResponse
	resp, httpResp, err = c.ProjectApi.GetProjectList(ctx, c.AccountId, &nextgen.ProjectApiGetProjectListOpts{OrgIdentifier: optional.NewString(orgId)})
	var output []nextgen.ProjectResponse = resp.Data.Content
	var projects []map[string]interface{}
	for _, v := range output {
		newProject := map[string]interface{}{
			"identifier": v.Project.Identifier,
			"name":       v.Project.Name,
		}

		projects = append(projects, newProject)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if projects == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	d.SetId(resp.CorrelationId)
	d.Set("projects", projects)

	return nil
}
