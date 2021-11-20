package connector

import (
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getJiraSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Jira connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "jira"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the Jira server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"username": {
					Description:   "Username to use for authentication.",
					Type:          schema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"jira.0.username_ref"},
					ExactlyOneOf:  []string{"jira.0.username", "jira.0.username_ref"},
				},
				"username_ref": {
					Description:   "Reference to a secret containing the username to use for authentication.",
					Type:          schema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"jira.0.username"},
					ExactlyOneOf:  []string{"jira.0.username", "jira.0.username_ref"},
				},
				"password_ref": {
					Description: "Reference to a secret containing the password to use for authentication.",
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

func expandJiraConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Jira
	connector.Jira = &nextgen.JiraConnector{}

	if attr, ok := config["url"]; ok {
		connector.Jira.JiraUrl = attr.(string)
	}

	if attr, ok := config["username"]; ok {
		connector.Jira.Username = attr.(string)
	}

	if attr, ok := config["username_ref"]; ok {
		connector.Jira.UsernameRef = attr.(string)
	}

	if attr, ok := config["password_ref"]; ok {
		connector.Jira.PasswordRef = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Jira.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenJiraConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Jira {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Jira.JiraUrl
	results["username"] = connector.Jira.Username
	results["username_ref"] = connector.Jira.UsernameRef
	results["password_ref"] = connector.Jira.PasswordRef
	results["delegate_selectors"] = connector.Jira.DelegateSelectors

	d.Set("jira", []interface{}{results})

	return nil
}
