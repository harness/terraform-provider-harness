package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorKubernetesCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Kubernetes Cloud Cost connector.",
		ReadContext:   resourceConnectorKubernetesCloudCostRead,
		CreateContext: resourceConnectorKubernetesCloudCostCreateOrUpdate,
		UpdateContext: resourceConnectorKubernetesCloudCostCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"connector_ref": {
				Description: "Referenve of the Connector.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"features_enabled": {
				Description: "Which feature to enable among BILLING, OPTIMIZATION, VISIBILITY",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"BILLING", "OPTIMIZATION", "VISIBILITY"}, false),
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorKubernetesCloudCostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.CEK8sCluster)
	if err != nil {
		return err
	}

	if err := readConnectorKubernetesCloudCost(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorKubernetesCloudCostCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorKubernetesCloudCost(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorKubernetesCloudCost(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorKubernetesCloudCost(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:               nextgen.ConnectorTypes.CEK8sCluster,
		K8sClusterCloudCost: &nextgen.CeKubernetesClusterConfigDto{},
	}

	if attr, ok := d.GetOk("features_enabled"); ok {
		connector.K8sClusterCloudCost.FeaturesEnabled = helpers.ExpandField(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("connector_ref"); ok {
		connector.K8sClusterCloudCost.ConnectorRef = attr.(string)
	}

	return connector
}

func readConnectorKubernetesCloudCost(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("connector_ref", connector.K8sClusterCloudCost.ConnectorRef)
	d.Set("features_enabled", connector.K8sClusterCloudCost.FeaturesEnabled)

	return nil
}
