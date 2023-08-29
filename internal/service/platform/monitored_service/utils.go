package monitored_service

import (
	"encoding/json"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func getHealthSourceByType(hs map[string]interface{}) nextgen.HealthSource {
	healthSourceType := hs["type"].(string)
	healthSource := hs["spec"].(string)

	if healthSourceType == "AppDynamics" {
		data := nextgen.AppDynamicsHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:        hs["name"].(string),
			Identifier:  hs["identifier"].(string),
			Version:     hs["version"].(string),
			Type_:       nextgen.HealthSourceType(healthSourceType),
			AppDynamics: &data,
		}
	}
	if healthSourceType == "NewRelic" {
		data := nextgen.NewRelicHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			NewRelic:   &data,
		}
	}
	if healthSourceType == "StackdriverLog" {
		data := nextgen.StackdriverLogHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:           hs["name"].(string),
			Identifier:     hs["identifier"].(string),
			Version:        hs["version"].(string),
			Type_:          nextgen.HealthSourceType(healthSourceType),
			StackdriverLog: &data,
		}
	}
	if healthSourceType == "Splunk" {
		data := nextgen.SplunkHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Splunk:     &data,
		}
	}
	if healthSourceType == "Prometheus" {
		data := nextgen.PrometheusHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Prometheus: &data,
		}
	}
	if healthSourceType == "Stackdriver" {
		data := nextgen.StackdriverMetricHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:        hs["name"].(string),
			Identifier:  hs["identifier"].(string),
			Version:     hs["version"].(string),
			Type_:       nextgen.HealthSourceType(healthSourceType),
			Stackdriver: &data,
		}
	}
	if healthSourceType == "DatadogMetrics" {
		data := nextgen.DatadogMetricHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:           hs["name"].(string),
			Identifier:     hs["identifier"].(string),
			Version:        hs["version"].(string),
			Type_:          nextgen.HealthSourceType(healthSourceType),
			DatadogMetrics: &data,
		}
	}
	if healthSourceType == "DatadogLog" {
		data := nextgen.DatadogLogHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			DatadogLog: &data,
		}
	}
	if healthSourceType == "Dynatrace" {
		data := nextgen.DynatraceHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Dynatrace:  &data,
		}
	}
	if healthSourceType == "ErrorTracking" {
		data := nextgen.ErrorTrackingHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Version:       hs["version"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			ErrorTracking: &data,
		}
	}
	if healthSourceType == "CustomHealthMetric" {
		data := nextgen.CustomHealthSourceMetric{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:               hs["name"].(string),
			Identifier:         hs["identifier"].(string),
			Version:            hs["version"].(string),
			Type_:              nextgen.HealthSourceType(healthSourceType),
			CustomHealthMetric: &data,
		}
	}
	if healthSourceType == "CustomHealthLog" {
		data := nextgen.CustomHealthSourceLog{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:            hs["name"].(string),
			Identifier:      hs["identifier"].(string),
			Version:         hs["version"].(string),
			Type_:           nextgen.HealthSourceType(healthSourceType),
			CustomHealthLog: &data,
		}
	}
	if healthSourceType == "SplunkMetric" {
		data := nextgen.SplunkMetricHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:         hs["name"].(string),
			Identifier:   hs["identifier"].(string),
			Version:      hs["version"].(string),
			Type_:        nextgen.HealthSourceType(healthSourceType),
			SplunkMetric: &data,
		}
	}
	if healthSourceType == "ElasticSearch" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Version:       hs["version"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			ElasticSearch: &data,
		}
	}
	if healthSourceType == "CloudWatchMetrics" {
		data := nextgen.CloudWatchMetricsHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:              hs["name"].(string),
			Identifier:        hs["identifier"].(string),
			Version:           hs["version"].(string),
			Type_:             nextgen.HealthSourceType(healthSourceType),
			CloudWatchMetrics: &data,
		}
	}
	if healthSourceType == "AwsPrometheus" {
		data := nextgen.AwsPrometheusHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Version:       hs["version"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			AwsPrometheus: &data,
		}
	}
	if healthSourceType == "SumologicMetrics" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:             hs["name"].(string),
			Identifier:       hs["identifier"].(string),
			Version:          hs["version"].(string),
			Type_:            nextgen.HealthSourceType(healthSourceType),
			SumologicMetrics: &data,
		}
	}
	if healthSourceType == "SumologicLogs" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Version:       hs["version"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			SumologicLogs: &data,
		}
	}
	if healthSourceType == "SplunkSignalFXMetrics" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:                  hs["name"].(string),
			Identifier:            hs["identifier"].(string),
			Version:               hs["version"].(string),
			Type_:                 nextgen.HealthSourceType(healthSourceType),
			SplunkSignalFXMetrics: &data,
		}
	}
	if healthSourceType == "GrafanaLokiLogs" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:            hs["name"].(string),
			Identifier:      hs["identifier"].(string),
			Version:         hs["version"].(string),
			Type_:           nextgen.HealthSourceType(healthSourceType),
			GrafanaLokiLogs: &data,
		}
	}
	if healthSourceType == "AzureLogs" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Version:    hs["version"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			AzureLogs:  &data,
		}
	}
	if healthSourceType == "AzureMetrics" {
		data := nextgen.NextGenHealthSource{}
		json.Unmarshal([]byte(healthSource), &data)

		return nextgen.HealthSource{
			Name:         hs["name"].(string),
			Identifier:   hs["identifier"].(string),
			Version:      hs["version"].(string),
			Type_:        nextgen.HealthSourceType(healthSourceType),
			AzureMetrics: &data,
		}
	}
	panic(fmt.Sprintf("Invalid health source type for monitored service"))
}

func getChangeSourceByType(cs map[string]interface{}) nextgen.ChangeSourceDto {
	changeSourceType := cs["type"].(string)
	changeSourceSpec := cs["spec"].(string)

	if changeSourceType == "HarnessCDNextGen" {
		data := nextgen.HarnessCdChangeSourceSpec{}
		json.Unmarshal([]byte(changeSourceSpec), &data)

		return nextgen.ChangeSourceDto{
			Name:             cs["name"].(string),
			Identifier:       cs["identifier"].(string),
			Type_:            nextgen.ChangeSourceType(changeSourceType),
			HarnessCDNextGen: &data,
			Enabled:          cs["enabled"].(bool),
			Category:         cs["category"].(string),
		}
	}
	if changeSourceType == "PagerDuty" {
		data := nextgen.PagerDutyChangeSourceSpec{}
		json.Unmarshal([]byte(changeSourceSpec), &data)

		return nextgen.ChangeSourceDto{
			Name:       cs["name"].(string),
			Identifier: cs["identifier"].(string),
			Type_:      nextgen.ChangeSourceType(changeSourceType),
			PagerDuty:  &data,
			Enabled:    cs["enabled"].(bool),
			Category:   cs["category"].(string),
		}
	}
	if changeSourceType == "K8sCluster" {
		data := nextgen.KubernetesChangeSourceSpec{}
		json.Unmarshal([]byte(changeSourceSpec), &data)

		return nextgen.ChangeSourceDto{
			Name:       cs["name"].(string),
			Identifier: cs["identifier"].(string),
			Type_:      nextgen.ChangeSourceType(changeSourceType),
			K8sCluster: &data,
			Enabled:    cs["enabled"].(bool),
			Category:   cs["category"].(string),
		}
	}
	if changeSourceType == "HarnessCD" {
		data := nextgen.HarnessCdCurrentGenChangeSourceSpec{}
		json.Unmarshal([]byte(changeSourceSpec), &data)

		return nextgen.ChangeSourceDto{
			Name:       cs["name"].(string),
			Identifier: cs["identifier"].(string),
			Type_:      nextgen.ChangeSourceType(changeSourceType),
			HarnessCD:  &data,
			Enabled:    cs["enabled"].(bool),
			Category:   cs["category"].(string),
		}
	}

	panic(fmt.Sprintf("Invalid change source type for monitored service"))
}

func getServiceDependencyByType(sd map[string]interface{}) nextgen.ServiceDependencyDto {
	dependencyType := sd["type"].(string)
	dependencyMetadata := sd["dependency_metadata"].(string)

	if dependencyType == "KUBERNETES" {
		data := nextgen.KubernetesDependencyMetadata{}
		json.Unmarshal([]byte(dependencyMetadata), &data)

		return nextgen.ServiceDependencyDto{
			MonitoredServiceIdentifier: sd["monitored_service_identifier"].(string),
			Type_:                      nextgen.DependencyMetadataType(dependencyType),
			KUBERNETES:                 &data,
		}
	}

	panic(fmt.Sprintf("Invalid service dependency type for monitored service"))
}
