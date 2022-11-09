package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorDynatrace() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Dynatrace connector.",
		ReadContext:   resourceConnectorDynatraceRead,
		CreateContext: resourceConnectorDynatraceCreateOrUpdate,
		UpdateContext: resourceConnectorDynatraceCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Dynatrace server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"api_token_ref": {
				Description: "The reference to the Harness secret containing the api token." + secret_ref_text,
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

func resourceConnectorDynatraceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Dynatrace)
	if err != nil {
		return err
	}

	if err := readConnectorDynatrace(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorDynatraceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorDynatrace(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorDynatrace(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorDynatrace(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:     nextgen.ConnectorTypes.Dynatrace,
		Dynatrace: &nextgen.DynatraceConnectorDto{},
	}
	if attr, ok := d.GetOk("url"); ok {
		connector.Dynatrace.Url = attr.(string)
	}

	if attr, ok := d.GetOk("api_token_ref"); ok {
		connector.Dynatrace.ApiTokenRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Dynatrace.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector

}

func readConnectorDynatrace(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Dynatrace.Url)
	d.Set("api_token_ref", connector.Dynatrace.ApiTokenRef)
	d.Set("delegate_selectors", connector.Dynatrace.DelegateSelectors)
	return nil
}
