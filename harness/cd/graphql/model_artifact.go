package graphql

import "github.com/harness/harness-go-sdk/harness/time"

type Artifact struct {
	ArtifactSource ArtifactSource
	BuildNo        string
	CollectedAt    time.Time
	Id             string
}
