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
