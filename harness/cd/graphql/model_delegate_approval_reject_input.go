package graphql

type DelegateApprovalRejectInput struct {
	AccountId        string               `json:"accountId,omitempty"`
	ClientMutationId string               `json:"clientMutationId,omitempty"`
	DelegateApproval DelegateApprovalType `json:"delegateApproval,omitempty"`
	DelegateId       string               `json:"delegateId,omitempty"`
}
