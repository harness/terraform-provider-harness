package monitored_service

import (
	"encoding/json"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"strconv" //this package is used to convert the data type
)

func getAppDynamicsHealthSource(hs map[string]interface{}) nextgen.AppDynamicsHealthSource {
	healthSource := &nextgen.AppDynamicsHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.AppDMetricDefinitions, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.AppDMetricDefinitions{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getPrometheusHealthSource(hs map[string]interface{}) nextgen.PrometheusHealthSource {
	healthSource := &nextgen.PrometheusHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.PrometheusMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.PrometheusMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getNewRelicHealthSource(hs map[string]interface{}) nextgen.NewRelicHealthSource {
	healthSource := &nextgen.NewRelicHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["newRelicMetricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.NewRelicMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.NewRelicMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.NewRelicMetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getStackDriverHealthSource(hs map[string]interface{}) nextgen.StackdriverMetricHealthSource {
	healthSource := &nextgen.StackdriverMetricHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.StackdriverDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.StackdriverDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getDataDogHealthSource(hs map[string]interface{}) nextgen.DatadogMetricHealthSource {
	healthSource := &nextgen.DatadogMetricHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.DatadogMetricHealthDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.DatadogMetricHealthDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getDynatraceHealthSource(hs map[string]interface{}) nextgen.DynatraceHealthSource {
	healthSource := &nextgen.DynatraceHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.DynatraceMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.DynatraceMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getCustomHealthSource(hs map[string]interface{}) nextgen.CustomHealthSourceMetric {
	healthSource := &nextgen.CustomHealthSourceMetric{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.CustomHealthMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.CustomHealthMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getSplunkHealthSource(hs map[string]interface{}) nextgen.SplunkMetricHealthSource {
	healthSource := &nextgen.SplunkMetricHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.SplunkMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.SplunkMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getCloudWatchHealthSource(hs map[string]interface{}) nextgen.CloudWatchMetricsHealthSource {
	healthSource := &nextgen.CloudWatchMetricsHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.CloudWatchMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.CloudWatchMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getAwsPrometheusHealthSource(hs map[string]interface{}) nextgen.AwsPrometheusHealthSource {
	healthSource := &nextgen.AwsPrometheusHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	metricDefinitions := hs["metricDefinitions"].([]interface{})
	healthSourceMetricDefinitions := make([]nextgen.PrometheusMetricDefinition, len(metricDefinitions))
	for i, metricDefinition := range metricDefinitions {
		data := nextgen.PrometheusMetricDefinition{}
		metricDef, errMarshal := json.Marshal(metricDefinition)
		if errMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		errUnMarshal := json.Unmarshal(metricDef, &data)
		if errUnMarshal != nil {
			panic(fmt.Sprintf("Invalid Health source %s", hs))
		}
		healthSourceMetricDefinitions[i] = data
	}
	healthSource.MetricDefinitions = healthSourceMetricDefinitions

	if hs["metricPacks"] != nil {
		healthSource.MetricPacks = getMetricPacks(hs, "metricPacks")
	}
	return *healthSource
}

func getNextGenHealthSource(hs map[string]interface{}) nextgen.NextGenHealthSource {
	healthSource := &nextgen.NextGenHealthSource{}

	healthSource.ConnectorRef = hs["connectorRef"].(string)

	healthSourceParamDto := nextgen.HealthSourceParamsDto{}
	healthSourceParamData, errMarshal := json.Marshal(hs["healthSourceParams"])
	if errMarshal != nil {
		panic(fmt.Sprintf("Invalid health source param dto %s", hs))
	}
	errUnMarshal := json.Unmarshal(healthSourceParamData, &healthSourceParamDto)
	if errUnMarshal != nil {
		panic(fmt.Sprintf("Invalid health source param dto %s", hs))
	}

	queryDefinitions := hs["queryDefinitions"].([]interface{})
	queryDefinitionDtos := make([]nextgen.QueryDefinition, len(queryDefinitions))
	for i, queryDefinition := range queryDefinitions {
		data := queryDefinition.(map[string]interface{})

		queryParams := nextgen.QueryParamsDto{}
		queryParamsData, _ := json.Marshal(hs["queryParamsDto"])
		json.Unmarshal(queryParamsData, &queryParams)

		riskProfile := nextgen.RiskProfile{}
		riskProfileData, _ := json.Marshal(hs["riskProfile"])
		json.Unmarshal(riskProfileData, &riskProfile)

		query := ""
		if hs["query"] != nil {
			query = hs["query"].(string)
		}
		liveMonitoringEnabled := false
		if data["liveMonitoringEnabled"] != nil {
			liveMonitoringEnabled, _ = strconv.ParseBool(data["liveMonitoringEnabled"].(string))
		}
		continuousVerificationEnabled := false
		if data["continuousVerificationEnabled"] != nil {
			continuousVerificationEnabled, _ = strconv.ParseBool(data["continuousVerificationEnabled"].(string))
		}
		sliEnabled := false
		if data["sliEnabled"] != nil {
			sliEnabled, _ = strconv.ParseBool(data["sliEnabled"].(string))
		}
		queryDefinitionDto := &nextgen.QueryDefinition{
			Identifier:                    data["identifier"].(string),
			Name:                          data["name"].(string),
			GroupName:                     data["groupName"].(string),
			LiveMonitoringEnabled:         liveMonitoringEnabled,
			ContinuousVerificationEnabled: continuousVerificationEnabled,
			SliEnabled:                    sliEnabled,
			Query:                         query,
			MetricThresholds:              getMetricThreshold(data),
			QueryParams:                   &queryParams,
			RiskProfile:                   &riskProfile,
		}
		queryDefinitionDtos[i] = *queryDefinitionDto
	}
	healthSource.HealthSourceParams = &healthSourceParamDto

	return *healthSource
}

func getMetricPacks(hs map[string]interface{}, path string) []nextgen.TimeSeriesMetricPackDto {
	metricPacks := hs[path].([]interface{})
	metricPackDto := make([]nextgen.TimeSeriesMetricPackDto, len(metricPacks))
	for i, metricPack := range metricPacks {
		test := metricPack.(map[string]interface{})
		timeSeriesMetricPackDto := &nextgen.TimeSeriesMetricPackDto{
			Identifier:       test["identifier"].(string),
			MetricThresholds: getMetricThreshold(test),
		}
		metricPackDto[i] = *timeSeriesMetricPackDto
	}
	return metricPackDto
}

func getMetricThreshold(hs map[string]interface{}) []nextgen.MetricThreshold {
	if hs["metricThresholds"] != nil {
		metricThresholds := hs["metricThresholds"].([]interface{})
		metricThresholdDto := make([]nextgen.MetricThreshold, len(metricThresholds))
		for j, metricThreshold := range metricThresholds {
			metricThresholdDto[j] = getMetricThresholdByType(metricThreshold.(map[string]interface{}))
		}
		return metricThresholdDto
	} else {
		return make([]nextgen.MetricThreshold, 0)
	}
}

func getMetricThresholdByType(hs map[string]interface{}) nextgen.MetricThreshold {
	metricThresholdType := hs["type"].(string)
	spec, errMarshal := json.Marshal(hs["spec"])
	if errMarshal != nil {
		panic(fmt.Sprintf("Invalid metric threshold %s", hs))
	}
	criteria := nextgen.MetricThresholdCriteria{}
	criteriaData, errMarshal := json.Marshal(hs["criteria"])
	if errMarshal != nil {
		panic(fmt.Sprintf("Invalid metric threshold %s", hs))
	}
	errUnMarshal := json.Unmarshal(criteriaData, &criteria)
	if errUnMarshal != nil {
		panic(fmt.Sprintf("Invalid metric threshold %s", hs))
	}
	groupName := ""
	if hs["groupName"] != nil {
		groupName = hs["groupName"].(string)
	}
	identifier := ""
	if hs["identifier"] != nil {
		identifier = hs["identifier"].(string)
	}
	if metricThresholdType == "FailImmediately" {
		data := nextgen.FailMetricThresholdSpec{}
		err := json.Unmarshal(spec, &data)
		if err != nil {
			panic(fmt.Sprintf("Invalid metric threshold %s", hs))
		}

		return nextgen.MetricThreshold{
			MetricName:       hs["metricName"].(string),
			MetricType:       hs["metricType"].(string),
			MetricIdentifier: identifier,
			GroupName:        groupName,
			Type_:            nextgen.MetricThresholdType(metricThresholdType),
			FailImmediately:  &data,
			Criteria:         &criteria,
		}
	}
	if metricThresholdType == "IgnoreThreshold" {
		data := nextgen.IgnoreMetricThresholdSpec{}
		err := json.Unmarshal(spec, &data)
		if err != nil {
			panic(fmt.Sprintf("Invalid metric threshold %s", hs))
		}
		return nextgen.MetricThreshold{
			MetricName:       hs["metricName"].(string),
			MetricType:       hs["metricType"].(string),
			MetricIdentifier: identifier,
			GroupName:        groupName,
			Type_:            nextgen.MetricThresholdType(metricThresholdType),
			IgnoreThreshold:  &data,
			Criteria:         &criteria,
		}
	}
	panic(fmt.Sprintf("Invalid metric threshold %s", hs))
}
