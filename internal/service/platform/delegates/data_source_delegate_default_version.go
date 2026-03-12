package delegates

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDelegateDefaultVersion() *schema.Resource {

	return &schema.Resource{
		Description: "Data source for retrieving the latest supported Harness delegate version.",

		ReadContext: dataSourceDelegateDefaultVersionRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Organization identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version": {
				Description: "Latest supported delegate version.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"minimal_version": {
				Description: "Latest supported minimal delegate version.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceDelegateDefaultVersionRead(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {

	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	accountId := c.AccountId

	opts := &nextgen.DelegateSetupResourceApiLatestVersionOpts{}

	if orgId, ok := d.GetOk("org_id"); ok {
		opts.OrgIdentifier = optional.NewString(orgId.(string))
	}
	if projectId, ok := d.GetOk("project_id"); ok {
		opts.ProjectIdentifier = optional.NewString(projectId.(string))
	}

	resp, httpResp, err := c.DelegateSetupResourceApi.GetLatestSupportedDelegateVersion(ctx, accountId, opts)
	if err != nil {
		tflog.Error(ctx, "Error retrieving latest supported delegate version", map[string]interface{}{
			"error": err,
		})
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource.LatestSupportedVersion == "" {
		return diag.Errorf("delegate latest supported version response is empty")
	}

	d.Set("version", resp.Resource.LatestSupportedVersion)
	d.Set("minimal_version", resp.Resource.LatestSupportedMinimalVersion)

	d.SetId(fmt.Sprintf("%s/%s/%s", accountId, d.Get("org_id").(string), d.Get("project_id").(string)))

	return nil
}
