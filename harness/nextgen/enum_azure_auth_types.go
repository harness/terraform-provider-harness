package nextgen

type AzureAuthType string

var AzureAuthTypes = struct {
	Certificate AzureAuthType
	SecretKey   AzureAuthType
}{
	Certificate: "Certificate",
	SecretKey:   "Secret",
}

var AzureAuthTypeValues = []string{
	AzureAuthTypes.Certificate.String(),
	AzureAuthTypes.SecretKey.String(),
}

func (e AzureAuthType) String() string {
	return string(e)
}
