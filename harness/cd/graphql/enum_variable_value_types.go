package graphql

type VariableValueType string

var VariableValueTypes = struct {
	Expression VariableValueType
	Id         VariableValueType
	Name       VariableValueType
}{
	Expression: "EXPRESSION",
	Id:         "ID",
	Name:       "NAME",
}

func (d VariableValueType) String() string {
	return string(d)
}

var VariableValueTypeList = []string{
	VariableValueTypes.Expression.String(),
	VariableValueTypes.Id.String(),
	VariableValueTypes.Name.String(),
}
