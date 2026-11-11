package loadtest

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosLoadTest() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for reading an existing Harness load test definition. It can read a REFERENCE import and a JMeter test. The resource cannot manage either of those.",
		ReadContext: dataSourceChaosLoadTestRead,
		Schema:      dataSourceChaosLoadTestSchema(),
	}
}

func dataSourceChaosLoadTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	identity := d.Get("identifier").(string)
	d.SetId(identity)
	diags := readChaosLoadTest(ctx, d, meta, false, false)
	if !diags.HasError() && d.Id() == "" {
		return diag.Errorf("chaos load test %q was not found", identity)
	}
	return diags
}
