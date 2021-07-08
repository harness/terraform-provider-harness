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
