package connector

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getSplunkSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Splunk connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "splunk"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the Splunk server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"username": {
					Description: "The username used for connecting to Splunk.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"account_id": {
					Description: "Splunk account id.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"password_ref": {
					Description: "The reference to the Harness secret containing the Splunk password.",
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

func expandSplunkConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Splunk
	connector.Splunk = &nextgen.SplunkConnector{}

	if attr, ok := config["url"]; ok {
		connector.Splunk.SplunkUrl = attr.(string)
	}

	if attr, ok := config["account_id"]; ok {
		connector.Splunk.AccountId = attr.(string)
	}

	if attr, ok := config["username"]; ok {
		connector.Splunk.Username = attr.(string)
	}

	if attr, ok := config["password_ref"]; ok {
		connector.Splunk.PasswordRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Splunk.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenSplunkConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Splunk {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Splunk.SplunkUrl
	results["account_id"] = connector.Splunk.AccountId
	results["username"] = connector.Splunk.Username
	results["password_ref"] = connector.Splunk.PasswordRef
	results["delegate_selectors"] = connector.Splunk.DelegateSelectors

	d.Set("splunk", []interface{}{results})

	return nil
}
