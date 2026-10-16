package scan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosScan() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for reading a Harness Chaos scan and the risks it produced.",

		ReadContext: dataSourceChaosScanRead,
		Schema:      dataSourceChaosScanSchema(),
	}
}

func dataSourceChaosScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	identity := d.Get("identity").(string)
	d.SetId(identity)
	return readScan(ctx, d, meta, identity)
}
