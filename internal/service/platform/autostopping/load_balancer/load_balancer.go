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

const (
	KindALB               = "alb"
	KindAppGateway        = "app_gateway"
	KindAutostoppingProxy = "autostopping_proxy"
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
	consent, _ := d.GetOk("delete_cloud_resources_on_destroy")
	delConsent, ok := consent.(bool)
	return ok && delConsent, nil
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
	d.Set("cloud_connector_id", accessPoint.CloudAccountId)
	d.Set("region", accessPoint.Region)
	d.Set("vpc", accessPoint.Vpc)
	kind := accessPoint.Kind
	switch kind {
	case KindALB:
		setALBData(d, accessPoint)
	case KindAppGateway:
		setAppGatewayData(d, accessPoint)
	case KindAutostoppingProxy:
		setProxyData(d, accessPoint)
	}
}

func setALBData(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	setSecurityGroups(accessPoint, d)
	setCertificateId(d, accessPoint)
	albArn := ""
	if accessPoint.Metadata != nil && accessPoint.Metadata.AlbArn != "" {
		albArn = accessPoint.Metadata.AlbArn
	}
	d.Set("alb_arn", albArn)
}

func setAppGatewayData(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	setCertificateId(d, accessPoint)
}

func setProxyData(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	setSecurityGroups(accessPoint, d)
	setProxyCertificates(d, accessPoint)
}

func setCertificateId(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	certificateId := ""
	if accessPoint.Metadata != nil && accessPoint.Metadata.CertificateId != "" {
		certificateId = accessPoint.Metadata.CertificateId
	}
	// Always set the certificate_id to the value to consider case where certificate is removed
	d.Set("certificate_id", certificateId)
}

func setSecurityGroups(accessPoint *nextgen.AccessPoint, d *schema.ResourceData) {
	securityGroups := make([]string, 0)
	if accessPoint.Metadata != nil && len(accessPoint.Metadata.SecurityGroups) > 0 {
		securityGroups = accessPoint.Metadata.SecurityGroups
	}
	// Always set the security_groups to the value to consider case where security groups are removed
	d.Set("security_groups", securityGroups)
}

func setProxyCertificates(d *schema.ResourceData, accessPoint *nextgen.AccessPoint) {
	certificates := make([]map[string]interface{}, 0)
	if accessPoint.Metadata != nil && len(accessPoint.Metadata.Certificates) > 0 {
		for _, cert := range accessPoint.Metadata.Certificates {
			certMap := map[string]interface{}{
				"cert_secret_id": cert.CertSecretId,
				"key_secret_id":  cert.KeySecretId,
			}
			certificates = append(certificates, certMap)
		}
	}
	// Always set the certificates to the value to consider case where certificates are removed
	d.Set("certificates", certificates)
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
	if validateFunc := implementationSpecificLBValidators(type_, kind); validateFunc != nil {
		if err := validateFunc(*lb); err != nil {
			return *lb, err
		}
	}
	return *lb, nil
}
