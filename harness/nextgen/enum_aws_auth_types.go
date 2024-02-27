package nextgen

type AwsAuthType string

var AwsAuthTypes = struct {
	ManualConfig        AwsAuthType
	Irsa                AwsAuthType
	InheritFromDelegate AwsAuthType
	OidcAuthentication  AwsAuthType
}{
	ManualConfig:        "ManualConfig",
	Irsa:                "Irsa",
	InheritFromDelegate: "InheritFromDelegate",
	OidcAuthentication:  "OidcAuthentication",
}

func (e AwsAuthType) String() string {
	return string(e)
}
