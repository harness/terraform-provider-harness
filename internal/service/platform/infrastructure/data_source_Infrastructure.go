package infrastructure

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
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
		},
	}
	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	env_id := d.Get("env_id").(string)

	resp, httpResp, err := c.InfrastructuresApi.GetInfrastructure(ctx, d.Id(), c.AccountId, env_id, &nextgen.InfrastructuresApiGetInfrastructureOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readInfrastructure(d, &resp.Data)

	return nil
}
