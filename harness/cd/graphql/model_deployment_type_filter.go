package graphql

type DeploymentTypeFilter struct {
	Operator OperatorType     `json:"operator,omitempty"`
	Values   []DeploymentType `json:"values,omitempty"`
}
