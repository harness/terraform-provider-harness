package client

type ApplicationClient struct {
	APIClient *ApiClient
}

type Application struct {
	CommonMetadata
	Environments              []*EnvironmentConnection
	GitSyncConfig             *GitSyncConfig
	IsManualTriggerAuthorized bool
	Pipelines                 *PipelineConnection
	Services                  *ServiceConnection
	Workflows                 *WorkflowConnection
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

type CreateApplicationInput struct {
	ClientMutationId          string `json:"clientMutationId"`
	Description               string `json:"description"`
	IsManualTriggerAuthorized bool   `json:"isManualTriggerAuthorized"`
	Name                      string `json:"name"`
}

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
