package nextgen

type ConnectorType string

var ConnectorTypes = struct {
	K8sCluster       ConnectorType
	Git              ConnectorType
	Splunk           ConnectorType
	AppDynamics      ConnectorType
	Prometheus       ConnectorType
	Dynatrace        ConnectorType
	Vault            ConnectorType
	AzureKeyVault    ConnectorType
	DockerRegistry   ConnectorType
	Local            ConnectorType
	AwsKms           ConnectorType
	GcpKms           ConnectorType
	AwsSecretManager ConnectorType
	Gcp              ConnectorType
	Aws              ConnectorType
	Artifactory      ConnectorType
	Jira             ConnectorType
	Nexus            ConnectorType
	Github           ConnectorType
	Gitlab           ConnectorType
	Bitbucket        ConnectorType
	Codecommit       ConnectorType
	CEAws            ConnectorType
	CEAzure          ConnectorType
	GcpCloudCost     ConnectorType
	CEK8sCluster     ConnectorType
	HttpHelmRepo     ConnectorType
	NewRelic         ConnectorType
	Datadog          ConnectorType
	SumoLogic        ConnectorType
	PagerDuty        ConnectorType
}{
	K8sCluster:       "K8sCluster",
	Git:              "Git",
	Splunk:           "Splunk",
	AppDynamics:      "AppDynamics",
	Prometheus:       "Prometheus",
	Dynatrace:        "Dynatrace",
	Vault:            "Vault",
	AzureKeyVault:    "AzureKeyVault",
	DockerRegistry:   "DockerRegistry",
	Local:            "Local",
	AwsKms:           "AwsKms",
	GcpKms:           "GcpKms",
	AwsSecretManager: "AwsSecretManager",
	Gcp:              "Gcp",
	Aws:              "Aws",
	Artifactory:      "Artifactory",
	Jira:             "Jira",
	Nexus:            "Nexus",
	Github:           "Github",
	Gitlab:           "Gitlab",
	Bitbucket:        "Bitbucket",
	Codecommit:       "Codecommit",
	CEAws:            "CEAws",
	CEAzure:          "CEAzure",
	GcpCloudCost:     "GcpCloudCost",
	CEK8sCluster:     "CEK8sCluster",
	HttpHelmRepo:     "HttpHelmRepo",
	NewRelic:         "NewRelic",
	Datadog:          "Datadog",
	SumoLogic:        "SumoLogic",
	PagerDuty:        "PagerDuty",
}

var ConnectorTypesSlice = []string{
	ConnectorTypes.K8sCluster.String(),
	ConnectorTypes.Git.String(),
	ConnectorTypes.Splunk.String(),
	ConnectorTypes.AppDynamics.String(),
	ConnectorTypes.Prometheus.String(),
	ConnectorTypes.Dynatrace.String(),
	ConnectorTypes.Vault.String(),
	ConnectorTypes.AzureKeyVault.String(),
	ConnectorTypes.DockerRegistry.String(),
	ConnectorTypes.Local.String(),
	ConnectorTypes.AwsKms.String(),
	ConnectorTypes.GcpKms.String(),
	ConnectorTypes.AwsSecretManager.String(),
	ConnectorTypes.Gcp.String(),
	ConnectorTypes.Aws.String(),
	ConnectorTypes.Artifactory.String(),
	ConnectorTypes.Jira.String(),
	ConnectorTypes.Nexus.String(),
	ConnectorTypes.Github.String(),
	ConnectorTypes.Gitlab.String(),
	ConnectorTypes.Bitbucket.String(),
	ConnectorTypes.Codecommit.String(),
	ConnectorTypes.CEAws.String(),
	ConnectorTypes.CEAzure.String(),
	ConnectorTypes.GcpCloudCost.String(),
	ConnectorTypes.CEK8sCluster.String(),
	ConnectorTypes.HttpHelmRepo.String(),
	ConnectorTypes.NewRelic.String(),
	ConnectorTypes.Datadog.String(),
	ConnectorTypes.SumoLogic.String(),
	ConnectorTypes.PagerDuty.String(),
}

func (c ConnectorType) String() string {
	return string(c)
}
