package nextgen

type ConnectorType string

var ConnectorTypes = struct {
	K8sCluster          ConnectorType
	Git                 ConnectorType
	Splunk              ConnectorType
	AppDynamics         ConnectorType
	Azure               ConnectorType
	Prometheus          ConnectorType
	Dynatrace           ConnectorType
	Vault               ConnectorType
	AzureKeyVault       ConnectorType
	DockerRegistry      ConnectorType
	JDBC                ConnectorType
	Local               ConnectorType
	AwsKms              ConnectorType
	GcpKms              ConnectorType
	AwsSecretManager    ConnectorType
	Gcp                 ConnectorType
	Aws                 ConnectorType
	Artifactory         ConnectorType
	Jira                ConnectorType
	Jenkins             ConnectorType
	Nexus               ConnectorType
	Github              ConnectorType
	Gitlab              ConnectorType
	Bitbucket           ConnectorType
	Codecommit          ConnectorType
	CEAws               ConnectorType
	CEAzure             ConnectorType
	GcpCloudCost        ConnectorType
	CEK8sCluster        ConnectorType
	HttpHelmRepo        ConnectorType
	OciHelmRepo         ConnectorType
	NewRelic            ConnectorType
	Datadog             ConnectorType
	SumoLogic           ConnectorType
	PagerDuty           ConnectorType
	GcpSecretManager    ConnectorType
	Spot                ConnectorType
	ServiceNow          ConnectorType
	Tas                 ConnectorType
	TerraformCloud      ConnectorType
	ElasticSearch       ConnectorType
	Rancher             ConnectorType
	CustomHealth        ConnectorType
	Pdc                 ConnectorType
	CustomSecretManager ConnectorType
}{
	K8sCluster:          "K8sCluster",
	Git:                 "Git",
	Splunk:              "Splunk",
	AppDynamics:         "AppDynamics",
	Prometheus:          "Prometheus",
	Dynatrace:           "Dynatrace",
	Vault:               "Vault",
	AzureKeyVault:       "AzureKeyVault",
	DockerRegistry:      "DockerRegistry",
	JDBC:                "JDBC",
	Local:               "Local",
	AwsKms:              "AwsKms",
	GcpKms:              "GcpKms",
	AwsSecretManager:    "AwsSecretManager",
	Gcp:                 "Gcp",
	Aws:                 "Aws",
	Artifactory:         "Artifactory",
	Jira:                "Jira",
	Jenkins:             "Jenkins",
	Nexus:               "Nexus",
	Github:              "Github",
	Gitlab:              "Gitlab",
	Bitbucket:           "Bitbucket",
	Codecommit:          "Codecommit",
	CEAws:               "CEAws",
	CEAzure:             "CEAzure",
	GcpCloudCost:        "GcpCloudCost",
	CEK8sCluster:        "CEK8sCluster",
	HttpHelmRepo:        "HttpHelmRepo",
	OciHelmRepo:         "OciHelmRepo",
	NewRelic:            "NewRelic",
	Datadog:             "Datadog",
	SumoLogic:           "SumoLogic",
	PagerDuty:           "PagerDuty",
	GcpSecretManager:    "GcpSecretManager",
	Azure:               "Azure",
	Spot:                "Spot",
	ServiceNow:          "ServiceNow",
	Tas:                 "Tas",
	TerraformCloud:      "TerraformCloud",
	ElasticSearch:       "ElasticSearch",
	Rancher:             "Rancher",
	CustomHealth:        "CustomHealth",
	Pdc:                 "Pdc",
	CustomSecretManager: "CustomSecretManager",
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
	ConnectorTypes.JDBC.String(),
	ConnectorTypes.Local.String(),
	ConnectorTypes.AwsKms.String(),
	ConnectorTypes.GcpKms.String(),
	ConnectorTypes.AwsSecretManager.String(),
	ConnectorTypes.Gcp.String(),
	ConnectorTypes.Aws.String(),
	ConnectorTypes.Artifactory.String(),
	ConnectorTypes.Jira.String(),
	ConnectorTypes.Jenkins.String(),
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
	ConnectorTypes.OciHelmRepo.String(),
	ConnectorTypes.NewRelic.String(),
	ConnectorTypes.Datadog.String(),
	ConnectorTypes.SumoLogic.String(),
	ConnectorTypes.PagerDuty.String(),
	ConnectorTypes.GcpSecretManager.String(),
	ConnectorTypes.Azure.String(),
	ConnectorTypes.Spot.String(),
	ConnectorTypes.ServiceNow.String(),
	ConnectorTypes.Tas.String(),
	ConnectorTypes.TerraformCloud.String(),
	ConnectorTypes.ElasticSearch.String(),
	ConnectorTypes.Rancher.String(),
	ConnectorTypes.CustomHealth.String(),
	ConnectorTypes.Pdc.String(),
	ConnectorTypes.CustomSecretManager.String(),
}

func (c ConnectorType) String() string {
	return string(c)
}
