package graphql

type StartExecutionInput struct {
	ApplicationId                string           `json:"applicationId,omitempty"`
	ClientMutationId             string           `json:"clientMutationId,omitempty"`
	ContinueWithDefaultValues    bool             `json:"continueWithDefaultValues,omitempty"`
	EntityId                     string           `json:"entityId,omitempty"`
	ExcludeHostsWithSameArtifact bool             `json:"excludeHostsWithSameArtifact,omitempty"`
	ExecutionType                ExecutionType    `json:"executionType,omitempty"`
	Notes                        string           `json:"notes,omitempty"`
	ServiceInputs                []*ServiceInput  `json:"serviceInputs,omitempty"`
	SpecificHosts                []string         `json:"specificHosts,omitempty"`
	TargetToSpecificHosts        bool             `json:"targetToSpecificHosts,omitempty"`
	VariableInputs               []*VariableInput `json:"variableInputs,omitempty"`
}
