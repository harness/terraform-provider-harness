package nextgen

type GithubHttpCredentialType string

var GithubHttpCredentialTypes = struct {
	UsernameToken GithubHttpCredentialType
}{
	UsernameToken: "UsernameToken",
}

var GithubHttpCredentialTypeValues = []string{
	GithubHttpCredentialTypes.UsernameToken.String(),
}

func (e GithubHttpCredentialType) String() string {
	return string(e)
}
