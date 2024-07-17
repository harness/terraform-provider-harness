package nextgen

type ChangeSourceType string

var ChangeSourceTypes = struct {
	HarnessCDNextGen ChangeSourceType
	PagerDuty        ChangeSourceType
	K8sCluster       ChangeSourceType
	HarnessCD        ChangeSourceType
	CustomDeploy     ChangeSourceType
	CustomIncident   ChangeSourceType
	CustomFF         ChangeSourceType
}{
	HarnessCDNextGen: "HarnessCDNextGen",
	PagerDuty:        "PagerDuty",
	K8sCluster:       "K8sCluster",
	HarnessCD:        "HarnessCD",
	CustomDeploy:     "CustomDeploy",
	CustomIncident:   "CustomIncident",
	CustomFF:         "CustomFF",
}

var ChangeSourceTypesSlice = []string{
	ChangeSourceTypes.HarnessCDNextGen.String(),
	ChangeSourceTypes.PagerDuty.String(),
	ChangeSourceTypes.K8sCluster.String(),
	ChangeSourceTypes.HarnessCD.String(),
	ChangeSourceTypes.CustomDeploy.String(),
	ChangeSourceTypes.CustomIncident.String(),
	ChangeSourceTypes.CustomFF.String(),
}

func (c ChangeSourceType) String() string {
	return string(c)
}
