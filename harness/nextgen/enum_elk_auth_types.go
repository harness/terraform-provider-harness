package nextgen

type ElkAuthType string

var ElkAuthTypes = struct {
	None             ElkAuthType
	UsernamePassword ElkAuthType
	ApiClientToken   ElkAuthType
}{
	None:             "None",
	UsernamePassword: "UsernamePassword",
	ApiClientToken:   "ApiClientToken",
}

var ElkAuthTypeValues = []string{
	ElkAuthTypes.None.String(),
	ElkAuthTypes.UsernamePassword.String(),
	ElkAuthTypes.ApiClientToken.String(),
}

func (e ElkAuthType) String() string {
	return string(e)
}
