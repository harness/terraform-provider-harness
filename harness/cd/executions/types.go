package executions

type ResponseMetadata struct{}

type ResponseMessage struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type ChangedBy struct {
	Uuid           string `json:"uuid,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	ExternalUserId string `json:"externalUserId,omitempty"`
}

type ExecutionItem struct {
	Uuid                     string      `json:"uuid,omitempty"`
	ApplicationId            string      `json:"appId,omitempty"`
	CreatedAt                string      `json:"createAt,omitempty"`
	LastUpdateAt             string      `json:"lastUpdateAt,omitempty"`
	ExecutionInterruptType   string      `json:"executionInterruptType,omitempty"`
	Seized                   bool        `json:"seized,omitempty"`
	EnvironmentId            string      `json:"environmentId,omitempty"`
	ExecutionUuid            string      `json:"executionUuid,omitempty"`
	StateExecutionInstanceId string      `json:"stateExecutionInstanceId,omitempty"`
	AccountId                string      `json:"accountId,omitempty"`
	Properties               interface{} `json:"properties,omitempty"`

	CreatedBy     ChangedBy `json:"createdBy,omitempty"`
	LastUpdatedBy ChangedBy `json:"lastUpdatedBy,omitempty"`
}

type Response struct {
	Metadata         *ResponseMetadata `json:"metaData"`
	Resource         *ExecutionItem    `json:"resource"`
	ResponseMessages []ResponseMessage `json:"responseMessages"`
}

const (
	ABORT = "ABORT_ALL"
)
