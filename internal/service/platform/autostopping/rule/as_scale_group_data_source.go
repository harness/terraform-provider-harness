package as_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceScaleGroupRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping rule for Scaling Groups.",

		ReadContext: resourceASRuleRead,

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
			"custom_domains": {
				Description: "Custom URLs used to access the instances",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dry_run": {
				Description: "Boolean that indicates whether the AutoStopping rule should be created in DryRun mode",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"scale_group": {
				Description: "Scaling Group configuration",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the Scaling Group",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name of the Scaling Group",
							Type:        schema.TypeString,
							Required:    true,
						},
						"region": {
							Description: "Region of the Scaling Group",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"zone": {
							Description: "Zone of the Scaling Group. Needed for GCP only",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"desired": {
							Description: "Desired capacity of the Scaling Group",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"min": {
							Description: "Minimum capacity of the Scaling Group",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"max": {
							Description: "Maximum capacity of the Scaling Group",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"on_demand": {
							Description: "On-demand capacity of the Scaling Group",
							Type:        schema.TypeInt,
							Required:    true,
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
							Description: "Routing configuration used to access the scaling group",
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
										Description: "Action to take for the routing rule",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"path": {
										Description: "Path to use for the proxy",
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
