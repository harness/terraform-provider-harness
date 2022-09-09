package environment_group

import (
	"context"
	"errors"

	"github.com/antihax/optional"
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
			"type": {
				Description: "The type of environment group.",
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
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoEnvironmentGroupResponse
		resp, _, err = c.EnvironmentGroupApi.GetEnvironmentGroupV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.EnvironmentGroupApiGetEnvironmentGroupV2Opts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
		env = resp.Data.EnvironmentGroup
	} else if name != "" {
		env, err = c.EnvironmentGroupApi.GetEnvironmentGroupByName(ctx, c.AccountId, name, nextgen.GetEnvironmentGroupByNameOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
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
