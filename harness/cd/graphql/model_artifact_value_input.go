package graphql

type ArtifactValueInput struct {
	ArtifactId                  *ArtifactIdInput                  `json:"artifactId,omitempty"`
	BuildNumber                 *BuildNumberInput                 `json:"buildNumber,omitempty"`
	ParameterizedArtifactSource *ParameterizedArtifactSourceInput `json:"parameterizedArtifactSource,omitempty"`
	ValueType                   ArtifactInputType                 `json:"valueType,omitempty"`
}
