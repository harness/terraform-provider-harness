package graphql

import "github.com/harness/harness-go-sdk/harness/time"

type ExecutionBase struct {
	Application *Application `json:"application,omitempty"`
	// Cause
	CreatedAt      *time.Time          `json:"createdAt,omitempty"`
	EndedAt        *time.Time          `json:"endedAt,omitempty"`
	FailureDetails string              `json:"failureDetails,omitempty"`
	Id             string              `json:"id,omitempty"`
	Notes          string              `json:"notes,omitempty"`
	StartedAt      *time.Time          `json:"startedAt,omitempty"`
	Status         ExecutionStatusType `json:"status,omitempty"`
	Tags           []*DeploymentTag    `json:"tags,omitempty"`
}
