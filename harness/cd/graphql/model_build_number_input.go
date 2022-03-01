package graphql

type BuildNumberInput struct {
	ArtifactSourceName string `json:"artifactSourceName,omitempty"`
	BuildNumber        string `json:"buildNumber,omitempty"`
}
