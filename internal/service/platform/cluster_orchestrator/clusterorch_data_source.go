package cluster_orchestrator

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceClusterOrchestrator() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness ClusterOrchestrator.",

		ReadContext: resourceClusterOrchestratorRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Orchestrator",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_endpoint": {
				Description: "Endpoint of the k8s cluster being onboarded under the orchestrator",
				Type:        schema.TypeString,
				Required:    true,
			},
			"k8s_connector_id": {
				Description: "ID of the Harness Kubernetes Connector Being used",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Description: "Region of the k8s cluster",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	return resource
}

func resourceClusterOrchestratorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	Identifier := d.Id()

	resp, httpResp, err := c.CloudCostClusterOrchestratorApi.ClusterOrchestratorDetails(ctx, c.AccountId, Identifier)

	if err != nil {
		if isClusterOrchestratorNotFound(err, httpResp) {
			clearClusterOrchestratorFromState(d)
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Response == nil {
		clearClusterOrchestratorFromState(d)
		return nil
	}

	setId(d, resp.Response.ID)
	d.Set("name", resp.Response.Name)
	d.Set("k8s_connector_id", resp.Response.K8sConnectorID)
	if resp.Response.UserConfig != nil {
		d.Set("cluster_endpoint", resp.Response.UserConfig.ClusterEndPoint)
		d.Set("region", resp.Response.UserConfig.Region)
	}

	return nil
}

func clearClusterOrchestratorFromState(d *schema.ResourceData) {
	d.SetId("")
	d.MarkNewResource()
}
