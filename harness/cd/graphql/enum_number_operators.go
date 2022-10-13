package graphql

type NumberOperatorType string

var NumberOperatorTypes = struct {
	Equals              NumberOperatorType
	GreaterThan         NumberOperatorType
	GreaterThanOrEquals NumberOperatorType
	In                  NumberOperatorType
	LessThan            NumberOperatorType
	LessThanOrEquals    NumberOperatorType
	NotEquals           NumberOperatorType
}{
	Equals:              "EQUALS",
	GreaterThan:         "GREATER_THAN",
	GreaterThanOrEquals: "GREATER_THAN_OR_EQUALS",
	In:                  "IN",
	LessThan:            "LESS_THAN",
	LessThanOrEquals:    "LESS_THAN_OR_EQUALS",
	NotEquals:           "NOT_EQUALS",
}
