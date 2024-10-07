package cluster_orchestrator

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceClusterOrchestrator() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating ClusterOrchestrators.",

		CreateContext: resourceClusterOrchestratorCreate,
		UpdateContext: resourceClusterOrchestratorCreate,
		ReadContext:   resourceClusterOrchestratorCreate,
		DeleteContext: resourceClusterOrchestratorDelete,

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
		},
	}

	return resource
}

func resourceClusterOrchestratorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	body := buildClusterOrch(d)
	var err error
	var resp nextgen.ClusterOrchestratorResponse
	var httpResp *http.Response

	resp, httpResp, err = c.CloudCostClusterOrchestratorApi.CreateClusterOrchestrator(ctx, c.AccountId, body)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		setId(d, resp.Response.ID)
	}

	return nil
}

func resourceClusterOrchestratorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
