package graphql

type VersionNumberInput struct {
	AppManifestName string `json:"appManifestName,omitempty"`
	VersionNumber   string `json:"versionNumber,omitempty"`
}
