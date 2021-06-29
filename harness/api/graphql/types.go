package graphql

import "github.com/harness-io/harness-go-sdk/harness/time"

// import "time"

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

type Service struct{}

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

type sshCredentialAuthenticationType struct {
	SSHAuthentication      string
	KerberosAuthentication string
}

var SSHAuthenticationTypes = &sshCredentialAuthenticationType{
	SSHAuthentication:      "SSH_AUTHENTICATION",
	KerberosAuthentication: "KERBEROS_AUTHENTICATION",
}

type winRMAuthenticationType struct {
	NTLM string
	// Kerberos string
}

var WinRMAuthenticationTypes = &winRMAuthenticationType{
	NTLM: "NTLM",
}

// type SSHCredential struct {
// 	Secret
// 	AuthenticationType     string                  `json:"authenticationType,omitempty"`
// 	KerberosAuthentication *KerberosAuthentication `json:"kerberosAuthentication,omitempty"`
// 	SSHAuthentication      *SSHAuthentication      `json:"sshAuthentication,omitempty"`
// }

type SecretType string

var SecretTypes = struct {
	EncryptedFile   SecretType
	EncryptedText   SecretType
	SSHCredential   SecretType
	WinRMCredential SecretType
}{
	EncryptedFile:   "ENCRYPTED_FILE",
	EncryptedText:   "ENCRYPTED_TEXT",
	SSHCredential:   "SSH_CREDENTIAL",
	WinRMCredential: "WINRM_CREDENTIAL",
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
	Domain               string `json:"domain,omitempty"`
	Port                 int    `json:"port,omitempty"`
	SkipCertCheck        bool   `json:"skipCertCheck,omitempty"`
	UseSSL               bool   `json:"useSSL,omitempty"`
	UserName             string `json:"username,omitempty"`
	AuthenticationScheme string `json:"authenticationScheme,omitempty"`
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
	AuthenticationScheme   string                  `json:"authenticationScheme,omitempty"`
	KerberosAuthentication *KerberosAuthentication `json:"kerberosAuthentication,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	SSHAuthentication      *SSHAuthentication      `json:"sshAuthentication,omitempty"`
	UsageScope             *UsageScope             `json:"usageScope,omitempty"`
	AuthenticationType     string                  `json:"authenticationType,omitempty"`
}

type SSHAuthentication struct {
	Port                    int                      `json:"port,omitempty"`
	SSHAuthenticationMethod *SSHAuthenticationMethod `json:"sshAuthenticationMethod,omitempty"`
	Username                string                   `json:"userName,omitempty"`
}

type SSHAuthenticationMethod struct {
	InlineSSHKey      *InlineSSHKey `json:"inlineSSHKey,omitempty"`
	ServerPassword    *SSHPassword  `json:"serverPassword,omitempty"`
	SSHCredentialType string        `json:"sshCredentialType,omitempty"`
	SSHKeyFile        *SSHKeyFile   `json:"sshKeyFile,omitempty"`
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

type sshCredentialType struct {
	Password       string
	SSHKey         string
	SSHKeyFilePath string
}

var SSHCredentialTypes = &sshCredentialType{
	Password:       "PASSWORD",
	SSHKey:         "SSH_KEY",
	SSHKeyFilePath: "SSH_KEY_FILE_PATH",
}

type KerberosAuthentication struct {
	Port                int                  `json:"port,omitempty"`
	Principal           string               `json:"principal,omitempty"`
	Realm               string               `json:"realm,omitempty"`
	TGTGenerationMethod *TGTGenerationMethod `json:"tgtGenerationMethod,omitempty"`
}
type TGTGenerationMethod struct {
	KerberosPassword   *KerberosPassword `json:"kerberosPassword,omitempty"`
	KeyTabFile         *KeyTabFile       `json:"keyTabFile,omitempty"`
	TGTGenerationUsing string            `json:"tgtGenerationUsing,omitempty"`
}

type tgtGenerationUsingOption struct {
	KeyTabFile string
	Password   string
}

var TGTGenerationUsingOptions = &tgtGenerationUsingOption{
	KeyTabFile: "KEY_TAB_FILE",
	Password:   "PASSWORD",
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

type environmentFilterType struct {
	NonProduction string
	Production    string
}

type applicationFilterType struct {
	All string
}

var ApplicationFilterTypes = &applicationFilterType{
	All: "ALL",
}

var EnvironmentFilterTypes = &environmentFilterType{
	NonProduction: "NON_PRODUCTION_ENVIRONMENTS",
	Production:    "PRODUCTION_ENVIRONMENTS",
}

type sshAuthenticationScheme struct {
	Kerberos string
	SSH      string
}

var SSHAuthenticationSchemes = &sshAuthenticationScheme{
	Kerberos: "KERBEROS",
	SSH:      "SSH",
}

type DeleteSecretInput struct {
	SecretId   string     `json:"secretId,omitempty"`
	SecretType SecretType `json:"secretType,omitempty"`
}

type connectorType struct {
	AmazonS3         string
	AmazonS3HelmRepo string
	APMVerification  string
	AppDynamics      string
	Artifactory      string
	Bamboo           string
	BugSnag          string
	DataDog          string
	Docker           string
	DynaTrace        string
	ECR              string
	ELB              string
	ELK              string
	GCR              string
	GCS              string
	GCSHelmRepo      string
	Git              string
	HTTPHelpRepo     string
	Jenkins          string
	Jira             string
	Logz             string
	NewRelic         string
	Nexus            string
	Prometheus       string
	ServiceNow       string
	SFTP             string
	Slack            string
	SMB              string
	SMTP             string
	Splunk           string
	Sumo             string
}

var ConnectorTypes = &connectorType{
	AmazonS3:         "AMAZON_S3",
	AmazonS3HelmRepo: "AMAZON_S3_HELM_REPO",
	APMVerification:  "APM_VERIFICATION",
	AppDynamics:      "APP_DYNAMICS",
	Artifactory:      "ARTIFACTORY",
	Bamboo:           "BAMBOO",
	BugSnag:          "BUG_SNAG",
	DataDog:          "DATA_DOG",
	Docker:           "DOCKER",
	DynaTrace:        "DYNA_TRACE",
	ECR:              "ECR",
	ELB:              "ELB",
	ELK:              "ELK",
	GCR:              "GCR",
	GCS:              "GCS",
	GCSHelmRepo:      "GCS_HELM_REPO",
	Git:              "GIT",
	HTTPHelpRepo:     "HTTP_HELM_REPO",
	Jenkins:          "JENKINS",
	Jira:             "JIRA",
	Logz:             "LOGZ",
	NewRelic:         "NEW_RELIC",
	Nexus:            "NEXUS",
	Prometheus:       "PROMETHEUS",
	ServiceNow:       "SERVICENOW",
	SFTP:             "SFTP",
	Slack:            "SLACK",
	SMB:              "SMB",
	SMTP:             "SMTP",
	Splunk:           "SPLUNK",
	Sumo:             "SUMO",
}

type nexusVersion struct {
	V2 string
	v3 string
}

var NexusVersions = &nexusVersion{
	V2: "V2",
	v3: "V3",
}

type gitUrlType struct {
	Account string
	Repo    string
}

var GitUrlTypes = &gitUrlType{
	Account: "ACCOUNT",
	Repo:    "REPO",
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
	Url                 string               `json:"url,omitempty"`
	Branch              string               `json:"branch,omitempty"`
	CustomCommitDetails *CustomCommitDetails `json:"customCommitDetails,omitempty"`
	DelegateSelectors   []string             `json:"delegateSelectors"`
	GenerateWebhookUrl  bool                 `json:"generateWebhookUrl,omitempty"`
	PasswordSecretId    string               `json:"passwordSecretId,omitempty"`
	SSHSettingId        string               `json:"sshSettingId,omitempty"`
	UrlType             string               `json:"urlType,omitempty"`
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
	UrlType             string               `json:"urlType,omitempty"`
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
	ConnectorType    string                `json:"connectorType,omitempty"`
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
	CredentialsType        awsCredentialsType         `json:"credentialsType,omitempty"`
	CrossAccountAttributes *AwsCrossAccountAttributes `json:"awsCrossAccountAttributes,omitempty"`
	DefaultRegion          string                     `json:"defaultRegion,omitempty"`
	Ec2IamCredentials      *Ec2IamCredentials         `json:"ec2IamCredentials,omitempty"`
	ManualCredentials      *AwsManualCredentials      `json:"manualCredentials,omitempty"`
}

type Ec2IamCredentials struct {
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

type awsCredentialsType string

var AwsCredentialsTypes = struct {
	Ec2Iam awsCredentialsType
	Manual awsCredentialsType
}{
	Ec2Iam: "EC2_IAM",
	Manual: "MANUAL",
}

type AzureCloudProvider struct {
	CloudProvider
	ClientId    string `json:"clientId,omitempty"`
	KeySecretId string `json:"keySecretId,omitempty"`
	TenantId    string `json:"tenantId,omitempty"`
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

type KubernetesCloudProvider struct {
	CloudProvider
	CEHealthStatus         *CEHealthStatus        `json:"ceHealthStatus,omitempty"`
	SkipK8sEventCollection bool                   `json:"skipK8sEventCollection,omitempty"`
	ClusterDetailsType     clusterDetailsType     `json:"clusterDetailsType,omitempty"`
	InheritClusterDetails  *InheritClusterDetails `json:"inheritClusterDetails,omitempty"`
	ManualClusterDetails   *ManualClusterDetails  `json:"manualClusterDetails,omitempty"`
	SkipValidation         bool                   `json:"skipValidation,omitempty"`
}

type clusterDetailsType string

var ClusterDetailsTypes = struct {
	InheritClusterDetails clusterDetailsType
	ManualClusterDetails  clusterDetailsType
}{
	InheritClusterDetails: "INHERIT_CLUSTER_DETAILS",
	ManualClusterDetails:  "MANUAL_CLUSTER_DETAILS",
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
	Type                manualClusterDetailsAuthenticationType `json:"type,omitempty"`
	UsernameAndPassword *UsernameAndPasswordAuthentication     `json:"usernameAndPassword,omitempty"`
}

type UsernameAndPasswordAuthentication struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserNameSecretId string `json:"userNameSecretId,omitempty"`
}

type manualClusterDetailsAuthenticationType string

var ManualClusterDetailsAuthenticationTypes = struct {
	ClientKeyAndCertificate manualClusterDetailsAuthenticationType
	Custom                  manualClusterDetailsAuthenticationType
	OIDCToken               manualClusterDetailsAuthenticationType
	ServiceAccountToken     manualClusterDetailsAuthenticationType
	UsernameAndPassword     manualClusterDetailsAuthenticationType
}{
	ClientKeyAndCertificate: "CLIENT_KEY_AND_CERTIFICATE",
	Custom:                  "CUSTOM",
	OIDCToken:               "OIDC_TOKEN",
	ServiceAccountToken:     "SERVICE_ACCOUNT_TOKEN",
	UsernameAndPassword:     "USERNAME_AND_PASSWORD",
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

type PcfCloudProvider struct {
	CloudProvider
	EndpointUrl      string `json:"endpointUrl,omitempty"`
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
	SkipValidation   bool   `json:"skipValidation,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserNameSecretId string `json:"userNameSecretId,omitempty"`
}

type PhysicalDataCenterCloudProvider struct {
	CloudProvider
	UsageScope *UsageScope `json:"usageScope,omitempty"`
}

type SpotInstCloudProvider struct {
	CloudProvider
	AccountId     string `json:"accountId,omitempty"`
	TokenSecretId string `json:"tokenSecretId,omitempty"`
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

type CloudProviderType string

var CloudProviderTypes = struct {
	Aws                CloudProviderType
	Azure              CloudProviderType
	Gcp                CloudProviderType
	KubernetesCluster  CloudProviderType
	Pcf                CloudProviderType
	PhysicalDataCenter CloudProviderType
	SpotInst           CloudProviderType
}{
	Aws:                "AWS",
	Azure:              "AZURE",
	Gcp:                "GCP",
	KubernetesCluster:  "KUBERNETES_CLUSTER",
	Pcf:                "PCF",
	PhysicalDataCenter: "PHYSICAL_DATA_CENTER",
	SpotInst:           "SPOT_INST",
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
