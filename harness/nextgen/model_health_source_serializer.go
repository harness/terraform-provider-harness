package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *HealthSource) UnmarshalJSON(data []byte) error {

	type Alias HealthSource

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
	case HealthSourceTypes.AppDynamics:
		err = json.Unmarshal(aux.Spec, &a.AppDynamics)
	case HealthSourceTypes.NewRelic:
		err = json.Unmarshal(aux.Spec, &a.NewRelic)
	case HealthSourceTypes.StackdriverLog:
		err = json.Unmarshal(aux.Spec, &a.StackdriverLog)
	case HealthSourceTypes.Splunk:
		err = json.Unmarshal(aux.Spec, &a.Splunk)
	case HealthSourceTypes.Prometheus:
		err = json.Unmarshal(aux.Spec, &a.Prometheus)
	case HealthSourceTypes.Stackdriver:
		err = json.Unmarshal(aux.Spec, &a.Stackdriver)
	case HealthSourceTypes.DatadogMetrics:
		err = json.Unmarshal(aux.Spec, &a.DatadogMetrics)
	case HealthSourceTypes.DatadogLog:
		err = json.Unmarshal(aux.Spec, &a.DatadogLog)
	case HealthSourceTypes.Dynatrace:
		err = json.Unmarshal(aux.Spec, &a.Dynatrace)
	case HealthSourceTypes.ErrorTracking:
		err = json.Unmarshal(aux.Spec, &a.ErrorTracking)
	case HealthSourceTypes.CustomHealthMetric:
		err = json.Unmarshal(aux.Spec, &a.CustomHealthMetric)
	case HealthSourceTypes.CustomHealthLog:
		err = json.Unmarshal(aux.Spec, &a.CustomHealthLog)
	case HealthSourceTypes.SplunkMetric:
		err = json.Unmarshal(aux.Spec, &a.SplunkMetric)
	case HealthSourceTypes.ElasticSearch:
		err = json.Unmarshal(aux.Spec, &a.ElasticSearch)
	case HealthSourceTypes.CloudWatchMetrics:
		err = json.Unmarshal(aux.Spec, &a.CloudWatchMetrics)
	case HealthSourceTypes.AwsPrometheus:
		err = json.Unmarshal(aux.Spec, &a.AwsPrometheus)
	case HealthSourceTypes.SumologicMetrics:
		err = json.Unmarshal(aux.Spec, &a.SumologicMetrics)
	case HealthSourceTypes.SumologicLogs:
		err = json.Unmarshal(aux.Spec, &a.SumologicLogs)
	case HealthSourceTypes.SplunkSignalFXMetrics:
		err = json.Unmarshal(aux.Spec, &a.SplunkSignalFXMetrics)
	case HealthSourceTypes.GrafanaLokiLogs:
		err = json.Unmarshal(aux.Spec, &a.GrafanaLokiLogs)
	case HealthSourceTypes.AzureLogs:
		err = json.Unmarshal(aux.Spec, &a.AzureLogs)
	case HealthSourceTypes.AzureMetrics:
		err = json.Unmarshal(aux.Spec, &a.AzureMetrics)
	default:
		panic(fmt.Sprintf("unknown health source type %s", a.Type_))
	}

	return err
}

func (a *HealthSource) MarshalJSON() ([]byte, error) {
	type Alias HealthSource

	var spec []byte
	var err error

	switch a.Type_ {
	case HealthSourceTypes.AppDynamics:
		spec, err = json.Marshal(a.AppDynamics)
	case HealthSourceTypes.NewRelic:
		spec, err = json.Marshal(a.NewRelic)
	case HealthSourceTypes.StackdriverLog:
		spec, err = json.Marshal(a.StackdriverLog)
	case HealthSourceTypes.Splunk:
		spec, err = json.Marshal(a.Splunk)
	case HealthSourceTypes.Prometheus:
		spec, err = json.Marshal(a.Prometheus)
	case HealthSourceTypes.Stackdriver:
		spec, err = json.Marshal(a.Stackdriver)
	case HealthSourceTypes.DatadogMetrics:
		spec, err = json.Marshal(a.DatadogMetrics)
	case HealthSourceTypes.DatadogLog:
		spec, err = json.Marshal(a.DatadogLog)
	case HealthSourceTypes.Dynatrace:
		spec, err = json.Marshal(a.Dynatrace)
	case HealthSourceTypes.ErrorTracking:
		spec, err = json.Marshal(a.ErrorTracking)
	case HealthSourceTypes.CustomHealthMetric:
		spec, err = json.Marshal(a.CustomHealthMetric)
	case HealthSourceTypes.CustomHealthLog:
		spec, err = json.Marshal(a.CustomHealthLog)
	case HealthSourceTypes.SplunkMetric:
		spec, err = json.Marshal(a.SplunkMetric)
	case HealthSourceTypes.ElasticSearch:
		spec, err = json.Marshal(a.ElasticSearch)
	case HealthSourceTypes.CloudWatchMetrics:
		spec, err = json.Marshal(a.CloudWatchMetrics)
	case HealthSourceTypes.AwsPrometheus:
		spec, err = json.Marshal(a.AwsPrometheus)
	case HealthSourceTypes.SumologicMetrics:
		spec, err = json.Marshal(a.SumologicMetrics)
	case HealthSourceTypes.SumologicLogs:
		spec, err = json.Marshal(a.SumologicLogs)
	case HealthSourceTypes.SplunkSignalFXMetrics:
		spec, err = json.Marshal(a.SplunkSignalFXMetrics)
	case HealthSourceTypes.GrafanaLokiLogs:
		spec, err = json.Marshal(a.GrafanaLokiLogs)
	case HealthSourceTypes.AzureLogs:
		spec, err = json.Marshal(a.AzureLogs)
	case HealthSourceTypes.AzureMetrics:
		spec, err = json.Marshal(a.AzureMetrics)
	default:
		panic(fmt.Sprintf("unknown health source type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
