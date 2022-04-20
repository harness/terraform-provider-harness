package graphql

type ArtifactSource struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Artifacts ArtifactConnection
}
