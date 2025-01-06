package nextgen

type AzureArtifactsAuthType string

var AzureArtifactsAuthTypes = struct {
	PersonalAccessToken AzureArtifactsAuthType
}{
	PersonalAccessToken: "PersonalAccessToken",
}

var AzureArtifactsAuthTypeValues = []string{
	AzureArtifactsAuthTypes.PersonalAccessToken.String(),
}

func (e AzureArtifactsAuthType) String() string {
	return string(e)
}
