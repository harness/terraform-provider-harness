package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorGCPCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a GCP Cloud Cost connector in Harness.",
		ReadContext:   resourceConnectorGCPCloudCostRead,
		CreateContext: resourceConnectorGCPCloudCostCreateOrUpdate,
		UpdateContext: resourceConnectorGCPCloudCostCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"features_enabled": {
				Description: "Indicates which features to enable among Billing, Optimization, and Visibility.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"BILLING", "OPTIMIZATION", "VISIBILITY"}, false),
				},
			},
			"gcp_project_id": {
				Description: "GCP Project Id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service_account_email": {
				Description: "Email corresponding to the Service Account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"billing_export_spec": {
				Description: "Returns billing details.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_set_id": {
							Description: "Data Set Id.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"table_id": {
							Description: "Table Id.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorGCPCloudCostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.GcpCloudCost)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorGCPCloudCost(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGCPCloudCostCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGcpCloudCost(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGCPCloudCost(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGcpCloudCost(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:        nextgen.ConnectorTypes.GcpCloudCost,
		GcpCloudCost: &nextgen.GcpCloudCostConnectorDto{},
	}

	if attr, ok := d.GetOk("features_enabled"); ok {
		connector.GcpCloudCost.FeaturesEnabled = helpers.ExpandField(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("gcp_project_id"); ok {
		connector.GcpCloudCost.ProjectId = attr.(string)
	}

	if attr, ok := d.GetOk("service_account_email"); ok {
		connector.GcpCloudCost.ServiceAccountEmail = attr.(string)
	}

	if attr, ok := d.GetOk("billing_export_spec"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.GcpCloudCost.BillingExportSpec = &nextgen.GcpBillingExportSpecDto{}
		if attr, ok := config["data_set_id"]; ok {
			connector.GcpCloudCost.BillingExportSpec.DatasetId = attr.(string)
		}

		if attr, ok := config["table_id"]; ok {
			connector.GcpCloudCost.BillingExportSpec.TableId = attr.(string)
		}
	}

	return connector
}

func readConnectorGCPCloudCost(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("features_enabled", connector.GcpCloudCost.FeaturesEnabled)
	d.Set("gcp_project_id", connector.GcpCloudCost.ProjectId)
	d.Set("service_account_email", connector.GcpCloudCost.ServiceAccountEmail)
	d.Set("billing_export_spec", []interface{}{
		map[string]interface{}{
			"data_set_id": connector.GcpCloudCost.BillingExportSpec.DatasetId,
			"table_id":    connector.GcpCloudCost.BillingExportSpec.TableId,
		},
	})

	return nil
}
