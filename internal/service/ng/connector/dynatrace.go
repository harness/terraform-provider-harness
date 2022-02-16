package connector

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getDynatraceSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Dynatrace connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "dynatrace"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the Dynatrace server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"api_token_ref": {
					Description: "The reference to the Harness secret containing the api token.",
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

func expandDynatraceConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Dynatrace
	connector.Dynatrace = &nextgen.DynatraceConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.Dynatrace.Url = attr.(string)
	}

	if attr, ok := config["api_token_ref"]; ok {
		connector.Dynatrace.ApiTokenRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Dynatrace.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenDynatraceConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Dynatrace {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Dynatrace.Url
	results["api_token_ref"] = connector.Dynatrace.ApiTokenRef
	results["delegate_selectors"] = connector.Dynatrace.DelegateSelectors

	d.Set("dynatrace", []interface{}{results})

	return nil
}
