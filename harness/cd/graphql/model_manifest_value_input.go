package graphql

type ManifestValueInput struct {
	HelmChartId   string              `json:"helmChartId,omitempty"`
	ValueType     ManifestInputType   `json:"valueType,omitempty"`
	VersionNumber *VersionNumberInput `json:"versionNumber,omitempty"`
}
