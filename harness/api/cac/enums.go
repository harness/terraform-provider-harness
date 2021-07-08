package cac

type HarnessApiVersion string

var HarnessApiVersions = &struct {
	V1 HarnessApiVersion
}{
	V1: "1.0",
}

type ObjectType string

var ObjectTypes = &struct {
	Application                     ObjectType
	AwsCloudProvider                ObjectType
	AzureCloudProvider              ObjectType
	GcpCloudProvider                ObjectType
	KubernetesCloudProvider         ObjectType
	PcfCloudProvider                ObjectType
	PhysicalDataCenterCloudProvider ObjectType
	Service                         ObjectType
	SpotInstCloudProvider           ObjectType
}{
	Application:                     "APPLICATION",
	AwsCloudProvider:                "AWS",
	AzureCloudProvider:              "AZURE",
	GcpCloudProvider:                "GCP",
	KubernetesCloudProvider:         "KUBERNETES_CLUSTER",
	PcfCloudProvider:                "PCF",
	PhysicalDataCenterCloudProvider: "PHYSICAL_DATA_CENTER",
	Service:                         "SERVICE",
	SpotInstCloudProvider:           "SPOT_INST",
}

type AzureEnvironmentType string

var AzureEnvironmentTypes = struct {
	AzureGlobal  AzureEnvironmentType
	USGovernment AzureEnvironmentType
}{
	AzureGlobal:  "AZURE",
	USGovernment: "AZURE_US_GOVERNMENT",
}

type HelmVersion string

var HelmVersions = &struct {
	V2 HelmVersion
	V3 HelmVersion
}{
	V2: "V2",
	V3: "V3",
}

type DeploymentType string

var DeploymentTypes = &struct {
	AMI           DeploymentType
	ECS           DeploymentType
	AWSCodeDeploy DeploymentType
	AWSLambda     DeploymentType
	Custom        DeploymentType
	Kubernetes    DeploymentType
	Helm          DeploymentType
	PCF           DeploymentType
	SSH           DeploymentType
	WinRM         DeploymentType
}{
	AMI:           "AMI",
	ECS:           "ECS",
	AWSCodeDeploy: "AWS_CODEDEPLOY",
	AWSLambda:     "AWS_LAMBDA",
	Custom:        "Custom",
	Kubernetes:    "KUBERNETES",
	Helm:          "HELM",
	PCF:           "PCF",
	SSH:           "SSH",
	WinRM:         "WINRM",
}

var DeploymenTypesSlice = &[]DeploymentType{
	DeploymentTypes.AMI,
	DeploymentTypes.AWSCodeDeploy,
	DeploymentTypes.AWSLambda,
	DeploymentTypes.Custom,
	DeploymentTypes.Kubernetes,
	DeploymentTypes.Helm,
	DeploymentTypes.PCF,
	DeploymentTypes.SSH,
	DeploymentTypes.WinRM,
}

type KubernetesAuthType string

var KubernetesAuthTypes = struct {
	ServiceAccount      KubernetesAuthType
	UsernameAndPassword KubernetesAuthType
	Custom              KubernetesAuthType
}{
	ServiceAccount:      "SERVICE_ACCOUNT",
	UsernameAndPassword: "USER_PASSWORD",
	Custom:              "NONE",
}

type ArtifactType string

var ArtifactTypes = &struct {
	AMI                 ArtifactType
	AWSCodeDeploy       ArtifactType
	AWSLambda           ArtifactType
	Docker              ArtifactType
	Jar                 ArtifactType
	Other               ArtifactType
	PCF                 ArtifactType
	RPM                 ArtifactType
	Tar                 ArtifactType
	War                 ArtifactType
	IISVirtualDirectory ArtifactType
	IISApp              ArtifactType
	IISWebsite          ArtifactType
	Zip                 ArtifactType
}{
	AMI:                 "AMI",
	AWSCodeDeploy:       "AWS_CODEDEPLOY",
	AWSLambda:           "AWS_LAMBDA",
	Docker:              "DOCKER",
	Jar:                 "JAR",
	Other:               "OTHER",
	PCF:                 "PCF",
	RPM:                 "RPM",
	Tar:                 "TAR",
	War:                 "WAR",
	IISVirtualDirectory: "IIS_VirtualDirectory",
	IISApp:              "IIS_APP",
	IISWebsite:          "IIS",
	Zip:                 "ZIP",
}

var SSHArtifactTypes []ArtifactType = []ArtifactType{
	ArtifactTypes.Docker,
	ArtifactTypes.Jar,
	ArtifactTypes.Other,
	ArtifactTypes.Tar,
	ArtifactTypes.War,
	ArtifactTypes.Zip,
}

var WinRMArtifactTypesSlice []ArtifactType = []ArtifactType{
	ArtifactTypes.Docker,
	ArtifactTypes.IISApp,
	ArtifactTypes.IISVirtualDirectory,
	ArtifactTypes.IISWebsite,
	ArtifactTypes.Other,
}

type ClassType string

var ClassTypes = &struct {
	Account                   ClassType
	Application               ClassType
	ArtifactStream            ClassType
	ConfigFile                ClassType
	Defaults                  ClassType
	Environment               ClassType
	InfrastructureProvisioner ClassType
	NotificationGroup         ClassType
	Pipeline                  ClassType
	Service                   ClassType
	SettingAttribute          ClassType
	Tags                      ClassType
	Template                  ClassType
	Workflow                  ClassType
}{
	Account:                   "Account",
	Application:               "Application",
	ArtifactStream:            "ArtifactStream",
	ConfigFile:                "ConfigFile",
	Defaults:                  "Defaults",
	Environment:               "Environment",
	InfrastructureProvisioner: "InfrastructureProvisioner",
	NotificationGroup:         "NotificationGroup",
	Pipeline:                  "Pipeline",
	Service:                   "Service",
	SettingAttribute:          "SettingAttribute",
	Tags:                      "HarnessTag",
	Template:                  "Template",
	Workflow:                  "Workflow",
}

type ApplicationFilterType string

var ApplicationFilterTypes = &struct {
	All      ApplicationFilterType
	Selected ApplicationFilterType
}{
	All:      "ALL",
	Selected: "SELECTED",
}

type EnvironmentFilterType string

var EnvironmentFilterTypes = &struct {
	Prod     EnvironmentFilterType
	NonProd  EnvironmentFilterType
	Selected EnvironmentFilterType
}{
	Prod:     "PROD",
	NonProd:  "NON_PROD",
	Selected: "SELECTED",
}

type SecretManagerType string

var SecretManagerTypes = &struct {
	GcpKMS            SecretManagerType
	AwsSecretsManager SecretManagerType
	AwsKMS            SecretManagerType
	AzureKeyVault     SecretManagerType
	CyberArk          SecretManagerType
	HashicorpVault    SecretManagerType
}{
	GcpKMS:            "gcpkms",
	AwsSecretsManager: "awssecretsmanager",
	AwsKMS:            "amazonkms",
	AzureKeyVault:     "azurekeyvault",
	CyberArk:          "cyberark",
	HashicorpVault:    "hashicorpvault",
}
