package infrastructure

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfrastructure() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Infrastructure.",

		ReadContext: dataSourceInfrastructureRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the Infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"env_id": {
				Description: "environment identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: fmt.Sprintf("Type of Infrastructure. Valid values are %s.", strings.Join(nextgen.InfrastructureTypeValues, ", ")),
				Type:        schema.TypeString,
				Computed:    true,
			},
			"yaml": {
				Description: "Infrastructure YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"deployment_type": {
				Description: fmt.Sprintf("Infrastructure deployment type. Valid values are %s.", strings.Join(nextgen.InfrastructureDeploymentypeValues, ", ")),
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
							Description: "Load environment yaml from fallback branch",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"repo_name": {
							Description: "Repo name of remote environment",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"load_from_cache": {
							Description: "If the Entity is to be fetched from cache",
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

	// overwrite schema for tags
	if s, ok := resource.Schema["tags"]; ok {
		s.Computed = true
	}

	return resource
}

func dataSourceInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)

	resp, httpResp, err := c.InfrastructuresApi.GetInfrastructure(ctx, d.Get("identifier").(string), c.AccountId, env_id, &nextgen.InfrastructuresApiGetInfrastructureOpts{
		OrgIdentifier:          helpers.BuildField(d, "org_id"),
		ProjectIdentifier:      helpers.BuildField(d, "project_id"),
		RepoName:               helpers.BuildField(d, "git_details.0.repo_name"),
		Branch:                 helpers.BuildField(d, "git_details.0.branch"),
		LoadFromFallbackBranch: helpers.BuildFieldBool(d, "git_details.0.load_from_fallback_branch"),
		LoadFromCache:          helpers.BuildField(d, "git_details.0.load_from_cache"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readInfrastructure(d, resp.Data)

	return nil
}
