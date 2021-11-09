package nextgen

type GitAuthType string

var GitAuthTypes = struct {
	Http GitAuthType
	Ssh  GitAuthType
}{
	Http: "Http",
	Ssh:  "Ssh",
}

var GitAuthTypeValues = []string{
	GitAuthTypes.Http.String(),
	GitAuthTypes.Ssh.String(),
}

func (e GitAuthType) String() string {
	return string(e)
}
