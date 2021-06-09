package client

type ConnectorClient struct {
	APIClient *ApiClient
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
	CreatedAt   Time   `json:"createdAt,omitempty"`
	CreatedBy   *User  `json:"createdBy,omitempty"`
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
}

type GitConnector struct {
	Connector
	Url                 string               `json:"url,omitempty"`
	Branch              string               `json:"branch,omitempty"`
	CustomCommitDetails *CustomCommitDetails `json:"customCommitDetails,omitempty"`
	DelegateSelectors   []string             `json:"delegateSelectors,omitempty"`
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
	CustomCommitDetails *CustomCommitDetails `json:"customCommitDetails,omitempty"`
	DelegateSelectors   []string             `json:"delegateSelectors,omitempty"`
	GenerateWebhookUrl  bool                 `json:"generateWebhookUrl,omitempty"`
	Name                string               `json:"name,omitempty"`
	PasswordSecretId    string               `json:"passwordSecretId,omitempty"`
	SSHSettingId        string               `json:"sshSettingId,omitempty"`
	UrlType             string               `json:"urlType,omitempty"`
	UsageScope          *UsageScope          `json:"usageScope,omitempty"`
	UserName            string               `json:"userName,omitempty"`
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
