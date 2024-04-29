package nextgen

type SecretTextValueType string

var SecretTextValueTypes = struct {
	Inline                    SecretTextValueType
	Reference                 SecretTextValueType
	CustomSecretManagerValues SecretTextValueType
}{
	Inline:                    "Inline",
	Reference:                 "Reference",
	CustomSecretManagerValues: "CustomSecretManagerValues",
}

var SecretTextValueTypeValues = []string{
	SecretTextValueTypes.Inline.String(),
	SecretTextValueTypes.Reference.String(),
}

func (e SecretTextValueType) String() string {
	return string(e)
}
