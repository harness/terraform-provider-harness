package graphql

import "github.com/harness/harness-go-sdk/harness/time"

type ExportExecutionsPayload struct {
	ClientMutationId string                 `json:"clientMutationId,omitempty"`
	RequestId        string                 `json:"requestId,omitempty"`
	Status           ExportExecutionsStatus `json:"status,omitempty"`
	TotalExecutions  int                    `json:"totalExecutions,omitempty"`
	TriggeredAt      *time.Time             `json:"triggeredAt,omitempty"`
	StatusLink       string                 `json:"statusLink,omitempty"`
	DownloadLink     string                 `json:"downloadLink,omitempty"`
	ExpiresAt        *time.Time             `json:"expiresAt,omitempty"`
	ErrorMessage     string                 `json:"errorMessage,omitempty"`
}
