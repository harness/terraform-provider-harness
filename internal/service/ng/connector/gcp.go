package connector

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getGcpSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Gcp connector configuration.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "gcp"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"manual": {
					Description:   "Manual credential configuration.",
					Type:          schema.TypeList,
					MaxItems:      1,
					Optional:      true,
					ConflictsWith: []string{"gcp.0.inherit_from_delegate"},
					AtLeastOneOf: []string{
						"gcp.0.inherit_from_delegate",
						"gcp.0.manual",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"secret_key_ref": {
								Description: "Reference to the Harness secret containing the secret key.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"delegate_selectors": {
								Description: "The delegates to connect with.",
								Type:        schema.TypeSet,
								Required:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"inherit_from_delegate": {
					Type:          schema.TypeList,
					Description:   "Inherit configuration from delegate.",
					Optional:      true,
					ConflictsWith: []string{"gcp.0.manual"},
					AtLeastOneOf: []string{
						"gcp.0.inherit_from_delegate",
						"gcp.0.manual",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"delegate_selectors": {
								Description: "The delegates to inherit the credentials from.",
								Type:        schema.TypeSet,
								Required:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
			},
		},
	}
}

func expandGcpConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Gcp
	connector.Gcp = &nextgen.GcpConnector{
		Credential: &nextgen.GcpConnectorCredential{},
	}

	if attr := config["manual"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Gcp.Credential.Type_ = nextgen.GcpAuthTypes.ManualConfig
		connector.Gcp.Credential.ManualConfig = &nextgen.GcpManualDetails{}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Gcp.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

		if attr := config["secret_key_ref"].(string); attr != "" {
			connector.Gcp.Credential.ManualConfig.SecretKeyRef = attr
		}
	}

	if attr := config["inherit_from_delegate"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Gcp.Credential.Type_ = nextgen.GcpAuthTypes.InheritFromDelegate

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Gcp.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}
}

func flattenGcpConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Gcp {
		return nil
	}

	results := map[string]interface{}{}

	switch connector.Gcp.Credential.Type_ {
	case nextgen.GcpAuthTypes.ManualConfig:
		results["manual"] = []map[string]interface{}{
			{
				"secret_key_ref":     connector.Gcp.Credential.ManualConfig.SecretKeyRef,
				"delegate_selectors": connector.Gcp.DelegateSelectors,
			},
		}
	case nextgen.GcpAuthTypes.InheritFromDelegate:
		results["inherit_from_delegate"] = []map[string]interface{}{
			{
				"delegate_selectors": connector.Gcp.DelegateSelectors,
			},
		}
	default:
		return fmt.Errorf("invalid gcp credential type: %s", connector.Gcp.Credential.Type_)
	}

	d.Set("gcp", []interface{}{results})

	return nil
}
