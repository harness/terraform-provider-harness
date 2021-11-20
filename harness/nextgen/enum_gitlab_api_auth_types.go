package nextgen

type GitlabApiAuthType string

var GitlabApiAuthTypes = struct {
	Token GitlabApiAuthType
}{
	Token: "Token",
}

var GitlabApiAuthTypeValues = []string{
	GitlabApiAuthTypes.Token.String(),
}

func (e GitlabApiAuthType) String() string {
	return string(e)
}
