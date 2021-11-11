package graphql

type DelegateStatusType string

var DelegateStatusTypes = struct {
	Deleted            DelegateStatusType
	Enabled            DelegateStatusType
	WaitingForApproval DelegateStatusType
}{
	Deleted:            "DELETED",
	Enabled:            "ENABLED",
	WaitingForApproval: "WAITING_FOR_APPROVAL",
}

var DelegateStatusTypeValues = []string{
	DelegateStatusTypes.Deleted.String(),
}

func (e DelegateStatusType) String() string {
	return string(e)
}
