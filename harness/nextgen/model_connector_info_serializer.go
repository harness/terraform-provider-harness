package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ConnectorInfo) UnmarshalJSON(data []byte) error {

	type Alias ConnectorInfo

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case ConnectorTypes.AppDynamics:
		err = json.Unmarshal(aux.Spec, &a.AppDynamics)
	case ConnectorTypes.ElasticSearch:
		err = json.Unmarshal(aux.Spec, &a.ElasticSearch)
	case ConnectorTypes.Artifactory:
		err = json.Unmarshal(aux.Spec, &a.Artifactory)
	case ConnectorTypes.Aws:
		err = json.Unmarshal(aux.Spec, &a.Aws)
	case ConnectorTypes.AwsKms:
		err = json.Unmarshal(aux.Spec, &a.AwsKms)
	case ConnectorTypes.AwsSecretManager:
		err = json.Unmarshal(aux.Spec, &a.AwsSecretManager)
	case ConnectorTypes.Azure:
		err = json.Unmarshal(aux.Spec, &a.Azure)
	case ConnectorTypes.AzureKeyVault:
		err = json.Unmarshal(aux.Spec, &a.AzureKeyVault)
	case ConnectorTypes.CEAzure:
		err = json.Unmarshal(aux.Spec, &a.AzureCloudCost)
	case ConnectorTypes.CEAws:
		err = json.Unmarshal(aux.Spec, &a.AwsCC)
	case ConnectorTypes.CEK8sCluster:
		err = json.Unmarshal(aux.Spec, &a.K8sClusterCloudCost)
	case ConnectorTypes.Bitbucket:
		err = json.Unmarshal(aux.Spec, &a.BitBucket)
	case ConnectorTypes.Datadog:
		err = json.Unmarshal(aux.Spec, &a.Datadog)
	case ConnectorTypes.DockerRegistry:
		err = json.Unmarshal(aux.Spec, &a.DockerRegistry)
	case ConnectorTypes.Dynatrace:
		err = json.Unmarshal(aux.Spec, &a.Dynatrace)
	case ConnectorTypes.Gcp:
		err = json.Unmarshal(aux.Spec, &a.Gcp)
	case ConnectorTypes.GcpCloudCost:
		err = json.Unmarshal(aux.Spec, &a.GcpCloudCost)
	case ConnectorTypes.Git:
		err = json.Unmarshal(aux.Spec, &a.Git)
	case ConnectorTypes.Github:
		err = json.Unmarshal(aux.Spec, &a.Github)
	case ConnectorTypes.Gitlab:
		err = json.Unmarshal(aux.Spec, &a.Gitlab)
	case ConnectorTypes.Vault:
		err = json.Unmarshal(aux.Spec, &a.Vault)
	case ConnectorTypes.HttpHelmRepo:
		err = json.Unmarshal(aux.Spec, &a.HttpHelm)
	case ConnectorTypes.OciHelmRepo:
		err = json.Unmarshal(aux.Spec, &a.OciHelm)
	case ConnectorTypes.Jira:
		err = json.Unmarshal(aux.Spec, &a.Jira)
	case ConnectorTypes.Jenkins:
		err = json.Unmarshal(aux.Spec, &a.Jenkins)
	case ConnectorTypes.K8sCluster:
		err = json.Unmarshal(aux.Spec, &a.K8sCluster)
	case ConnectorTypes.Nexus:
		err = json.Unmarshal(aux.Spec, &a.Nexus)
	case ConnectorTypes.NewRelic:
		err = json.Unmarshal(aux.Spec, &a.NewRelic)
	case ConnectorTypes.PagerDuty:
		err = json.Unmarshal(aux.Spec, &a.PagerDuty)
	case ConnectorTypes.Prometheus:
		err = json.Unmarshal(aux.Spec, &a.Prometheus)
	case ConnectorTypes.Splunk:
		err = json.Unmarshal(aux.Spec, &a.Splunk)
	case ConnectorTypes.SumoLogic:
		err = json.Unmarshal(aux.Spec, &a.SumoLogic)
	case ConnectorTypes.GcpSecretManager:
		err = json.Unmarshal(aux.Spec, &a.GcpSecretManager)
	case ConnectorTypes.Spot:
		err = json.Unmarshal(aux.Spec, &a.Spot)
	case ConnectorTypes.ServiceNow:
		err = json.Unmarshal(aux.Spec, &a.ServiceNow)
	case ConnectorTypes.Tas:
		err = json.Unmarshal(aux.Spec, &a.Tas)
	case ConnectorTypes.TerraformCloud:
		err = json.Unmarshal(aux.Spec, &a.TerraformCloud)
	case ConnectorTypes.Rancher:
		err = json.Unmarshal(aux.Spec, &a.Rancher)
	case ConnectorTypes.CustomHealth:
		err = json.Unmarshal(aux.Spec, &a.CustomHealth)
	case ConnectorTypes.Pdc:
		err = json.Unmarshal(aux.Spec, &a.Pdc)
	case ConnectorTypes.CustomSecretManager:
		err = json.Unmarshal(aux.Spec, &a.CustomSecretManager)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	return err
}

func (a *ConnectorInfo) MarshalJSON() ([]byte, error) {
	type Alias ConnectorInfo

	var spec []byte
	var err error

	switch a.Type_ {
	case ConnectorTypes.AppDynamics:
		spec, err = json.Marshal(a.AppDynamics)
	case ConnectorTypes.ElasticSearch:
		spec, err = json.Marshal(a.ElasticSearch)
	case ConnectorTypes.Artifactory:
		spec, err = json.Marshal(a.Artifactory)
	case ConnectorTypes.Aws:
		spec, err = json.Marshal(a.Aws)
	case ConnectorTypes.AwsKms:
		spec, err = json.Marshal(a.AwsKms)
	case ConnectorTypes.AwsSecretManager:
		spec, err = json.Marshal(a.AwsSecretManager)
	case ConnectorTypes.CEAws:
		spec, err = json.Marshal(a.AwsCC)
	case ConnectorTypes.Bitbucket:
		spec, err = json.Marshal(a.BitBucket)
	case ConnectorTypes.Datadog:
		spec, err = json.Marshal(a.Datadog)
	case ConnectorTypes.DockerRegistry:
		spec, err = json.Marshal(a.DockerRegistry)
	case ConnectorTypes.Dynatrace:
		spec, err = json.Marshal(a.Dynatrace)
	case ConnectorTypes.Gcp:
		spec, err = json.Marshal(a.Gcp)
	case ConnectorTypes.GcpCloudCost:
		spec, err = json.Marshal(a.GcpCloudCost)
	case ConnectorTypes.Git:
		spec, err = json.Marshal(a.Git)
	case ConnectorTypes.Github:
		spec, err = json.Marshal(a.Github)
	case ConnectorTypes.Gitlab:
		spec, err = json.Marshal(a.Gitlab)
	case ConnectorTypes.Vault:
		spec, err = json.Marshal(a.Vault)
	case ConnectorTypes.HttpHelmRepo:
		spec, err = json.Marshal(a.HttpHelm)
	case ConnectorTypes.OciHelmRepo:
		spec, err = json.Marshal(a.OciHelm)
	case ConnectorTypes.Jira:
		spec, err = json.Marshal(a.Jira)
	case ConnectorTypes.Jenkins:
		spec, err = json.Marshal(a.Jenkins)
	case ConnectorTypes.K8sCluster:
		spec, err = json.Marshal(a.K8sCluster)
	case ConnectorTypes.NewRelic:
		spec, err = json.Marshal(a.NewRelic)
	case ConnectorTypes.Nexus:
		spec, err = json.Marshal(a.Nexus)
	case ConnectorTypes.PagerDuty:
		spec, err = json.Marshal(a.PagerDuty)
	case ConnectorTypes.Prometheus:
		spec, err = json.Marshal(a.Prometheus)
	case ConnectorTypes.Splunk:
		spec, err = json.Marshal(a.Splunk)
	case ConnectorTypes.Azure:
		spec, err = json.Marshal(a.Azure)
	case ConnectorTypes.AzureKeyVault:
		spec, err = json.Marshal(a.AzureKeyVault)
	case ConnectorTypes.CEAzure:
		spec, err = json.Marshal(a.AzureCloudCost)
	case ConnectorTypes.CEK8sCluster:
		spec, err = json.Marshal(a.K8sClusterCloudCost)
	case ConnectorTypes.SumoLogic:
		spec, err = json.Marshal(a.SumoLogic)
	case ConnectorTypes.GcpSecretManager:
		spec, err = json.Marshal(a.GcpSecretManager)
	case ConnectorTypes.Spot:
		spec, err = json.Marshal(a.Spot)
	case ConnectorTypes.ServiceNow:
		spec, err = json.Marshal(a.ServiceNow)
	case ConnectorTypes.Tas:
		spec, err = json.Marshal(a.Tas)
	case ConnectorTypes.TerraformCloud:
		spec, err = json.Marshal(a.TerraformCloud)
	case ConnectorTypes.Rancher:
		spec, err = json.Marshal(a.Rancher)
	case ConnectorTypes.CustomHealth:
		spec, err = json.Marshal(a.CustomHealth)
	case ConnectorTypes.Pdc:
		spec, err = json.Marshal(a.Pdc)
	case ConnectorTypes.CustomSecretManager:
		spec, err = json.Marshal(a.CustomSecretManager)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
