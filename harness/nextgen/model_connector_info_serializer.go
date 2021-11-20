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
	case ConnectorTypes.Artifactory:
		err = json.Unmarshal(aux.Spec, &a.Artifactory)
	case ConnectorTypes.Aws:
		err = json.Unmarshal(aux.Spec, &a.Aws)
	case ConnectorTypes.AwsKms:
		err = json.Unmarshal(aux.Spec, &a.AwsKms)
	case ConnectorTypes.AwsSecretManager:
		err = json.Unmarshal(aux.Spec, &a.AwsSecretManager)
	case ConnectorTypes.CEAws:
		err = json.Unmarshal(aux.Spec, &a.AwsCC)
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
	case ConnectorTypes.Git:
		err = json.Unmarshal(aux.Spec, &a.Git)
	case ConnectorTypes.Github:
		err = json.Unmarshal(aux.Spec, &a.Github)
	case ConnectorTypes.Gitlab:
		err = json.Unmarshal(aux.Spec, &a.Gitlab)
	case ConnectorTypes.HttpHelmRepo:
		err = json.Unmarshal(aux.Spec, &a.HttpHelm)
	case ConnectorTypes.Jira:
		err = json.Unmarshal(aux.Spec, &a.Jira)
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
	case ConnectorTypes.Git:
		spec, err = json.Marshal(a.Git)
	case ConnectorTypes.Github:
		spec, err = json.Marshal(a.Github)
	case ConnectorTypes.Gitlab:
		spec, err = json.Marshal(a.Gitlab)
	case ConnectorTypes.HttpHelmRepo:
		spec, err = json.Marshal(a.HttpHelm)
	case ConnectorTypes.Jira:
		spec, err = json.Marshal(a.Jira)
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
	case ConnectorTypes.SumoLogic:
		spec, err = json.Marshal(a.SumoLogic)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
