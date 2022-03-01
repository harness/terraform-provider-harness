package graphql

type PipelineExecution struct {
	ExecutionBase
	Pipeline                *Pipeline                 `json:"pipeline,omitempty"`
	PipelineStageExecutions []*PipelineStageExecution `json:"pipelineStageExecutions,omitempty"`
}
