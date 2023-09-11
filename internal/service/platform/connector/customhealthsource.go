package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorCustomHealthSource() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Custom Health source connector.",
		ReadContext:   resourceConnectorCustomHealthSourceRead,
		CreateContext: resourceConnectorCustomHealthSourceCreateOrUpdate,
		UpdateContext: resourceConnectorCustomHealthSourceCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Custom Healthsource controller.",
				Type:        schema.TypeString,
				Required:    true,
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
							Description: "Reference to the Harness secret containing the encrypted value." + secret_ref_text,
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
			"params": {
				Description: "Parameters",
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
							Description: "Reference to the Harness secret containing the encrypted value." + secret_ref_text,
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
			"method": {
				Description: "HTTP Verb Method for the API Call",
				Type:        schema.TypeString,
				Required:    true,
			},
			"validation_body": {
				Description: "Body to be sent with the API Call",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"validation_path": {
				Description: "Path to be added to the base URL for the API Call",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorCustomHealthSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.CustomHealth)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorCustomHealthSource(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorCustomHealthSourceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorCustomHealthSource(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorCustomHealthSource(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorCustomHealthSource(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:        nextgen.ConnectorTypes.CustomHealth,
		CustomHealth: &nextgen.CustomHealthConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.CustomHealth.BaseURL = attr.(string)
	}

	if attr, ok := d.GetOk("method"); ok {
		connector.CustomHealth.Method = attr.(string)
	}

	if attr, ok := d.GetOk("validation_body"); ok {
		connector.CustomHealth.ValidationBody = attr.(string)
	}

	if attr, ok := d.GetOk("validation_path"); ok {
		connector.CustomHealth.ValidationPath = attr.(string)
	}

	if attr, ok := d.GetOk("headers"); ok {
		connector.CustomHealth.Headers = expandHeaders(attr.(*schema.Set).List())
	}
	if attr, ok := d.GetOk("params"); ok {
		connector.CustomHealth.Params = expandHeaders(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.CustomHealth.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorCustomHealthSource(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.CustomHealth.BaseURL)
	d.Set("delegate_selectors", connector.CustomHealth.DelegateSelectors)
	d.Set("method", connector.CustomHealth.Method)
	d.Set("validation_body", connector.CustomHealth.ValidationBody)
	d.Set("validation_path", connector.CustomHealth.ValidationPath)
	d.Set("headers", readHeaders(connector.CustomHealth.Headers))
	d.Set("params", readHeaders(connector.CustomHealth.Params))
	return nil
}
