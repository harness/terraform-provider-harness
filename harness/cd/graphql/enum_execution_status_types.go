package graphql

type ExecutionStatusType string

var ExecutionStatusTypes = struct {
	Aborted  ExecutionStatusType
	Error    ExecutionStatusType
	Expired  ExecutionStatusType
	Failed   ExecutionStatusType
	Paused   ExecutionStatusType
	Queued   ExecutionStatusType
	Rejected ExecutionStatusType
	Resumed  ExecutionStatusType
	Running  ExecutionStatusType
	Skipped  ExecutionStatusType
	Success  ExecutionStatusType
	Waiting  ExecutionStatusType
}{
	Aborted:  "ABORTED",
	Error:    "ERROR",
	Expired:  "EXPIRED",
	Failed:   "FAILED",
	Paused:   "PAUSED",
	Queued:   "QUEUED",
	Rejected: "REJECTED",
	Resumed:  "RESUMED",
	Running:  "RUNNING",
	Skipped:  "SKIPPED",
	Success:  "SUCCESS",
	Waiting:  "WAITING",
}

func (d ExecutionStatusType) String() string {
	return string(d)
}

var ExecutionStatusTypeList = []string{
	ExecutionStatusTypes.Aborted.String(),
	ExecutionStatusTypes.Error.String(),
	ExecutionStatusTypes.Expired.String(),
	ExecutionStatusTypes.Failed.String(),
	ExecutionStatusTypes.Paused.String(),
	ExecutionStatusTypes.Queued.String(),
	ExecutionStatusTypes.Rejected.String(),
	ExecutionStatusTypes.Resumed.String(),
	ExecutionStatusTypes.Running.String(),
	ExecutionStatusTypes.Skipped.String(),
	ExecutionStatusTypes.Success.String(),
	ExecutionStatusTypes.Waiting.String(),
}
