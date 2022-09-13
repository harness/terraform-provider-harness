package environment_group

import (
	"context"
	"errors"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironmentGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness environment group.",

		ReadContext: dataSourceEnvironmentGroupRead,

		Schema: map[string]*schema.Schema{
			"color": {
				Description: "Color of the environment group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceEnvironmentGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var env *nextgen.EnvironmentGroupResponse

	id := d.Get("identifier").(string)

	if id != "" {
		var resp nextgen.ResponseDtoEnvironmentGroup

		orgIdentifier :=     (d.Get("org_id").(string))
		projectIdentifier := (d.Get("project_id").(string))

		resp, _, err = c.EnvironmentGroupApi.GetEnvironmentGroup(ctx, d.Get("identifier").(string), c.AccountId, orgIdentifier, projectIdentifier, &nextgen.EnvironmentGroupApiGetEnvironmentGroupOpts{
			Branch:     helpers.BuildField(d, "brach"),
			RepoIdentifier: helpers.BuildField(d, "repo_id"),
		})
		env = resp.Data.EnvGroup
	} else {
		return diag.FromErr(errors.New("identifier must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if env == nil {
		return nil
	}

	readEnvironmentGroup(d, env)

	return nil
}
