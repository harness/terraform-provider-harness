package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorDatadog() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Datadog connector.",
		ReadContext:   resourceConnectorDatadogRead,
		CreateContext: resourceConnectorDatadogCreateOrUpdate,
		UpdateContext: resourceConnectorDatadogCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Datadog server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"application_key_ref": {
				Description: "Reference to the Harness secret containing the application key." + secret_ref_text,
				Type:        schema.TypeString,
				Required:    true,
			},
			"api_key_ref": {
				Description: "Reference to the Harness secret containing the api key." + secret_ref_text,
				Type:        schema.TypeString,
				Required:    true,
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

func resourceConnectorDatadogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Datadog)
	if err != nil {
		return err
	}

	if err := readConnectorDatadog(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorDatadogCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorDatadog(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorDatadog(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorDatadog(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:   nextgen.ConnectorTypes.Datadog,
		Datadog: &nextgen.DatadogConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Datadog.Url = attr.(string)
	}

	if attr, ok := d.GetOk("application_key_ref"); ok {
		connector.Datadog.ApplicationKeyRef = attr.(string)
	}

	if attr, ok := d.GetOk("api_key_ref"); ok {
		connector.Datadog.ApiKeyRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Datadog.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorDatadog(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Datadog.Url)
	d.Set("application_key_ref", connector.Datadog.ApplicationKeyRef)
	d.Set("api_key_ref", connector.Datadog.ApiKeyRef)
	d.Set("delegate_selectors", connector.Datadog.DelegateSelectors)

	return nil
}
