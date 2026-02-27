package as_rule_test

// Static entity IDs used by autostopping rule tests (connectors, proxies, etc. in the test account).
//
// Dependent fields to review when changing cloud connectors:
//   - K8s: k8sConnectorID must be a valid K8s connector in the account (used with cloudConnectorIDK8s).
//   - VM:  vmFilterVMID, vmFilterRegion must be valid Azure resource IDs/regions; proxyIDVM must be an Azure proxy for doNotDeleteAzureConnector.
//   - ECS: uses cloudConnectorIDVM (Azure) and proxyIDVM; container cluster/service/region should match the connector.
//   - RDS: cloudConnectorIDRDS, proxyIDRDS; database id and region should exist in the RDS connector account.
//   - Scale group: scaleGroupARN, scaleGroupName, scaleGroupRegion must be a real ASG in the account for cloudConnectorIDScaleGroup (AWS).
const (
	cloudConnectorIDK8s        = "donotdeleteautomationhemanth"
	cloudConnectorIDVM         = "doNotDeleteAzureConnector"
	cloudConnectorIDRDS        = "DoNotDelete_LightwingNonProd"
	cloudConnectorIDScaleGroup = "DoNotDelete_LightwingNonProd"
)

// K8s connector ID for rule_k8s tests
const k8sConnectorID = "account.k8s_connector"

// Proxy IDs used by VM/ECS and RDS rules. Not defined in load_balancer (those tests create proxies);
// use IDs of existing proxies in the test account for doNotDeleteAzureConnector (VM) and DoNotDelete_LightwingNonProd (RDS).
const (
	proxyIDVM  = "ap-chdpf8f83v0c1aj69oog"
	proxyIDRDS = "ap-ciun1635us1fhpjiotfg"
)

// Rule dependency ID used by VM/ECS rules. Must be an existing autostopping rule ID in the account.
const ruleIDDependency = 24576

// VM filter and scale group entities. Aligned with load_balancer where applicable:
// - Azure: same subscription/resource_group pattern and region "eastus2" as azure_proxy_test.
// - AWS scale group: region "us-east-1" matches aws_alb_test; ARN/name must be a real ASG in the account.
const (
	vmFilterVMID     = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Compute/virtualMachines/virtual_machine"
	vmFilterRegion   = "eastus2"
	scaleGroupARN    = "arn:aws:autoscaling:us-east-1:1234:autoScalingGroup:abcd:autoScalingGroupName/demo-asg"
	scaleGroupName   = "demo-asg"
	scaleGroupRegion = "us-east-1"
)
