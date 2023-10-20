package environment

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

func DataSourceEnvironment() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness environment.",

		ReadContext: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"yaml": {
				Description: "Environment YAML." + helpers.Descriptions.YamlText.String(),
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var env *nextgen.EnvironmentResponseDetails
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoEnvironmentResponse
		resp, httpResp, err = c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
		env = resp.Data.Environment
	} else if name != "" {
		env, httpResp, err = c.EnvironmentsApi.GetEnvironmentByName(ctx, c.AccountId, name, nextgen.GetEnvironmentByNameOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if env == nil {
		return nil
	}

	readDataSourceEnvironment(d, env)

	return nil
}

func readDataSourceEnvironment(d *schema.ResourceData, env *nextgen.EnvironmentResponseDetails) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("project_id", env.ProjectIdentifier)
	d.Set("name", env.Name)
	d.Set("color", env.Color)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
	d.Set("type", env.Type_.String())
	d.Set("yaml", env.Yaml)
}
