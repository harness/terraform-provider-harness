package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var defaultHostnameConvention = "${host.ec2Instance.privateDnsName.split('\\.')[0]}"
var defaultK8sReleaseName = "release-${infra.kubernetes.infraId}"
var defaultHelmReleaseName = "${infra.kubernetes.infraId}"

var infraDetailTypes = []string{
	"kubernetes",
	"kubernetes_gcp",
	"aws_ssh",
	"aws_ami",
	"aws_ecs",
	"aws_lambda",
	"aws_winrm",
	"azure_vmss",
	"azure_webapp",
	"tanzu",
	"datacenter_winrm",
	"datacenter_ssh",
}

func infraDefSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "The unique id of the infrastructure definition.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "The name of the infrastructure definition",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"app_id": {
			Description: "The id of the application the infrastructure definition belongs to.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"env_id": {
			Description: "The id of the environment the infrastructure definition belongs to.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"cloud_provider_type": {
			Description:  fmt.Sprintf("The type of the cloud provider to connect with. Valid options are %s", strings.Join(cac.CloudProviderTypesSlice, ", ")),
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(cac.CloudProviderTypesSlice, false),
		},
		"deployment_type": {
			Description:  fmt.Sprintf("The type of the deployment to use. Valid options are %s", strings.Join(cac.DeploymenTypesSlice, ", ")),
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(cac.DeploymenTypesSlice, false),
		},
		"provisioner_name": {
			Description: "The name of the infrastructure provisioner to use.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"deployment_template_uri": {
			Description: "The URI of the deployment template to use. Only used if deployment_type is `CUSTOM`.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"scoped_services": {
			Description: "The list of service names to scope this infrastructure definition to.",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"kubernetes": {
			Description: "The configuration details for Kubernetes deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsK8sDirectSchema(),
			ConflictsWith: []string{
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"kubernetes_gcp": {
			Description: "The configuration details for Kubernetes on GCP deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsK8sGcp(),
			ConflictsWith: []string{
				"kubernetes",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"aws_ssh": {
			Description: "The configuration details for AWS SSH deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAwsSSH(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"aws_ami": {
			Description: "The configuration details for Aws AMI deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAwsAmi(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"aws_ecs": {
			Description: "The configuration details for Aws AMI deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAwsEcs(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ami",
				"aws_ssh",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"aws_lambda": {
			Description: "The configuration details for Aws Lambda deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAwsLambda(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ami",
				"aws_ssh",
				"aws_ecs",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"aws_winrm": {
			Description: "The configuration details for AWS WinRM deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAwsWinRM(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ami",
				"aws_ssh",
				"aws_ecs",
				"aws_lambda",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
		},
		"azure_vmss": {
			Description: "The configuration details for Azure VMSS deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAzureVmss(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"azure_webapp": {
			Description: "The configuration details for Azure WebApp deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsAzureWebApp(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"tanzu",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"tanzu": {
			Description: "The configuration details for PCF deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsPcf(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"datacenter_winrm",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"datacenter_winrm": {
			Description: "The configuration details for WinRM datacenter deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsDatacenterWinRM(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_ssh",
			},
			ExactlyOneOf: infraDetailTypes,
		},
		"datacenter_ssh": {
			Description: "The configuration details for SSH datacenter deployments.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem:        infraDetailsDatacenterSSH(),
			ConflictsWith: []string{
				"kubernetes",
				"kubernetes_gcp",
				"aws_ssh",
				"aws_ami",
				"aws_ecs",
				"aws_lambda",
				"aws_winrm",
				"azure_vmss",
				"azure_webapp",
				"tanzu",
				"datacenter_winrm",
			},
			ExactlyOneOf: infraDetailTypes,
		},
	}
}

func infraDetailsK8sDirectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description:  "The namespace in Kubernetes to deploy to.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`([a-z\d])([a-z\d-])`), "namespaces may only contain lowercase letters, numbers, and '-'"),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			"release_name": {
				Description: fmt.Sprintf("The naming convention of the release. When using Helm Native the default is %[1]s. For standard Kubernetes manifests the default is %[2]s", defaultHelmReleaseName, defaultK8sReleaseName),
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsK8sGcp() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_name": {
				Description: "The name of the cluster being deployed to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description: "The namespace in Kubernetes to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"release_name": {
				Description: "The naming convention of the release.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsAwsSSH() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag": {
				Description: "The tags to use when selecting the instances.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "The key of the tag.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "The value of the tag.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"vpc_ids": {
				Description: "The VPC ids to use when selecting the instances.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"desired_capacity": {
				Description: "The desired capacity of the auto scaling group.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host_connection_type": {
				Description: fmt.Sprintf("The type of host connection to use. Valid options are %s", strings.Join(cac.HostConnectionTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"loadbalancer_name": {
				Description: "The name of the load balancer to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname_convention": {
				Description: fmt.Sprintf("The naming convention to use for the hostname. Defaults to %s", defaultHostnameConvention),
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultHostnameConvention,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsAwsEcs() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_public_ip": {
				Description: "Flag to assign a public IP address.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_name": {
				Description: "The name of the ECS cluster to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"execution_role": {
				Description: "The ARN of the role to use for execution.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"launch_type": {
				Description: fmt.Sprintf("The type of launch configuration to use. Valid options are %s", strings.Join(cac.AwsEcsLaunchTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"security_group_ids": {
				Description: "The security group ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"subnet_ids": {
				Description: "The subnet ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Description: "The VPC ids to use when selecting the instances.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func infraDetailsAwsAmi() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ami_deployment_type": {
				Description:  fmt.Sprintf("The ami deployment type to use. Valid options are %s", strings.Join(cac.AmiDeploymentTypesSlice, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cac.AmiDeploymentTypesSlice, false),
			},
			"asg_identifies_workload": {
				Description: "Flag to indicate whether the autoscaling group identifies the workload.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"classic_loadbalancers": {
				Description: "The classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname_convention": {
				Description: fmt.Sprintf("The naming convention to use for the hostname. Defaults to %s", defaultHostnameConvention),
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultHostnameConvention,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"spotinst_cloud_provider_name": {
				Description: "The name of the SpotInst cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"spotinst_config_json": {
				Description: "The SpotInst configuration to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"stage_classic_loadbalancers": {
				Description: "The staging classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"stage_target_group_arns": {
				Description: "The staging classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"target_group_arns": {
				Description: "The ARN's of the target groups.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"use_traffic_shift": {
				Description: "Flag to enable traffic shifting.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func infraDetailsAwsLambda() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"iam_role": {
				Description: "The IAM role to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"security_group_ids": {
				Description: "The security group ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"subnet_ids": {
				Description: "The subnet ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Description: "The VPC ids to use when selecting the instances.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func infraDetailsAwsWinRM() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"desired_capacity": {
				Description: "The desired capacity of the autoscaling group.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_connection_type": {
				Description: fmt.Sprintf("The type of host connection to use. Valid options are %s", strings.Join(cac.HostConnectionTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname_convention": {
				Description: fmt.Sprintf("The naming convention to use for the hostname. Defaults to %s", defaultHostnameConvention),
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultHostnameConvention,
			},
			"loadbalancer_name": {
				Description: "The name of the load balancer to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsAzureVmss() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_name": {
				Description: "Base name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_group_name": {
				Description: "The name of the resource group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "The unique id of the azure subscription.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description: "The username to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"auth_type": {
				Description: fmt.Sprintf("The type of authentication to use. Valid options are %s.", strings.Join(cac.VmssAuthTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"deployment_type": {
				Description: fmt.Sprintf("The type of deployment. Valid options are %s", strings.Join(cac.VmssDeploymentTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsAzureWebApp() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_group": {
				Description: "The name of the resource group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "The unique id of the azure subscription.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsPcf() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"organization": {
				Description: "The PCF organization to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"space": {
				Description: "The PCF space to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsDatacenterWinRM() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostnames": {
				Description: "A list of hosts to deploy to.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"winrm_connection_attributes_name": {
				Description: "The name of the WinRM connection attributes to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func infraDetailsDatacenterSSH() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostnames": {
				Description: "A list of hosts to deploy to.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"host_connection_attributes_name": {
				Description: "The name of the SSH connection attributes to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
