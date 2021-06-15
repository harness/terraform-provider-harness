package client

import "time"

type CommonMetadata struct {
	CreatedAt   *Time  `json:"createdAt,omitempty"`
	CreatedBy   *User  `json:"createdBy,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
}

type WorkflowConnection struct {
	Nodes    []Workflow
	PageInfo PageInfo
}

type Workflow struct {
	CommonMetadata
	WorkflowVariables []Variable
}

type ServiceConnection struct {
	Nodes    []Service
	PageInfo PageInfo
}

type Service struct {
	CommonMetadata
	ArtifactSources []ArtifactSource
	ArtifactType    string
}

type ArtifactSource struct {
	CommonMetadata
	Artifacts ArtifactConnection
}

type ArtifactConnection struct {
	Nodes    []Artifact
	PageInfo PageInfo
}

type Artifact struct {
	ArtifactSource ArtifactSource
	BuildNo        string
	CollectedAt    time.Time
	Id             string
}

type PipelineConnection struct {
	Nodes    []Pipeline
	PageInfo PageInfo
}

type Pipeline struct {
	CommonMetadata
	PipelineVariables []Variable
}

type Variable struct {
	AllowMultipleVariables bool
	AllowedValues          []string
	DefaultValue           string
	Fixed                  bool
	Name                   string
	Required               bool
	Type                   string
}

type EnvironmentConnection struct {
	Nodes    []Environment
	PageInfo PageInfo
}

type Environment struct {
	Application               Application
	CreatedAt                 string
	CreatedBy                 User
	Description               string
	Id                        string
	InfrastructureDefinitions InfrastructureDefinitionConnection
	Name                      string
	Tags                      []Tag
	Type                      string
}

type InfrastructureDefinitionConnection struct {
	Nodes    []InfrastructureDefinition
	PageInfo PageInfo
}

type PageInfo struct {
	HasMore bool
	Limit   int
	Offset  int
	Total   int
}

type InfrastructureDefinition struct {
	CreatedAt        string
	DeploymentType   string
	Environment      Environment
	Id               string
	Name             string
	ScopedToServices []string
}

type Tag struct {
	Name  string
	Value string
}

type User struct {
	Email                            string `json:"email,omitempty"`
	Id                               string `json:"id,omitempty"`
	IsEmailVerified                  bool   `json:"isEmailVerified,omitempty"`
	IsImportedFromIdentityProvider   bool   `json:"isImportedFromIdentityProvider,omitempty"`
	IsPasswordExpired                bool   `json:"isPasswordExpired,omitempty"`
	IsTwoFactorAuthenticationEnabled bool   `json:"isTwoFactorAuthenticationEnabled,omitempty"`
	IsUserLocked                     bool   `json:"isUserLocked,omitempty"`
	Name                             string `json:"name,omitempty"`
}

type GitSyncConfig struct {
	Branch         string
	GitConnector   *GitConnector
	RepositoryName string
	SyncEnabled    bool
}

type UsageScope struct {
	AppEnvScopes []*AppEnvScope `json:"appEnvScopes,omitempty"`
}

type AppEnvScope struct {
	Application *AppScopeFilter `json:"application,omitempty"`
	Environment *EnvScopeFilter `json:"environment,omitempty"`
}

type AppScopeFilter struct {
	AppId      string `json:"appId,omitempty"`
	FilterType string `json:"filterType,omitempty"`
}

type EnvScopeFilter struct {
	EnvId      string `json:"envId,omitempty"`
	FilterType string `json:"filterType,omitempty"`
}
