package nextgen

type AzureRepoApiAuthType string

var AzureRepoApiAuthTypes = struct {
	Token AzureRepoApiAuthType
}{
	Token: "Token",
}

var AzureRepoApiAuthTypeValues = []string{
	AzureRepoApiAuthTypes.Token.String(),
}

func (e AzureRepoApiAuthType) String() string {
	return string(e)
}
