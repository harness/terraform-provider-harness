package graphql

type SSHAuthenticationType string

var SSHAuthenticationTypes = &struct {
	SSHAuthentication      SSHAuthenticationType
	KerberosAuthentication SSHAuthenticationType
}{
	SSHAuthentication:      "SSH_AUTHENTICATION",
	KerberosAuthentication: "KERBEROS_AUTHENTICATION",
}

type WinRMAuthenticationScheme string

var WinRMAuthenticationSchemes = &struct {
	NTLM WinRMAuthenticationScheme
}{
	NTLM: "NTLM",
}

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

type SSHCredentialType string

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

var TGTGenerationUsingOptions = &struct {
	KeyTabFile TGTGenerationUsingOption
	Password   TGTGenerationUsingOption
}{
	KeyTabFile: "KEY_TAB_FILE",
	Password:   "PASSWORD",
}

type EnvironmentFilterType string

var EnvironmentFilterTypes = &struct {
	NonProduction EnvironmentFilterType
	Production    EnvironmentFilterType
}{
	NonProduction: "NON_PRODUCTION_ENVIRONMENTS",
	Production:    "PRODUCTION_ENVIRONMENTS",
}

type ApplicationFilterType string

var ApplicationFilterTypes = &struct {
	All ApplicationFilterType
}{
	All: "ALL",
}

type SSHAuthenticationScheme string

var SSHAuthenticationSchemes = &struct {
	Kerberos SSHAuthenticationScheme
	SSH      SSHAuthenticationScheme
}{
	Kerberos: "KERBEROS",
	SSH:      "SSH",
}

type ConnectorType string

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

var NexusVersions = &struct {
	V2 NexusVersion
	v3 NexusVersion
}{
	V2: "V2",
	v3: "V3",
}

type GitUrlType string

var GitUrlTypes = &struct {
	Account GitUrlType
	Repo    GitUrlType
}{
	Account: "ACCOUNT",
	Repo:    "REPO",
}

type AwsCredentialsType string

var AwsCredentialsTypes = struct {
	Ec2Iam AwsCredentialsType
	Manual AwsCredentialsType
}{
	Ec2Iam: "EC2_IAM",
	Manual: "MANUAL",
}

type ClusterDetailsType string

var ClusterDetailsTypes = struct {
	InheritClusterDetails ClusterDetailsType
	ManualClusterDetails  ClusterDetailsType
}{
	InheritClusterDetails: "INHERIT_CLUSTER_DETAILS",
	ManualClusterDetails:  "MANUAL_CLUSTER_DETAILS",
}

type ManualClusterDetailsAuthenticationType string

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
