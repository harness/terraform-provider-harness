package load_balancer

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAwsALB() *schema.Resource {
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
			"security_groups": {
				Description: "Security Group to define the security rules that determine the inbound and outbound traffic",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpc": {
				Description: "VPC in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"certificate_id": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"route53_hosted_zone_id": {
				Description: "Route 53 hosted zone id",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
	return resource
}
