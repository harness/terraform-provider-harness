package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorPdc() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Pdc connector.",
		ReadContext:   resourceConnectorPdcRead,
		CreateContext: resourceConnectorPdcCreateOrUpdate,
		UpdateContext: resourceConnectorPdcCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"host": {
				Description: "Host of the Physical data centers.",
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Description: "hostname",
							Type:        schema.TypeString,
							Required:    true,
						},
						"attributes": {
							Description: "attributes for current host",
							Type:        schema.TypeMap,
							Optional:    true,
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
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorPdcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Pdc)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorPdc(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorPdcCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorPdc(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorPdc(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorPdc(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Pdc,
		Pdc:   &nextgen.PhysicalDataCenterConnectorDto{},
		// Pdc: &nextgen.PdcConnector{},
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Pdc.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	connector.Pdc.Hosts = buildConnectorPdcHosts(d)

	return connector
}

func readConnectorPdc(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("delegate_selectors", connector.Pdc.DelegateSelectors)

	var hosts []map[string]interface{}
	for _, x := range connector.Pdc.Hosts {
		itemAsMap := make(map[string]interface{})
		itemAsMap["hostname"] = x.Hostname
		itemAsMap["attributes"] = x.HostAttributes

		hosts = append(hosts, itemAsMap)
	}

	d.Set("host", hosts)

	return nil
}

func buildConnectorPdcHosts(d *schema.ResourceData) []nextgen.HostDto {
	if attr, ok := d.GetOk("host"); ok {
		listOfHosts := attr.(*schema.Set).List()

		var pdcHosts []nextgen.HostDto

		for _, x := range listOfHosts {
			itemAsMap := x.(map[string]interface{})
			attributes := itemAsMap["attributes"]

			var hostAttributes map[string]string
			if attributes != nil {
				hostAttributes = convertMapStrInterfaceToStrStr(attributes.(map[string]interface{}))
			}

			pdcHost := nextgen.HostDto{
				Hostname:       itemAsMap["hostname"].(string),
				HostAttributes: hostAttributes,
			}
			pdcHosts = append(pdcHosts, pdcHost)
		}

		return pdcHosts
	}

	return nil
}

func convertMapStrInterfaceToStrStr(m map[string]interface{}) map[string]string {
	if m == nil {
		return nil
	}

	m2 := make(map[string]string)

	for key, value := range m {
		switch value := value.(type) {
		case string:
			m2[key] = value
		}
	}

	return m2
}
