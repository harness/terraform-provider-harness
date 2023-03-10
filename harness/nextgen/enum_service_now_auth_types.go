package nextgen

type ServiceNowAuthType string

var ServiceNowAuthTypes = struct {
	ServiceNowUserNamePassword ServiceNowAuthType
	ServiceNowAdfs             ServiceNowAuthType
}{
	ServiceNowUserNamePassword: "UsernamePassword",
	ServiceNowAdfs:             "AdfsClientCredentialsWithCertificate",
}

var ServiceNowAuthTypeValues = []string{
	ServiceNowAuthTypes.ServiceNowUserNamePassword.String(),
	ServiceNowAuthTypes.ServiceNowAdfs.String(),
}

func (e ServiceNowAuthType) String() string {
	return string(e)
}
