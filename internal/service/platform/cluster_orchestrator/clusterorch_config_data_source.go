package cluster_orchestrator

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceClusterOrchestratorConfig() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data Source for retrieving Harness CCM ClusterOrchestrator Config.",
		ReadContext: resourceClusterOrchestratorConfigRead,

		Schema: map[string]*schema.Schema{
			"orchestrator_id": {
				Description: "ID of the Cluster Orchestrator Object",
				Type:        schema.TypeString,
				Required:    true,
			},
			"distribution": {
				Description: "Spot and Ondemand Distribution Preferences for workload replicas",
				Type:        schema.TypeList,
				Optional:    true,
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
			"commitment_integration": {
				Description: "Commitment integration configuration for Cluster Orchestrator",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Flag to enable Commitment Integration",
						},
						"master_account_id": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "Master AWS account id for commitment integration",
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},
			"replacement_schedule": {
				Description: "Replacement schedule for Cluster Orchestrator",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"window_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Window type for replacement schedule",
							ValidateFunc: validation.StringInSlice([]string{
								string(nextgen.AlwaysReplacementWindow),
								string(nextgen.NeverReplacementWindow),
								string(nextgen.CustomReplacementWindow),
							}, true),
						},
						"applies_to": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Defines the scope of the replacement schedule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"harness_pod_eviction": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"consolidation": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"reverse_fallback": {
										Type:     schema.TypeBool,
										Required: true,
									},
								},
							},
						},
						"window_details": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeList,
										Description: "List of days on which schedule need to be active. Valid values are SUN, MON, TUE, WED, THU, FRI and SAT.",
										Required:    true,
										MinItems:    1,
										MaxItems:    7,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validateDayValueDiag,
										},
									},
									"all_day": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"start_time": {
										Type:             schema.TypeString,
										Optional:         true,
										Description:      "Start time of schedule in the format HH:MM. Eg : 13:15 for 01:15pm",
										ValidateDiagFunc: timeValidation,
									},
									"end_time": {
										Type:             schema.TypeString,
										Optional:         true,
										Description:      "End time of schedule in the format HH:MM. Eg : 13:15 for 01:15pm",
										ValidateDiagFunc: timeValidation,
									},
									"time_zone": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Time zone in which the schedule needs to be executed. Example Valid values: UTC, America/New_York, Europe/London, Asia/Kolkata, Asia/Tokyo, Asia/Hong_Kong, Asia/Singapore, Australia/Melbourne and Australia/Sydney.",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceClusterOrchestratorConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	Identifier := d.Get("orchestrator_id").(string)

	resp, httpResp, err := c.CloudCostClusterOrchestratorApi.ClusterOrchestratorDetails(ctx, c.AccountId, Identifier)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		readClusterOrchConfig(d, resp.Response)
	}

	return nil
}
