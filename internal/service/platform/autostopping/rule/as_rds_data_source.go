package as_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRDSRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness AutoStopping rule for RDS databases.",
		ReadContext: resourceASRuleRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"depends": {
				Description: "Dependent rules",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Description: "Rule id of the dependent rule",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"delay_in_sec": {
							Description: "Number of seconds the rule should wait after warming up the dependent rule",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}
