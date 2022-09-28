package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorPrometheus() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Prometheus connector.",
		ReadContext:   resourceConnectorPrometheusRead,
		CreateContext: resourceConnectorPrometheusCreateOrUpdate,
		UpdateContext: resourceConnectorPrometheusCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
			"user_name": {
				Description: "User name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"password_ref": {
				Description: "Password reference.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"headers": {
				Description: "Headers.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Key.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"encrypted_value_ref": {
							Description: "Encrypted value reference.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "Value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value_encrypted": {
							Description: "Encrypted value.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					}},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorPrometheusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Prometheus)
	if err != nil {
		return err
	}

	if err := readConnectorPrometheus(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorPrometheusCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorPrometheus(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorPrometheus(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorPrometheus(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:      nextgen.ConnectorTypes.Prometheus,
		Prometheus: &nextgen.PrometheusConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Prometheus.Url = attr.(string)
	}

	if attr, ok := d.GetOk("user_name"); ok {
		connector.Prometheus.Username = attr.(string)
	}

	if attr, ok := d.GetOk("password_ref"); ok {
		connector.Prometheus.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("headers"); ok {
		connector.Prometheus.Headers = expandHeaders(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Prometheus.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func expandHeaders(headers []interface{}) []nextgen.CustomHealthKeyAndValue {
	var result []nextgen.CustomHealthKeyAndValue
	for _, header := range headers {
		v := header.(map[string]interface{})

		var resultHeader nextgen.CustomHealthKeyAndValue
		resultHeader.EncryptedValueRef = v["encrypted_value_ref"].(string)
		resultHeader.Key = v["key"].(string)
		resultHeader.Value = v["value"].(string)
		resultHeader.ValueEncrypted = v["value_encrypted"].(bool)
		result = append(result, resultHeader)
	}

	return result
}

func readHeaders(headers []nextgen.CustomHealthKeyAndValue) []interface{} {
	var result []interface{}
	for _, header := range headers {
		result = append(result, map[string]interface{}{
			"encrypted_value_ref": header.EncryptedValueRef,
			"key":                 header.Key,
			"value":               header.Value,
			"value_encrypted":     header.ValueEncrypted,
		})
	}

	return result
}

func readConnectorPrometheus(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Prometheus.Url)
	d.Set("delegate_selectors", connector.Prometheus.DelegateSelectors)
	d.Set("user_name", connector.Prometheus.Username)
	d.Set("password_ref", connector.Prometheus.PasswordRef)
	d.Set("headers", readHeaders(connector.Prometheus.Headers))
	return nil
}
