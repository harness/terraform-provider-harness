package load_balancer

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAzureProxy() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure autostopping proxy",
		ReadContext:   resourceLoadBalancerRead,
		CreateContext: resourceAzureProxyCreateOrUpdate,
		UpdateContext: resourceAzureProxyCreateOrUpdate,
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
				Description: "Region in which cloud resources are hosted",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subnet_id": {
				Description: "VPC in which cloud resources are hosted",
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
			"allocate_static_ip": {
				Description: "Boolean value to indicate if proxy vm needs to have static IP",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"machine_type": {
				Description: "Machine instance type",
				Type:        schema.TypeString,
				Required:    true,
			},
			"api_key": {
				Description: "Harness NG API key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"keypair": {
				Description: "",
				Type:        schema.TypeString,
				Required:    true,
			},
			"certificate_id": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_secret_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Certificate secret ID",
						},
						"key_secret_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Private key secret ID",
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceLoadBalancerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.CloudCostAutoStoppingLoadBalancersApi.DescribeLoadBalancer(ctx, c.AccountId, id, c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readLoadBalancer(d, resp.Response)

	return nil
}

func resourceAzureProxyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	lb := buildLoadBalancer(d, c.AccountId, "azure", "autostopping_proxy")
	return resourceLoadBalancerCreateOrUpdate(ctx, d, meta, lb)
}

func resourceLoadBalancerCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, lb nextgen.AccessPoint) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.CreateAccessPointResponse
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.CloudCostAutoStoppingLoadBalancersApi.CreateLoadBalancer(ctx, lb, c.AccountId, c.AccountId)
	} else {
		resp, httpResp, err = c.CloudCostAutoStoppingLoadBalancersApi.EditLoadBalancer(ctx, lb, c.AccountId, c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readLoadBalancer(d, resp.Response)

	return nil
}

func resourceLoadBalancerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	httpResp, err := c.CloudCostAutoStoppingLoadBalancersApi.DeleteLoadBalancer(ctx, nextgen.DeleteAccessPointPayload{
		Ids:           []string{d.Id()},
		WithResources: true,
	}, c.AccountId, c.AccountId)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func readLoadBalancer(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	d.SetId(accessPoint.Id)
	d.Set("identifier", accessPoint.Id)
	d.Set("name", accessPoint.Name)
	d.Set("host_name", accessPoint.HostName)
	d.Set("cloud_connector_id", accessPoint.CloudAccountId)
	d.Set("region", accessPoint.Region)
	d.Set("vpc", accessPoint.Vpc)
}

func buildLoadBalancer(d *schema.ResourceData, accountId, type_, kind string) nextgen.AccessPoint {
	lb := &nextgen.AccessPoint{
		Metadata: &nextgen.AccessPointMeta{},
	}
	lb.AccountId = accountId

	if attr, ok := d.GetOk("identifier"); ok {
		lb.Id = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		lb.Name = attr.(string)
	}

	if attr, ok := d.GetOk("host_name"); ok {
		lb.HostName = attr.(string)
	}

	if attr, ok := d.GetOk("cloud_connector_id"); ok {
		lb.CloudAccountId = attr.(string)
	}

	lb.Type_ = type_
	if kind != "" {
		lb.Kind = kind
	}

	if attr, ok := d.GetOk("region"); ok {
		lb.Region = attr.(string)
	}

	if attr, ok := d.GetOk("vpc"); ok {
		lb.Vpc = attr.(string)
	}

	if attr, ok := d.GetOk("resource_group"); ok {
		lb.Metadata.ResourceGroup = attr.(string)
	}

	if attr, ok := d.GetOk("security_groups"); ok {
		groups := make([]string, 0)
		for _, v := range attr.([]interface{}) {
			groups = append(groups, v.(string))
		}
		lb.Metadata.SecurityGroups = groups
	}

	if attr, ok := d.GetOk("subnet_id"); ok {
		if type_ == "gcp" {
			lb.Metadata.SubnetName = attr.(string)
		} else {
			lb.Metadata.SubnetId = attr.(string)
		}
	}

	if attr, ok := d.GetOk("zone"); ok {
		lb.Metadata.Zone = attr.(string)
	}

	if attr, ok := d.GetOk("certificate_id"); ok {
		lb.Metadata.CertificateId = attr.(string)
	}

	if attr, ok := d.GetOk("machine_type"); ok {
		lb.Metadata.MachineType = attr.(string)
	}

	if attr, ok := d.GetOk("api_key"); ok {
		lb.Metadata.ApiKey = attr.(string)
	}

	if attr, ok := d.GetOk("keypair"); ok {
		lb.Metadata.Keypair = attr.(string)
	}

	lb.Metadata.AllocateStaticIp = false
	if attr, ok := d.GetOk("allocate_static_ip"); ok {
		lb.Metadata.AllocateStaticIp = attr.(bool)
	}
	if attr, ok := d.GetOk("certificates"); ok {
		certificates := make([]nextgen.CertificatesData, 0)
		certificateDetails := &nextgen.CertificatesData{}
		config := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := config["cert_secret_id"]; ok {
			certificateDetails.CertSecretId = attr.(string)
		}

		if attr, ok := config["key_secret_id"]; ok {
			certificateDetails.KeySecretId = attr.(string)
		}
		lb.Metadata.Certificates = append(certificates, *certificateDetails)
	}
	if type_ == "aws" {
		if attr, ok := d.GetOk("route53_hosted_zone_id"); ok {
			route53 := &nextgen.AccessPointMetaDnsRoute53{
				HostedZoneId: attr.(string),
			}
			lb.Metadata.Dns = &nextgen.AccessPointMetaDns{
				Route53: route53,
			}
		} else {
			lb.Metadata.Dns = &nextgen.AccessPointMetaDns{
				Others: lb.HostName,
			}
		}
	}
	return *lb
}
