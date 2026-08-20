package idp

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceScorecardCheck() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving an IDP scorecard check.",
		ReadContext: dataSourceScorecardCheckRead,
		Schema:      scorecardCheckSchema(true),
	}
}

func dataSourceScorecardCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	id := d.Get("identifier").(string)

	resp, httpResp, err := c.ChecksApi.GetCheck(ctx, id, &idp.ChecksApiGetCheckOpts{
		HarnessAccount: optional.NewString(c.AccountId),
		Custom:         optional.NewBool(true),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	if err := readScorecardCheck(d, resp.CheckDetails); err != nil {
		return diag.Errorf("failed to read IDP scorecard check from data source: %v", err)
	}
	return nil
}
