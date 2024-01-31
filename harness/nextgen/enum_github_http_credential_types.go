package nextgen

type GithubHttpCredentialType string

var GithubHttpCredentialTypes = struct {
	UsernameToken GithubHttpCredentialType
	GithubApp     GithubHttpCredentialType
	Anonymous     GithubHttpCredentialType
}{
	UsernameToken: "UsernameToken",
	GithubApp:     "GithubApp",
	Anonymous:     "Anonymous",
}

var GithubHttpCredentialTypeValues = []string{
	GithubHttpCredentialTypes.UsernameToken.String(),
	GithubHttpCredentialTypes.GithubApp.String(),
	GithubHttpCredentialTypes.Anonymous.String(),
}

func (e GithubHttpCredentialType) String() string {
	return string(e)
}
