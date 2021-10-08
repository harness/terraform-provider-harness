package cac

import (
	"errors"
	"reflect"

	"github.com/harness-io/harness-go-sdk/harness/utils"
)

type Entity interface {
	IsEmpty() bool
	Validate() (bool, error)
}

type YamlEntity struct {
	Name          string
	Id            string
	Content       string
	ApplicationId string
	Path          YamlPath
}

type Application struct {
	HarnessApiVersion HarnessApiVersion `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              ObjectType        `yaml:"type" json:"type"`
	Id                string            `yaml:"-"`
	Name              string            `yaml:"-"`
	Description       string            `yaml:"description"`
}

func (a *Application) IsEmpty() bool {
	return reflect.DeepEqual(a, &Application{})
}

type Tag struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type Response struct {
	Metadata         ResponseMetadata  `json:"metaData"`
	Resource         ConfigAsCodeItem  `json:"resource"`
	ResponseMessages []ResponseMessage `json:"responseMessages"`
}

type ConfigAsCodeItem struct {
	AccountId       string              `json:"accountId,omitempty"`
	Type            string              `json:"type,omitempty"`
	Name            string              `json:"name,omitempty"`
	ClassName       string              `json:"className,omitempty"`
	ShortClassName  string              `json:"shortClassName,omitempty"`
	RestName        string              `json:"restName,omitempty"`
	DirectoryPath   *DirectoryPath      `json:"directoryPath,omitempty"`
	DefaultToClosed bool                `json:"defaultToClosed,omitempty"`
	Children        []*ConfigAsCodeItem `json:"children,omitempty"`
	AppId           string              `json:"appId,omitempty"`
	YamlGitConfig   interface{}         `json:"yamlGitConfig,omitempty"`
	UUID            string              `json:"uuid,omitempty"`
	YamlVersionType string              `json:"yamlVersionType,omitempty"`
	YamlFilePath    string              `json:"yamlFilePath,omitempty"`
	Status          string              `json:"status,omitempty"`
	ErrorMessage    string              `json:"errorMssg,omitempty"`
	Yaml            string              `json:"yaml"`
	EntityId        string              `json:"entityId,omitempty"`
}

type DirectoryPath struct {
	Path string `json:"path,omitempty"`
}

type ResponseMetadata struct{}

type ResponseMessage struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type Service struct {
	HarnessApiVersion         HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type                      ObjectType         `yaml:"type" json:"type"`
	Id                        string             `yaml:"-"`
	Name                      string             `yaml:"-"`
	ArtifactType              ArtifactType       `yaml:"artifactType,omitempty"`
	DeploymentType            DeploymentType     `yaml:"deploymentType,omitempty"`
	Description               string             `yaml:"description,omitempty"`
	Tags                      map[string]string  `yaml:"tags,omitempty"`
	HelmVersion               HelmVersion        `yaml:"helmVersion,omitempty"`
	ApplicationId             string             `yaml:"-"`
	DeploymentTypeTemplateUri string             `yaml:"deploymentTypeTemplateUri,omitempty"`
	ConfigVariables           []*ServiceVariable `yaml:"configVariables,omitempty"`
}

func (a *Service) IsEmpty() bool {
	return reflect.DeepEqual(a, &Service{})
}

func (s *Service) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(s, []string{"ApplicationId"})
}

type ServiceVariable struct {
	Name      string            `yaml:"name,omitempty"`
	Value     string            `yaml:"value,omitempty"`
	ValueType VariableValueType `yaml:"valueType,omitempty"`
}

type AwsCloudProvider struct {
	HarnessApiVersion      HarnessApiVersion          `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type                   ObjectType                 `yaml:"type" json:"type"`
	Id                     string                     `yaml:"-"`
	Name                   string                     `yaml:"-"`
	AccessKey              string                     `yaml:"accessKey,omitempty"`
	AccessKeySecretId      *SecretRef                 `yaml:"accessKeySecretId,omitempty"`
	AssumeCrossAccountRole bool                       `yaml:"assumeCrossAccountRole,omitempty"`
	CrossAccountAttributes *AwsCrossAccountAttributes `yaml:"crossAccountAttributes,omitempty"`
	SecretKey              *SecretRef                 `yaml:"secretKey,omitempty"`
	UseIRSA                bool                       `yaml:"useIRSA,omitempty"`
	UseEc2IamCredentials   bool                       `yaml:"useEc2IamCredentials,omitempty"`
	UsageRestrictions      *UsageRestrictions         `yaml:"usageRestrictions,omitempty"`
	DelegateSelector       string                     `yaml:"tag,omitempty"`
}

func (a *AwsCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &AwsCloudProvider{})
}

type AwsCrossAccountAttributes struct {
	CrossAccountRoleArn string `yaml:"crossAccountRoleArn,omitempty"`
	ExternalId          string `yaml:"externalId,omitempty"`
}

type PcfCloudProvider struct {
	HarnessApiVersion HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              ObjectType         `yaml:"type" json:"type"`
	Id                string             `yaml:"-"`
	Name              string             `yaml:"-"`
	EndpointUrl       string             `yaml:"endpointUrl,omitempty"`
	Password          *SecretRef         `yaml:"password,omitempty"`
	SkipValidation    bool               `yaml:"skipValidation,omitempty"`
	Username          string             `yaml:"username,omitempty"`
	UsernameSecretId  *SecretRef         `yaml:"usernameSecretId,omitempty"`
	UsageRestrictions *UsageRestrictions `yaml:"usageRestrictions,omitempty"`
}

func (a *PcfCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &PcfCloudProvider{})
}

type PhysicalDatacenterCloudProvider struct {
	HarnessApiVersion HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              ObjectType         `yaml:"type" json:"type"`
	Name              string             `yaml:"-"`
	Id                string             `yaml:"-"`
	UsageRestrictions *UsageRestrictions `yaml:"usageRestrictions,omitempty"`
}

func (a *PhysicalDatacenterCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &PhysicalDatacenterCloudProvider{})
}

func (cp *PhysicalDatacenterCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

type AzureCloudProvider struct {
	HarnessApiVersion    HarnessApiVersion    `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Id                   string               `yaml:"-"`
	Name                 string               `yaml:"-"`
	Type                 ObjectType           `yaml:"type,omitempty"`
	AzureEnvironmentType AzureEnvironmentType `yaml:"azureEnvironmentType,omitempty"`
	ClientId             string               `yaml:"clientId,omitempty"`
	TenantId             string               `yaml:"tenantId,omitempty"`
	Key                  *SecretRef           `yaml:"key,omitempty"`
	UsageRestrictions    *UsageRestrictions   `yaml:"usageRestrictions,omitempty"`
}

func (a *AzureCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &AzureCloudProvider{})
}

type GcpCloudProvider struct {
	HarnessApiVersion            HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Id                           string             `yaml:"-"`
	Name                         string             `yaml:"-"`
	Type                         ObjectType         `yaml:"type,omitempty"`
	CertValidationRequired       bool               `yaml:"certValidationRequired,omitempty"`
	DelegateSelectors            []string           `yaml:"delegateSelectors,omitempty"`
	SkipValidation               bool               `yaml:"skipValidation,omitempty"`
	ServiceAccountKeyFileContent *SecretRef         `yaml:"serviceAccountKeyFileContent,omitempty"`
	UseDelegate                  bool               `yaml:"useDelegate,omitempty"`
	UseDelegateSelectors         bool               `yaml:"useDelegateSelectors,omitempty"`
	UsageRestrictions            *UsageRestrictions `yaml:"usageRestrictions,omitempty"`
}

func (a *GcpCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &GcpCloudProvider{})
}

func (cp *GcpCloudProvider) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(cp, []string{"Name"})
}

type KubernetesCloudProvider struct {
	HarnessApiVersion          HarnessApiVersion           `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Id                         string                      `yaml:"-"`
	Name                       string                      `yaml:"-"`
	Type                       ObjectType                  `yaml:"type,omitempty"`
	AuthType                   KubernetesAuthType          `yaml:"authType,omitempty"`
	CACert                     *SecretRef                  `yaml:"caCert,omitempty"`
	ClientCert                 *SecretRef                  `yaml:"clientCert,omitempty"`
	ClientKey                  *SecretRef                  `yaml:"clientKey,omitempty"`
	ClientKeyAlgorithm         string                      `yaml:"clientKeyAlgorithm,omitempty"`
	ClientKeyPassPhrase        *SecretRef                  `yaml:"clientKeyPassPhrase,omitempty"`
	DelegateSelectors          []string                    `yaml:"delegateSelectors,omitempty"`
	Username                   string                      `yaml:"username,omitempty"`
	UsernameSecretId           *SecretRef                  `yaml:"usernameSecretId,omitempty"`
	ContinuousEfficiencyConfig *ContinuousEfficiencyConfig `yaml:"continuousEfficiencyConfig,omitempty"`
	MasterUrl                  string                      `yaml:"masterUrl,omitempty"`
	ServiceAccountToken        *SecretRef                  `yaml:"serviceAccountToken,omitempty"`
	SkipValidation             bool                        `yaml:"skipValidation,omitempty"`
	UseKubernetesDelegate      bool                        `yaml:"useKubernetesDelegate,omitempty"`
	UseEncryptedUsername       bool                        `yaml:"useEncryptedUsername,omitempty"`
	Password                   *SecretRef                  `yaml:"password,omitempty"`
	OIDCClientId               *SecretRef                  `yaml:"oidcClientId,omitempty"`
	OIDCIdentityProviderUrl    string                      `yaml:"oidcIdentityProviderUrl,omitempty"`
	OIDCPassword               *SecretRef                  `yaml:"oidcPassword,omitempty"`
	OIDCSecret                 *SecretRef                  `yaml:"oidcSecret,omitempty"`
	OIDCScopes                 string                      `yaml:"oidcScopes,omitempty"`
	OIDCUsername               string                      `yaml:"oidcUsername,omitempty"`
	UsageRestrictions          *UsageRestrictions          `yaml:"usageRestrictions,omitempty"`
}

func (a *KubernetesCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &KubernetesCloudProvider{})
}

type ContinuousEfficiencyConfig struct {
	ContinuousEfficiencyEnabled bool `json:"continuousEfficiencyEnabled,omitempty"`
}

type SpotInstCloudProvider struct {
	HarnessApiVersion HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Id                string             `yaml:"id,omitempty"`
	Name              string             `yaml:"-"`
	Type              ObjectType         `yaml:"type,omitempty"`
	AccountId         string             `yaml:"spotInstAccountId,omitempty"`
	Token             *SecretRef         `yaml:"spotInstToken,omitempty"`
	UsageRestrictions *UsageRestrictions `yaml:"usageRestrictions,omitempty"`
}

func (a *SpotInstCloudProvider) IsEmpty() bool {
	return reflect.DeepEqual(a, &SpotInstCloudProvider{})
}

type UsageRestrictions struct {
	AppEnvRestrictions []*AppEnvRestriction `yaml:"appEnvRestrictions,omitempty"`
}

type AppEnvRestriction struct {
	AppFilter *AppFilter `yaml:"appFilter,omitempty"`
	EnvFilter *EnvFilter `yaml:"envFilter,omitempty"`
}

type AppFilter struct {
	FilterType  ApplicationFilterType `yaml:"filterType,omitempty"`
	EntityNames []string              `yaml:"entityNames,omitempty"`
}

type EnvFilter struct {
	FilterTypes []EnvironmentFilterType `yaml:"filterTypes,omitempty"`
	EntityNames []string                `yaml:"entityNames,omitempty"`
}

type SecretRef struct {
	SecretManagerType SecretManagerType
	Name              string
}

type YamlPath string

type Environment struct {
	HarnessApiVersion                  HarnessApiVersion       `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type                               ObjectType              `yaml:"type" json:"type"`
	Id                                 string                  `yaml:"-"`
	Name                               string                  `yaml:"-"`
	ConfigMapYamlByServiceTemplateName *map[string]interface{} `yaml:"configMapYamlByServiceTemplateName,omitempty"`
	Description                        string                  `yaml:"description,omitempty"`
	EnvironmentType                    EnvironmentType         `yaml:"environmentType,omitempty"`
	VariableOverrides                  []*VariableOverride     `yaml:"variableOverrides,omitempty"`
	ApplicationId                      string                  `yaml:"-"`
}

func (a *Environment) IsEmpty() bool {
	return reflect.DeepEqual(a, &Environment{})
}

func (e *Environment) Validate() (bool, error) {
	return utils.RequiredStringFieldsSet(e, []string{"Name", "EnvironmentType", "ApplicationId"})
}

type VariableOverride struct {
	Name        string            `yaml:"name,omitempty"`
	ServiceName string            `yaml:"serviceName,omitempty"`
	Value       string            `yaml:"value,omitempty"`
	ValueType   VariableValueType `yaml:"valueType,omitempty"`
}

type InfrastructureDefinition struct {
	HarnessApiVersion         HarnessApiVersion       `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type                      ObjectType              `yaml:"type" json:"type"`
	Id                        string                  `yaml:"-"`
	Name                      string                  `yaml:"-"`
	ApplicationId             string                  `yaml:"-"`
	EnvironmentId             string                  `yaml:"-"`
	CloudProviderType         CloudProviderType       `yaml:"cloudProviderType,omitempty"`
	DeploymentType            DeploymentType          `yaml:"deploymentType,omitempty"`
	InfrastructureDetail      []*InfrastructureDetail `yaml:"infrastructure,omitempty"`
	Provisioner               string                  `yaml:"provisioner,omitempty"`
	DeploymentTypeTemplateUri string                  `yaml:"deploymentTypeTemplateUri,omitempty"`
	ScopedServices            []string                `yaml:"scopedServices,omitempty"`
}

func (a *InfrastructureDefinition) IsEmpty() bool {
	return reflect.DeepEqual(a, &InfrastructureDefinition{})
}

func (i *InfrastructureDefinition) Validate() (bool, error) {
	if _, err := utils.RequiredStringFieldsSet(i, []string{"ApplicationId", "EnvironmentId"}); err != nil {
		return false, err
	}

	if len(i.InfrastructureDetail) != 1 {
		return false, errors.New("expect one infrastructure detail to be set")
	}

	return true, nil
}

type InfrastructureDetail struct {
	Type                          InfrastructureType `yaml:"type,omitempty"`
	AmiDeploymentType             AmiDeploymentType  `yaml:"amiDeploymentType,omitempty"`
	ASGIdentifiesWorkload         bool               `yaml:"asgIdentifiesWorkload,omitempty"`
	AssignPublicIp                bool               `yaml:"assignPublicIp,omitempty"`
	AutoscalingGroupName          string             `yaml:"autoScalingGroupName,omitempty"`
	AwsInstanceFilter             *AwsInstanceFilter `yaml:"awsInstanceFilter,omitempty"`
	BaseVMSSName                  string             `yaml:"baseVMSSName,omitempty"`
	ClassicLoadBalancers          []string           `yaml:"classicLoadBalancers,omitempty"`
	CloudProviderName             string             `yaml:"cloudProviderName,omitempty"`
	ClusterName                   string             `yaml:"clusterName,omitempty"`
	DesiredCapacity               int                `yaml:"desiredCapacity,omitempty"`
	ExecutionRole                 string             `yaml:"executionRole,omitempty"`
	Expressions                   map[string]string  `yaml:"expressions,omitempty"`
	HostConnectionAttrs           string             `yaml:"hostConnectionAttrs,omitempty"`
	HostConnectionAttrsName       string             `yaml:"hostConnectionAttrsName,omitempty"`
	HostConnectionType            string             `yaml:"hostConnectionType,omitempty"`
	HostNameConvention            string             `yaml:"hostNameConvention,omitempty"`
	HostNames                     []string           `yaml:"hostNames,omitempty"`
	IamRole                       string             `yaml:"iamRole,omitempty"`
	InfraVariables                *InfraVariable     `yaml:"infraVariables,omitempty"`
	LaunchType                    string             `yaml:"launchType,omitempty"`
	LoadBalancerName              string             `yaml:"loadBalancerName,omitempty"`
	Namespace                     string             `yaml:"namespace,omitempty"`
	Organization                  string             `yaml:"organization,omitempty"`
	Region                        string             `yaml:"region,omitempty"`
	ReleaseName                   string             `yaml:"releaseName,omitempty"`
	ResourceGroup                 string             `yaml:"resourceGroup,omitempty"`
	ResourceGroupName             string             `yaml:"resourceGroupName,omitempty"`
	SecurityGroupIds              []string           `yaml:"securityGroupIds,omitempty"`
	SetDesiredCapacity            bool               `yaml:"setDesiredCapacity,omitempty"`
	Space                         string             `yaml:"space,omitempty"`
	SpotinstCloudProviderName     string             `yaml:"spotinstCloudProviderName,omitempty"`
	SpotinstElastiGroupJson       string             `yaml:"spotinstElastiGroupJson,omitempty"`
	StageClassicLoadBalancers     []string           `yaml:"stageClassicLoadBalancers,omitempty"`
	StageTargetGroupArns          []string           `yaml:"stageTargetGroupArns,omitempty"`
	SubnetIds                     []string           `yaml:"subnetIds,omitempty"`
	SubscriptionId                string             `yaml:"subscriptionId,omitempty"`
	TargetGroupArns               []string           `yaml:"targetGroupArns,omitempty"`
	UseAutoScalingGroup           bool               `yaml:"useAutoScalingGroup,omitempty"`
	UsePublicDns                  bool               `yaml:"usePublicDns,omitempty"`
	Username                      string             `yaml:"username,omitempty"`
	UseTrafficShift               bool               `yaml:"useTrafficShift,omitempty"`
	VmssAuthType                  VmssAuthType       `yaml:"vmssAuthType,omitempty"`
	VmssDeploymentType            VmssDeploymentType `yaml:"vmssDeploymentType,omitempty"`
	VpcId                         string             `yaml:"vpcId,omitempty"`
	WinRmConnectionAttributesName string             `yaml:"winRmConnectionAttributesName,omitempty"`
}

type InfrastructureAwsSSH struct {
	AwsInstanceFilter       *AwsInstanceFilter `yaml:"awsInstanceFilter,omitempty"`
	CloudProviderName       string             `yaml:"cloudProviderName,omitempty"`
	AutoscalingGroupName    string             `yaml:"autoScalingGroupName,omitempty"`
	DesiredCapacity         int                `yaml:"desiredCapacity,omitempty"`
	HostConnectionAttrsName string             `yaml:"hostConnectionAttrsName,omitempty"`
	HostConnectionType      HostConnectionType `yaml:"hostConnectionType,omitempty"`
	LoadBalancerName        string             `yaml:"loadBalancerName,omitempty"`
	HostNameConvention      string             `yaml:"hostNameConvention,omitempty"`
	Region                  string             `yaml:"region,omitempty"`
	SetDesiredCapacity      bool               `yaml:"setDesiredCapacity,omitempty"`
	UseAutoScalingGroup     bool               `yaml:"useAutoScalingGroup,omitempty"`
	UsePublicDns            bool               `yaml:"usePublicDns,omitempty"`
	Expressions             AwsSSHExpressions  `yaml:"expressions,omitempty"`
}

type AwsSSHExpressions struct {
	LoadBalancerId       string `yaml:"loadBalancerId,omitempty"`
	AutoscalingGroupName string `yaml:"autoScalingGroupName,omitempty"`
	VpcIds               string `yaml:"vpcIds,omitempty"`
	Region               string `yaml:"region,omitempty"`
	Tags                 string `yaml:"tags,omitempty"`
}

type AwsInstanceFilter struct {
	Tags   []*AwsTag `yaml:"tags,omitempty"`
	VpcIds []string  `yaml:"vpcIds,omitempty"`
}

type AwsTag struct {
	Key   string `yaml:"key,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type InfrastructureAwsAmi struct {
	AmiDeploymentType         AmiDeploymentType `yaml:"amiDeploymentType,omitempty"`
	ASGIdentifiesWorkload     bool              `yaml:"asgIdentifiesWorkload,omitempty"`
	AutoscalingGroupName      string            `yaml:"autoScalingGroupName,omitempty"`
	ClassicLoadBalancers      []string          `yaml:"classicLoadBalancers,omitempty"`
	CloudProviderName         string            `yaml:"cloudProviderName,omitempty"`
	HostNameConvention        string            `yaml:"hostNameConvention,omitempty"`
	Region                    string            `yaml:"region,omitempty"`
	SpotinstCloudProviderName string            `yaml:"spotinstCloudProviderName,omitempty"`
	SpotinstElastiGroupJson   string            `yaml:"spotinstElastiGroupJson,omitempty"`
	StageClassicLoadBalancers []string          `yaml:"stageClassicLoadBalancers,omitempty"`
	StageTargetGroupArns      []string          `yaml:"stageTargetGroupArns,omitempty"`
	TargetGroupArns           []string          `yaml:"targetGroupArns,omitempty"`
	UseTrafficShift           bool              `yaml:"useTrafficShift,omitempty"`
}

type InfrastructureKubernetesDirect struct {
	CloudProviderName string `yaml:"cloudProviderName,omitempty"`
	Namespace         string `yaml:"namespace,omitempty"`
	ReleaseName       string `yaml:"releaseName,omitempty"`
}

type InfrastructureKubernetesGcp struct {
	CloudProviderName string `yaml:"cloudProviderName,omitempty"`
	ClusterName       string `yaml:"clusterName,omitempty"`
	Namespace         string `yaml:"namespace,omitempty"`
	ReleaseName       string `yaml:"releaseName,omitempty"`
}

type InfrastructureAzureVmss struct {
	BaseVMSSName        string             `yaml:"baseVMSSName,omitempty"`
	CloudProviderName   string             `yaml:"cloudProviderName,omitempty"`
	HostConnectionAttrs string             `yaml:"hostConnectionAttrs,omitempty"`
	ResourceGroupName   string             `yaml:"resourceGroupName,omitempty"`
	SubscriptionId      string             `yaml:"subscriptionId,omitempty"`
	Username            string             `yaml:"username,omitempty"`
	VmssAuthType        VmssAuthType       `yaml:"vmssAuthType,omitempty"`
	VmssDeploymentType  VmssDeploymentType `yaml:"vmssDeploymentType,omitempty"`
}

type InfrastructureAzureWebApp struct {
	CloudProviderName string `yaml:"cloudProviderName,omitempty"`
	ResourceGroup     string `yaml:"resourceGroup,omitempty"`
	SubscriptionId    string `yaml:"subscriptionId,omitempty"`
}

type InfrastructureTanzu struct {
	CloudProviderName string `yaml:"cloudProviderName,omitempty"`
	Organization      string `yaml:"organization,omitempty"`
	Space             string `yaml:"space,omitempty"`
}

type InfrastructureAwsEcs struct {
	AssignPublicIp    bool             `yaml:"assignPublicIp,omitempty"`
	CloudProviderName string           `yaml:"cloudProviderName,omitempty"`
	ClusterName       string           `yaml:"clusterName,omitempty"`
	ExecutionRole     string           `yaml:"executionRole,omitempty"`
	LaunchType        AwsEcsLaunchType `yaml:"launchType,omitempty"`
	Region            string           `yaml:"region,omitempty"`
	SecurityGroupIds  []string         `yaml:"securityGroupIds,omitempty"`
	SubnetIds         []string         `yaml:"subnetIds,omitempty"`
	VpcId             string           `yaml:"vpcId,omitempty"`
}

type InfrastructureDataCenterWinRM struct {
	CloudProviderName             string   `yaml:"cloudProviderName,omitempty"`
	HostNames                     []string `yaml:"hostNames,omitempty"`
	WinRmConnectionAttributesName string   `yaml:"winRmConnectionAttributesName,omitempty"`
}

type InfrastructureDataCenterSSH struct {
	CloudProviderName       string   `yaml:"cloudProviderName,omitempty"`
	HostConnectionAttrsName string   `yaml:"hostConnectionAttrsName,omitempty"`
	HostNames               []string `yaml:"hostNames,omitempty"`
}

type InfrastructureAwsLambda struct {
	CloudProviderName string            `yaml:"cloudProviderName,omitempty"`
	IamRole           string            `yaml:"iamRole,omitempty"`
	Region            string            `yaml:"region,omitempty"`
	SecurityGroupIds  []string          `yaml:"securityGroupIds,omitempty"`
	SubnetIds         []string          `yaml:"subnetIds,omitempty"`
	VpcId             string            `yaml:"vpcId,omitempty"`
	Expressions       map[string]string `yaml:"expressions,omitempty"`
}

type InfrastructureAwsWinRM struct {
	AutoscalingGroupName    string             `yaml:"autoScalingGroupName,omitempty"`
	CloudProviderName       string             `yaml:"cloudProviderName,omitempty"`
	DesiredCapacity         int                `yaml:"desiredCapacity,omitempty"`
	HostConnectionAttrsName string             `yaml:"hostConnectionAttrsName,omitempty"`
	HostConnectionType      HostConnectionType `yaml:"hostConnectionType,omitempty"`
	HostNameConvention      string             `yaml:"hostNameConvention,omitempty"`
	LoadBalancerName        string             `yaml:"loadBalancerName,omitempty"`
	Region                  string             `yaml:"region,omitempty"`
	SetDesiredCapacity      bool               `yaml:"setDesiredCapacity,omitempty"`
	UseAutoScalingGroup     bool               `yaml:"useAutoScalingGroup,omitempty"`
	UsePublicDns            bool               `yaml:"usePublicDns,omitempty"`
}

type InfrastructureCustom struct {
	InfraVariables *InfraVariable `yaml:"infraVariables,omitempty"`
}

type InfraVariable struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}
