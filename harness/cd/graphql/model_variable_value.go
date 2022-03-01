package graphql

type VariableValue struct {
	Type  VariableValueType `json:"type,omitempty"`
	Value string            `json:"value,omitempty"`
}
