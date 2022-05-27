package graphql

type DetachTagInput struct {
	ClientMutationId string        `json:"clientMutationId,omitempty"`
	EntityId         string        `json:"entityId,omitempty"`
	EntityType       TagEntityType `json:"entityType,omitempty"`
	Name             string        `json:"name,omitempty"`
}
