package graphql

type ExecutionFilter struct {
	Execution        *IdFilter            `json:"execution,omitempty"`
	Application      *IdFilter            `json:"application,omitempty"`
	Service          *IdFilter            `json:"service,omitempty"`
	CloudProvider    *IdFilter            `json:"cloudProvider,omitempty"`
	Environment      *IdFilter            `json:"environment,omitempty"`
	Status           *IdFilter            `json:"status,omitempty"`
	EndTime          *TimeFilter          `json:"endTime,omitempty"`
	StartTime        *TimeFilter          `json:"startTime,omitempty"`
	Duration         *NumberFilter        `json:"duration,omitempty"`
	TriggeredBy      *IdFilter            `json:"triggeredBy,omitempty"`
	Trigger          *IdFilter            `json:"trigger,omitempty"`
	Workflow         *IdFilter            `json:"workflow,omitempty"`
	Pipeline         *IdFilter            `json:"pipeline,omitempty"`
	CreationTime     *TimeFilter          `json:"creationTime,omitempty"`
	Tag              *DeploymentTagFilter `json:"tag,omitempty"`
	ArtifactBuildNo  *IdFilter            `json:"artifactBuildNo,omitempty"`
	HelmChartVersion *IdFilter            `json:"helmChartVersion,omitempty"`
}
