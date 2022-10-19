package graphql

type TimeOperatorType string

var TimeOperatorTypes = struct {
	Equals TimeOperatorType
	Before TimeOperatorType
	After  TimeOperatorType
}{
	Equals: "EQUALS",
	Before: "BEFORE",
	After:  "AFTER",
}
