package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorAzureCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure Cloud Cost connector in Harness.",
		ReadContext:   resourceConnectorAzureCloudCostRead,
		CreateContext: resourceConnectorAzureCloudCostCreateOrUpdate,
		UpdateContext: resourceConnectorAzureCloudCostCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"features_enabled": {
				Description: "Indicates which feature to enable among Billing, Optimization, Visibility and Governance.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"BILLING", "OPTIMIZATION", "VISIBILITY", "GOVERNANCE"}, false),
				},
			},
			"tenant_id": {
				Description: "Tenant id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "Subsription id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"billing_export_spec": {
				Description: "Returns billing details for the Azure account.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_name": {
							Description: "Name of the storage account.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"container_name": {
							Description: "Name of the container.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"directory_name": {
							Description: "Name of the directory.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"report_name": {
							Description: "Name of the report.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"subscription_id": {
							Description: "Subsription Id.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"billing_type": {
							Description: "Billing type.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"billing_export_spec2": {
				Description: "Returns billing details for the Azure account.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_name": {
							Description: "Name of the storage account.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"container_name": {
							Description: "Name of the container.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"directory_name": {
							Description: "Name of the directory.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"report_name": {
							Description: "Name of the report.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"subscription_id": {
							Description: "Subsription Id.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"billing_type": {
							Description: "Billing type.",
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

func resourceConnectorAzureCloudCostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.CEAzure)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAzureCloudCost(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAzureCloudCostCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAzureCloudCost(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAzureCloudCost(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAzureCloudCost(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:          nextgen.ConnectorTypes.CEAzure,
		AzureCloudCost: &nextgen.CeAzureConnector{},
	}

	if attr, ok := d.GetOk("features_enabled"); ok {
		connector.AzureCloudCost.FeaturesEnabled = helpers.ExpandField(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("tenant_id"); ok {
		connector.AzureCloudCost.TenantId = attr.(string)
	}

	if attr, ok := d.GetOk("subscription_id"); ok {
		connector.AzureCloudCost.SubscriptionId = attr.(string)
	}

	if attr, ok := d.GetOk("billing_export_spec"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.AzureCloudCost.BillingExportSpec = &nextgen.BillingExportSpec{}
		if attr, ok := config["container_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec.ContainerName = attr.(string)
		}

		if attr, ok := config["report_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec.ReportName = attr.(string)
		}

		if attr, ok := config["storage_account_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec.StorageAccountName = attr.(string)
		}

		if attr, ok := config["subscription_id"]; ok {
			connector.AzureCloudCost.BillingExportSpec.SubscriptionId = attr.(string)
		}

		if attr, ok := config["directory_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec.DirectoryName = attr.(string)
		}

		if attr, ok := config["billing_type"]; ok {
			connector.AzureCloudCost.BillingExportSpec.BillingType = attr.(string)
		}
	}
	if attr, ok := d.GetOk("billing_export_spec2"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.AzureCloudCost.BillingExportSpec2 = &nextgen.BillingExportSpec{}
		if attr, ok := config["container_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.ContainerName = attr.(string)
		}

		if attr, ok := config["report_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.ReportName = attr.(string)
		}

		if attr, ok := config["storage_account_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.StorageAccountName = attr.(string)
		}

		if attr, ok := config["subscription_id"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.SubscriptionId = attr.(string)
		}

		if attr, ok := config["directory_name"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.DirectoryName = attr.(string)
		}

		if attr, ok := config["billing_type"]; ok {
			connector.AzureCloudCost.BillingExportSpec2.BillingType = attr.(string)
		}
	}
	return connector
}

func readConnectorAzureCloudCost(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("features_enabled", connector.AzureCloudCost.FeaturesEnabled)
	d.Set("tenant_id", connector.AzureCloudCost.TenantId)
	d.Set("subscription_id", connector.AzureCloudCost.SubscriptionId)
	if isFeatureEnabled("BILLING", connector.AzureCloudCost.FeaturesEnabled) {
		d.Set("billing_export_spec", []interface{}{
			map[string]interface{}{
				"storage_account_name": connector.AzureCloudCost.BillingExportSpec.StorageAccountName,
				"container_name":       connector.AzureCloudCost.BillingExportSpec.ContainerName,
				"directory_name":       connector.AzureCloudCost.BillingExportSpec.DirectoryName,
				"report_name":          connector.AzureCloudCost.BillingExportSpec.ReportName,
				"subscription_id":      connector.AzureCloudCost.BillingExportSpec.SubscriptionId,
				"billing_type":         connector.AzureCloudCost.BillingExportSpec.BillingType,
			},
		})
		d.Set("billing_export_spec2", []interface{}{
			map[string]interface{}{
				"storage_account_name": connector.AzureCloudCost.BillingExportSpec2.StorageAccountName,
				"container_name":       connector.AzureCloudCost.BillingExportSpec2.ContainerName,
				"directory_name":       connector.AzureCloudCost.BillingExportSpec2.DirectoryName,
				"report_name":          connector.AzureCloudCost.BillingExportSpec2.ReportName,
				"subscription_id":      connector.AzureCloudCost.BillingExportSpec2.SubscriptionId,
				"billing_type":         connector.AzureCloudCost.BillingExportSpec2.BillingType,
			},
		})
	}

	return nil
}
