package nextgen

type AzureRepoConnectorType string

var AzureRepoConnectorTypes = struct {
	Project AzureRepoConnectorType
	Repo    AzureRepoConnectorType
}{
	Project: "Project",
	Repo:    "Repo",
}

var AzureRepoConnectorTypeValues = []string{
	AzureRepoConnectorTypes.Project.String(),
	AzureRepoConnectorTypes.Repo.String(),
}

func (e AzureRepoConnectorType) String() string {
	return string(e)
}
