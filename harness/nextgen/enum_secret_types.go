package nextgen

type SecretType string

var SecretTypes = struct {
	SecretFile       SecretType
	SecretText       SecretType
	SSHKey           SecretType
	WinRmCredentials SecretType
}{
	SecretFile:       "SecretFile",
	SecretText:       "SecretText",
	SSHKey:           "SSHKey",
	WinRmCredentials: "WinRmCredentials",
}

var SecretTypeValues = []string{
	SecretTypes.SecretFile.String(),
	SecretTypes.SecretText.String(),
	SecretTypes.SSHKey.String(),
	SecretTypes.WinRmCredentials.String(),
}

func (e SecretType) String() string {
	return string(e)
}
