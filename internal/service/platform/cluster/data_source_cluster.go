package cluster

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCluster() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Cluster.",

		ReadContext: dataSourceClusterRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"env_id": {
				Description: "environment identifier of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "org_id of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project_id of the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	return resource
}

func dataSourceClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get("identifier").(string)

	resp, httpResp, err := c.ClustersApi.GetCluster(ctx, id, c.AccountId, id, &nextgen.ClustersApiGetClusterOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readCluster(d, resp.Data)

	return nil
}


