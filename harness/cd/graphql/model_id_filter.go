package graphql

type IdFilter struct {
	Operator IdOperatorType `json:"operator,omitempty"`
	Values   []string       `json:"values,omitempty"`
}
