package cac

type Validation interface {
	Validate() (bool, error)
}

// type Entity interface {
// 	GetPath() (string, error)
// }

type Application struct {
	HarnessApiVersion HarnessApiVersion `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              ObjectType        `yaml:"type" json:"type"`
	Id                string            `yaml:"-"`
	Name              string            `yaml:"-"`
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
	AccessKeySecretId      *SecretRef                 `yaml:"AccessKeySecretId,omitempty"`
	AssumeCrossAccountRole bool                       `yaml:"assumeCrossAccountRole,omitempty"`
	CrossAccountAttributes *AwsCrossAccountAttributes `yaml:"crossAccountAttributes,omitempty"`
	SecretKey              *SecretRef                 `yaml:"secretKey,omitempty"`
	UseIRSA                bool                       `yaml:"useIRSA,omitempty"`
	UseEc2IamCredentials   bool                       `yaml:"useEc2IamCredentials,omitempty"`
	UsageRestrictions      *UsageRestrictions         `yaml:"usageRestrictions,omitempty"`
	DelegateSelector       string                     `yaml:"tag,omitempty"`
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

type PhysicalDatacenterCloudProvider struct {
	HarnessApiVersion HarnessApiVersion  `yaml:"harnessApiVersion" json:"harnessApiVersion"`
	Type              ObjectType         `yaml:"type" json:"type"`
	Name              string             `yaml:"-"`
	Id                string             `yaml:"-"`
	UsageRestrictions *UsageRestrictions `yaml:"usageRestrictions,omitempty"`
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
	EnvironmentType                    EnvironmentType         `yaml:"environmentType,omitempty"`
	VariableOverrides                  []*VariableOverride     `yaml:"variableOverrides,omitempty"`
	ApplicationId                      string                  `yaml:"-"`
}

type VariableOverride struct {
	Name        string            `yaml:"name,omitempty"`
	ServiceName string            `yaml:"serviceName,omitempty"`
	Value       string            `yaml:"value,omitempty"`
	ValueType   VariableValueType `yaml:"valueType,omitempty"`
}
