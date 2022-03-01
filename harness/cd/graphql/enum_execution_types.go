package graphql

type ExecutionType string

var ExecutionTypes = struct {
	Workflow ExecutionType
	Pipeline ExecutionType
}{
	Workflow: "WORKFLOW",
	Pipeline: "PIPELINE",
}

func (d ExecutionType) String() string {
	return string(d)
}

var ExecutionStatusSlice = []string{
	ExecutionTypes.Pipeline.String(),
	ExecutionTypes.Workflow.String(),
}
