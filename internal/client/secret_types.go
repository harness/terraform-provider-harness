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

type sshAuthenticationType struct {
	KerberosAuthentication string
	SSHAuthentication      string
}

var SSHAuthenticationTypes *sshAuthenticationType = &sshAuthenticationType{
	KerberosAuthentication: "KerberosAuthentication",
	SSHAuthentication:      "SSHAuthentication",
}

type secretTypes struct {
	EncryptedFile   string
	EncryptedText   string
	SSHCredential   string
	WinRMCredential string
}

var SecretTypes *secretTypes = &secretTypes{
	EncryptedFile:   "ENCRYPTED_FILE",
	EncryptedText:   "ENCRYPTED_TEXT",
	SSHCredential:   "SSH_CREDENTIAL",
	WinRMCredential: "WINRM_CREDENTIAL",
}

type SecretManager struct {
	Id         string      `json:"id"`
	Name       string      `json:""`
	UsageScope *UsageScope `json:"usageScope"`
}

type Secret struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	SecretType string      `json:"secretType"`
	UsageScope *UsageScope `json:"usageScope"`
}

// type EncryptedTextByName struct {
// 	Data struct {
// 		SecretByName struct {
// 			EncryptedText
// 		} `json:"secretByName"`
// 	} `json:"data"`
// 	GraphQLStandardResponse
// }

// type EncryptedFileByName struct {
// 	Data struct {
// 		SecretByName struct {
// 			EncryptedText
// 		} `json:"secretByName"`
// 	} `json:"data"`
// 	GraphQLStandardResponse
// }

type EncryptedText struct {
	Secret
	InheritScopesFromSM bool        `json:"inheritScopesFromSM"`
	ScopedToAccount     bool        `json:"scopedToAccount"`
	SecretManagerId     string      `json:"secretManagerId"`
	UsageScope          *UsageScope `json:"usageScope"`
}

type EncryptedFile struct {
	Secret
	ScopedToAccount bool   `json:"scopedToAccount"`
	SecretManagerId string `json:"secretManagerId"`
}

type SSHCredential struct {
	Secret
	AuthenticationType string `json:"authenticationType"`
}

type WinRMCredential struct {
	Secret
	Domain        string `json:"domain"`
	Port          int    `json:"port"`
	SkipCertCheck bool   `json:"skipCertCheck"`
	UseSSL        bool   `json:"useSSL"`
	UserName      string `json:"userName"`
}
