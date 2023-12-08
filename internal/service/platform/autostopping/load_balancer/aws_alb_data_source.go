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
			"alb_arn": {
				Description: "Arn of AWS ALB to be imported. Required only for importing existing ALB",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"delete_cloud_resources_on_destroy": {
				Description: "Governs how the loadabalancer entity will be deleted on Terraform destroy. When set to true, the associated ALB will be deleted permanently from AWS account. Be fully aware of the consequneces of settting this to true, as the action is irreversible. When set to false, solely the Harness LB representation will be deleted, leaving the cloud resources intact.",
				Type:        schema.TypeBool,
				Required:    true,
			},
		},
	}
	return resource
}
