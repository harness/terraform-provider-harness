package connector

import (
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getPrometheusSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Prometheus connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "prometheus"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the Prometheus server.",
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

func expandPrometheusConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Prometheus
	connector.Prometheus = &nextgen.PrometheusConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.Prometheus.Url = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Prometheus.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenPrometheusConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Prometheus {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Prometheus.Url
	results["delegate_selectors"] = connector.Prometheus.DelegateSelectors

	d.Set("prometheus", []interface{}{results})

	return nil
}
