package graphql

type PipelineStageExecution struct {
	PipelineStageElementId string              `json:"pipelineStageElementId,omitempty"`
	PipelineStageName      string              `json:"pipelineStageName,omitempty"`
	PipelineStepName       string              `json:"pipelineStepName,omitempty"`
	Status                 ExecutionStatusType `json:"status,omitempty"`
	ApprovalStepType       ApprovalStepType    `json:"approvalStepType,omitempty"`
	RuntimeInputVariables  []*Variable         `json:"runtimeInputVariables,omitempty"`
	WorkflowExecutionId    string              `json:"workflowExecutionId,omitempty"`
}
