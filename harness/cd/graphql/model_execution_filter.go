package graphql

type ExecutionFilter struct {
	Execution        IdFilter
	Application      IdFilter
	Service          IdFilter
	CloudProvider    IdFilter
	Environment      IdFilter
	Status           IdFilter
	EndTime          TimeFilter
	StartTime        TimeFilter
	Duration         NumberFilter
	TriggeredBy      IdFilter
	Trigger          IdFilter
	Workflow         IdFilter
	Pipeline         IdFilter
	CreationTime     TimeFilter
	Tag              DeploymentTagFilter
	ArtifactBuildNo  IdFilter
	HelmChartVersion IdFilter
}
