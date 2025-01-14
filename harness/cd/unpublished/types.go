package unpublished

import "encoding/json"

type Package struct {
	MetaData         MetaData         `json:"metaData"`
	Resource         *json.RawMessage `json:"resource"`
	ResponseMessages []interface{}    `json:"responseMessages"`
}

type MetaData struct{}

type EncryptedTextResource struct {
	Offset         string           `json:"offset"`
	Start          int64            `json:"start"`
	Limit          string           `json:"limit"`
	Filters        []interface{}    `json:"filters"`
	Orders         []interface{}    `json:"orders"`
	FieldsIncluded []interface{}    `json:"fieldsIncluded"`
	FieldsExcluded []interface{}    `json:"fieldsExcluded"`
	Secrets        []*EncryptedText `json:"response"`
	Total          int64            `json:"total"`
	CurrentPage    int64            `json:"currentPage"`
	Empty          bool             `json:"empty"`
	PageSize       int64            `json:"pageSize"`
	Or             bool             `json:"or"`
}

type SecretManager struct {
	Name                                          string             `json:"name"`
	AccessKey                                     string             `json:"accessKey,omitempty"`
	SecretKey                                     string             `json:"secretKey,omitempty"`
	KmsArn                                        string             `json:"kmsArn,omitempty"`
	Region                                        string             `json:"region"`
	AssumeIamRoleOnDelegate                       bool               `json:"assumeIamRoleOnDelegate,omitempty"`
	AssumeStsRoleOnDelegate                       bool               `json:"assumeStsRoleOnDelegate,omitempty"`
	AssumeStsRoleDuration                         int64              `json:"assumeStsRoleDuration,omitempty"`
	RoleArn                                       interface{}        `json:"roleArn"`
	ExternalName                                  interface{}        `json:"externalName"`
	DelegateSelectors                             interface{}        `json:"delegateSelectors"`
	UUID                                          string             `json:"uuid"`
	EncryptionType                                string             `json:"encryptionType"`
	AccountID                                     string             `json:"accountId"`
	NumOfEncryptedValue                           int64              `json:"numOfEncryptedValue"`
	EncryptedBy                                   *User              `json:"encryptedBy"`
	CreatedBy                                     *User              `json:"createdBy"`
	CreatedAt                                     int64              `json:"createdAt"`
	LastUpdatedBy                                 *User              `json:"lastUpdatedBy"`
	LastUpdatedAt                                 int64              `json:"lastUpdatedAt"`
	NextTokenRenewIteration                       interface{}        `json:"nextTokenRenewIteration"`
	ManuallyEnteredSecretEngineMigrationIteration interface{}        `json:"manuallyEnteredSecretEngineMigrationIteration"`
	UsageRestrictions                             *UsageRestrictions `json:"usageRestrictions"`
	ScopedToAccount                               bool               `json:"scopedToAccount"`
	TemplatizedFields                             interface{}        `json:"templatizedFields"`
	Default                                       bool               `json:"default"`
	Templatized                                   bool               `json:"templatized"`
	SecretNamePrefix                              string             `json:"secretNamePrefix,omitempty"`
	ProjectID                                     string             `json:"projectId,omitempty"`
	KeyRing                                       string             `json:"keyRing,omitempty"`
	KeyName                                       string             `json:"keyName,omitempty"`
	Credentials                                   string             `json:"credentials,omitempty"`
	UsePutSecret                                  bool               `json:"usePutSecret,omitempty"`
	ForceDeleteWithoutRecovery                    bool               `json:"forceDeleteWithoutRecovery,omitempty"`
	RecoveryWindowInDays                          int64              `json:"recoveryWindowInDays,omitempty"`
}

type User struct {
	UUID  string `json:"uuid,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type EncryptedText struct {
	UUID                                 string             `json:"uuid"`
	CreatedBy                            *User              `json:"createdBy"`
	CreatedAt                            int64              `json:"createdAt"`
	LastUpdatedBy                        *User              `json:"lastUpdatedBy"`
	LastUpdatedAt                        int64              `json:"lastUpdatedAt"`
	Name                                 string             `json:"name"`
	EncryptionKey                        string             `json:"encryptionKey"`
	EncryptedValue                       string             `json:"encryptedValue"`
	Path                                 interface{}        `json:"path"`
	Parameters                           []interface{}      `json:"parameters"`
	Type                                 string             `json:"type"`
	Parents                              []Parent           `json:"parents"`
	AccountID                            string             `json:"accountId"`
	Enabled                              bool               `json:"enabled"`
	KmsID                                string             `json:"kmsId"`
	AdditionalMetadata                   AdditionalMetadata `json:"additionalMetadata"`
	EncryptionType                       string             `json:"encryptionType"`
	FileSize                             int64              `json:"fileSize"`
	AppIDS                               []interface{}      `json:"appIds"`
	ServiceIDS                           []interface{}      `json:"serviceIds"`
	EnvIDS                               []interface{}      `json:"envIds"`
	BackupEncryptedValue                 string             `json:"backupEncryptedValue"`
	BackupEncryptionKey                  string             `json:"backupEncryptionKey"`
	BackupKmsID                          interface{}        `json:"backupKmsId"`
	BackupEncryptionType                 interface{}        `json:"backupEncryptionType"`
	ServiceVariableIDS                   interface{}        `json:"serviceVariableIds"`
	SearchTags                           map[string]int64   `json:"searchTags"`
	ScopedToAccount                      bool               `json:"scopedToAccount"`
	UsageRestrictions                    *UsageRestrictions `json:"usageRestrictions"`
	InheritScopesFromSM                  bool               `json:"inheritScopesFromSM"`
	NextMigrationIteration               interface{}        `json:"nextMigrationIteration"`
	NextAwsToGcpKmsMigrationIteration    interface{}        `json:"nextAwsToGcpKmsMigrationIteration"`
	NextLocalToGcpKmsMigrationIteration  int64              `json:"nextLocalToGcpKmsMigrationIteration"`
	NextAwsKmsToGcpKmsMigrationIteration interface{}        `json:"nextAwsKmsToGcpKmsMigrationIteration"`
	Base64Encoded                        bool               `json:"base64Encoded"`
	EncryptedBy                          string             `json:"encryptedBy"`
	SetupUsage                           int64              `json:"setupUsage"`
	RunTimeUsage                         int64              `json:"runTimeUsage"`
	ChangeLog                            int64              `json:"changeLog"`
	Keywords                             []string           `json:"keywords"`
	NgMetadata                           interface{}        `json:"ngMetadata"`
	HideFromListing                      bool               `json:"hideFromListing"`
	ReferencedSecret                     bool               `json:"referencedSecret"`
	ParameterizedSecret                  bool               `json:"parameterizedSecret"`
	InlineSecret                         bool               `json:"inlineSecret"`
}

type Credential struct {
	EnvID                        string             `json:"envId"`
	UUID                         string             `json:"uuid"`
	AppID                        string             `json:"appId"`
	CreatedBy                    *User              `json:"createdBy"`
	CreatedAt                    int64              `json:"createdAt"`
	LastUpdatedAt                int64              `json:"lastUpdatedAt"`
	AccountID                    string             `json:"accountId"`
	Name                         string             `json:"name"`
	Value                        *CredentialValue   `json:"value"`
	ValidationAttributes         interface{}        `json:"validationAttributes"`
	Category                     string             `json:"category"`
	AppIDS                       interface{}        `json:"appIds"`
	UsageRestrictions            *UsageRestrictions `json:"usageRestrictions"`
	ArtifactStreamCount          int64              `json:"artifactStreamCount"`
	ArtifactStreams              interface{}        `json:"artifactStreams"`
	Sample                       bool               `json:"sample"`
	NextIteration                interface{}        `json:"nextIteration"`
	NextSecretMigrationIteration int64              `json:"nextSecretMigrationIteration"`
	SecretsMigrated              bool               `json:"secretsMigrated"`
	ConnectivityError            interface{}        `json:"connectivityError"`
	EncryptionType               string             `json:"encryptionType"`
	EncryptedBy                  string             `json:"encryptedBy"`
}

type CredentialValue struct {
	Type                   string          `json:"type"`
	ConnectionType         string          `json:"connectionType,omitempty"`
	AccessType             string          `json:"accessType,omitempty"`
	UserName               string          `json:"userName"`
	SSHPassword            string          `json:"sshPassword"`
	SSHPort                int64           `json:"sshPort,omitempty"`
	Key                    string          `json:"key"`
	AccountID              string          `json:"accountId"`
	Keyless                bool            `json:"keyless,omitempty"`
	KeyPath                interface{}     `json:"keyPath"`
	Passphrase             string          `json:"passphrase"`
	AuthenticationScheme   string          `json:"authenticationScheme"`
	Role                   interface{}     `json:"role"`
	PublicKey              interface{}     `json:"publicKey"`
	SignedPublicKey        interface{}     `json:"signedPublicKey"`
	SSHVaultConfigID       interface{}     `json:"sshVaultConfigId"`
	SSHVaultConfig         interface{}     `json:"sshVaultConfig"`
	KerberosConfig         *KerberosConfig `json:"kerberosConfig"`
	KerberosPassword       string          `json:"kerberosPassword"`
	VaultSSH               bool            `json:"vaultSSH,omitempty"`
	SettingType            string          `json:"settingType"`
	CERTValidationRequired bool            `json:"certValidationRequired"`
	Domain                 string          `json:"domain,omitempty"`
	Username               string          `json:"username,omitempty"`
	Password               string          `json:"password,omitempty"`
	UseSSL                 bool            `json:"useSSL,omitempty"`
	Port                   int64           `json:"port,omitempty"`
	SkipCERTChecks         bool            `json:"skipCertChecks,omitempty"`
	UseKeyTab              bool            `json:"useKeyTab,omitempty"`
	KeyTabFilePath         string          `json:"keyTabFilePath,omitempty"`
	UseNoProfile           bool            `json:"useNoProfile,omitempty"`
}

type KerberosConfig struct {
	Principal      string      `json:"principal"`
	GenerateTGT    bool        `json:"generateTGT"`
	Realm          string      `json:"realm"`
	KeyTabFilePath interface{} `json:"keyTabFilePath"`
}

type AdditionalMetadata struct {
	Values MetaData `json:"values"`
}

type Parent struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	FieldName string `json:"fieldName"`
}

type UsageRestrictions struct {
	AppEnvRestrictions []AppEnvRestriction `json:"appEnvRestrictions"`
}

type AppEnvRestriction struct {
	AppFilter AppFilter `json:"appFilter"`
	EnvFilter EnvFilter `json:"envFilter"`
}

type AppFilter struct {
	Type       string      `json:"type"`
	IDS        interface{} `json:"ids"`
	FilterType string      `json:"filterType"`
}

type EnvFilter struct {
	Type        string      `json:"type"`
	IDS         interface{} `json:"ids"`
	FilterTypes []string    `json:"filterTypes"`
}
