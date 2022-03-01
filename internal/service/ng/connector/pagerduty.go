package connector

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getPagerDutySchema() *schema.Schema {
	return &schema.Schema{
		Description:   "PagerDuty connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "pagerduty"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"api_token_ref": {
					Description: "Reference to the Harness secret containing the api token.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"delegate_selectors": {
					Description: "Connect using only the delegates which have these tags.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func expandPagerDutyConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.PagerDuty
	connector.PagerDuty = &nextgen.PagerDutyConnectorDto{}

	if attr, ok := config["api_token_ref"]; ok {
		connector.PagerDuty.ApiTokenRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.PagerDuty.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenPagerDutyConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.PagerDuty {
		return nil
	}

	results := map[string]interface{}{}

	results["api_token_ref"] = connector.PagerDuty.ApiTokenRef
	results["delegate_selectors"] = connector.PagerDuty.DelegateSelectors

	d.Set("pagerduty", []interface{}{results})

	return nil
}
