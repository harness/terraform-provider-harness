package graphql

type ApplicationClient struct {
	APIClient *ApiClient
}

type Application struct {
	CommonMetadata
	ClientMutationId          string                   `json:"clientMutationId,omitempty"`
	Description               string                   `json:"description,omitempty"`
	Name                      string                   `json:"name,omitempty"`
	Environments              []*EnvironmentConnection `json:"environments,omitempty"`
	GitSyncConfig             *GitSyncConfig           `json:"gitSyncConfig,omitempty"`
	IsManualTriggerAuthorized bool                     `json:"isManualTriggerAuthorized"`
	Pipelines                 *PipelineConnection      `json:"pipelines,omitempty"`
	Services                  *ServiceConnection       `json:"services,omitempty"`
	Workflows                 *WorkflowConnection      `json:"workflows,omitempty"`
}

type Applications struct {
	PageInfo `json:"pageInfo"`
	Nodes    []Application `json:"nodes"`
}

type UpdateApplicationInput struct {
	ApplicationId             string `json:"applicationId"`
	ClientMutationId          string `json:"clientMutationId"`
	Description               string `json:"description"`
	IsManualTriggerAuthorized bool   `json:"isManualTriggerAuthorized"`
	Name                      string `json:"name"`
}

type UpdateApplicationPayload struct {
	Application      *Application `json:"application"`
	ClientMutationId string       `json:"clientMutationId"`
}

// type CreateApplicationInput struct {
// 	ClientMutationId          string `json:"clientMutationId"`
// 	Description               string `json:"description"`
// 	IsManualTriggerAuthorized bool   `json:"isManualTriggerAuthorized"`
// 	Name                      string `json:"name"`
// }

type CreateApplicationPayload struct {
	Application      *Application `json:"application"`
	ClientMutationId string       `json:"clientMutationId"`
}

type DeleteApplicationInput struct {
	ApplicationId    string `json:"applicationId"`
	ClientMutationId string `json:"clientMutationId"`
}

type DeleteApplicationPayload struct {
	ClientMutationId string `json:"clientMutationId"`
}
