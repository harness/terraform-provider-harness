package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorSplunk() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Splunk connector.",
		ReadContext:   resourceConnectorSplunkRead,
		CreateContext: resourceConnectorSplunkCreateOrUpdate,
		UpdateContext: resourceConnectorSplunkCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorSplunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Splunk)
	if err != nil {
		return err
	}

	if err := readConnectorSplunk(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorSplunkCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorSplunk(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorSplunk(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorSplunk(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.Splunk,
		Splunk: &nextgen.SplunkConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Splunk.SplunkUrl = attr.(string)
	}

	if attr, ok := d.GetOk("account_id"); ok {
		connector.Splunk.AccountId = attr.(string)
	}

	if attr, ok := d.GetOk("username"); ok {
		connector.Splunk.Username = attr.(string)
	}

	if attr, ok := d.GetOk("password_ref"); ok {
		connector.Splunk.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Splunk.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorSplunk(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Splunk.SplunkUrl)
	d.Set("account_id", connector.Splunk.AccountId)
	d.Set("username", connector.Splunk.Username)
	d.Set("password_ref", connector.Splunk.PasswordRef)
	d.Set("delegate_selectors", connector.Splunk.DelegateSelectors)

	return nil
}
