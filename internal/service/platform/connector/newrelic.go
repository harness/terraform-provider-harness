package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorNewRelic() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a New Relic connector.",
		ReadContext:   resourceConnectorNewRelicRead,
		CreateContext: resourceConnectorNewRelicCreateOrUpdate,
		UpdateContext: resourceConnectorNewRelicCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the NewRelic server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Account ID of the NewRelic account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"api_key_ref": {
				Description: "Reference to the Harness secret containing the api key.",
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

func resourceConnectorNewRelicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.NewRelic)
	if err != nil {
		return err
	}

	if err := readConnectorNewRelic(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorNewRelicCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorNewRelic(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorNewRelic(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorNewRelic(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:    nextgen.ConnectorTypes.NewRelic,
		NewRelic: &nextgen.NewRelicConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.NewRelic.Url = attr.(string)
	}

	if attr, ok := d.GetOk("account_id"); ok {
		connector.NewRelic.NewRelicAccountId = attr.(string)
	}

	if attr, ok := d.GetOk("api_key_ref"); ok {
		connector.NewRelic.ApiKeyRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.NewRelic.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorNewRelic(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.NewRelic.Url)
	d.Set("account_id", connector.NewRelic.NewRelicAccountId)
	d.Set("api_key_ref", connector.NewRelic.ApiKeyRef)
	d.Set("delegate_selectors", connector.NewRelic.DelegateSelectors)

	return nil
}
