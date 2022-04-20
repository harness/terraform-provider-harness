package graphql

type ServiceTagFilter struct {
	EntityType ServiceTagType `json:"entityType,omitempty"`
	Tags       []*TagInput    `json:"tags,omitempty"`
}
