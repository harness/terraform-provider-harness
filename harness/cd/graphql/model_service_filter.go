package graphql

type ServiceFilter struct {
	Application    *IdFilter             `json:"application,omitempty"`
	DeploymentType *DeploymentTypeFilter `json:"deploymentType,omitempty"`
	Service        *IdFilter             `json:"service,omitempty"`
	Tag            *ServiceTagFilter     `json:"tag,omitempty"`
}
