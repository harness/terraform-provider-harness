package nextgen

type SLIEvaluationType string

var SLIEvaluationTypes = struct {
	Window     SLIEvaluationType
	Request    SLIEvaluationType
	MetricLess SLIEvaluationType
}{
	Window:     "Window",
	Request:    "Request",
	MetricLess: "MetricLess",
}

var SLIEvaluationTypesSlice = []string{
	SLIEvaluationTypes.Window.String(),
	SLIEvaluationTypes.Request.String(),
	SLIEvaluationTypes.MetricLess.String(),
}

func (c SLIEvaluationType) String() string {
	return string(c)
}
