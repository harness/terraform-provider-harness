package as_rule

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceVMRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Variables.",

		ReadContext:   resourceASRuleRead,
		UpdateContext: resourceVMRuleCreateOrUpdate,
		DeleteContext: resourceASRuleDelete,
		CreateContext: resourceVMRuleCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"name": {
				Description: "Name of the rule",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_connector_id": {
				Description: "Id of the cloud connector",
				Type:        schema.TypeString,
				Required:    true,
			},
			"idle_time_mins": {
				Description: "Idle time in minutes. This is the time that the AutoStopping rule waits before stopping the idle instances.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     15,
			},
			"use_spot": {
				Description: "Boolean that indicates whether the selected instances should be converted to spot vm",
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
			},
			"custom_domains": {
				Description: "Custom URLs used to access the instances",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_ids": {
							Description: "Ids of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Description: "Tags of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"regions": {
							Description: "Regions of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"zones": {
							Description: "Zones of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"http": {
				Description: "Http routing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Id of the proxy",
							Type:        schema.TypeString,
							Required:    true,
						},
						"routing": {
							Description: "Routing configuration used to access the instances",
							Type:        schema.TypeList,
							MinItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_protocol": {
										Description: "Source protocol of the proxy can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"target_protocol": {
										Description: "Target protocol of the instance can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"source_port": {
										Description: "Port on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"target_port": {
										Description: "Port on the VM",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"action": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"health": {
							Description: "Health Check Details",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Description: "Protocol can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"port": {
										Description: "Health check port on the VM",
										Type:        schema.TypeInt,
										Required:    true,
									},
									"path": {
										Description: "API path to use for health check",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"timeout": {
										Description: "Health check timeout",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_from": {
										Description: "Lower limit for acceptable status code",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_to": {
										Description: "Upper limit for acceptable status code",
										Type:        schema.TypeInt,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"tcp": {
				Description: "TCP routing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Id of the Proxy",
							Type:        schema.TypeString,
							Required:    true,
						},
						"ssh": {
							Description: "SSH configuration",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     22,
									},
								},
							},
						},
						"rdp": {
							Description: "RDP configuration",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3389,
									},
								},
							},
						},
						"forward_rule": {
							Description: "Additional tcp forwarding rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"depends": {
				Description: "Dependent rules",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Description: "Rule id of the dependent rule",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"delay_in_sec": {
							Description: "Number of seconds the rule should wait after warming up the dependent rule",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceVMRuleCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	saveServiceRequestV2 := buildASRule(d, Database, c.AccountId)
	return resourceASRuleCreateOrUpdate(ctx, d, meta, saveServiceRequestV2)
}
