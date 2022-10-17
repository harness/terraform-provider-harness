package graphql

type ExportExecutionsInput struct {
	ClientMutationId         string             `json:"clientMutationId,omitempty"`
	NotifyOnlyTriggeringUser bool               `json:"notifyOnlyTriggeringUser,omitempty"`
	UserGroupIds             []string           `json:"userGroupIds,omitempty"`
	Filters                  []*ExecutionFilter `json:"filters"`
}
