package nextgen

type AzureCredentialType string

var AzureCredentialTypes = struct {
	InheritFromDelegate AzureCredentialType
	ManualConfig        AzureCredentialType
}{
	InheritFromDelegate: "InheritFromDelegate",
	ManualConfig:        "ManualConfig",
}

var AzureCredentialTypeValues = []string{
	AzureCredentialTypes.InheritFromDelegate.String(),
	AzureCredentialTypes.ManualConfig.String(),
}

func (e AzureCredentialType) String() string {
	return string(e)
}
