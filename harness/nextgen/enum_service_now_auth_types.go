package nextgen

type ServiceNowAuthType string

var ServiceNowAuthTypes = struct {
	ServiceNowUserNamePassword ServiceNowAuthType
	ServiceNowAdfs             ServiceNowAuthType
	ServiceNowRefreshToken     ServiceNowAuthType
}{
	ServiceNowUserNamePassword: "UsernamePassword",
	ServiceNowAdfs:             "AdfsClientCredentialsWithCertificate",
	ServiceNowRefreshToken:     "RefreshTokenGrantType",  
}

var ServiceNowAuthTypeValues = []string{
	ServiceNowAuthTypes.ServiceNowUserNamePassword.String(),
	ServiceNowAuthTypes.ServiceNowAdfs.String(),
	ServiceNowAuthTypes.ServiceNowRefreshToken.String(),
}

func (e ServiceNowAuthType) String() string {
	return string(e)
}
