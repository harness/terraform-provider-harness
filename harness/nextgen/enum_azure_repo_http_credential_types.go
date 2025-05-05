package nextgen

type AzureRepoHttpCredentialType string

var AzureRepoHttpCredentialTypes = struct {
	UsernameToken AzureRepoHttpCredentialType
}{
	UsernameToken: "UsernameToken",
}

var AzureRepoHttpCredentialTypeValues = []string{
	AzureRepoHttpCredentialTypes.UsernameToken.String(),
}

func (e AzureRepoHttpCredentialType) String() string {
	return string(e)
}
