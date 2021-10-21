package nextgen

type AwsAuthType string

var AwsAuthTypes = struct {
	ManualConfig        AwsAuthType
	Irsa                AwsAuthType
	InheritFromDelegate AwsAuthType
}{
	ManualConfig:        "ManualConfig",
	Irsa:                "Irsa",
	InheritFromDelegate: "InheritFromDelegate",
}

func (e AwsAuthType) String() string {
	return string(e)
}
