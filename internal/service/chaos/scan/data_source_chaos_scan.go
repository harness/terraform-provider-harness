package scan

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosPipelineScan() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for reading an existing Harness Chaos pipeline scan and the risks it produced.",
		ReadContext: dataSourcePipelineScanRead,
		Schema:      dataSourceChaosPipelineScanSchema(),
	}
}

func DataSourceChaosInfrastructureScan() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for reading an existing Harness Chaos infrastructure scan and the risks it produced.",
		ReadContext: dataSourceInfrastructureScanRead,
		Schema:      dataSourceChaosInfrastructureScanSchema(),
	}
}

func dataSourcePipelineScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceChaosScanRead(ctx, d, meta, chaos.ScanTypePipelineExecution)
}

func dataSourceInfrastructureScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceChaosScanRead(ctx, d, meta, chaos.ScanTypeDiscoveryAgent)
}

func dataSourceChaosScanRead(ctx context.Context, d *schema.ResourceData, meta interface{}, scanType string) diag.Diagnostics {
	identity := d.Get("identity").(string)
	d.SetId(identity)
	diags := readScan(ctx, d, meta, identity, scanType, false)
	if !diags.HasError() && d.Id() == "" {
		return diag.Errorf("chaos scan %q was not found", identity)
	}
	return diags
}
