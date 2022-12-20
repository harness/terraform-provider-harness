package graphql

import (
	"github.com/harness/harness-go-sdk/harness/time"
)

type CommonMetadata struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	CreatedBy   *User      `json:"createdBy,omitempty"`
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Tags        []Tag      `json:"tags,omitempty"`
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

type InputVariable struct {
	Name  string
	Value string
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

type InfrastructureDefinition struct {
	CreatedAt        string
	DeploymentType   string
	Environment      Environment
	Id               string
	Name             string
	ScopedToServices []string
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

type CreateUserInput struct {
	ClientMutationID string   `json:"clientMutationId,omitempty"`
	Email            string   `json:"email,omitempty"`
	Name             string   `json:"name,omitempty"`
	UserGroupIds     []string `json:"userGroupIds,omitempty"`
}

type UserGroup struct {
	Id                   string                `json:"id,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Description          string                `json:"description,omitempty"`
	ImportedBySCIM       bool                  `json:"importedByScim,omitempty"`
	IsSSOLinked          bool                  `json:"isSSOLinked,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	Permissions          *UserGroupPermissions `json:"permissions,omitempty"`
	LDAPSettings         *LDAPSettings         `json:"ldapSettings,omitempty"`
	SAMLSettings         *SAMLSettings         `json:"samlSettings,omitempty"`
}

type CreateUserGroupInput struct {
	Name                 string                `json:"name,omitempty"`
	Description          string                `json:"description,omitempty"`
	ImportedBySCIM       bool                  `json:"importedByScim,omitempty"`
	IsSSOLinked          bool                  `json:"isSSOLinked,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	Permissions          *UserGroupPermissions `json:"permissions,omitempty"`
	SSOSetting           *SSOSettingInput      `json:"ssoSetting,omitempty"`
}

type UpdateUserGroupInput struct {
	Id                   string                `json:"userGroupId,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Description          string                `json:"description,omitempty"`
	ImportedBySCIM       bool                  `json:"importedByScim,omitempty"`
	IsSSOLinked          bool                  `json:"isSSOLinked,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	Permissions          *UserGroupPermissions `json:"permissions,omitempty"`
	SSOSetting           *SSOSettingInput      `json:"ssoSetting,omitempty"`
}

type SSOSettingInput struct {
	LDAPSettings *LDAPSettings `json:"ldapSettings,omitempty"`
	SAMLSettings *SAMLSettings `json:"samlSettings,omitempty"`
}

type LDAPSettings struct {
	GroupDN       string `json:"groupDN,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	SSOProviderId string `json:"ssoProviderId,omitempty"`
}

type SAMLSettings struct {
	GroupName     string `json:"groupName,omitempty"`
	SSOProviderId string `json:"ssoProviderId,omitempty"`
}

type LinkedSSOSettings struct {
	GroupDN       string `json:"groupDN,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	SsoProviderId string `json:"ssoProviderId,omitempty"`
}

type SSOProvider struct {
	Id      string  `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	SSOType SSOType `json:"ssoType,omitempty"`
}

type NotificationSettings struct {
	GroupEmailAddresses       []string                  `json:"groupEmailAddresses,omitempty"`
	MicrosoftTeamsWebhookUrl  string                    `json:"microsoftTeamsWebhookUrl,omitempty"`
	SendMailToNewMembers      bool                      `json:"sendMailToNewMembers,omitempty"`
	SendNotificationToMembers bool                      `json:"sendNotificationToMembers,omitempty"`
	SlackNotificationSetting  *SlackNotificationSetting `json:"slackNotificationSetting,omitempty"`
	// PagerDutyIntegrationKey   string                    `json:"pagerDutyIntegrationKey,omitempty"`
}

type UserGroupPermissions struct {
	AccountPermissions *AccountPermissions `json:"accountPermissions,omitempty"`
	AppPermissions     []*AppPermission    `json:"appPermissions,omitempty"`
}

type AppPermission struct {
	Actions        []Action                     `json:"actions,omitempty"`
	Applications   *AppFilter                   `json:"applications,omitempty"`
	Deployments    *DeploymentPermissionFilter  `json:"deployments,omitempty"`
	Environments   *EnvPermissionFilter         `json:"environments,omitempty"`
	PermissionType AppPermissionType            `json:"permissionType,omitempty"`
	Pipelines      *PipelinePermissionFilter    `json:"pipelines,omitempty"`
	Provisioners   *ProvisionerPermissionFilter `json:"provisioners,omitempty"`
	Services       *ServicePermissionFilter     `json:"services,omitempty"`
	Templates      *TemplatePermissionFilter    `json:"templates,omitempty"`
	Workflows      *WorkflowPermissionFilter    `json:"workflows,omitempty"`
}

type WorkflowPermissionFilter struct {
	EnvIds      []string                       `json:"envIds,omitempty"`
	FilterTypes []WorkflowPermissionFilterType `json:"filterTypes,omitempty"`
}

type AddUserToUserGroupInput struct {
	ClientMutationId string `json:"clientMutationId,omitempty"`
	UserGroupId      string `json:"userGroupId,omitempty"`
	UserId           string `json:"userId,omitempty"`
}

type ServicePermissionFilter struct {
	FilterType FilterType `json:"filterType,omitempty"`
	ServiceIds []string   `json:"serviceIds,omitempty"`
}

type ProvisionerPermissionFilter struct {
	FilterType     FilterType `json:"filterType,omitempty"`
	ProvisionerIds []string   `json:"provisionerIds,omitempty"`
}

type PipelinePermissionFilter struct {
	EnvIds      []string                       `json:"envIds,omitempty"`
	FilterTypes []PipelinePermissionFilterType `json:"filterTypes,omitempty"`
}

type AppFilter struct {
	AppIds     []string   `json:"appIds,omitempty"`
	FilterType FilterType `json:"filterType,omitempty"`
}

type DeploymentPermissionFilter struct {
	EnvIds      []string                         `json:"envIds,omitempty"`
	FilterTypes []DeploymentPermissionFilterType `json:"filterTypes,omitempty"`
}

type EnvPermissionFilter struct {
	EnvIds      []string        `json:"envIds,omitempty"`
	FilterTypes []EnvFilterType `json:"filterTypes,omitempty"`
}

type TemplatePermissionFilter struct {
	TemplateIds []string   `json:"templateIds,omitempty"`
	FilterType  FilterType `json:"filterType,omitempty"`
}

type AccountPermissions struct {
	AccountPermissionTypes []AccountPermissionType `json:"accountPermissionTypes,omitempty"`
}

type SlackNotificationSetting struct {
	SlackChannelName string `json:"slackChannelName,omitempty"`
	SlackWebhookUrl  string `json:"slackWebhookURL,omitempty"`
}

type UpdateUserInput struct {
	ClientMutationID string `json:"clientMutationId,omitempty"`
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
}

type GitSyncConfig struct {
	Branch         string        `json:"branch,omitempty"`
	GitConnector   *GitConnector `json:"gitConnector,omitempty"`
	RepositoryName string        `json:"repositoryName,omitempty"`
	SyncEnabled    bool          `json:"syncEnabled,omitempty"`
}

type UpdateApplicationGitSyncConfigInput struct {
	ClientMutationID string `json:"clientMutationId,omitempty"`
	ApplicationId    string `json:"applicationId,omitempty"`
	Branch           string `json:"branch,omitempty"`
	GitConnectorId   string `json:"gitConnectorId,omitempty"`
	RepositoryName   string `json:"repositoryName,omitempty"`
	SyncEnabled      bool   `json:"syncEnabled"`
}

type UsageScope struct {
	AppEnvScopes []*AppEnvScope `json:"appEnvScopes,omitempty"`
}

type AppEnvScope struct {
	Application *AppScopeFilter `json:"application,omitempty"`
	Environment *EnvScopeFilter `json:"environment,omitempty"`
}

type AppScopeFilter struct {
	AppId      string                `json:"appId,omitempty"`
	FilterType ApplicationFilterType `json:"filterType,omitempty"`
}

type EnvScopeFilter struct {
	EnvId      string                `json:"envId,omitempty"`
	FilterType EnvironmentFilterType `json:"filterType,omitempty"`
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
	Tags                      []*Tag                   `json:"tags,omitempty"`
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

type SecretManager struct {
	Id         string
	Name       string
	UsageScope *UsageScope
}

type EncryptedFile struct {
	Secret
	EncryptedText
}

type WinRMCredential struct {
	Secret
	Domain               string                    `json:"domain,omitempty"`
	Port                 int                       `json:"port,omitempty"`
	SkipCertCheck        bool                      `json:"skipCertCheck,omitempty"`
	UseSSL               bool                      `json:"useSSL,omitempty"`
	UserName             string                    `json:"username,omitempty"`
	AuthenticationScheme WinRMAuthenticationScheme `json:"authenticationScheme,omitempty"`
}

func (c *WinRMCredential) IsEmpty() bool {
	return c.Id == ""
}

type CreateSecretInput struct {
	ClientMutationId string                `json:"clientMutationId,omitempty"`
	EncryptedText    *EncryptedTextInput   `json:"encryptedText,omitempty"`
	SecretType       SecretType            `json:"secretType,omitempty"`
	SSHCredential    *SSHCredential        `json:"sshCredential,omitempty"`
	WinRMCredential  *WinRMCredentialInput `json:"winRMCredential,omitempty"`
}

type UpdateSecretInput struct {
	ClientMutationId string                 `json:"clientMutationId,omitempty"`
	EncryptedText    *UpdateEncryptedText   `json:"encryptedText,omitempty"`
	SecretId         string                 `json:"secretId,omitempty"`
	SecretType       SecretType             `json:"secretType,omitempty"`
	SSHCredential    *SSHCredential         `json:"sshCredential,omitempty"`
	WinRMCredential  *UpdateWinRMCredential `json:"winRMCredential,omitempty"`
}

type Secret struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	SecretType SecretType  `json:"secretType,omitempty"`
	UsageScope *UsageScope `json:"usageScope,omitempty"`
}

func (s *Secret) IsEmpty() bool {
	return s.Id == "" && s.Name == ""
}

type EncryptedText struct {
	Secret
	InheritScopesFromSM bool   `json:"inheritScopesFromSM,omitempty"`
	ScopedToAccount     bool   `json:"scopedToAccount,omitempty"`
	SecretManagerId     string `json:"secretManagerId,omitempty"`
}

type EncryptedTextInput struct {
	InheritScopesFromSM bool        `json:"inheritScopesFromSM,omitempty"`
	Name                string      `json:"name,omitempty"`
	ScopedToAccount     bool        `json:"scopedToAccount,omitempty"`
	SecretManagerId     string      `json:"secretManagerId,omitempty"`
	SecretReference     string      `json:"secretReference,omitempty"`
	UsageScope          *UsageScope `json:"usageScope,omitempty"`
	Value               string      `json:"value,omitempty"`
}

type UpdateEncryptedText struct {
	InheritScopesFromSM bool        `json:"inheritScopesFromSM,omitempty"`
	Name                string      `json:"name,omitempty"`
	ScopedToAccount     bool        `json:"scopedToAccount,omitempty"`
	SecretReference     string      `json:"secretReference,omitempty"`
	UsageScope          *UsageScope `json:"usageScope,omitempty"`
	Value               string      `json:"value,omitempty"`
}

type UpdateWinRMCredential struct {
	AuthenticationScheme string      `json:"authenticationScheme,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Name                 string      `json:"name,omitempty"`
	PasswordSecretId     string      `json:"passwordSecretID,omitempty"`
	Port                 int         `json:"port,omitempty"`
	SkipCertCheck        bool        `json:"skipCertCheck,omitempty"`
	UsageScope           *UsageScope `json:"usageScope,omitempty"`
	UseSSL               bool        `json:"useSSL,omitempty"`
	Username             string      `json:"username,omitempty"`
}

type SSHCredential struct {
	Secret
	AuthenticationScheme   SSHAuthenticationScheme `json:"authenticationScheme,omitempty"`
	KerberosAuthentication *KerberosAuthentication `json:"kerberosAuthentication,omitempty"`
	SSHAuthentication      *SSHAuthentication      `json:"sshAuthentication,omitempty"`
	UsageScope             *UsageScope             `json:"usageScope,omitempty"`
	AuthenticationType     SSHAuthenticationType   `json:"authenticationType,omitempty"`
}

type SSHAuthentication struct {
	Port                    int                      `json:"port,omitempty"`
	SSHAuthenticationMethod *SSHAuthenticationMethod `json:"sshAuthenticationMethod,omitempty"`
	Username                string                   `json:"userName,omitempty"`
}

type SSHAuthenticationMethod struct {
	InlineSSHKey      *InlineSSHKey     `json:"inlineSSHKey,omitempty"`
	ServerPassword    *SSHPassword      `json:"serverPassword,omitempty"`
	SSHCredentialType SSHCredentialType `json:"sshCredentialType,omitempty"`
	SSHKeyFile        *SSHKeyFile       `json:"sshKeyFile,omitempty"`
}

type InlineSSHKey struct {
	PassphraseSecretId string `json:"passphraseSecretId,omitempty"`
	SSHKeySecretFileId string `json:"sshKeySecretFileId,omitempty"`
}

type SSHPassword struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
}

type SSHKeyFile struct {
	PassphraseSecretId string `json:"passphraseSecretId,omitempty"`
	Path               string `json:"path,omitempty"`
}

type KerberosAuthentication struct {
	Port                int                  `json:"port,omitempty"`
	Principal           string               `json:"principal,omitempty"`
	Realm               string               `json:"realm,omitempty"`
	TGTGenerationMethod *TGTGenerationMethod `json:"tgtGenerationMethod,omitempty"`
}
type TGTGenerationMethod struct {
	KerberosPassword   *KerberosPassword        `json:"kerberosPassword,omitempty"`
	KeyTabFile         *KeyTabFile              `json:"keyTabFile,omitempty"`
	TGTGenerationUsing TGTGenerationUsingOption `json:"tgtGenerationUsing,omitempty"`
}

type KeyTabFile struct {
	FilePath string `json:"filePath,omitempty"`
}
type KerberosPassword struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
}
type WinRMCredentialInput struct {
	AuthenticationScheme string      `json:"authenticationScheme,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Name                 string      `json:"name,omitempty"`
	PasswordSecretId     string      `json:"passwordSecretId,omitempty"`
	Port                 int         `json:"port,omitempty"`
	SkipCertCheck        bool        `json:"skipCertCheck,omitempty"`
	UsageScope           *UsageScope `json:"usageScope,omitempty"`
	UseSSL               bool        `json:"useSSL,omitempty"`
	Username             string      `json:"username,omitempty"`
}

type DeleteSecretInput struct {
	SecretId   string     `json:"secretId,omitempty"`
	SecretType SecretType `json:"secretType,omitempty"`
}

type Connector struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	CreatedBy   *User      `json:"createdBy,omitempty"`
	Description string     `json:"description,omitempty"`
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
}

type GitConnector struct {
	Connector
	Url                 string               `json:"URL,omitempty"`
	Branch              string               `json:"branch,omitempty"`
	CustomCommitDetails *CustomCommitDetails `json:"customCommitDetails,omitempty"`
	Description         string               `json:"description,omitempty"`
	DelegateSelectors   []string             `json:"delegateSelectors"`
	GenerateWebhookUrl  bool                 `json:"generateWebhookUrl,omitempty"`
	PasswordSecretId    string               `json:"passwordSecretId,omitempty"`
	SSHSettingId        string               `json:"sshSettingId,omitempty"`
	UrlType             GitUrlType           `json:"urlType,omitempty"`
	UsageScope          *UsageScope          `json:"usageScope,omitempty"`
	UserName            string               `json:"userName,omitempty"`
	WebhookUrl          string               `json:"webhookUrl,omitempty"`
}

type GitConnectorInput struct {
	Url                 string               `json:"URL,omitempty"`
	Branch              string               `json:"branch,omitempty"`
	CustomCommitDetails *CustomCommitDetails `json:"customCommitDetails"`
	DelegateSelectors   []string             `json:"delegateSelectors"`
	GenerateWebhookUrl  bool                 `json:"generateWebhookUrl,omitempty"`
	Name                string               `json:"name,omitempty"`
	PasswordSecretId    string               `json:"passwordSecretId,omitempty"`
	SSHSettingId        string               `json:"sshSettingId,omitempty"`
	UrlType             GitUrlType           `json:"urlType,omitempty"`
	UsageScope          *UsageScope          `json:"usageScope,omitempty"`
	UserName            string               `json:"userName,omitempty"`
}

type CustomCommitDetails struct {
	AuthorEmailId string `json:"authorEmailId"`
	AuthorName    string `json:"authorName"`
	CommitMessage string `json:"commitMessage"`
}

type CreateConnectorInput struct {
	ClientMutationId string                `json:"clientMutationId,omitempty"`
	ConnectorType    ConnectorType         `json:"connectorType,omitempty"`
	DockerConnector  *DockerConnectorInput `json:"dockerConnector,omitempty"`
	GitConnector     *GitConnectorInput    `json:"gitConnector,omitempty"`
	HelmConnector    *HelmConnectorInput   `json:"helmConnector,omitempty"`
	NexusConnector   *NexusConnectorInput  `json:"nexusConnector,omitempty"`
}

type UpdateConnectorInput struct {
	CreateConnectorInput
	ConnectorId string `json:"connectorId,omitempty"`
}

type HelmConnectorInput struct {
	AmazonS3PlatformDetails   *AmazonS3PlatformInput   `json:"amazonS3PlatformDetails,omitempty"`
	GCSPlatformDetails        *GCSPlatformInput        `json:"gcsPlatformDetails,omitempty"`
	HTTPServerPlatformDetails *HTTPServerPlatformInput `json:"httpServerPlatformDetails,omitempty"`
	Name                      string                   `json:"name,omitempty"`
}

type AmazonS3PlatformInput struct {
	AWSCloudProvider string `json:"awsCloudProvider,omitempty"`
	BucketName       string `json:"bucketName,omitempty"`
	Region           string `json:"region,omitempty"`
}

type GCSPlatformInput struct {
	BucketName          string `json:"bucketName,omitempty"`
	GoogleCloudProvider string `json:"googleCloudProvider,omitempty"`
}

type HTTPServerPlatformInput struct {
	Url              string `json:"url,omitempty"`
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	Username         string `json:"userName,omitempty"`
}

type NexusConnectorInput struct {
	Url               string   `json:"url,omitempty"`
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	Name              string   `json:"name,omitempty"`
	PasswordSecretId  string   `json:"passwordSecretId,omitempty"`
	Username          string   `json:"userName,omitempty"`
	Version           string   `json:"version,omitempty"`
}

type AmazonS3HelmRepoConnector struct {
	Connector
}

type GCSHelmRepoConnector struct {
	Connector
}

type HTTPHelmRepoConnector struct {
	Connector
}

type DockerConnector struct {
	Connector
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
}

type DockerConnectorInput struct {
	Connector
	Url               string   `json:"url,omitempty"`
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	PassswordSecretId string   `json:"passwordSecretId"`
	Username          string   `json:"username"`
}

type Tag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CloudProvider struct {
	CreatedAt                     *time.Time `json:"createdAt,omitempty"`
	CreatedBy                     *User      `json:"createdBy,omitempty"`
	Description                   string     `json:"description,omitempty"`
	Id                            string     `json:"id,omitempty"`
	IsContinuousEfficiencyEnabled bool       `json:"isContinuousEfficiencyEnabled,omitempty"`
	Name                          string     `json:"name,omitempty"`
	Type                          string     `json:"type,omitempty"`
}

type AwsCloudProvider struct {
	CloudProvider
	CEHealthStatus         *CEHealthStatus            `json:"ceHealthStatus,omitempty"`
	CredentialsType        AwsCredentialsType         `json:"credentialsType,omitempty"`
	CrossAccountAttributes *AwsCrossAccountAttributes `json:"crossAccountAttributes,omitempty"`
	DefaultRegion          string                     `json:"defaultRegion,omitempty"`
	Ec2IamCredentials      *Ec2IamCredentials         `json:"ec2IamCredentials,omitempty"`
	ManualCredentials      *AwsManualCredentials      `json:"manualCredentials,omitempty"`
	IrsaCredentials        *AwsIrsaCredentials        `json:"irsaCredentials,omitempty"`
}

type UpdateAwsCloudProviderInput struct {
	CredentialsType        AwsCredentialsType         `json:"credentialsType,omitempty"`
	CrossAccountAttributes *AwsCrossAccountAttributes `json:"crossAccountAttributes,omitempty"`
	DefaultRegion          string                     `json:"defaultRegion,omitempty"`
	Ec2IamCredentials      *Ec2IamCredentials         `json:"ec2IamCredentials,omitempty"`
	ManualCredentials      *AwsManualCredentials      `json:"manualCredentials,omitempty"`
	IrsaCredentials        *AwsIrsaCredentials        `json:"irsaCredentials,omitempty"`
	Name                   string                     `json:"name,omitempty"`
}

type Ec2IamCredentials struct {
	DelegateSelector string      `json:"delegateSelector"`
	UsageScope       *UsageScope `json:"usageScope,omitempty"`
}

type AwsIrsaCredentials struct {
	DelegateSelector string      `json:"delegateSelector"`
	UsageScope       *UsageScope `json:"usageScope,omitempty"`
}

type AwsManualCredentials struct {
	AccessKey         string `json:"accessKey,omitempty"`
	AccessKeySecretId string `json:"accessKeySecretId,omitempty"`
	SecretKeySecretId string `json:"secretKeySecretId,omitempty"`
}

type AwsCrossAccountAttributes struct {
	AssumeCrossAccountRole bool   `json:"assumeCrossAccountRole,omitempty"`
	CrossAccountRoleArm    string `json:"crossAccountRoleArn,omitempty"`
	ExternalId             string `json:"externalId,omitempty"`
}

type AzureCloudProvider struct {
	CloudProvider
	ClientId    string `json:"clientId,omitempty"`
	KeySecretId string `json:"keySecretId,omitempty"`
	TenantId    string `json:"tenantId,omitempty"`
}

type UpdateAzureCloudProviderInput struct {
	ClientId    string `json:"clientId,omitempty"`
	KeySecretId string `json:"keySecretId,omitempty"`
	Name        string `json:"name,omitempty"`
	TenantID    string `json:"tenantId,omitempty"`
}

type GcpCloudProvider struct {
	CloudProvider
	DelegateSelector          string   `json:"delegateSelector,omitempty"`
	DelegateSelectors         []string `json:"delegateSelectors,omitempty"`
	Description               string   `json:"description,omitempty"`
	ServiceAccountKeySecretId string   `json:"serviceAccountKeySecretId,omitempty"`
	SkipValidation            bool     `json:"skipValidation,omitempty"`
	UseDelegate               bool     `json:"useDelegate,omitempty"`
	UseDelegateSelectors      bool     `json:"useDelegateSelectors,omitempty"`
}

type UpdateGcpCloudProviderInput struct {
	DelegateSelector          string   `json:"delegateSelector,omitempty"`
	DelegateSelectors         []string `json:"delegateSelectors,omitempty"`
	Name                      string   `json:"name,omitempty"`
	ServiceAccountKeySecretId string   `json:"serviceAccountKeySecretId,omitempty"`
	SkipValidation            bool     `json:"skipValidation,omitempty"`
	UseDelegate               bool     `json:"useDelgate,omitempty"`
	UseDelegateSelectors      bool     `json:"useDelegateSelectors,omitempty"`
}

type KubernetesCloudProvider struct {
	CloudProvider
	CEHealthStatus         *CEHealthStatus        `json:"ceHealthStatus,omitempty"`
	SkipK8sEventCollection bool                   `json:"skipK8sEventCollection,omitempty"`
	ClusterDetailsType     ClusterDetailsType     `json:"clusterDetailsType,omitempty"`
	InheritClusterDetails  *InheritClusterDetails `json:"inheritClusterDetails,omitempty"`
	ManualClusterDetails   *ManualClusterDetails  `json:"manualClusterDetails,omitempty"`
	SkipValidation         bool                   `json:"skipValidation,omitempty"`
}

type UpdateKubernetesCloudProviderInput struct {
	ClusterDetailsType    ClusterDetailsType     `json:"clusterDetailsType,omitempty"`
	InheritClusterDetails *InheritClusterDetails `json:"inheritClusterDetails,omitempty"`
	ManualClusterDetails  *ManualClusterDetails  `json:"manualClusterDetails,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	SkipValidation        bool                   `json:"skipValidation,omitempty"`
}

type InheritClusterDetails struct {
	DelegateName      string      `json:"delegateName,omitempty"`
	DelegateSelectors []string    `json:"delegateSelectors,omitempty"`
	UsageScope        *UsageScope `json:"usageScope,omitempty"`
}

type ManualClusterDetails struct {
	MasterUrl           string                                 `json:"masterUrl,omitempty"`
	None                *None                                  `json:"none,omitempty"`
	OIDCToken           *OIDCToken                             `json:"oidcToken,omitempty"`
	ServiceAccountToken *ServiceAccountToken                   `json:"serviceAccountToken,omitempty"`
	Type                ManualClusterDetailsAuthenticationType `json:"type,omitempty"`
	UsernameAndPassword *UsernameAndPasswordAuthentication     `json:"usernameAndPassword,omitempty"`
}

type UsernameAndPasswordAuthentication struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserNameSecretId string `json:"userNameSecretId,omitempty"`
}

type ServiceAccountToken struct {
	ServiceAccountTokenSecretId string `json:"serviceAccountTokenSecretId,omitempty"`
}

type OIDCToken struct {
	ClientIdSecretId     string `json:"clientIdSecretId,omitempty"`
	ClientSecretSecretId string `json:"clientSecretSecretId,omitempty"`
	IdentityProviderUrl  string `json:"identityProviderUrl,omitempty"`
	PasswordSecretId     string `json:"passwordSecretId,omitempty"`
	Scopes               string `json:"scopes,omitempty"`
	UserName             string `json:"userName,omitempty"`
}

type None struct {
	CaCertificateSecretId       string      `json:"caCertificateSecretId,omitempty"`
	ClientCertificateSecretId   string      `json:"clientCertificateSecretId,omitempty"`
	ClientKeyAlgorithm          string      `json:"clientKeyAlgorithm,omitempty"`
	ClientKeyPassphraseSecretId string      `json:"clientKeyPassphraseSecretId,omitempty"`
	ClientKeySecretId           string      `json:"clientKeySecretId,omitempty"`
	PasswordSecretId            string      `json:"passwordSecretId,omitempty"`
	ServiceAccountTokenSecretId string      `json:"serviceAccountTokenSecretId,omitempty"`
	UsageScope                  *UsageScope `json:"usageScope,omitempty"`
	Username                    string      `json:"username,omitempty"`
}

type UpdateCloudProvider struct {
	AwsCloudProvider                *UpdateAwsCloudProviderInput                `json:"awsCloudProvider,omitempty"`
	AzureCloudProvider              *UpdateAzureCloudProviderInput              `json:"azureCloudProvider,omitempty"`
	ClientMutationID                string                                      `json:"clientMutationId,omitempty"`
	CloudProviderId                 string                                      `json:"cloudProviderId,omitempty"`
	CloudProviderType               *CloudProviderType                          `json:"cloudProviderType,omitempty"`
	GcpCloudProvider                *UpdateGcpCloudProviderInput                `json:"gcpCloudProvider,omitempty"`
	K8sCloudProvider                *UpdateKubernetesCloudProviderInput         `json:"k8sCloudProvider,omitempty"`
	PcfCloudProvider                *UpdatePcfCloudProviderInput                `json:"pcfCloudProvider,omitempty"`
	PhysicalDataCenterCloudProvider *UpdatePhysicalDataCenterCloudProviderInput `json:"physicalDataCenterCloudProvider,omitempty"`
	SpotInstCloudProvider           *UpdateSpotInstCloudProviderInst            `json:"spotInstCloudProvider,omitempty"`
}

type PcfCloudProvider struct {
	CloudProvider
	EndpointUrl      string `json:"endpointUrl,omitempty"`
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	SkipValidation   bool   `json:"skipValidation,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserNameSecretId string `json:"userNameSecretId,omitempty"`
}

type UpdatePcfCloudProviderInput struct {
	EndpointUrl      string `json:"endpointUrl,omitempty"`
	Name             string `json:"name,omitempty"`
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	SkipValidation   bool   `json:"skipValidation,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserNameSecretId string `json:"userNameSecretId,omitempty"`
}
type PhysicalDataCenterCloudProvider struct {
	CloudProvider
	UsageScope *UsageScope `json:"usageScope,omitempty"`
}

type UpdatePhysicalDataCenterCloudProviderInput struct {
	Name       string      `json:"name,omitempty"`
	UsageScope *UsageScope `json:"usageScope"`
}

type SpotInstCloudProvider struct {
	CloudProvider
	AccountId     string `json:"accountId,omitempty"`
	TokenSecretId string `json:"tokenSecretId,omitempty"`
}

type UpdateSpotInstCloudProviderInst struct {
	AccountId     string `json:"accountId,omitempty"`
	Name          string `json:"name,omitempty"`
	TokenSecretId string `json:"tokenSecretId"`
}

type CreateCloudProviderInput struct {
	AwsCloudProvider                *AwsCloudProvider                `json:"awsCloudProvider,omitempty"`
	AzureCloudProvider              *AzureCloudProvider              `json:"azureCloudProvider,omitempty"`
	ClientMutationId                string                           `json:"clientMutationId,omitempty"`
	CloudProviderType               CloudProviderType                `json:"cloudProviderType,omitempty"`
	GCPCloudProvider                *GcpCloudProvider                `json:"gcpCloudProvider,omitempty"`
	K8sCloudProvider                *KubernetesCloudProvider         `json:"k8sCloudProvider,omitempty"`
	PcfCloudProvider                *PcfCloudProvider                `json:"pcfCloudProvider,omitempty"`
	PhysicalDataCenterCloudProvider *PhysicalDataCenterCloudProvider `json:"physicalDataCenterCloudProvider,omitempty"`
	SpotInstCloudProvider           *SpotInstCloudProvider           `json:"spotInstCloudProvider,omitempty"`
}

type CEHealthStatus struct {
	ClusterHealthStatusList []*CEClusterHealth `json:"ClusterHealthStatusList,omitempty"`
	IsCEConnector           bool               `json:"isCEConnector,omitempty"`
	IsHealthy               bool               `json:"isHealthy,omitempty"`
	Messages                []string           `json:"messages,omitempty"`
}

type CEClusterHealth struct {
	ClusterId          string   `json:"clusterId,omitempty"`
	ClusterName        string   `json:"clusterName,omitempty"`
	Errors             []string `json:"errors,omitempty"`
	LastEventTimestamp float64  `json:"lastEventTimestamp,omitempty"`
	Messages           []string `json:"messages"`
}

type ApprovalVariable struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ApprovalDetailsPayload struct {
	ApprovalDetails []*Approval `json:"approvalDetails,omitempty"`
}

type Approval struct {
	ApprovalId   string              `json:"approvalId,omitempty"`
	ApprovalType ApprovalStepType    `json:"approvalType,omitempty"`
	StepName     string              `json:"stepName,omitempty"`
	StageName    string              `json:"stageName,omitempty"`
	StartedAt    *time.Time          `json:"startedAt,omitempty"`
	TriggeredBy  User                `json:"triggeredBy,omitempty"`
	WillExpireAt *time.Time          `json:"willExpireAt,omitempty"`
	Approvers    []string            `json:"approvers,omitempty"`
	ExecutionId  string              `json:"executionId,omitempty"`
	Variables    []*ApprovalVariable `json:"variables,omitempty"`
}

type ApproveOrRejectApprovals struct {
	ClientMutationId string `json:"clientMutationId,omitempty"`
	Success          bool   `json:"success,omitempty"`
}

type Trigger struct {
	CommonMetadata
	Action    *TriggerAction    `json:"triggerAction,omitempty"`
	Condition *TriggerCondition `json:"condition,omitempty"`
}

type TriggerCondition struct {
	TriggerConditionType string          `json:"triggerConditionType,omitempty"`
	WebhookDetails       *WebhookDetails `json:"webhookDetails,omitempty"`
}

type WebhookDetails struct {
	Payload    string `json:"payload,omitempty"`
	WebhookUrl string `json:"webhookUrl,omitempty"`
	Method     string `json:"method,omitempty"`
	Header     string `json:"header,omitempty"`
}

type TriggerAction struct {
	ArtifactSelections *ArtifactSelection `json:"artifactSelection,omitempty"`
}

type ArtifactSelection struct {
	Service_id   string `json:"serviceId,omitempty"`
	Service_name string `json:"serviceName,omitempty"`
}
