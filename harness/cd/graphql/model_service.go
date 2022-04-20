package graphql

import "github.com/harness/harness-go-sdk/harness/time"

type Service struct {
	Id              string            `json:"id"`
	ArtifactSources []*ArtifactSource `json:"artifactSources"`
	ArtifactType    ArtifactType      `json:"artifactType"`
	CreatedAt       *time.Time        `json:"createdAt"`
	CreatedBy       *User             `json:"createdBy,omitempty"`
	DeploymentType  DeploymentType    `json:"deploymentType"`
	Description     string            `json:"description"`
	Name            string            `json:"name"`
	Tags            []Tag             `json:"tags"`
}
