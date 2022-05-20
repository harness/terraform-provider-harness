package nextgen

type SecretSpecType string

var SecretSpecTypes = struct {
	File   SecretSpecType
	SSHKey SecretSpecType
	Text   SecretSpecType
}{
	File:   "SecretFileSpe",
	SSHKey: "SSHKeySpec",
	Text:   "SecretTextSpec",
}

var SecretSpecTypeValues = []string{
	SecretSpecTypes.File.String(),
	SecretSpecTypes.SSHKey.String(),
	SecretSpecTypes.Text.String(),
}

func (e SecretSpecType) String() string {
	return string(e)
}
