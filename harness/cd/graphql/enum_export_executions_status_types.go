package graphql

type ExportExecutionsStatus string

var ExportExecutionsStatusTypes = struct {
	Queued, Ready, Failed, Expired ExportExecutionsStatus
}{
	Queued:  "QUEUED",
	Ready:   "READY",
	Failed:  "FAILED",
	Expired: "EXPIRED",
}

func (e ExportExecutionsStatus) String() string {
	return string(e)
}
