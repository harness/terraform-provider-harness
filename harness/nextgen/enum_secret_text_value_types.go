package nextgen

type SecretTextValueType string

var SecretTextValueTypes = struct {
	Inline    SecretTextValueType
	Reference SecretTextValueType
}{
	Inline:    "Inline",
	Reference: "Reference",
}

var SecretTextValueTypeValues = []string{
	SecretTextValueTypes.Inline.String(),
	SecretTextValueTypes.Reference.String(),
}

func (e SecretTextValueType) String() string {
	return string(e)
}
