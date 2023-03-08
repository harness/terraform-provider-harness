package nextgen

type JiraAuthType string

var JiraAuthTypes = struct {
	UsernamePassword JiraAuthType
}{
	UsernamePassword: "UsernamePassword",
}

func (e JiraAuthType) String() string {
	return string(e)
}
