package nextgen

type SecretType string

var SecretTypes = struct {
	SecretFile SecretType
	SecretText SecretType
	SSHKey     SecretType
}{
	SecretFile: "SecretFile",
	SecretText: "SecretText",
	SSHKey:     "SSHKey",
}

var SecretTypeValues = []string{
	SecretTypes.SecretFile.String(),
	SecretTypes.SecretText.String(),
	SecretTypes.SSHKey.String(),
}

func (e SecretType) String() string {
	return string(e)
}
