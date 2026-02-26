package cluster_orchestrator

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	dayIndex = map[string]time.Weekday{
		"SUN": 0,
		"MON": 1,
		"TUE": 2,
		"WED": 3,
		"THU": 4,
		"FRI": 5,
		"SAT": 6,
	}

	reverseDayIndex = map[time.Weekday]string{
		0: "SUN",
		1: "MON",
		2: "TUE",
		3: "WED",
		4: "THU",
		5: "FRI",
		6: "SAT",
	}
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
			"disabled": {
				Description: "Whether the cluster orchestrator is disabled",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
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
							Description: "Strategy for choosing spot nodes for cluster. Allowed values: CostOptimized, LeastInterrupted",
							ValidateFunc: validation.StringInSlice([]string{"CostOptimized", "LeastInterrupted"}, false),
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
						"enable_spot_to_spot": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Enable spot-to-spot consolidation",
						},
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

func validateDayValueDiag(val interface{}, path cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics
	day, valid := val.(string)
	if !valid {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Day should be string value. Valid values are SUN, MON, TUE, WED, THU, FRI and SAT",
		})
		return diags
	}
	if _, ok := dayIndex[day]; !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Valid values are SUN, MON, TUE, WED, THU, FRI and SAT",
		})
		return diags
	}
	return nil
}

func timeValidation(timeVal interface{}, p cty.Path) diag.Diagnostics {
	diags := diag.Diagnostics{}
	if timeVal == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Non empty value is mandatory",
		})
		return diags
	}
	v, ok := timeVal.(string)
	if !ok || v == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Non empty value is mandatory",
		})
		return diags
	}
	v = strings.TrimSpace(v)
	parts := strings.Split(v, ":")
	if len(parts) != 2 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	hh, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	if hh < 0 || hh > 24 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Hour value should be between 0 and 24",
		})
	}
	mm, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value should be in HH:MM format. Eg : 13:15 for 01:15pm",
		})
		return diags
	}
	if mm < 0 || mm > 59 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Minute value should be between 0 and 59",
		})
	}
	return diags
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

	if !resp.Success {
		return diag.Errorf("%s", fmt.Sprintf("update failed: %s", strings.Join(resp.Errors, ",")))
	}

	disabled := d.Get("disabled").(bool)
	if diags := updateClusterOrchestratorStatus(ctx, c, orchID, disabled); diags.HasError() {
		return diags
	}

	d.SetId(orchID)
	return nil
}

func updateClusterOrchestratorStatus(ctx context.Context, c *nextgen.APIClient, orchID string, disabled bool) diag.Diagnostics {
	_, httpResp, err := c.CloudCostClusterOrchestratorApi.ToggleClusterOrchestratorState(ctx, c.AccountId, orchID, disabled)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}
	return nil
}

func resourceClusterOrchestratorConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
