package graphql

type DelegateApprovalType string

var DelegateApprovalTypes = struct {
	Activate DelegateApprovalType
	Reject   DelegateApprovalType
}{
	Activate: "ACTIVATE",
	Reject:   "REJECT",
}

var DelegateApprovalTypeValues = []string{
	DelegateApprovalTypes.Activate.String(),
	DelegateApprovalTypes.Reject.String(),
}

func (e DelegateApprovalType) String() string {
	return string(e)
}
