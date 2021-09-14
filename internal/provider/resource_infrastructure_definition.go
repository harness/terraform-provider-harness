package provider

import (
	"context"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInfraDefinition() *schema.Resource {

	infraDefSchema := infraDefSchema()

	return &schema.Resource{
		Description:   "Resource for creating am infrastructure definition",
		CreateContext: resourceInfraDefinitionCreateOrUpdate,
		ReadContext:   resourceInfraDefinitionRead,
		UpdateContext: resourceInfraDefinitionCreateOrUpdate,
		DeleteContext: resourceInfraDefinitionDelete,
		Schema:        infraDefSchema,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				// <app_id>/<env_id>/<id>
				parts := strings.Split(d.Id(), "/")

				d.Set("app_id", parts[0])
				d.Set("env_id", parts[1])
				d.SetId(parts[2])

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceInfraDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)
	envId := d.Get("env_id").(string)

	infraDef, err := c.ConfigAsCode().GetInfraDefinitionById(appId, envId, id)
	if err != nil {
		return diag.FromErr(err)
	} else if infraDef == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readInfraDefinition(d, infraDef)

	return nil
}

func resourceInfraDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)
	envId := d.Get("env_id").(string)

	err := c.ConfigAsCode().DeleteInfraDefinition(appId, envId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceInfraDefinitionCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.InfrastructureDefinition
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.InfrastructureDefinition).(*cac.InfrastructureDefinition)
	} else {
		id := d.Get("id").(string)
		appId := d.Get("app_id").(string)
		envId := d.Get("env_id").(string)
		input, err = c.ConfigAsCode().GetInfraDefinitionById(appId, envId, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if attr := d.Get("app_id"); attr != "" {
		input.ApplicationId = attr.(string)
	}

	if attr := d.Get("env_id"); attr != "" {
		input.EnvironmentId = attr.(string)
	}

	if attr := d.Get("name"); attr != "" {
		input.Name = attr.(string)
	}

	if attr := d.Get("cloud_provider_type"); attr != "" {
		input.CloudProviderType = cac.CloudProviderType(attr.(string))
	}

	if attr := d.Get("deployment_type"); attr != "" {
		input.DeploymentType = cac.DeploymentType(attr.(string))
	}

	if attr := d.Get("provisioner_name"); attr != "" {
		input.Provisioner = attr.(string)
	}

	if attr := d.Get("deployment_template_uri"); attr != "" {
		input.DeploymentTypeTemplateUri = attr.(string)
	}

	expandScopedServices(d.Get("scoped_services").(*schema.Set).List(), input)
	expandKubernetesConfiguration(d.Get("kubernetes").([]interface{}), input)
	expandKubernetesGcpConfiguration(d.Get("kubernetes_gcp").([]interface{}), input)
	expandAwsSSHConfiguration(d.Get("aws_ssh").([]interface{}), input)
	expandAwsAmiConfiguration(d.Get("aws_ami").([]interface{}), input)
	expandAwsEcsConfiguration(d.Get("aws_ecs").([]interface{}), input)
	expandAwsLambdaConfiguration(d.Get("aws_lambda").([]interface{}), input)
	expandAwsWinRMConfiguration(d.Get("aws_winrm").([]interface{}), input)
	expandTanzuConfiguration(d.Get("tanzu").([]interface{}), input)
	expandAzureWebAppConfiguration(d.Get("azure_webapp").([]interface{}), input)

	infraDef, err := c.ConfigAsCode().UpsertInfraDefinition(input)
	if err != nil {
		return diag.FromErr(err)
	}

	readInfraDefinition(d, infraDef)

	return nil
}

func readInfraDefinition(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) {
	d.SetId(infraDef.Id)
	d.Set("app_id", infraDef.ApplicationId)
	d.Set("env_id", infraDef.EnvironmentId)
	d.Set("name", infraDef.Name)
	d.Set("cloud_provider_type", infraDef.CloudProviderType)
	d.Set("deployment_type", infraDef.DeploymentType)
	d.Set("provisioner_name", infraDef.Provisioner)
	d.Set("deployment_template_uri", infraDef.DeploymentTypeTemplateUri)

	if services := flattenScopedServices(d, infraDef); len(services) > 0 {
		d.Set("scoped_services", services)
	}

	if config := flattenKubernetesConfiguration(d, infraDef); len(config) > 0 {
		d.Set("kubernetes", config)
	}

	if config := flattenKubernetesGcpConfiguration(d, infraDef); len(config) > 0 {
		d.Set("kubernetes_gcp", config)
	}

	if config := flattenAwsSSHConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ssh", config)
	}

	if config := flattenAwsAmiConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ami", config)
	}

	if config := flattenAwsEcsConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ecs", config)
	}

	if config := flattenAwsLambdaConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_lambda", config)
	}

	if config := flattenAwsWinRMConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_winrm", config)
	}

	if config := flattenTanzuConfiguration(d, infraDef); len(config) > 0 {
		d.Set("tanzu", config)
	}

	if config := flattenAzureWebAppConfiguration(d, infraDef); len(config) > 0 {
		d.Set("azure_webapp", config)
	}

}

func flattenScopedServices(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	for _, v := range infraDef.ScopedServices {
		results = append(results, v)
	}

	return results
}

func expandScopedServices(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	results := []string{}

	for _, v := range d {
		results = append(results, v.(string))
	}

	infraDef.ScopedServices = results
}

func flattenKubernetesConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.KubernetesDirect {
		return results
	}

	detailConfig := map[string]interface{}{}
	k8sDetail := detail.ToKubernetesDirect()

	detailConfig["cloud_provider_name"] = k8sDetail.CloudProviderName
	detailConfig["namespace"] = k8sDetail.Namespace
	detailConfig["release_name"] = k8sDetail.ReleaseName

	return append(results, detailConfig)
}

func expandKubernetesConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := cac.InfrastructureKubernetesDirect{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["namespace"]; attr != "" {
		details.Namespace = attr.(string)
	}

	if attr := config["release_name"]; attr != "" {
		details.ReleaseName = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenKubernetesGcpConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.KubernetesGcp {
		return results
	}

	detailConfig := map[string]interface{}{}
	k8sDetail := detail.ToKubernetesGcp()

	detailConfig["cloud_provider_name"] = k8sDetail.CloudProviderName
	detailConfig["namespace"] = k8sDetail.Namespace
	detailConfig["cluster_name"] = k8sDetail.ClusterName
	detailConfig["release_name"] = k8sDetail.ReleaseName

	return append(results, detailConfig)
}

func expandKubernetesGcpConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := cac.InfrastructureKubernetesGcp{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["namespace"]; attr != "" {
		details.Namespace = attr.(string)
	}

	if attr := config["cluster_name"]; attr != "" {
		details.ClusterName = attr.(string)
	}

	if attr := config["release_name"]; attr != "" {
		details.ReleaseName = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func expandAwsSSHConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsSSH{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["autoscaling_group_name"]; attr != "" {
		details.AutoscalingGroupName = attr.(string)
		details.UseAutoScalingGroup = true
	}

	if attr := config["desired_capacity"]; attr != nil {
		details.DesiredCapacity = attr.(int)
		details.SetDesiredCapacity = details.DesiredCapacity > 0
	}

	if attr := config["host_connection_attrs_name"]; attr != nil {
		details.HostConnectionAttrsName = attr.(string)
	}

	if attr := config["host_connection_type"]; attr != nil {
		details.HostConnectionType = cac.HostConnectionType(attr.(string))
	}

	if attr := config["loadbalancer_name"]; attr != nil {
		details.LoadBalancerName = attr.(string)
	}

	if attr := config["hostname_convention"]; attr != nil {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["region"]; attr != nil {
		details.Region = attr.(string)
	}

	expandAwsInstanceFilter(config, details)

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func expandAwsInstanceFilter(config map[string]interface{}, details *cac.InfrastructureAwsSSH) {
	details.AwsInstanceFilter = &cac.AwsInstanceFilter{}

	if attr := config["tag"]; attr != nil {
		tagConfig := attr.(*schema.Set).List()
		if len(tagConfig) > 0 {
			details.AwsInstanceFilter.Tags = make([]*cac.AwsTag, len(tagConfig))
			for i, tag := range tagConfig {
				tConfig := tag.(map[string]interface{})
				t := &cac.AwsTag{
					Key:   tConfig["key"].(string),
					Value: tConfig["value"].(string),
				}
				details.AwsInstanceFilter.Tags[i] = t
			}
		}
	}

	if attr := config["vpc_ids"]; attr != nil {
		vpcIds := attr.(*schema.Set).List()
		if len(vpcIds) > 0 {
			details.AwsInstanceFilter.VpcIds = make([]string, len(vpcIds))
			for i, vpcId := range vpcIds {
				details.AwsInstanceFilter.VpcIds[i] = vpcId.(string)
			}
		}
	}
}

func flattenAwsSSHConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type == cac.InfrastructureTypes.AwsSSH && infraDef.DeploymentType == cac.DeploymentTypes.SSH {

		detailConfig := map[string]interface{}{}
		infraDetail := detail.ToAwsSSH()

		detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
		detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
		detailConfig["desired_capacity"] = infraDetail.DesiredCapacity
		detailConfig["host_connection_attrs_name"] = infraDetail.HostConnectionAttrsName
		detailConfig["host_connection_type"] = infraDetail.HostConnectionType
		detailConfig["loadbalancer_name"] = infraDetail.LoadBalancerName
		detailConfig["hostname_convention"] = infraDetail.HostNameConvention
		detailConfig["region"] = infraDetail.Region

		if infraDetail.AwsInstanceFilter != nil {
			if tags := flattenAwsTags(infraDetail.AwsInstanceFilter.Tags); len(tags) > 0 {
				detailConfig["tag"] = tags
			}

			if vpcIds := flattenVpcIds(infraDetail.AwsInstanceFilter.VpcIds); len(vpcIds) > 0 {
				detailConfig["vpc_ids"] = vpcIds
			}
		}

		results = append(results, detailConfig)
	}

	return results
}

func expandAwsAmiConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsAmi{}

	if attr := config["ami_deployment_type"]; attr != "" {
		details.AmiDeploymentType = cac.AmiDeploymentType(attr.(string))
	}

	if attr := config["asg_identifies_workload"]; attr != nil {
		details.ASGIdentifiesWorkload = attr.(bool)
	}

	if attr := config["autoscaling_group_name"]; attr != "" {
		details.AutoscalingGroupName = attr.(string)
	}

	if attr := config["classic_loadbalancers"]; attr != nil {
		lbs := []string{}
		for _, lb := range attr.(*schema.Set).List() {
			lbs = append(lbs, lb.(string))
		}
		if len(lbs) > 0 {
			details.ClassicLoadBalancers = lbs
		}
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["hostname_convention"]; attr != "" {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	if attr := config["spotinst_cloud_provider_name"]; attr != "" {
		details.SpotinstCloudProviderName = attr.(string)
	}

	if attr := config["spotinst_config_json"]; attr != "" {
		details.SpotinstElastiGroupJson = attr.(string)
	}

	if attr := config["stage_classic_loadbalancers"]; attr != nil {
		lbs := []string{}
		for _, lb := range attr.(*schema.Set).List() {
			lbs = append(lbs, lb.(string))
		}
		if len(lbs) > 0 {
			details.StageClassicLoadBalancers = lbs
		}
	}

	if attr := config["stage_target_group_arns"]; attr != nil {
		groupsArns := []string{}
		for _, arn := range attr.(*schema.Set).List() {
			groupsArns = append(groupsArns, arn.(string))
		}
		if len(groupsArns) > 0 {
			details.StageTargetGroupArns = groupsArns
		}
	}

	if attr := config["target_group_arns"]; attr != nil {
		groupsArns := []string{}
		for _, arn := range attr.(*schema.Set).List() {
			groupsArns = append(groupsArns, arn.(string))
		}
		if len(groupsArns) > 0 {
			details.TargetGroupArns = groupsArns
		}
	}

	if attr := config["use_traffic_shift"]; attr != nil {
		details.UseTrafficShift = attr.(bool)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsAmiConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsAmi {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsAmi()

	detailConfig["ami_deployment_type"] = infraDetail.AmiDeploymentType.String()
	detailConfig["asg_identifies_workload"] = infraDetail.ASGIdentifiesWorkload
	detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
	detailConfig["classic_loadbalancers"] = infraDetail.ClassicLoadBalancers
	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["hostname_convention"] = infraDetail.HostNameConvention
	detailConfig["region"] = infraDetail.Region
	detailConfig["spotinst_cloud_provider_name"] = infraDetail.SpotinstCloudProviderName
	detailConfig["spotinst_config_json"] = infraDetail.SpotinstElastiGroupJson
	detailConfig["stage_classic_loadbalancers"] = infraDetail.StageClassicLoadBalancers
	detailConfig["stage_target_group_arns"] = infraDetail.StageClassicLoadBalancers
	detailConfig["target_group_arns"] = infraDetail.TargetGroupArns
	detailConfig["use_traffic_shift"] = infraDetail.UseTrafficShift

	return append(results, detailConfig)
}

func expandAwsEcsConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsEcs{}

	if attr := config["assign_public_ip"]; attr != nil {
		details.AssignPublicIp = attr.(bool)
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["cluster_name"]; attr != "" {
		details.ClusterName = attr.(string)
	}

	if attr := config["execution_role"]; attr != "" {
		details.ExecutionRole = attr.(string)
	}

	if attr := config["launch_type"]; attr != "" {
		details.LaunchType = cac.AwsEcsLaunchType(attr.(string))
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	if attr := config["security_group_ids"]; attr != nil {
		sgs := []string{}
		for _, sg := range attr.(*schema.Set).List() {
			sgs = append(sgs, sg.(string))
		}
		if len(sgs) > 0 {
			details.SecurityGroupIds = sgs
		}
	}

	if attr := config["subnet_ids"]; attr != nil {
		subnets := []string{}
		for _, subnet := range attr.(*schema.Set).List() {
			subnets = append(subnets, subnet.(string))
		}
		if len(subnets) > 0 {
			details.SubnetIds = subnets
		}
	}

	if attr := config["vpc_id"]; attr != "" {
		details.VpcId = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsEcsConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsEcs {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsEcs()

	detailConfig["assign_public_ip"] = infraDetail.AssignPublicIp
	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["cluster_name"] = infraDetail.ClusterName
	detailConfig["execution_role"] = infraDetail.ExecutionRole
	detailConfig["launch_type"] = infraDetail.LaunchType
	detailConfig["region"] = infraDetail.Region
	detailConfig["security_group_ids"] = infraDetail.SecurityGroupIds
	detailConfig["subnet_ids"] = infraDetail.SubnetIds
	detailConfig["vpc_id"] = infraDetail.VpcId

	return append(results, detailConfig)
}

func expandAwsLambdaConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsLambda{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["iam_role"]; attr != "" {
		details.IamRole = attr.(string)
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	if attr := config["security_group_ids"]; attr != nil {
		sgs := []string{}
		for _, sg := range attr.(*schema.Set).List() {
			sgs = append(sgs, sg.(string))
		}
		if len(sgs) > 0 {
			details.SecurityGroupIds = sgs
		}
	}

	if attr := config["subnet_ids"]; attr != nil {
		subnets := []string{}
		for _, subnet := range attr.(*schema.Set).List() {
			subnets = append(subnets, subnet.(string))
		}
		if len(subnets) > 0 {
			details.SubnetIds = subnets
		}
	}

	if attr := config["vpc_id"]; attr != "" {
		details.VpcId = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsLambdaConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsLambda {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsLambda()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["iam_role"] = infraDetail.IamRole
	detailConfig["region"] = infraDetail.Region
	detailConfig["security_group_ids"] = infraDetail.SecurityGroupIds
	detailConfig["subnet_ids"] = infraDetail.SubnetIds
	detailConfig["vpc_id"] = infraDetail.VpcId

	return append(results, detailConfig)
}

func expandAwsWinRMConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsWinRM{}

	if attr := config["autoscaling_group_name"]; attr != nil {
		details.AutoscalingGroupName = attr.(string)
		details.UseAutoScalingGroup = true
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["desired_capacity"]; attr != "" {
		details.DesiredCapacity = attr.(int)
		if details.DesiredCapacity > 0 {
			details.SetDesiredCapacity = true
		}
	}

	if attr := config["host_connection_attrs_name"]; attr != "" {
		details.HostConnectionAttrsName = attr.(string)
	}

	if attr := config["host_connection_type"]; attr != "" {
		details.HostConnectionType = cac.HostConnectionType(attr.(string))
	}

	if attr := config["hostname_convention"]; attr != "" {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["loadbalancer_name"]; attr != "" {
		details.LoadBalancerName = attr.(string)
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsWinRMConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type == cac.InfrastructureTypes.AwsSSH && infraDef.DeploymentType == cac.DeploymentTypes.WinRM {
		detailConfig := map[string]interface{}{}
		infraDetail := detail.ToAwsWinRm()

		detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
		detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
		detailConfig["desired_capacity"] = infraDetail.DesiredCapacity
		detailConfig["host_connection_attrs_name"] = infraDetail.HostConnectionAttrsName
		detailConfig["host_connection_type"] = infraDetail.HostConnectionType
		detailConfig["hostname_convention"] = infraDetail.HostNameConvention
		detailConfig["loadbalancer_name"] = infraDetail.LoadBalancerName
		detailConfig["region"] = infraDetail.Region

		results = append(results, detailConfig)
	}

	return results
}

func expandTanzuConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureTanzu{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["organization"]; attr != nil {
		details.Organization = attr.(string)
	}

	if attr := config["space"]; attr != nil {
		details.Space = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenTanzuConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.Pcf {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToPcf()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["organization"] = infraDetail.Organization
	detailConfig["space"] = infraDetail.Space

	return append(results, detailConfig)
}

func expandAzureWebAppConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAzureWebApp{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["resource_group"]; attr != nil {
		details.ResourceGroup = attr.(string)
	}

	if attr := config["subscription_id"]; attr != nil {
		details.SubscriptionId = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAzureWebAppConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AzureWebApp {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAzureWebApp()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["resource_group"] = infraDetail.ResourceGroup
	detailConfig["subscription_id"] = infraDetail.SubscriptionId

	return append(results, detailConfig)
}

func flattenAwsTags(tags []*cac.AwsTag) []interface{} {
	results := []interface{}{}

	for _, tag := range tags {
		t := map[string]interface{}{
			"key":   tag.Key,
			"value": tag.Value,
		}
		results = append(results, t)
	}

	return results
}

func flattenVpcIds(vpcIds []string) []interface{} {
	var results []interface{}

	for _, vpcId := range vpcIds {
		results = append(results, vpcId)
	}

	return results
}
