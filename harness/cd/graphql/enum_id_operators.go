package graphql

type IdOperatorType string

var IdOperatorTypes = struct {
	Equals  IdOperatorType
	In      IdOperatorType
	Like    IdOperatorType
	NotIn   IdOperatorType
	NotNull IdOperatorType
}{
	Equals:  "EQUALS",
	In:      "IN",
	Like:    "LIKE",
	NotIn:   "NOT_IN",
	NotNull: "NOT_NULL",
}

var IdOperatorTypeValues = []string{
	IdOperatorTypes.Equals.String(),
	IdOperatorTypes.In.String(),
	IdOperatorTypes.Like.String(),
	IdOperatorTypes.NotIn.String(),
	IdOperatorTypes.NotNull.String(),
}

func (e IdOperatorType) String() string {
	return string(e)
}
