package nextgen

type ChangeSourceType string

var ChangeSourceTypes = struct {
	HarnessCDNextGen    ChangeSourceType
	PagerDuty           ChangeSourceType
	K8sCluster          ChangeSourceType
	HarnessCD           ChangeSourceType
}{
	HarnessCDNextGen:   "HarnessCDNextGen",
	PagerDuty:          "PagerDuty",
	K8sCluster:         "K8sCluster",
	HarnessCD:          "HarnessCD",
}

var ChangeSourceTypesSlice = []string{
	ChangeSourceTypes.HarnessCDNextGen.String(),
	ChangeSourceTypes.PagerDuty.String(),
	ChangeSourceTypes.K8sCluster.String(),
	ChangeSourceTypes.HarnessCD.String(),
}

func (c ChangeSourceType) String() string {
	return string(c)
}
