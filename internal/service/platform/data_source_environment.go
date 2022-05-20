package platform

import (
	"context"

	"github.com/antihax/optional"
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
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, _, err := c.EnvironmentsApi.GetEnvironmentV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.EnvironmentsApiGetEnvironmentV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil || resp.Data.Environment == nil {
		return nil
	}

	readEnvironment(d, resp.Data.Environment)

	return nil
}
