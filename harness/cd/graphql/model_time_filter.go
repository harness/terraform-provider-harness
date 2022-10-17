package graphql

type TimeFilter struct {
	Operator    TimeOperatorType `json:"operator,omitempty"`
	ValueMillis int64            `json:"value,omitempty"`
}
