package nextgen

type ConnectorStatus string

var ConnectorStatuses = struct {
	Success ConnectorStatus
	Failure ConnectorStatus
	Partial ConnectorStatus
	Unknown ConnectorStatus
}{
	Success: "SUCCESS",
	Failure: "FAILURE",
	Partial: "PARTIAL",
	Unknown: "UNKNOWN",
}

var ConnectorStatusSlice = []string{
	ConnectorStatuses.Success.String(),
	ConnectorStatuses.Failure.String(),
	ConnectorStatuses.Partial.String(),
	ConnectorStatuses.Unknown.String(),
}

func (c ConnectorStatus) String() string {
	return string(c)
}
