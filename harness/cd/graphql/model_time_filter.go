package graphql

type TimeFilter struct {
	Operator TimeOperatorType `json:"operator,omitempty"`
	Value    int64            `json:"value,omitempty"`
}
