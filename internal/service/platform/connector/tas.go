package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorTas() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Tas in Harness.",
		ReadContext:   resourceConnectorTasRead,
		CreateContext: resourceConnectorTasCreateOrUpdate,
		UpdateContext: resourceConnectorTasCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"credentials": {
				Description: "Contains Tas connector credentials.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description:  "Type can be ManualConfig.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"ManualConfig"}, false),
						},
						"tas_manual_details": {
							Description: "Authenticate to Tas using manual details.",
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_url": {
										Description: "URL of the Tas server.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.tas_manual_details.0.username_ref"},
										ExactlyOneOf:  []string{"credentials.0.tas_manual_details.0.username", "credentials.0.tas_manual_details.0.username_ref"},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.tas_manual_details.0.username"},
										ExactlyOneOf:  []string{"credentials.0.tas_manual_details.0.username", "credentials.0.tas_manual_details.0.username_ref"},
									},
									"password_ref": {
										Description: "Reference of the secret for the password." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"reference_token": {
										Description: "Reference token for authentication.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorTasRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Tas)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorTas(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorTasCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorTas(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorTas(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorTas(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Tas,
		Tas: &nextgen.TasConnector{
			ConnectorType: nextgen.ConnectorTypes.Tas.String(),
			Credential:    &nextgen.TasCredential{},
		},
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		if attr, ok := config["type"]; ok {
			connector.Tas.Credential.Type_ = nextgen.TasCredentialType(attr.(string))

			if attr.(string) == nextgen.TasCredentialTypes.ManualConfig.String() {
				if attr, ok := config["tas_manual_details"]; ok {
					configCredentials := attr.([]interface{})[0].(map[string]interface{})

					connector.Tas.Credential.TasManualDetails = &nextgen.TasManualDetails{}
					if attr, ok := configCredentials["endpoint_url"]; ok {
						connector.Tas.Credential.TasManualDetails.EndpointUrl = attr.(string)
					}

					if attr, ok := configCredentials["username"]; ok {
						connector.Tas.Credential.TasManualDetails.Username = attr.(string)
					}

					if attr, ok := configCredentials["username_ref"]; ok {
						connector.Tas.Credential.TasManualDetails.UsernameRef = attr.(string)
					}

					if attr, ok := configCredentials["password_ref"]; ok {
						connector.Tas.Credential.TasManualDetails.PasswordRef = attr.(string)
					}
					if attr, ok := configCredentials["reference_token"]; ok {
						connector.Tas.Credential.TasManualDetails.ReferenceToken = attr.(string)
					}
				}
			}
		}
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		delegate_selectors := attr.(*schema.Set).List()
		if len(delegate_selectors) > 0 {
			connector.Tas.DelegateSelectors = utils.InterfaceSliceToStringSlice(delegate_selectors)
		}
	}

	return connector
}

func readConnectorTas(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("credentials", []interface{}{
		map[string]interface{}{
			"type":               connector.Tas.Credential.Type_,
			"tas_manual_details": readTasManualDetails(connector),
		},
	})
	d.Set("delegate_selectors", connector.Tas.DelegateSelectors)

	return nil
}

func readTasManualDetails(connector *nextgen.ConnectorInfo) []map[string]interface{} {
	var spec []map[string]interface{}
	switch connector.Tas.Credential.Type_ {
	case nextgen.TasCredentialTypes.ManualConfig:
		spec = []map[string]interface{}{
			{
				"endpoint_url":    connector.Tas.Credential.TasManualDetails.EndpointUrl,
				"username":        connector.Tas.Credential.TasManualDetails.Username,
				"username_ref":    connector.Tas.Credential.TasManualDetails.UsernameRef,
				"password_ref":    connector.Tas.Credential.TasManualDetails.PasswordRef,
				"reference_token": connector.Tas.Credential.TasManualDetails.ReferenceToken,
			},
		}
	}

	return spec
}
