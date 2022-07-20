package nextgen

type AttributeFilter struct {
	AttributeName   string   `json:"attributeName,omitempty"`
	AttributeValues []string `json:"attributeValues,omitempty"`
}
