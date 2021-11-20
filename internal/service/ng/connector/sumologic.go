package connector

import (
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getSumoLogicSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "SumoLogic connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "sumologic"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the SumoLogic server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"access_id_ref": {
					Description: "Reference to the Harness secret containing the access id.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"access_key_ref": {
					Description: "Reference to the Harness secret containing the access key.",
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

func expandSumoLogicConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.SumoLogic
	connector.SumoLogic = &nextgen.SumoLogicConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.SumoLogic.Url = attr.(string)
	}

	if attr, ok := config["access_id_ref"]; ok {
		connector.SumoLogic.AccessIdRef = attr.(string)
	}

	if attr, ok := config["access_key_ref"]; ok {
		connector.SumoLogic.AccessKeyRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.SumoLogic.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenSumoLogicConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.SumoLogic {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.SumoLogic.Url
	results["access_id_ref"] = connector.SumoLogic.AccessIdRef
	results["access_key_ref"] = connector.SumoLogic.AccessKeyRef
	results["delegate_selectors"] = connector.SumoLogic.DelegateSelectors

	d.Set("sumologic", []interface{}{results})

	return nil
}
