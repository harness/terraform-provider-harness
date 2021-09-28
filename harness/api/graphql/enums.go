package graphql

type SSHAuthenticationType string

func (s SSHAuthenticationType) String() string {
	return string(s)
}

var SSHAuthenticationTypes = &struct {
	SSHAuthentication      SSHAuthenticationType
	KerberosAuthentication SSHAuthenticationType
}{
	SSHAuthentication:      "SSH_AUTHENTICATION",
	KerberosAuthentication: "KERBEROS_AUTHENTICATION",
}

type WinRMAuthenticationScheme string

func (s WinRMAuthenticationScheme) String() string {
	return string(s)
}

var WinRMAuthenticationSchemes = &struct {
	NTLM WinRMAuthenticationScheme
}{
	NTLM: "NTLM",
}

type SecretType string

func (s SecretType) String() string {
	return string(s)
}

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

type SSHCredentialType string

func (s SSHCredentialType) String() string {
	return string(s)
}

var SSHCredentialTypes = &struct {
	Password       SSHCredentialType
	SSHKey         SSHCredentialType
	SSHKeyFilePath SSHCredentialType
}{
	Password:       "PASSWORD",
	SSHKey:         "SSH_KEY",
	SSHKeyFilePath: "SSH_KEY_FILE_PATH",
}

type TGTGenerationUsingOption string

func (s TGTGenerationUsingOption) String() string {
	return string(s)
}

var TGTGenerationUsingOptions = &struct {
	KeyTabFile TGTGenerationUsingOption
	Password   TGTGenerationUsingOption
}{
	KeyTabFile: "KEY_TAB_FILE",
	Password:   "PASSWORD",
}

type EnvironmentFilterType string

func (s EnvironmentFilterType) String() string {
	return string(s)
}

var EnvironmentFilterTypes = &struct {
	NonProduction EnvironmentFilterType
	Production    EnvironmentFilterType
}{
	NonProduction: "NON_PRODUCTION_ENVIRONMENTS",
	Production:    "PRODUCTION_ENVIRONMENTS",
}

type ApplicationFilterType string

func (s ApplicationFilterType) String() string {
	return string(s)
}

var ApplicationFilterTypes = &struct {
	All ApplicationFilterType
}{
	All: "ALL",
}

type SSHAuthenticationScheme string

func (s SSHAuthenticationScheme) String() string {
	return string(s)
}

var SSHAuthenticationSchemes = &struct {
	Kerberos SSHAuthenticationScheme
	SSH      SSHAuthenticationScheme
}{
	Kerberos: "KERBEROS",
	SSH:      "SSH",
}

type ConnectorType string

func (s ConnectorType) String() string {
	return string(s)
}

var ConnectorTypes = &struct {
	AmazonS3         ConnectorType
	AmazonS3HelmRepo ConnectorType
	APMVerification  ConnectorType
	AppDynamics      ConnectorType
	Artifactory      ConnectorType
	Bamboo           ConnectorType
	BugSnag          ConnectorType
	DataDog          ConnectorType
	Docker           ConnectorType
	DynaTrace        ConnectorType
	ECR              ConnectorType
	ELB              ConnectorType
	ELK              ConnectorType
	GCR              ConnectorType
	GCS              ConnectorType
	GCSHelmRepo      ConnectorType
	Git              ConnectorType
	HTTPHelpRepo     ConnectorType
	Jenkins          ConnectorType
	Jira             ConnectorType
	Logz             ConnectorType
	NewRelic         ConnectorType
	Nexus            ConnectorType
	Prometheus       ConnectorType
	ServiceNow       ConnectorType
	SFTP             ConnectorType
	Slack            ConnectorType
	SMB              ConnectorType
	SMTP             ConnectorType
	Splunk           ConnectorType
	Sumo             ConnectorType
}{
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

type NexusVersion string

func (s NexusVersion) String() string {
	return string(s)
}

var NexusVersions = &struct {
	V2 NexusVersion
	v3 NexusVersion
}{
	V2: "V2",
	v3: "V3",
}

type GitUrlType string

func (s GitUrlType) String() string {
	return string(s)
}

var GitUrlTypes = &struct {
	Account GitUrlType
	Repo    GitUrlType
}{
	Account: "ACCOUNT",
	Repo:    "REPO",
}

type AwsCredentialsType string

func (s AwsCredentialsType) String() string {
	return string(s)
}

var AwsCredentialsTypes = struct {
	Ec2Iam AwsCredentialsType
	Manual AwsCredentialsType
}{
	Ec2Iam: "EC2_IAM",
	Manual: "MANUAL",
}

type ClusterDetailsType string

func (s ClusterDetailsType) String() string {
	return string(s)
}

var ClusterDetailsTypes = struct {
	InheritClusterDetails ClusterDetailsType
	ManualClusterDetails  ClusterDetailsType
}{
	InheritClusterDetails: "INHERIT_CLUSTER_DETAILS",
	ManualClusterDetails:  "MANUAL_CLUSTER_DETAILS",
}

type ManualClusterDetailsAuthenticationType string

func (s ManualClusterDetailsAuthenticationType) String() string {
	return string(s)
}

var ManualClusterDetailsAuthenticationTypes = struct {
	ClientKeyAndCertificate ManualClusterDetailsAuthenticationType
	Custom                  ManualClusterDetailsAuthenticationType
	OIDCToken               ManualClusterDetailsAuthenticationType
	ServiceAccountToken     ManualClusterDetailsAuthenticationType
	UsernameAndPassword     ManualClusterDetailsAuthenticationType
}{
	ClientKeyAndCertificate: "CLIENT_KEY_AND_CERTIFICATE",
	Custom:                  "CUSTOM",
	OIDCToken:               "OIDC_TOKEN",
	ServiceAccountToken:     "SERVICE_ACCOUNT_TOKEN",
	UsernameAndPassword:     "USERNAME_AND_PASSWORD",
}

type CloudProviderType string

func (s CloudProviderType) String() string {
	return string(s)
}

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

type AccountPermissionType string

var AccountPermissionTypes = struct {
	// ADMINISTER_CE                            AccountPermissionType
	ADMINISTER_OTHER_ACCOUNT_FUNCTIONS       AccountPermissionType
	CREATE_AND_DELETE_APPLICATION            AccountPermissionType
	CREATE_CUSTOM_DASHBOARDS                 AccountPermissionType
	MANAGE_ALERT_NOTIFICATION_RULES          AccountPermissionType
	MANAGE_API_KEYS                          AccountPermissionType
	MANAGE_APPLICATION_STACKS                AccountPermissionType
	MANAGE_AUTHENTICATION_SETTINGS           AccountPermissionType
	MANAGE_CLOUD_PROVIDERS                   AccountPermissionType
	MANAGE_CONFIG_AS_CODE                    AccountPermissionType
	MANAGE_CONNECTORS                        AccountPermissionType
	MANAGE_CUSTOM_DASHBOARDS                 AccountPermissionType
	MANAGE_DELEGATE_PROFILES                 AccountPermissionType
	MANAGE_DELEGATES                         AccountPermissionType
	MANAGE_DEPLOYMENT_FREEZES                AccountPermissionType
	MANAGE_IP_WHITELIST                      AccountPermissionType
	MANAGE_PIPELINE_GOVERNANCE_STANDARDS     AccountPermissionType
	MANAGE_RESTRICTED_ACCESS                 AccountPermissionType
	MANAGE_SECRET_MANAGERS                   AccountPermissionType
	MANAGE_SECRETS                           AccountPermissionType
	MANAGE_SSH_AND_WINRM                     AccountPermissionType
	MANAGE_TAGS                              AccountPermissionType
	MANAGE_TEMPLATE_LIBRARY                  AccountPermissionType
	MANAGE_USER_AND_USER_GROUPS_AND_API_KEYS AccountPermissionType
	MANAGE_USERS_AND_GROUPS                  AccountPermissionType
	READ_USERS_AND_GROUPS                    AccountPermissionType
	VIEW_AUDITS                              AccountPermissionType
	VIEW_CE                                  AccountPermissionType
	VIEW_USER_AND_USER_GROUPS_AND_API_KEYS   AccountPermissionType
}{
	// ADMINISTER_CE:                            "ADMINISTER_CE",
	ADMINISTER_OTHER_ACCOUNT_FUNCTIONS:       "ADMINISTER_OTHER_ACCOUNT_FUNCTIONS",
	CREATE_AND_DELETE_APPLICATION:            "CREATE_AND_DELETE_APPLICATION",
	CREATE_CUSTOM_DASHBOARDS:                 "CREATE_CUSTOM_DASHBOARDS",
	MANAGE_ALERT_NOTIFICATION_RULES:          "MANAGE_ALERT_NOTIFICATION_RULES",
	MANAGE_API_KEYS:                          "MANAGE_API_KEYS",
	MANAGE_APPLICATION_STACKS:                "MANAGE_APPLICATION_STACKS",
	MANAGE_AUTHENTICATION_SETTINGS:           "MANAGE_AUTHENTICATION_SETTINGS",
	MANAGE_CLOUD_PROVIDERS:                   "MANAGE_CLOUD_PROVIDERS",
	MANAGE_CONFIG_AS_CODE:                    "MANAGE_CONFIG_AS_CODE",
	MANAGE_CONNECTORS:                        "MANAGE_CONNECTORS",
	MANAGE_CUSTOM_DASHBOARDS:                 "MANAGE_CUSTOM_DASHBOARDS",
	MANAGE_DELEGATE_PROFILES:                 "MANAGE_DELEGATE_PROFILES",
	MANAGE_DELEGATES:                         "MANAGE_DELEGATES",
	MANAGE_DEPLOYMENT_FREEZES:                "MANAGE_DEPLOYMENT_FREEZES",
	MANAGE_IP_WHITELIST:                      "MANAGE_IP_WHITELIST",
	MANAGE_PIPELINE_GOVERNANCE_STANDARDS:     "MANAGE_PIPELINE_GOVERNANCE_STANDARDS",
	MANAGE_RESTRICTED_ACCESS:                 "MANAGE_RESTRICTED_ACCESS",
	MANAGE_SECRET_MANAGERS:                   "MANAGE_SECRET_MANAGERS",
	MANAGE_SECRETS:                           "MANAGE_SECRETS",
	MANAGE_SSH_AND_WINRM:                     "MANAGE_SSH_AND_WINRM",
	MANAGE_TAGS:                              "MANAGE_TAGS",
	MANAGE_TEMPLATE_LIBRARY:                  "MANAGE_TEMPLATE_LIBRARY",
	MANAGE_USER_AND_USER_GROUPS_AND_API_KEYS: "MANAGE_USER_AND_USER_GROUPS_AND_API_KEYS",
	MANAGE_USERS_AND_GROUPS:                  "MANAGE_USERS_AND_GROUPS",
	READ_USERS_AND_GROUPS:                    "READ_USERS_AND_GROUPS",
	VIEW_AUDITS:                              "VIEW_AUDITS",
	VIEW_CE:                                  "VIEW_CE",
	VIEW_USER_AND_USER_GROUPS_AND_API_KEYS:   "VIEW_USER_AND_USER_GROUPS_AND_API_KEYS",
}

var AccountPermissionTypeValues = []string{
	"ADMINISTER_OTHER_ACCOUNT_FUNCTIONS",
	"CREATE_AND_DELETE_APPLICATION",
	"CREATE_CUSTOM_DASHBOARDS",
	"MANAGE_ALERT_NOTIFICATION_RULES",
	"MANAGE_API_KEYS",
	"MANAGE_APPLICATION_STACKS",
	"MANAGE_AUTHENTICATION_SETTINGS",
	"MANAGE_CLOUD_PROVIDERS",
	"MANAGE_CONFIG_AS_CODE",
	"MANAGE_CONNECTORS",
	"MANAGE_CUSTOM_DASHBOARDS",
	"MANAGE_DELEGATE_PROFILES",
	"MANAGE_DELEGATES",
	"MANAGE_DEPLOYMENT_FREEZES",
	"MANAGE_IP_WHITELIST",
	"MANAGE_PIPELINE_GOVERNANCE_STANDARDS",
	"MANAGE_RESTRICTED_ACCESS",
	"MANAGE_SECRET_MANAGERS",
	"MANAGE_SECRETS",
	"MANAGE_SSH_AND_WINRM",
	"MANAGE_TAGS",
	"MANAGE_TEMPLATE_LIBRARY",
	"MANAGE_USER_AND_USER_GROUPS_AND_API_KEYS",
	"MANAGE_USERS_AND_GROUPS",
	"READ_USERS_AND_GROUPS",
	"VIEW_AUDITS",
	"VIEW_USER_AND_USER_GROUPS_AND_API_KEYS",
	// "ADMINISTER_CE",
	// "VIEW_CE",
}

func (p AccountPermissionType) String() string {
	return string(p)
}

type Action string

var Actions = struct {
	CREATE            Action
	DELETE            Action
	EXECUTE           Action
	EXECUTE_PIPELINE  Action
	EXECUTE_WORKFLOW  Action
	READ              Action
	ROLLBACK_WORKFLOW Action
	UPDATE            Action
}{
	CREATE:            "CREATE",
	DELETE:            "DELETE",
	EXECUTE:           "EXECUTE",
	EXECUTE_PIPELINE:  "EXECUTE_PIPELINE",
	EXECUTE_WORKFLOW:  "EXECUTE_WORKFLOW",
	READ:              "READ",
	ROLLBACK_WORKFLOW: "ROLLBACK_WORKFLOW",
	UPDATE:            "UPDATE",
}

type FilterType string

var FilterTypes = struct {
	All FilterType
}{
	All: "ALL",
}

type DeploymentPermissionFilterType string

var DeploymentPermissionFilterTypes = struct {
	NonProductionEnvironments DeploymentPermissionFilterType
	ProductionEnvironments    DeploymentPermissionFilterType
}{
	NonProductionEnvironments: "NON_PRODUCTION_ENVIRONMENTS",
	ProductionEnvironments:    "PRODUCTION_ENVIRONMENTS",
}

var DeploymentPermissionFiltersSlice = []string{
	"NON_PRODUCTION_ENVIRONMENTS",
	"PRODUCTION_ENVIRONMENTS",
}

func (f DeploymentPermissionFilterType) String() string {
	return string(f)
}

type EnvFilterType string

var EnvFilterTypes = struct {
	NonProductionEnvironments EnvFilterType
	ProductionEnvironments    EnvFilterType
}{
	NonProductionEnvironments: "NON_PRODUCTION_ENVIRONMENTS",
	ProductionEnvironments:    "PRODUCTION_ENVIRONMENTS",
}

func (f EnvFilterType) String() string {
	return string(f)
}

var EnvFiltersSlice = []string{
	"NON_PRODUCTION_ENVIRONMENTS",
	"PRODUCTION_ENVIRONMENTS",
}

type AppPermissionType string

var AppPermissionTypes = struct {
	All         AppPermissionType
	Deployment  AppPermissionType
	Env         AppPermissionType
	Pipeline    AppPermissionType
	Provisioner AppPermissionType
	Service     AppPermissionType
	Template    AppPermissionType
	Workflow    AppPermissionType
}{
	All:         "ALL",
	Deployment:  "DEPLOYMENT",
	Env:         "ENV",
	Pipeline:    "PIPELINE",
	Provisioner: "PROVISIONER",
	Service:     "SERVICE",
	Template:    "TEMPLATE",
	Workflow:    "WORKFLOW",
}

type PipelinePermissionFilterType string

var PipelinePermissionFilterTypes = struct {
	NonProductionPipelines PipelinePermissionFilterType
	ProductionPipelines    PipelinePermissionFilterType
}{
	NonProductionPipelines: "NON_PRODUCTION_PIPELINES",
	ProductionPipelines:    "PRODUCTION_PIPELINES",
}

func (f PipelinePermissionFilterType) String() string {
	return string(f)
}

var PipelinePermissionFiltersSlice = []string{
	"NON_PRODUCTION_PIPELINES",
	"PRODUCTION_PIPELINES",
}

type WorkflowPermissionFilterType string

var WorkflowPermissionFilterTypes = struct {
	NonProductionWorkflows WorkflowPermissionFilterType
	ProductionWorkflows    WorkflowPermissionFilterType
	WorkflowTemplates      WorkflowPermissionFilterType
}{
	NonProductionWorkflows: "NON_PRODUCTION_WORKFLOWS",
	ProductionWorkflows:    "PRODUCTION_WORKFLOWS",
	WorkflowTemplates:      "WORKFLOW_TEMPLATES",
}

var WorkflowPermissionFiltersSlice = []string{
	"NON_PRODUCTION_WORKFLOWS",
	"PRODUCTION_WORKFLOWS",
	"WORKFLOW_TEMPLATES",
}

func (f WorkflowPermissionFilterType) String() string {
	return string(f)
}

type SSOType string

var SSOTypes = struct {
	LDAP SSOType
	SAML SSOType
}{
	LDAP: "LDAP",
	SAML: "SAML",
}
