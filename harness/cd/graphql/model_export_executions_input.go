package graphql

type ExportExecutionsInput struct {
	ClientMutationId         string
	NotifyOnlyTriggeringUser bool
	UserGroupIds             []string
	Filters                  ExecutionFilter
}
