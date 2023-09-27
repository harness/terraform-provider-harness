package nextgen

type GithubHttpCredentialType string

var GithubHttpCredentialTypes = struct {
	UsernameToken GithubHttpCredentialType
	GithubApp     GithubHttpCredentialType
}{
	UsernameToken: "UsernameToken",
	GithubApp:     "GithubApp",
}

var GithubHttpCredentialTypeValues = []string{
	GithubHttpCredentialTypes.UsernameToken.String(),
	GithubHttpCredentialTypes.GithubApp.String(),
}

func (e GithubHttpCredentialType) String() string {
	return string(e)
}
