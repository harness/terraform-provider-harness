package nextgen

type AwsKmsAuthType string

var AwsKmsAuthTypes = struct {
	AssumeIAMRole AwsKmsAuthType
	AssumeSTSRole AwsKmsAuthType
	ManualConfig  AwsKmsAuthType
}{
	AssumeIAMRole: "AssumeIAMRole",
	AssumeSTSRole: "AssumeSTSRole",
	ManualConfig:  "ManualConfig",
}

var AwsKmsAuthTypeValues = []string{
	AwsKmsAuthTypes.AssumeIAMRole.String(),
	AwsKmsAuthTypes.AssumeSTSRole.String(),
	AwsKmsAuthTypes.ManualConfig.String(),
}

func (e AwsKmsAuthType) String() string {
	return string(e)
}
