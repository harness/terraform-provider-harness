package monitored_service

import (
	"encoding/json"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getHealthSourceByType(hs map[string]interface{}) nextgen.HealthSource {
	healthSourceType := hs["type"].(string)
	healthSource := hs["spec"].(string)

	if healthSourceType == "AppDynamics" {
		data := getAppDynamicsHealthSource(hs["spec"].(map[string]interface{}))

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
		data := getPrometheusHealthSource(hs["spec"].(map[string]interface{}))

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
	return nextgen.ServiceDependencyDto{
		MonitoredServiceIdentifier: sd["monitored_service_identifier"].(string),
	}
}

func getMetricThresholdByType(hs map[string]interface{}) nextgen.MetricThreshold {
	metricThresholdType := hs["type"].(string)
	spec := hs["spec"].(string)

	if metricThresholdType == "FailImmediately" {
		data := nextgen.FailMetricThresholdSpec{}
		json.Unmarshal([]byte(spec), &data)

		return nextgen.MetricThreshold{
			GroupName:        hs["groupName"].(string),
			MetricName:       hs["metricName"].(string),
			MetricIdentifier: hs["metricIdentifier"].(string),
			MetricType:       hs["metricType"].(string),
			Type_:            nextgen.MetricThresholdType(metricThresholdType),
			FailImmediately:  &data,
		}
	}
	if metricThresholdType == "IgnoreThreshold" {
		data := nextgen.IgnoreMetricThresholdSpec{}
		json.Unmarshal([]byte(spec), &data)

		return nextgen.MetricThreshold{
			GroupName:        hs["groupName"].(string),
			MetricName:       hs["metricName"].(string),
			MetricIdentifier: hs["metricIdentifier"].(string),
			MetricType:       hs["metricType"].(string),
			Type_:            nextgen.MetricThresholdType(metricThresholdType),
			IgnoreThreshold:  &data,
		}
	}
	panic(fmt.Sprintf("Invalid metric threshold"))
}

func getAppDynamicsHealthSource(hs map[string]interface{}) nextgen.AppDynamicsHealthSource {
	appDynamicHealthSource := &nextgen.AppDynamicsHealthSource{}

	appDynamicHealthSource.ConnectorRef = hs["connectorRef"].(string)
	appDynamicHealthSource.Feature = hs["feature"].(string)
	appDynamicHealthSource.ApplicationName = hs["applicationName"].(string)
	appDynamicHealthSource.TierName = hs["tierName"].(string)

	metricDefinitions := hs["metricDefinitions"].(*schema.Set).List()
	appDMetricDefinitions := make([]nextgen.AppDMetricDefinitions, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.AppDMetricDefinitions{}
		md := metricDefinition.(string)
		json.Unmarshal([]byte(md), &data)
		appDMetricDefinitions[i] = data
	}
	appDynamicHealthSource.MetricDefinitions = appDMetricDefinitions

	metricPacks := hs["metricPacks"].(*schema.Set).List()
	metricPackDto := make([]nextgen.TimeSeriesMetricPackDto, len(metricPacks))
	for i, metricPack := range metricPacks {
		test := metricPack.(map[string]interface{})
		metricThresholds := test["metricThresholds"].(*schema.Set).List()
		metricThresholdDto := make([]nextgen.MetricThreshold, len(metricPacks))
		for j, metricThreshold := range metricThresholds {
			metricThresholdDto[j] = getMetricThresholdByType(metricThreshold.(map[string]interface{}))
		}
		timeSeriesMetricPackDto := &nextgen.TimeSeriesMetricPackDto{
			Identifier:       test["identifier"].(string),
			MetricThresholds: metricThresholdDto,
		}
		metricPackDto[i] = *timeSeriesMetricPackDto
	}
	appDynamicHealthSource.MetricPacks = metricPackDto
	return *appDynamicHealthSource
}

func getPrometheusHealthSource(hs map[string]interface{}) nextgen.PrometheusHealthSource {
	appDynamicHealthSource := &nextgen.PrometheusHealthSource{}

	appDynamicHealthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].(*schema.Set).List()
	appDMetricDefinitions := make([]nextgen.PrometheusMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.PrometheusMetricDefinition{}
		md := metricDefinition.(string)
		json.Unmarshal([]byte(md), &data)
		appDMetricDefinitions[i] = data
	}
	appDynamicHealthSource.MetricDefinitions = appDMetricDefinitions

	metricPacks := hs["metricPacks"].(*schema.Set).List()
	metricPackDto := make([]nextgen.TimeSeriesMetricPackDto, len(metricPacks))
	for i, metricPack := range metricPacks {
		test := metricPack.(map[string]interface{})
		metricThresholds := test["metricThresholds"].(*schema.Set).List()
		metricThresholdDto := make([]nextgen.MetricThreshold, len(metricPacks))
		for j, metricThreshold := range metricThresholds {
			metricThresholdDto[j] = getMetricThresholdByType(metricThreshold.(map[string]interface{}))
		}
		timeSeriesMetricPackDto := &nextgen.TimeSeriesMetricPackDto{
			Identifier:       test["identifier"].(string),
			MetricThresholds: metricThresholdDto,
		}
		metricPackDto[i] = *timeSeriesMetricPackDto
	}
	appDynamicHealthSource.MetricPacks = metricPackDto
	return *appDynamicHealthSource
}
