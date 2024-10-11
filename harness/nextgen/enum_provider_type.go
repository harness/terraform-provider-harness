package nextgen

type ProviderType string

var ProviderTypes = struct {
	BitbucketServer ProviderType
}{
	BitbucketServer: "BITBUCKET_SERVER",
}

var ProviderTypeValues = []string{
	ProviderTypes.BitbucketServer.String(),
}

func (e ProviderType) String() string {
	return string(e)
}
