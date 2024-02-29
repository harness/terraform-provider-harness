package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness service.",

		ReadContext: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"git_details": {
				Description: "Contains parameters related to Git Experience for remote entities",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Description: "Name of the branch.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"load_from_fallback_branch": {
							Description: "Load service yaml from fallback branch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"repo_name": {
							Description: "Repo name of remote service",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var svc *nextgen.ServiceResponseDetails
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoServiceResponse
		resp, httpResp, err = c.ServicesApi.GetServiceV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
			OrgIdentifier:          helpers.BuildField(d, "org_id"),
			ProjectIdentifier:      helpers.BuildField(d, "project_id"),
			RepoName:               helpers.BuildField(d, "repo_name"),
			Branch:                 helpers.BuildField(d, "branch"),
			LoadFromFallbackBranch: helpers.BuildFieldBool(d, "load_from_fallback_branch"),
		})
		svc = resp.Data.Service
	} else if name != "" {
		svc, httpResp, err = c.ServicesApi.GetServiceByName(ctx, c.AccountId, name, nextgen.GetServiceByNameOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readService(d, svc)

	return nil
}
