package graphql

type AttachTagInput struct {
	ClientMutationId string        `json:"clientMutationId,omitempty"`
	EntityId         string        `json:"entityId,omitempty"`
	EntityType       TagEntityType `json:"entityType,omitempty"`
	Name             string        `json:"name,omitempty"`
	Value            string        `json:"value,omitempty"`
}
