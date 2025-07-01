package cluster_orchestrator

import (
	"context"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceClusterOrchestratorConfig() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for ClusterOrchestrator Config.",

		CreateContext: resourceClusterOrchestratorConfigCreateOrUpdate,
		UpdateContext: resourceClusterOrchestratorConfigCreateOrUpdate,
		ReadContext:   resourceClusterOrchestratorConfigRead,
		DeleteContext: resourceClusterOrchestratorConfigDelete,

		Schema: map[string]*schema.Schema{
			"orchestrator_id": {
				Description: "ID of the Cluster Orchestrator Object",
				Type:        schema.TypeString,
				Required:    true,
			},
			"distribution": {
				Description: "Spot and Ondemand Distribution Preferences for workload replicas",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base_ondemand_capacity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of minimum ondemand replicas required for workloads",
						},
						"ondemand_replica_percentage": {
							Type:        schema.TypeFloat,
							Required:    true,
							Description: "Percentage of on-demand replicas required for workloads",
						},
						"selector": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Selector for choosing workloads for distribution",
						},
						"strategy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Strategy for choosing spot nodes for cluster",
						},
					},
				},
			},
			"binpacking": {
				Description: "Binpacking preferences for Cluster Orchestrator",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pod_eviction": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Harness Pod Evictor Configuration",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Minimum Threshold for considering a node as underutilized",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cpu": {
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "CPU percentage for considering a node as underutilized",
												},
												"memory": {
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "Memory percentage for considering a node as underutilized",
												},
											},
										},
									},
								},
							},
						},
						"disruption": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Harness disruption configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"criteria": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Criteria for considering a nodes for disruption",
									},
									"delay": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Deletion delay",
									},
									"budget": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Budgets for disruption",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"reasons": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Reasons for disruption",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"nodes": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Number or percentage of Nodes to consider for disruption",
												},
												"schedule": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Schedule for disruption budget",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Frequency for disruption budget",
															},
															"duration": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Duration for disruption budget",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"node_preferences": {
				Description: "Node preferences for Cluster Orchestrator",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "TTL for nodes",
						},
						"reverse_fallback_interval": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Reverse fallback interval",
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceClusterOrchestratorConfigCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	body := buildClusterOrchConfig(d)
	var err error
	var resp nextgen.UpdateClusterOrchestratorConfigResponse
	var httpResp *http.Response
	orchID := d.Get("orchestrator_id").(string)
	resp, httpResp, err = c.CloudCostClusterOrchestratorApi.UpdateClusterOrchestratorConfig(ctx, c.AccountId, orchID, body)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Success != true {
		return diag.Errorf(strings.Join(resp.Errors, ","))
	}
	d.SetId(orchID)
	return nil
}

func resourceClusterOrchestratorConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
