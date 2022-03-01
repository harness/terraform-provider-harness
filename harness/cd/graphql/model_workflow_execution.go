package graphql

type WorkflowExecution struct {
	ExecutionBase
	Artifacts         []*Artifact `json:"artifacts,omitempty"`
	RollbackArtifacts []*Artifact `json:"rollbackArtifacts,omitempty"`
}
