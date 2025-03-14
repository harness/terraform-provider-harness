package infrastructure

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosInfrastructureService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a chaos infrastructure.",

		ReadContext: dataSourceChaosInfrastructureRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier of the organization in which the chaos infrastructure is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Identifier of the project in which the chaos infrastructure is configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment_id": {
				Description: "Environment identifier of the chaos infrastructure.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	return resource
}

func dataSourceChaosInfrastructureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)
	var accountIdentifier, orgIdentifier, projectIdentifier, identifier, envIdentifier string
	accountIdentifier = c.AccountId
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		identifier = attr.(string)
	}
	if attr, ok := d.GetOk("environment_id"); ok {
		envIdentifier = attr.(string)
	}
	resp, httpResp, err := c.ChaosSdkApi.GetInfraV2(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier, envIdentifier)

	if err != nil {
		if err.Error() == "404 Not Found" {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}
	readChaosInfrastructure(d, resp)

	return nil
}
