package nextgen

type SLIMetricType string

var SLIMetricTypes = struct {
	Threshold SLIMetricType
	Ratio     SLIMetricType
}{
	Threshold: "Threshold",
	Ratio:     "Ratio",
}

var SLIMetricTypesSlice = []string{
	SLIMetricTypes.Threshold.String(),
	SLIMetricTypes.Ratio.String(),
}

func (c SLIMetricType) String() string {
	return string(c)
}
