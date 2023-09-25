package as_rule

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceECSRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Variables.",

		ReadContext:   resourceASRuleRead,
		CreateContext: resourceECSRuleCreateOrUpdate,
		UpdateContext: resourceECSRuleCreateOrUpdate,
		DeleteContext: resourceASRuleDelete,
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
							Type:        schema.TypeString,
							Description: "Name of cluster in which service belong to",
							Required:    true,
						},
						"service": {
							Type:        schema.TypeString,
							Description: "Name of service to be onboarded",
							Required:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "Region of cluster",
							Required:    true,
						},
						"task_count": {
							Type:        schema.TypeInt,
							Description: "Desired number of tasks on warming up a rule",
							Optional:    true,
							Default:     1,
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

func resourceECSRuleCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	saveServiceRequestV2 := buildASRule(d, ECS, c.AccountId)
	return resourceASRuleCreateOrUpdate(ctx, d, meta, saveServiceRequestV2)
}
