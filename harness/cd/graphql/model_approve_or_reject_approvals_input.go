package graphql

type ApproveOrRejectApprovalsInput struct {
	Action           ApprovalActionType  `json:"action,omitempty"`
	ApplicationId    string              `json:"applicationId,omitempty"`
	ApprovalId       string              `json:"approvalId,omitempty"`
	ClientMutationId string              `json:"clientMutationId,omitempty"`
	Comments         string              `json:"comments,omitempty"`
	ExecutionId      string              `json:"executionId,omitempty"`
	VariableInputs   []*ApprovalVariable `json:"variableInputs,omitempty"`
}
