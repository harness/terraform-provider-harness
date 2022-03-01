package environment

import (
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
