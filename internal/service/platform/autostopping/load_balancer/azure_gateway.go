package load_balancer

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAzureGateway() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure Application Gateway",
		ReadContext:   resourceLoadBalancerRead,
		CreateContext: resourceAzureGatewayCreateOrUpdate,
		UpdateContext: resourceAzureGatewayCreateOrUpdate,
		DeleteContext: resourceLoadBalancerDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the proxy",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_name": {
				Description: "Hostname for the proxy",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_connector_id": {
				Description: "Id of the cloud connector",
				Type:        schema.TypeString,
				Required:    true,
			},
			"region": {
				Description: "Region in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_group": {
				Description: "Resource group in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"vpc": {
				Description: "VPC in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subnet_id": {
				Description: "Subnet in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"azure_func_region": {
				Description: "Region in which azure cloud function will be provisioned",
				Type:        schema.TypeString,
				Required:    true,
			},
			"frontend_ip": {
				Description: "",
				Type:        schema.TypeString,
				Required:    true,
			},
			"sku_size": {
				Description: "Size of machine used for the gateway",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	return resource
}

func resourceAzureGatewayCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	lb := buildLoadBalancer(d, c.AccountId, "azure", "app_gateway")
	return resourceLoadBalancerCreateOrUpdate(ctx, d, meta, lb)
}
