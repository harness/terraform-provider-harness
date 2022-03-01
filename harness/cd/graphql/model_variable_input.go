package graphql

type VariableInput struct {
	Name          string         `json:"name,omitempty"`
	VariableValue *VariableValue `json:"variableValue,omitempty"`
}
