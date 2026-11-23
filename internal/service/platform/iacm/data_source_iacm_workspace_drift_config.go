package iacm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIacmWorkspaceDriftConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving IaCM workspace-level drift detection configuration.",

		ReadContext: dataSourceIacmWorkspaceDriftConfigRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier where the project resides.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier where the workspace resides.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"workspace_id": {
				Description: "Workspace identifier to retrieve drift configuration for.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"provisioner_type": {
				Description: "Provisioner type to retrieve drift configuration for (terraform, opentofu, terragrunt, awscdk).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pipeline_id": {
				Description: "Pipeline configured for drift detection.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"run_every": {
				Description: "Repeat interval — run drift detection every N units.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"run_every_unit": {
				Description: "Unit for the run_every interval. Valid values: seconds (1–59), minutes (1–59), hours (1–24), days (1–31).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"enabled": {
				Description: "Whether workspace-level drift automation is currently enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"trigger_id": {
				Description: "Generated trigger identifier assigned by the backend when drift is enabled.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceIacmWorkspaceDriftConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	org := d.Get("org_id").(string)
	project := d.Get("project_id").(string)
	workspace := d.Get("workspace_id").(string)
	provisioner := d.Get("provisioner_type").(string)

	result, httpResp, err := c.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, c.AccountId, org, project, workspace, provisioner,
	)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return diag.Errorf("workspace drift config not found for org=%s project=%s workspace=%s provisioner=%s", org, project, workspace, provisioner)
		}
		return parseError(err, httpResp)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", org, project, workspace, provisioner))
	if err := d.Set("enabled", result.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("pipeline_id", result.PipelineId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("run_every", int(result.RunEvery)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("run_every_unit", result.RunEveryUnit); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("trigger_id", result.TriggerId); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
