package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorPagerDuty() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a PagerDuty connector.",
		ReadContext:   resourceConnectorPagerDutyRead,
		CreateContext: resourceConnectorPagerDutyCreateOrUpdate,
		UpdateContext: resourceConnectorPagerDutyCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"api_token_ref": {
				Description: "Reference to the Harness secret containing the api token." + secret_ref_text,
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

func resourceConnectorPagerDutyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.PagerDuty)
	if err != nil {
		return err
	}

	if err := readConnectorPagerDuty(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorPagerDutyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorPagerDuty(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorPagerDuty(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorPagerDuty(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:     nextgen.ConnectorTypes.PagerDuty,
		PagerDuty: &nextgen.PagerDutyConnectorDto{},
	}

	if attr, ok := d.GetOk("api_token_ref"); ok {
		connector.PagerDuty.ApiTokenRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.PagerDuty.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorPagerDuty(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("api_token_ref", connector.PagerDuty.ApiTokenRef)
	d.Set("delegate_selectors", connector.PagerDuty.DelegateSelectors)

	return nil
}
