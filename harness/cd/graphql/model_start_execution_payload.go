package graphql

type StartExecutionPayload struct {
	ClientMutationId  string             `json:"clientMutationId,omitempty"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution,omitempty"`
	PipelineExecution *PipelineExecution `json:"pipelineExecution,omitempty"`
	WarningMessage    string             `json:"warningMessage,omitempty"`
}
