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

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Prometheus.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorPrometheus(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Prometheus.Url)
	d.Set("delegate_selectors", connector.Prometheus.DelegateSelectors)

	return nil
}
