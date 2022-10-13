package graphql

type NumberFilter struct {
	Operator NumberOperatorType
	Values   []int
}
