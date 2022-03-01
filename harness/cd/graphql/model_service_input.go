package graphql

type ServiceInput struct {
	ArtifactValueInput *ArtifactValueInput `json:"artifactValueInput,omitempty"`
	ManifestValueInput *ManifestValueInput `json:"manifestValueInput,omitempty"`
	Name               string              `json:"name,omitempty"`
}
