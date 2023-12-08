package load_balancer

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAwsALB() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an AWS application load balancer",
		ReadContext:   resourceLoadBalancerRead,
		CreateContext: resourceAwsALBCreateOrUpdate,
		UpdateContext: resourceAwsALBCreateOrUpdate,
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

func resourceAwsALBCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	lb, err := buildLoadBalancer(d, c.AccountId, "aws", "")
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceLoadBalancerCreateOrUpdate(ctx, d, meta, lb)
}
