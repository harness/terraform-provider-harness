package client

import "time"

type CommonMetadata struct {
	CreatedAt   time.Time
	CreatedBy   User
	Id          string
	Name        string
	Description string
	Tags        []Tag
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
	Email                            string
	Id                               string
	IsEmailVerified                  bool
	IsImportedFromIdentityProvider   bool
	IsPasswordExpired                bool
	IsTwoFactorAuthenticationEnabled bool
	IsUserLocked                     bool
	Name                             string
}

type GitSyncConfig struct {
	Branch         string
	GitConnector   *GitConnector
	RepositoryName string
	SyncEnabled    bool
}

type GitConnector struct {
	Url                 string
	Branch              string
	CreatedAt           string
	CreatedBy           *User
	CustomCommitDetails *CustomCommitDetails
	DelegateSelectors   []string
	Description         string
	GenerateWebhookUrl  bool
	Id                  string
	Name                string
	PasswordSecretId    string
	SshSettingId        string
	UrlType             string
	UsageScope          *UsageScope
	UserName            string
	WebhookUrl          string
}

type CustomCommitDetails struct {
	AuthorEmailId string
	AuthorName    string
	CommitMessage string
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
