package as_rule

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceK8sRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping rule for K8s services.",
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
			"dry_run": {
				Description: "Boolean that indicates whether the AutoStopping rule should be created in DryRun mode",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"k8s_connector_id": {
				Description: "Id of the K8s connector",
				Type:        schema.TypeString,
				Required:    true,
			},
			"k8s_namespace": {
				Description: "Namespace of the cluster",
				Type:        schema.TypeString,
				Required:    true,
			},
			"rule_yaml": {
				Description:      "YAML definition of the K8s AutoStopping rule (workload selector, ingress, etc.).",
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
				ValidateFunc:     validateRuleYAML,
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
