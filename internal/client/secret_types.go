package client

type SecretClient struct {
	APIClient *ApiClient
}

// Get the client for interacting with Harness Applications
func (c *ApiClient) Secrets() *SecretClient {
	return &SecretClient{
		APIClient: c,
	}
}

type sshCredentialAuthenticationType struct {
	SSHAuthentication      string
	KerberosAuthentication string
}

var SSHAuthenticationTypes = &sshCredentialAuthenticationType{
	SSHAuthentication:      "SSH_AUTHENTICATION",
	KerberosAuthentication: "KERBEROS_AUTHENTICATION",
}

type winRMAuthenticationType struct {
	NTLM string
	// Kerberos string
}

var WinRMAuthenticationTypes = &winRMAuthenticationType{
	NTLM: "NTLM",
}

type SSHAuthentication struct {
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
}

type SSHCredential struct {
	Secret
	AuthenticationType     string
	KerberosAuthentication *KerberosAuthentication
	SSHAuthentication      *SSHAuthentication
}

type secretType struct {
	EncryptedFile   string
	EncryptedText   string
	SSHCredential   string
	WinRMCredential string
}

var SecretTypes *secretType = &secretType{
	EncryptedFile:   "ENCRYPTED_FILE",
	EncryptedText:   "ENCRYPTED_TEXT",
	SSHCredential:   "SSH_CREDENTIAL",
	WinRMCredential: "WINRM_CREDENTIAL",
}

type SecretManager struct {
	Id         string
	Name       string
	UsageScope *UsageScope
}

type EncryptedFile struct {
	Secret
	EncryptedText
}

type WinRMCredential struct {
	Secret
	Domain               string `json:"domain,omitempty"`
	Port                 int    `json:"port,omitempty"`
	SkipCertCheck        bool   `json:"skipCertCheck,omitempty"`
	UseSSL               bool   `json:"useSSL,omitempty"`
	UserName             string `json:"username,omitempty"`
	AuthenticationScheme string `json:"authenticationScheme,omitempty"`
}

type CreateSecretInput struct {
	ClientMutationId string                `json:"clientMutationId,omitempty"`
	EncryptedText    *EncryptedTextInput   `json:"encryptedText,omitempty"`
	SecretType       string                `json:"secretType,omitempty"`
	SSHCredential    *SSHCredentialInput   `json:"sshCredential,omitempty"`
	WinRMCredential  *WinRMCredentialInput `json:"winRMCredential,omitempty"`
}

type UpdateSecretInput struct {
	ClientMutationId string                 `json:"clientMutationId,omitempty"`
	EncryptedText    *UpdateEncryptedText   `json:"encryptedText,omitempty"`
	SecretId         string                 `json:"secretId,omitempty"`
	SecretType       string                 `json:"secretType,omitempty"`
	SSHCredential    *UpdateSSHCredential   `json:"sshCredential,omitempty"`
	WinRMCredential  *UpdateWinRMCredential `json:"winRMCredential,omitempty"`
}

type Secret struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	SecretType string      `json:"secretType,omitempty"`
	UsageScope *UsageScope `json:"usageScope,omitempty"`
}

type EncryptedText struct {
	Secret
	InheritScopesFromSM bool   `json:"inheritScopesFromSM,omitempty"`
	ScopedToAccount     bool   `json:"scopedToAccount,omitempty"`
	SecretManagerId     string `json:"secretManagerId,omitempty"`
}

type EncryptedTextInput struct {
	InheritScopesFromSM bool        `json:"inheritScopesFromSM,omitempty"`
	Name                string      `json:"name,omitempty"`
	ScopedToAccount     bool        `json:"scopedToAccount,omitempty"`
	SecretManagerId     string      `json:"secretManagerId,omitempty"`
	SecretReference     string      `json:"secretReference,omitempty"`
	UsageScope          *UsageScope `json:"usageScope,omitempty"`
	Value               string      `json:"value,omitempty"`
}

type UpdateEncryptedText struct {
	InheritScopesFromSM bool        `json:"inheritScopesFromSM,omitempty"`
	Name                string      `json:"name,omitempty"`
	ScopedToAccount     bool        `json:"scopedToAccount,omitempty"`
	SecretReference     string      `json:"secretReference,omitempty"`
	UsageScope          *UsageScope `json:"usageScope,omitempty"`
	Value               string      `json:"value,omitempty"`
}

type UpdateSSHCredential struct {
	AuthenticationScheme   string                  `json:"authenticationScheme,omitempty"`
	KerberosAuthentication *KerberosAuthentication `json:"kerberosAuthentication,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	SSHAuthentication      *SSHAuthenticationInput `json:"sshAuthentication,omitempty"`
	UsageScope             *UsageScope             `json:"usageScope,omitempty"`
}

type UpdateWinRMCredential struct {
	AuthenticationScheme string      `json:"authenticationScheme,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Name                 string      `json:"name,omitempty"`
	PasswordSecretId     string      `json:"passwordSecretID,omitempty"`
	Port                 int         `json:"port,omitempty"`
	SkipCertCheck        bool        `json:"skipCertCheck,omitempty"`
	UsageScope           *UsageScope `json:"usageScope,omitempty"`
	UseSSL               bool        `json:"useSSL,omitempty"`
	Username             string      `json:"username,omitempty"`
}

type SSHCredentialInput struct {
	SSHAuthenticationScheme string                 `json:"-"`
	KerberosAuthentication  KerberosAuthentication `json:"kerberosAuthentication,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	SSHAuthentication       SSHAuthenticationInput `json:"sshAuthentication,omitempty"`
	UsageScope              UsageScope             `json:"usageScope,omitempty"`
}
type SSHAuthenticationInput struct {
	Port                    int                     `json:"port,omitempty"`
	SSHAuthenticationMethod SSHAuthenticationMethod `json:"sshAuthenticationMethod,omitempty"`
	Username                string                  `json:"username,omitempty"`
}

type SSHAuthenticationMethod struct {
	InlineSSHKey      InlineSSHKey `json:"inlineSSHKey,omitempty"`
	ServerPassword    SSHPassword  `json:"serverPassword,omitempty"`
	SSHCredentialType string       `json:"sshCredentialType,omitempty"`
	SSHKeyFile        SSHKeyFile   `json:"sshKeyFile,omitempty"`
}

type InlineSSHKey struct {
	PassphraseSecretId string `json:"passphraseSecretId,omitempty"`
	SSHKeySecretFileId string `json:"sshKeySecretField,omitempty"`
}

type SSHPassword struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
}

type SSHKeyFile struct {
	PassphraseSecretId string `json:"passphraseSecretId,omitempty"`
	Path               string `json:"path,omitempty"`
}

type sshCredentialType struct {
	Password       string
	SSHKey         string
	SSHKeyFilePath string
}

var SSHCredentialTypes = &sshCredentialType{
	Password:       "PASSWORD",
	SSHKey:         "SSH_KEY",
	SSHKeyFilePath: "SSH_KEY_FILE_PATH",
}

type KerberosAuthentication struct {
	Port                int                 `json:"port,omitempty"`
	Principal           string              `json:"principal,omitempty"`
	Realm               string              `json:"realm,omitempty"`
	TGTGenerationMethod TGTGenerationMethod `json:"tgtGenerationMethod,omitempty"`
}
type TGTGenerationMethod struct {
	KerberosPassword   KerberosPassword `json:"kerberosPassword,omitempty"`
	KeyTabFile         KeyTabFile       `json:"keyTabFile,omitempty"`
	TGTGenerationUsing string           `json:"tgtGenerationUsing,omitempty"`
}

type tgtGenerationUsingOption struct {
	KeyTabFile string
	Password   string
}

var TGTGenerationUsingOptions = &tgtGenerationUsingOption{
	KeyTabFile: "KEY_TAB_FILE",
	Password:   "PASSWORD",
}

type KeyTabFile struct {
	FilePath string `json:"filePath,omitempty"`
}
type KerberosPassword struct {
	PasswordSecretId string `json:"passwordSecretId,omitempty"`
}
type WinRMCredentialInput struct {
	AuthenticationScheme string      `json:"authenticationScheme,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Name                 string      `json:"name,omitempty"`
	PasswordSecretId     string      `json:"passwordSecretId,omitempty"`
	Port                 int         `json:"port,omitempty"`
	SkipCertCheck        bool        `json:"skipCertCheck,omitempty"`
	UsageScope           *UsageScope `json:"usageScope,omitempty"`
	UseSSL               bool        `json:"useSSL,omitempty"`
	Username             string      `json:"username,omitempty"`
}

type environmentFilterType struct {
	NonProduction string
	Production    string
}

type applicationFilterType struct {
	All string
}

var ApplicationFilterTypes = &applicationFilterType{
	All: "ALL",
}

var EnvironmentFilterTypes = &environmentFilterType{
	NonProduction: "NON_PRODUCTION_ENVIRONMENTS",
	Production:    "PRODUCTION_ENVIRONMENTS",
}

type DeleteSecretInput struct {
	SecretId   string `json:"secretId,omitempty"`
	SecretType string `json:"secretType,omitempty"`
}
