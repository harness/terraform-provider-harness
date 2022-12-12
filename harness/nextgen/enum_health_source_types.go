package nextgen

type HealthSourceType string

var HealthSourceTypes = struct {
	AppDynamics        HealthSourceType
	NewRelic           HealthSourceType
	StackdriverLog     HealthSourceType
	Splunk             HealthSourceType
	Prometheus         HealthSourceType
	Stackdriver        HealthSourceType
	DatadogMetrics     HealthSourceType
	DatadogLog         HealthSourceType
	Dynatrace          HealthSourceType
	ErrorTracking      HealthSourceType
	CustomHealthMetric HealthSourceType
	CustomHealthLog    HealthSourceType
	SplunkMetric       HealthSourceType
	ElasticSearch      HealthSourceType
	CloudWatchMetrics  HealthSourceType
	AwsPrometheus      HealthSourceType
}{
	AppDynamics:        "AppDynamics",
	NewRelic:           "NewRelic",
	StackdriverLog:     "StackdriverLog",
	Splunk:             "Splunk",
	Prometheus:         "Prometheus",
	Stackdriver:        "Stackdriver",
	DatadogMetrics:     "DatadogMetrics",
	DatadogLog:         "DatadogLog",
	Dynatrace:          "Dynatrace",
	ErrorTracking:      "ErrorTracking",
	CustomHealthMetric: "CustomHealthMetric",
	CustomHealthLog:    "CustomHealthLog",
	SplunkMetric:       "SplunkMetric",
	ElasticSearch:      "ElasticSearch",
	CloudWatchMetrics:  "CloudWatchMetrics",
	AwsPrometheus:      "AwsPrometheus",
}

var HealthSourceTypesSlice = []string{
	HealthSourceTypes.AppDynamics.String(),
	HealthSourceTypes.NewRelic.String(),
	HealthSourceTypes.StackdriverLog.String(),
	HealthSourceTypes.Splunk.String(),
	HealthSourceTypes.Prometheus.String(),
	HealthSourceTypes.Stackdriver.String(),
	HealthSourceTypes.DatadogMetrics.String(),
	HealthSourceTypes.DatadogLog.String(),
	HealthSourceTypes.Dynatrace.String(),
	HealthSourceTypes.ErrorTracking.String(),
	HealthSourceTypes.CustomHealthMetric.String(),
	HealthSourceTypes.CustomHealthLog.String(),
	HealthSourceTypes.SplunkMetric.String(),
	HealthSourceTypes.ElasticSearch.String(),
	HealthSourceTypes.CloudWatchMetrics.String(),
	HealthSourceTypes.AwsPrometheus.String(),
}

func (c HealthSourceType) String() string {
	return string(c)
}
