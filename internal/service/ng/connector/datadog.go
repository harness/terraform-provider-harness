package connector

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getDatadogSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Datadog connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "datadog"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the Datadog server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"application_key_ref": {
					Description: "Reference to the Harness secret containing the application key.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"api_key_ref": {
					Description: "Reference to the Harness secret containing the api key.",
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

func expandDatadogConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Datadog
	connector.Datadog = &nextgen.DatadogConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.Datadog.Url = attr.(string)
	}

	if attr, ok := config["application_key_ref"]; ok {
		connector.Datadog.ApplicationKeyRef = attr.(string)
	}

	if attr, ok := config["api_key_ref"]; ok {
		connector.Datadog.ApiKeyRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Datadog.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenDatadogConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Datadog {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Datadog.Url
	results["application_key_ref"] = connector.Datadog.ApplicationKeyRef
	results["api_key_ref"] = connector.Datadog.ApiKeyRef
	results["delegate_selectors"] = connector.Datadog.DelegateSelectors

	d.Set("datadog", []interface{}{results})

	return nil
}
