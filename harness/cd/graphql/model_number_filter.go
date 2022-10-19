package graphql

type NumberFilter struct {
	Operator NumberOperatorType `json:"operator,omitempty"`
	Values   []int              `json:"values,omitempty"`
}
