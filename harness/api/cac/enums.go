package cac

type HarnessApiVersion string

func (v HarnessApiVersion) String() string {
	return string(v)
}

var HarnessApiVersions = &struct {
	V1 HarnessApiVersion
}{
	V1: "1.0",
}

type ObjectType string

func (v ObjectType) String() string {
	return string(v)
}

var ObjectTypes = &struct {
	Application                     ObjectType
	AwsCloudProvider                ObjectType
	AzureCloudProvider              ObjectType
	Environment                     ObjectType
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
	Environment:                     "ENVIRONMENT",
	GcpCloudProvider:                "GCP",
	KubernetesCloudProvider:         "KUBERNETES_CLUSTER",
	PcfCloudProvider:                "PCF",
	PhysicalDataCenterCloudProvider: "PHYSICAL_DATA_CENTER",
	Service:                         "SERVICE",
	SpotInstCloudProvider:           "SPOT_INST",
}

type AzureEnvironmentType string

func (v AzureEnvironmentType) String() string {
	return string(v)
}

var AzureEnvironmentTypes = struct {
	AzureGlobal  AzureEnvironmentType
	USGovernment AzureEnvironmentType
}{
	AzureGlobal:  "AZURE",
	USGovernment: "AZURE_US_GOVERNMENT",
}

var AzureEnvironmentTypesSlice = []string{
	AzureEnvironmentTypes.AzureGlobal.String(),
	AzureEnvironmentTypes.USGovernment.String(),
}

type HelmVersion string

func (v HelmVersion) String() string {
	return string(v)
}

var HelmVersions = &struct {
	V2 HelmVersion
	V3 HelmVersion
}{
	V2: "V2",
	V3: "V3",
}

type DeploymentType string

func (v DeploymentType) String() string {
	return string(v)
}

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

func (v KubernetesAuthType) String() string {
	return string(v)
}

var KubernetesAuthTypes = struct {
	ServiceAccount      KubernetesAuthType
	UsernameAndPassword KubernetesAuthType
	Custom              KubernetesAuthType
	OIDC                KubernetesAuthType
}{
	ServiceAccount:      "SERVICE_ACCOUNT",
	UsernameAndPassword: "USER_PASSWORD",
	Custom:              "NONE",
	OIDC:                "OIDC",
}

type ArtifactType string

func (v ArtifactType) String() string {
	return string(v)
}

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

var SSHArtifactTypes = []string{
	string(ArtifactTypes.Docker),
	string(ArtifactTypes.Jar),
	string(ArtifactTypes.Other),
	string(ArtifactTypes.Tar),
	string(ArtifactTypes.War),
	string(ArtifactTypes.Zip),
}

var WinRMArtifactTypesSlice = []string{
	string(ArtifactTypes.Docker),
	string(ArtifactTypes.IISApp),
	string(ArtifactTypes.IISVirtualDirectory),
	string(ArtifactTypes.IISWebsite),
	string(ArtifactTypes.Other),
}

type ClassType string

func (v ClassType) String() string {
	return string(v)
}

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

func (v ApplicationFilterType) String() string {
	return string(v)
}

var ApplicationFilterTypes = &struct {
	All      ApplicationFilterType
	Selected ApplicationFilterType
}{
	All:      "ALL",
	Selected: "SELECTED",
}

type EnvironmentFilterType string

func (v EnvironmentFilterType) String() string {
	return string(v)
}

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

func (v SecretManagerType) String() string {
	return string(v)
}

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

type EnvironmentType string

func (v EnvironmentType) String() string {
	return string(v)
}

var EnvironmentTypes = struct {
	NonProd EnvironmentType
	Prod    EnvironmentType
}{
	NonProd: "NON_PROD",
	Prod:    "PROD",
}

type VariableValueType string

func (v VariableValueType) String() string {
	return string(v)
}

var VariableOverrideValueTypes = struct {
	EncryptedText VariableValueType
	Text          VariableValueType
}{
	EncryptedText: "ENCRYPTED_TEXT",
	Text:          "TEXT",
}

type RestName string

func (v RestName) String() string {
	return string(v)
}

var RestNames = struct {
	Settings RestName
	Services RestName
	Tags     RestName
	Folders  RestName
}{
	Settings: "settings",
	Services: "services",
	Tags:     "tags",
	Folders:  "folders",
}
