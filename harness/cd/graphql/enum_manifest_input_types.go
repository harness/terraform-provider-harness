package graphql

type ManifestInputType string

var ManifestInputTypes = struct {
	HelmChartId   ManifestInputType
	VersionNumber ManifestInputType
}{
	HelmChartId:   "HELM_CHART_ID",
	VersionNumber: "VERSION_NUMBER",
}

var ManifestInputTypeList = []string{
	ManifestInputTypes.HelmChartId.String(),
	ManifestInputTypes.VersionNumber.String(),
}

func (d ManifestInputType) String() string {
	return string(d)
}
