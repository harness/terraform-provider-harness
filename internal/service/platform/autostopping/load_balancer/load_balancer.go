package load_balancer

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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

func deletionConsent(d *schema.ResourceData) (bool, error) {
	attr, ok := d.GetOk("delete_cloud_resources_on_destroy")
	if !ok {
		return false, fmt.Errorf("delete_cloud_resources_on_destroy attribute should be set for destroying Loadabalancer")
	}
	delConsent, ok := attr.(bool)
	if !ok {
		return false, fmt.Errorf("delete_cloud_resources_on_destroy should be of bool type. Value can be true or false")
	}
	return delConsent, nil
}

func resourceLoadBalancerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	consent, err := deletionConsent(d)
	if err != nil {
		return diag.Diagnostics{
			{
				Severity: diag.Error,
				Summary:  err.Error(),
			},
		}
	}
	httpResp, err := c.CloudCostAutoStoppingLoadBalancersApi.DeleteLoadBalancer(ctx, nextgen.DeleteAccessPointPayload{
		Ids:           []string{d.Id()},
		WithResources: consent,
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

func nonEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}

func azureAppGwValidator(ap nextgen.AccessPoint) error {
	importing := nonEmptyString(ap.Metadata.AppGatewayId)
	if importing {
		return nil
	}
	if !nonEmptyString(ap.Vpc) {
		return fmt.Errorf("vpc is required for creating Azure AppGateway")
	}
	if !nonEmptyString(ap.Metadata.SubnetId) {
		return fmt.Errorf("subnet_id is required for creating Azure AppGateway")
	}
	if !nonEmptyString(ap.Metadata.FeIpId) {
		return fmt.Errorf("frontend_ip is required for creating Azure AppGateway")
	}
	if !nonEmptyString(ap.Metadata.Size) {
		return fmt.Errorf("sku_size is required for creating Azure AppGateway")
	}
	return nil
}

func implementationSpecificLBValidators(type_, kind string) func(ap nextgen.AccessPoint) error {
	impl := fmt.Sprintf("%s:%s", strings.ToLower(type_), strings.ToLower(kind))
	switch impl {
	case "azure:app_gateway":
		return azureAppGwValidator
	default:
		return nil
	}
}

func buildLoadBalancer(d *schema.ResourceData, accountId, type_, kind string) (nextgen.AccessPoint, error) {
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

	if attr, ok := d.GetOk("app_gateway_id"); ok {
		lb.Metadata.AppGatewayId = attr.(string)
	}

	if attr, ok := d.GetOk("azure_func_region"); ok {
		lb.Metadata.FuncRegion = attr.(string)
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

	if attr, ok := d.GetOk("alb_arn"); ok {
		lb.Metadata.AlbArn = attr.(string)
	}

	if attr, ok := d.GetOk("certificate_id"); ok {
		lb.Metadata.CertificateId = attr.(string)
	}

	if attr, ok := d.GetOk("machine_type"); ok {
		lb.Metadata.MachineType = attr.(string)
	}

	if attr, ok := d.GetOk("sku_size"); ok {
		lb.Metadata.Size = attr.(string)
	}

	if attr, ok := d.GetOk("frontend_ip"); ok {
		lb.Metadata.FeIpId = attr.(string)
	}

	if attr, ok := d.GetOk("api_key"); ok {
		lb.Metadata.ApiKey = attr.(string)
	}

	if attr, ok := d.GetOk("keypair"); ok {
		lb.Metadata.Keypair = attr.(string)
	}

	if attr, ok := d.GetOk("resource_group"); ok {
		lb.Metadata.ResourceGroup = attr.(string)
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
	if validateFunc := implementationSpecificLBValidators(type_, kind); validateFunc != nil {
		if err := validateFunc(*lb); err != nil {
			return *lb, err
		}
	}
	return *lb, nil
}
