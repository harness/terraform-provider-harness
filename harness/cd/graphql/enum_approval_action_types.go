package graphql

type ApprovalActionType string

var ApprovalActionTypes = struct {
	Approve ApprovalActionType
	Reject  ApprovalActionType
}{
	Approve: "APPROVE",
	Reject:  "REJECT",
}

var ApprovalActionTypeList = []string{
	ApprovalActionTypes.Approve.String(),
	ApprovalActionTypes.Reject.String(),
}

func (d ApprovalActionType) String() string {
	return string(d)
}
