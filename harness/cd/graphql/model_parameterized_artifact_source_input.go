package graphql

type ParameterizedArtifactSourceInput struct {
	ArtifactSourceName   string                 `json:"artifactSourceName,omitempty"`
	BuildNumber          string                 `json:"buildNumber,omitempty"`
	ParameterValueInputs []*ParameterValueInput `json:"parameterValueInputs,omitempty"`
}
