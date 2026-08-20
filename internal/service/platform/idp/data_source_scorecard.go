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

func DataSourceScorecard() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving an IDP scorecard.",
		ReadContext: dataSourceScorecardRead,
		Schema:      scorecardSchema(true),
	}
}

func dataSourceScorecardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	id := d.Get("identifier").(string)

	resp, httpResp, err := c.ScorecardsApi.GetScorecard(ctx, id, &idp.ScorecardsApiGetScorecardOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	if err := readScorecard(d, resp); err != nil {
		return diag.Errorf("failed to read IDP scorecard from data source: %v", err)
	}
	return nil
}
