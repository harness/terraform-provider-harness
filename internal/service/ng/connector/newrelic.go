package connector

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getNewRelicSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "NewRelic connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "newrelic"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the NewRelic server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"account_id": {
					Description: "Account ID of the NewRelic account.",
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

func expandNewRelicConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.NewRelic
	connector.NewRelic = &nextgen.NewRelicConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.NewRelic.Url = attr.(string)
	}

	if attr, ok := config["account_id"]; ok {
		connector.NewRelic.NewRelicAccountId = attr.(string)
	}

	if attr, ok := config["api_key_ref"]; ok {
		connector.NewRelic.ApiKeyRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.NewRelic.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenNewRelicConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.NewRelic {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.NewRelic.Url
	results["account_id"] = connector.NewRelic.NewRelicAccountId
	results["api_key_ref"] = connector.NewRelic.ApiKeyRef
	results["delegate_selectors"] = connector.NewRelic.DelegateSelectors

	d.Set("newrelic", []interface{}{results})

	return nil
}
