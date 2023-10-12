package load_balancer

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGCPProxy() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an GCP Autostopping proxy",
		ReadContext:   resourceLoadBalancerRead,
		CreateContext: resourceGCPProxyCreateOrUpdate,
		UpdateContext: resourceGCPProxyCreateOrUpdate,
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
			"zone": {
				Description: "Zone in which cloud resources are hosted",
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
				Sensitive:   true,
				Type:        schema.TypeString,
				Required:    true,
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

func resourceGCPProxyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	lb, err := buildLoadBalancer(d, c.AccountId, "gcp", "autostopping_proxy")
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceLoadBalancerCreateOrUpdate(ctx, d, meta, lb)
}
