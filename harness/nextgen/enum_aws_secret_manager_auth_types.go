package nextgen

type AwsSecretManagerAuthType string

var AwsSecretManagerAuthTypes = struct {
	AssumeIAMRole AwsSecretManagerAuthType
	AssumeSTSRole AwsSecretManagerAuthType
	ManualConfig  AwsSecretManagerAuthType
}{
	AssumeIAMRole: "AssumeIAMRole",
	AssumeSTSRole: "AssumeSTSRole",
	ManualConfig:  "ManualConfig",
}

var AwsSecretManagerAuthTypeValues = []string{
	AwsSecretManagerAuthTypes.AssumeIAMRole.String(),
	AwsSecretManagerAuthTypes.AssumeSTSRole.String(),
	AwsSecretManagerAuthTypes.ManualConfig.String(),
}

func (e AwsSecretManagerAuthType) String() string {
	return string(e)
}
