package as_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceECSRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Variable.",
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
			"container": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"task_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
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
