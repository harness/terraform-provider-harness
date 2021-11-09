package environment

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
				Description: fmt.Sprintf("The naming convention of the release. When using Helm Native the default is %[1]s. For standard Kubernetes manifests the default is %[2]s", DefaultHelmReleaseName, DefaultK8sReleaseName),
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
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
