package nextgen

type TasCredentialType string

var TasCredentialTypes = struct {
	ManualConfig TasCredentialType
}{
	ManualConfig: "ManualConfig",
}

var TasCredentialTypeValues = []string{
	TasCredentialTypes.ManualConfig.String(),
}

func (e TasCredentialType) String() string {
	return string(e)
}
