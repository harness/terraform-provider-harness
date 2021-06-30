package graphql

import (
	"errors"
)

func (d *CustomCommitDetails) IsEmpty() bool {
	if d == nil {
		return true
	}

	return (d.AuthorEmailId + d.AuthorName + d.CommitMessage) == ""
}

func (c *SSHCredential) SetSSHAuthenticationType() error {

	if c.SSHAuthentication != nil && c.SSHAuthentication.isValid() {
		c.AuthenticationType = SSHAuthenticationTypes.SSHAuthentication
		c.KerberosAuthentication = nil
	} else if c.KerberosAuthentication != nil && c.KerberosAuthentication.isValid() {
		c.AuthenticationType = SSHAuthenticationTypes.KerberosAuthentication
		c.SSHAuthentication = nil
	} else {
		return errors.New("invalid SSH Authentication type")
	}

	return nil
}

func (auth *SSHAuthentication) isValid() bool {
	return auth.Username != ""
}

func (auth *KerberosAuthentication) isValid() bool {
	return auth.Principal != "" && auth.Realm != ""
}

func (e Enum) ToString() string {
	return string(e)
}

// func (t ApplicationFilterType) ToString() string {
// 	return string(t)
// }

// func (t EnvironmentFilterType) ToString() string {
// 	return string(t)
// }

// func (t CloudProviderType) ToString() string {
// 	return string(t)
// }

// func (t SSHAuthenticationType) ToString() string {
// 	return string(t)
// }

// func (t SSHCredentialTypes) ToString() string {
// 	return string(t)
// }

// func (t WinRMAuthenticationType) ToString() string {
// 	return string(t)
// }

// func (t AwsCredentialsType) ToString() string {
// 	return string(t)
// }

// func (t ClusterDetailsType) ToString() string {
// 	return string(t)
// }

// func (t ManualClusterDetailsAuthenticationType) ToString() string {
// 	return string(t)
// }

// func (t SecretType) ToString() string {
// 	return string(t)
// }
