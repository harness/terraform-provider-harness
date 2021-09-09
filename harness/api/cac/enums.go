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
	InfrastructureDefinition        ObjectType
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
	InfrastructureDefinition:        "INFRA_DEFINITION",
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
	AWSCodeDeploy DeploymentType
	AWSLambda     DeploymentType
	AzureVMSS     DeploymentType
	AzureWebApp   DeploymentType
	Custom        DeploymentType
	ECS           DeploymentType
	Helm          DeploymentType
	Kubernetes    DeploymentType
	PCF           DeploymentType
	SSH           DeploymentType
	WinRM         DeploymentType
}{
	AMI:           "AMI",
	AWSCodeDeploy: "AWS_CODEDEPLOY",
	AWSLambda:     "AWS_LAMBDA",
	AzureVMSS:     "AZURE_VMSS",
	AzureWebApp:   "AZURE_WEBAPP",
	Custom:        "Custom",
	ECS:           "ECS",
	Helm:          "HELM",
	Kubernetes:    "KUBERNETES",
	PCF:           "PCF",
	SSH:           "SSH",
	WinRM:         "WINRM",
}

var DeploymenTypesSlice = []string{
	"AMI",
	"AWS_CODEDEPLOY",
	"AWS_LAMBDA",
	"AZURE_VMSS",
	"AZURE_WEBAPP",
	"Custom",
	"ECS",
	"HELM",
	"KUBERNETES",
	"PCF",
	"SSH",
	"WINRM",
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

type CloudProviderType string

var CloudProviderTypes = struct {
	Aws               CloudProviderType
	Azure             CloudProviderType
	Custom            CloudProviderType
	DataCenter        CloudProviderType
	KubernetesCluster CloudProviderType
	Pcf               CloudProviderType
	Spot              CloudProviderType
}{
	Aws:               "AWS",
	Azure:             "AZURE",
	Custom:            "CUSTOM",
	DataCenter:        "PHYSICAL_DATA_CENTER",
	KubernetesCluster: "KUBERNETES_CLUSTER",
	Pcf:               "PCF",
	Spot:              "SPOT_INST",
}

var CloudProviderTypesSlice = []string{
	"AWS",
	"AZURE",
	"CUSTOM",
	"PHYSICAL_DATA_CENTER",
	"KUBERNETES_CLUSTER",
	"PCF",
	"SPOT_INST",
}

func (v CloudProviderType) String() string {
	return string(v)
}

type InfrastructureType string

var InfrastructureTypes = struct {
	AwsAmi           InfrastructureType
	AwsEcs           InfrastructureType
	AwsLambda        InfrastructureType
	AwsSSH           InfrastructureType
	AzureVmss        InfrastructureType
	AzureWebApp      InfrastructureType
	Custom           InfrastructureType
	DataCenterSSH    InfrastructureType
	DataCenterWinRM  InfrastructureType
	KubernetesDirect InfrastructureType
	KubernetesGcp    InfrastructureType
	Pcf              InfrastructureType
}{
	AwsAmi:           "AWS_AMI",
	AwsEcs:           "AWS_ECS",
	AwsLambda:        "AWS_AWS_LAMBDA",
	AwsSSH:           "AWS_SSH",
	AzureVmss:        "AZURE_VMSS",
	AzureWebApp:      "AZURE_WEBAPP",
	Custom:           "CUSTOM",
	DataCenterSSH:    "PHYSICAL_DATA_CENTER_SSH",
	DataCenterWinRM:  "PHYSICAL_DATA_CENTER_WINRM",
	KubernetesDirect: "DIRECT_KUBERNETES",
	KubernetesGcp:    "KUBERNETES_GCP",
	Pcf:              "PCF_PCF",
}

func (v InfrastructureType) String() string {
	return string(v)
}

type LaunchType string

var LaunchTypes = struct {
	EC2     LaunchType
	Fargate LaunchType
}{
	EC2:     "EC2",
	Fargate: "FARGATE",
}

type VmssAuthType string

var VmssAuthTypes = struct {
	SSHPublicKey VmssAuthType
}{
	SSHPublicKey: "SSH_PUBLIC_KEY",
}

var VmssAuthTypesSlice = []string{
	"SSH_PUBLIC_KEY",
}

func (v VmssAuthType) String() string {
	return string(v)
}

type VmssDeploymentType string

var VmssDeploymentTypes = struct {
	NativeVmss VmssDeploymentType
}{
	NativeVmss: "NATIVE_VMSS",
}

var VmssDeploymentTypesSlice = []string{
	"NATIVE_VMSS",
}

func (v VmssDeploymentType) String() string {
	return string(v)
}

type AmiDeploymentType string

var AmiDeploymentTypes = struct {
	ASG      AmiDeploymentType
	SpotInst AmiDeploymentType
}{
	ASG:      "AWS_ASG",
	SpotInst: "SPOTINST",
}

var AmiDeploymentTypesSlice = []string{
	"AWS_ASG",
	"SPOTINST",
}

func (v AmiDeploymentType) String() string {
	return string(v)
}

type AwsEcsLaunchType string

var AwsEcsLaunchTypes = struct {
	Fargate AwsEcsLaunchType
}{
	Fargate: "FARGATE",
}

var AwsEcsLaunchTypesSlice = []string{
	"FARGATE",
}

type HostConnectionType string

var HostConnectionTypes = struct {
	PrivateDns HostConnectionType
	PublicDns  HostConnectionType
	PrivateIp  HostConnectionType
	PublicIp   HostConnectionType
}{
	PrivateDns: "PRIVATE_DNS",
	PublicDns:  "PUBLIC_DNS",
	PrivateIp:  "PRIVATE_IP",
	PublicIp:   "PUBLIC_IP",
}

var HostConnectionTypesSlice = []string{
	"PRIVATE_DNS",
	"PUBLIC_DNS",
	"PRIVATE_IP",
	"PUBLIC_IP",
}

func (v HostConnectionType) String() string {
	return string(v)
}
