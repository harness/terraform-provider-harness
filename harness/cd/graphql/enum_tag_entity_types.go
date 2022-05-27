package graphql

type TagEntityType string

var TagEntityTypes = struct {
	Application TagEntityType
	Environment TagEntityType
	Pipeline    TagEntityType
	Provisioner TagEntityType
	Service     TagEntityType
	Workflow    TagEntityType
}{
	Application: "APPLICATION",
	Environment: "ENVIRONMENT",
	Pipeline:    "PIPELINE",
	Provisioner: "PROVISIONER",
	Service:     "SERVICE",
	Workflow:    "WORKFLOW",
}

var TagEntityTypeValues = []string{
	TagEntityTypes.Application.String(),
	TagEntityTypes.Environment.String(),
	TagEntityTypes.Pipeline.String(),
	TagEntityTypes.Provisioner.String(),
	TagEntityTypes.Service.String(),
	TagEntityTypes.Workflow.String(),
}

func (e TagEntityType) String() string {
	return string(e)
}
