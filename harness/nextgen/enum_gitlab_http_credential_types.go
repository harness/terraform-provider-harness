package nextgen

type GitlabHttpCredentialType string

var GitlabHttpCredentialTypes = struct {
	UsernamePassword GitlabHttpCredentialType
	UsernameToken    GitlabHttpCredentialType
}{
	UsernamePassword: "UsernamePassword",
	UsernameToken:    "UsernameToken",
}

var GitlabHttpCredentialTypeValues = []string{
	GitlabHttpCredentialTypes.UsernamePassword.String(),
	GitlabHttpCredentialTypes.UsernameToken.String(),
}

func (e GitlabHttpCredentialType) String() string {
	return string(e)
}
