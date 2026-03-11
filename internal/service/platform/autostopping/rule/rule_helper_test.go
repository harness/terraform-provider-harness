package as_rule_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// randAlnum returns a short random string with no underscore or space (for resource names).
func randAlnum(n int) string {
	s := strings.ToLower(utils.RandStringBytes(n))
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

// extractAttr captures a Terraform state attribute into a *string variable
// so it can be used in subsequent test steps (e.g. in PreConfig closures).
func extractAttr(resourceName, key string, dest *string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(resourceName, key, func(value string) error {
		*dest = value
		return nil
	})
}

// waitForProxyReady polls the proxy status until it reaches a terminal state.
// Returns nil if the proxy reaches "created"; returns an error if it reaches
// "errored" or the timeout expires.
func waitForProxyReady(proxyID string, timeout time.Duration) error {
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		resp, _, err := c.CloudCostAutoStoppingLoadBalancersApi.DescribeLoadBalancer(ctx, c.AccountId, proxyID, c.AccountId)
		if err != nil {
			return fmt.Errorf("reading proxy %s: %w", proxyID, err)
		}
		if resp.Response == nil {
			return fmt.Errorf("proxy %s not found", proxyID)
		}
		switch resp.Response.Status {
		case "created":
			return nil
		case "errored":
			return fmt.Errorf("proxy %s provisioning failed", proxyID)
		}
		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("timeout waiting %v for proxy %s to provision", timeout, proxyID)
}

// Static entity IDs used by autostopping rule tests (connectors, proxies, etc. in the test account).
//
// Dependent fields to review when changing cloud connectors:
//   - K8s: k8sConnectorID must be a valid K8s connector; uses cloudConnectorIDAWS.
//   - VM:  vmFilterVMID, vmFilterRegion must be valid Azure resource IDs/regions; proxyIDVM must be an Azure proxy.
//   - ECS, RDS, Scale group, AWS proxy: all use cloudConnectorIDAWS.
const (
	cloudConnectorIDAWS = "DoNotDelete_LightwingNonProd"
	cloudConnectorIDVM  = "automation_azure_connector"
)

// AWS VPC and security group for dynamically created proxies (ECS/RDS rule tests).
// These must exist in the AWS account behind cloudConnectorIDAWS (357919113896).
const (
	awsProxyVPC = "vpc-08f63488a1e3c1bf1"
	awsProxySG  = "sg-0e1f9ee9896d96583"
)

// K8s connector ID for rule_k8s tests
const k8sConnectorID = "account.k8s_connector"

// Azure networking for dynamically created proxies (VM rule tests).
// These must exist in the Azure subscription behind cloudConnectorIDVM.
const (
	azureProxyResourceGroup = "ccm-terraform-rg"
	azureProxyVNet          = "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/virtualNetworks/DoNotDelete-Terraform-VNet"
	azureProxySubnet        = "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/virtualNetworks/DoNotDelete-Terraform-VNet/subnets/default"
	azureProxyNSG           = "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourceGroups/ccm-terraform-rg/providers/Microsoft.Network/networkSecurityGroups/DoNotDelete-Terraform-NSG"
)

// Rule dependency ID used by VM/ECS rules. Must be an existing autostopping rule ID in the account.
const ruleIDDependency = 24576

// VM filter and scale group entities.
// - Azure: VM ID, region must match resources in cloudConnectorIDVM subscription.
// - AWS scale group: region "us-east-1" matches aws_alb_test; ARN/name must be a real ASG in the account.
const (
	vmFilterVMID     = "/subscriptions/20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0/resourcegroups/ccm-terraform-rg/providers/Microsoft.Compute/virtualMachines/DoNotDelete-Terraform-AS-Test-VM"
	vmFilterRegion   = "eastus"
	scaleGroupARN    = "arn:aws:autoscaling:us-east-1:357919113896:autoScalingGroup:c4a7b124-8366-40bb-beed-0c3f1ecab3e2:autoScalingGroupName/DoNotDelete-Terrform-AS-Test-ASG"
	scaleGroupName   = "DoNotDelete-Terrform-AS-Test-ASG"
	scaleGroupRegion = "us-east-1"
)
