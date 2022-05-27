package graphql

type TagLink struct {
	ApplicationId string        `json:"appId,omitempty"`
	EntityId      string        `json:"entityId,omitempty"`
	EntityType    TagEntityType `json:"entityType,omitempty"`
	Name          string        `json:"name,omitempty"`
	Value         string        `json:"value,omitempty"`
}
