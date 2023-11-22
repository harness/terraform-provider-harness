package nextgen

type MetricThresholdType string

var MetricThresholdTypes = struct {
	FailImmediately MetricThresholdType
	IgnoreThreshold MetricThresholdType
}{
	FailImmediately: "FailImmediately",
	IgnoreThreshold: "IgnoreThreshold",
}

var MetricThresholdTypesSlice = []string{
	MetricThresholdTypes.FailImmediately.String(),
	MetricThresholdTypes.IgnoreThreshold.String(),
}

func (c MetricThresholdType) String() string {
	return string(c)
}
