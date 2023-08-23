package monitored_service

import (
	"encoding/json"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func getHealthSourceByType(hs map[string]interface{}) nextgen.HealthSource {
	healthSourceType := hs["type"].(string)
	healthSourceSpec := hs["spec"].(string)

	if healthSourceType == "AppDynamics" {
		data := nextgen.AppDynamicsHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:        hs["name"].(string),
			Identifier:  hs["identifier"].(string),
			Type_:       nextgen.HealthSourceType(healthSourceType),
			AppDynamics: &data,
		}
	}
	if healthSourceType == "NewRelic" {
		data := nextgen.NewRelicHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			NewRelic:   &data,
		}
	}
	if healthSourceType == "StackdriverLog" {
		data := nextgen.StackdriverLogHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:           hs["name"].(string),
			Identifier:     hs["identifier"].(string),
			Type_:          nextgen.HealthSourceType(healthSourceType),
			StackdriverLog: &data,
		}
	}
	if healthSourceType == "Splunk" {
		data := nextgen.SplunkHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Splunk:     &data,
		}
	}
	if healthSourceType == "Prometheus" {
		data := nextgen.PrometheusHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Prometheus: &data,
		}
	}
	if healthSourceType == "Stackdriver" {
		data := nextgen.StackdriverMetricHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:        hs["name"].(string),
			Identifier:  hs["identifier"].(string),
			Type_:       nextgen.HealthSourceType(healthSourceType),
			Stackdriver: &data,
		}
	}
	if healthSourceType == "DatadogMetrics" {
		data := nextgen.DatadogMetricHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:           hs["name"].(string),
			Identifier:     hs["identifier"].(string),
			Type_:          nextgen.HealthSourceType(healthSourceType),
			DatadogMetrics: &data,
		}
	}
	if healthSourceType == "DatadogLog" {
		data := nextgen.DatadogLogHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			DatadogLog: &data,
		}
	}
	if healthSourceType == "Dynatrace" {
		data := nextgen.DynatraceHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:       hs["name"].(string),
			Identifier: hs["identifier"].(string),
			Type_:      nextgen.HealthSourceType(healthSourceType),
			Dynatrace:  &data,
		}
	}
	if healthSourceType == "ErrorTracking" {
		data := nextgen.ErrorTrackingHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			ErrorTracking: &data,
		}
	}
	if healthSourceType == "CustomHealthMetric" {
		data := nextgen.CustomHealthSourceMetricSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:               hs["name"].(string),
			Identifier:         hs["identifier"].(string),
			Type_:              nextgen.HealthSourceType(healthSourceType),
			CustomHealthMetric: &data,
		}
	}
	if healthSourceType == "CustomHealthLog" {
		data := nextgen.CustomHealthSourceLogSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:            hs["name"].(string),
			Identifier:      hs["identifier"].(string),
			Type_:           nextgen.HealthSourceType(healthSourceType),
			CustomHealthLog: &data,
		}
	}
	if healthSourceType == "SplunkMetric" {
		data := nextgen.SplunkMetricHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:         hs["name"].(string),
			Identifier:   hs["identifier"].(string),
			Type_:        nextgen.HealthSourceType(healthSourceType),
			SplunkMetric: &data,
		}
	}
	if healthSourceType == "ElasticSearch" {
		data := nextgen.ElkHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:          hs["name"].(string),
			Identifier:    hs["identifier"].(string),
			Type_:         nextgen.HealthSourceType(healthSourceType),
			ElasticSearch: &data,
		}
	}
	if healthSourceType == "CloudWatchMetrics" {
		data := nextgen.CloudWatchMetricsHealthSourceSpec{}
		json.Unmarshal([]byte(healthSourceSpec), &data)

		return nextgen.HealthSource{
			Name:              hs["name"].(string),
			Identifier:        hs["identifier"].(string),
			Type_:             nextgen.HealthSourceType(healthSourceType),
			CloudWatchMetrics: &data,
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
	return nextgen.ServiceDependencyDto{
		MonitoredServiceIdentifier: sd["monitored_service_identifier"].(string),
	}
}
