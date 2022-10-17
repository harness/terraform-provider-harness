package graphql

type DeploymentTagFilter struct {
	EntityType DeploymentTagType `json:"entityType,omitempty"`
	Tags       []DeploymentTag   `json:"tags,omitempty"`
}
