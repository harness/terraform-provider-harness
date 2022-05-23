package nextgen

type ResponseStatusType string

var ResponseStatusTypes = struct {
	Error   ResponseStatusType
	Failure ResponseStatusType
	Success ResponseStatusType
}{
	Error:   "ERROR",
	Failure: "FAILURE",
	Success: "SUCCESS",
}

var ResponseStatusTypeValues = []string{
	ResponseStatusTypes.Error.String(),
	ResponseStatusTypes.Failure.String(),
	ResponseStatusTypes.Success.String(),
}

func (e ResponseStatusType) String() string {
	return string(e)
}
