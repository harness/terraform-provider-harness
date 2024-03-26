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
			"page": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
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
	page := d.Get("page").(int)

	var opt *nextgen.ProjectApiGetProjectListOpts

	if limit, ok := d.GetOk("limit"); ok {
		// Include the limit parameter in the API call
		opt = &nextgen.ProjectApiGetProjectListOpts{
			OrgIdentifier: optional.NewString(orgId),
			PageIndex:     optional.NewInt32(int32(page)),
			PageSize:      optional.NewInt32(int32(limit.(int))),
		}
	} else {
		// Exclude the limit parameter if it's not provided
		opt = &nextgen.ProjectApiGetProjectListOpts{
			OrgIdentifier: optional.NewString(orgId),
			PageIndex:     optional.NewInt32(int32(page)),
		}
	}

	var resp nextgen.ResponseDtoPageResponseProjectResponse
	resp, httpResp, err = c.ProjectApi.GetProjectList(ctx, c.AccountId, opt)
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
