package nextgen

type AwsSecretManagerAuthType string

var AwsSecretManagerAuthTypes = struct {
	AssumeIAMRole      AwsSecretManagerAuthType
	AssumeSTSRole      AwsSecretManagerAuthType
	ManualConfig       AwsSecretManagerAuthType
	OidcAuthentication AwsSecretManagerAuthType
}{
	AssumeIAMRole:      "AssumeIAMRole",
	AssumeSTSRole:      "AssumeSTSRole",
	ManualConfig:       "ManualConfig",
	OidcAuthentication: "OidcAuthentication",
}

var AwsSecretManagerAuthTypeValues = []string{
	AwsSecretManagerAuthTypes.AssumeIAMRole.String(),
	AwsSecretManagerAuthTypes.AssumeSTSRole.String(),
	AwsSecretManagerAuthTypes.ManualConfig.String(),
	AwsSecretManagerAuthTypes.OidcAuthentication.String(),
}

func (e AwsSecretManagerAuthType) String() string {
	return string(e)
}
