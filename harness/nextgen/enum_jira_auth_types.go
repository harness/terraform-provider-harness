package nextgen

type JiraAuthType string

var JiraAuthTypes = struct {
	UsernamePassword JiraAuthType
	PersonalAccessToken JiraAuthType
}{
	UsernamePassword: "UsernamePassword",
	PersonalAccessToken: "PersonalAccessToken",
}

var JiraAuthTypeValues = []string{
	JiraAuthTypes.UsernamePassword.String(),
	JiraAuthTypes.PersonalAccessToken.String(),
}

func (e JiraAuthType) String() string {
	return string(e)
}
