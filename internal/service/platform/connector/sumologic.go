package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/gitsync"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorSumologic() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Sumologic connector.",
		ReadContext:   resourceConnectorSumologicRead,
		CreateContext: resourceConnectorSumologicCreateOrUpdate,
		UpdateContext: resourceConnectorSumologicCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the SumoLogic server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"access_id_ref": {
				Description: "Reference to the Harness secret containing the access id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"access_key_ref": {
				Description: "Reference to the Harness secret containing the access key.",
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
	gitsync.SetGitSyncSchema(resource.Schema, false)

	return resource
}

func resourceConnectorSumologicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.SumoLogic)
	if err != nil {
		return err
	}

	if err := readConnectorSumologic(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorSumologicCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorSumologic(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorSumologic(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorSumologic(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:     nextgen.ConnectorTypes.SumoLogic,
		SumoLogic: &nextgen.SumoLogicConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.SumoLogic.Url = attr.(string)
	}

	if attr, ok := d.GetOk("access_id_ref"); ok {
		connector.SumoLogic.AccessIdRef = attr.(string)
	}

	if attr, ok := d.GetOk("access_key_ref"); ok {
		connector.SumoLogic.AccessKeyRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.SumoLogic.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorSumologic(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.SumoLogic.Url)
	d.Set("access_id_ref", connector.SumoLogic.AccessIdRef)
	d.Set("access_key_ref", connector.SumoLogic.AccessKeyRef)
	d.Set("delegate_selectors", connector.SumoLogic.DelegateSelectors)

	return nil
}
