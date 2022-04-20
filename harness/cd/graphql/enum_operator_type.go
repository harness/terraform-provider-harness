package graphql

type OperatorType string

var OperatorTypes = struct {
	Equals OperatorType
	In     OperatorType
}{
	Equals: "EQUALS",
	In:     "IN",
}

var OperatorTypeValues = []string{
	OperatorTypes.Equals.String(),
	OperatorTypes.In.String(),
}

func (e OperatorType) String() string {
	return string(e)
}
