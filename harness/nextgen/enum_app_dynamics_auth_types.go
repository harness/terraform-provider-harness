package nextgen

type AppDynamicsAuthType string

var AppDynamicsAuthTypes = struct {
	ApiClientToken   AppDynamicsAuthType
	UsernamePassword AppDynamicsAuthType
}{
	ApiClientToken:   "ApiClientToken",
	UsernamePassword: "UsernamePassword",
}

var AppDynamicsAuthTypeValues = []string{
	AppDynamicsAuthTypes.ApiClientToken.String(),
	AppDynamicsAuthTypes.UsernamePassword.String(),
}

func (e AppDynamicsAuthType) String() string {
	return string(e)
}
