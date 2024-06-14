package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAzureGateway() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for AWS Autostopping proxy",
		ReadContext: resourceLoadBalancerRead,
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
				Optional:    true,
			},
			"azure_func_region": {
				Description: "Region in which azure cloud function will be provisioned",
				Type:        schema.TypeString,
				Required:    true,
			},
			"frontend_ip": {
				Description: "ID of IP address to be used. Required only for creating new AppGateway. See https://learn.microsoft.com/en-us/azure/application-gateway/application-gateway-components#static-versus-dynamic-public-ip-address for more details",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sku_size": {
				Description: "Size of machine used for the gateway",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"app_gateway_id": {
				Description: "ID of Azure AppGateway for importing",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"certificate_id": {
				Description: "ID of existing SSL certificate from AppGateway being imported. Required only for SSL based rules",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delete_cloud_resources_on_destroy": {
				Description: "Governs how the loadabalancer entity will be deleted on Terraform destroy. When set to true, the associated Application Gateway will be deleted permanently from Azure account. Be fully aware of the consequneces of settting this to true, as the action is irreversible. When set to false, solely the Harness LB representation will be deleted, leaving the cloud resources intact.",
				Type:        schema.TypeBool,
				Required:    true,
			},
		},
	}

	return resource
}
