package nextgen

type AwsKmsAuthType string

var AwsKmsAuthTypes = struct {
	AssumeIAMRole      AwsKmsAuthType
	AssumeSTSRole      AwsKmsAuthType
	ManualConfig       AwsKmsAuthType
	OidcAuthentication AwsKmsAuthType
}{
	AssumeIAMRole:      "AssumeIAMRole",
	AssumeSTSRole:      "AssumeSTSRole",
	ManualConfig:       "ManualConfig",
	OidcAuthentication: "OidcAuthentication",
}

var AwsKmsAuthTypeValues = []string{
	AwsKmsAuthTypes.AssumeIAMRole.String(),
	AwsKmsAuthTypes.AssumeSTSRole.String(),
	AwsKmsAuthTypes.ManualConfig.String(),
	AwsKmsAuthTypes.OidcAuthentication.String(),
}

func (e AwsKmsAuthType) String() string {
	return string(e)
}
