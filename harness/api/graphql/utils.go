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

func (s *LDAPSettings) IsEmpty() bool {
	return s.GroupDN+s.GroupName+s.SSOProviderId == ""
}

func (s *SAMLSettings) IsEmpty() bool {
	return s.GroupName+s.GroupName+s.SSOProviderId == ""
}

func (s *UserGroupPermissions) IsEmpty() bool {
	return (s.AccountPermissions == nil || len(s.AccountPermissions.AccountPermissionTypes) == 0) && len(s.AppPermissions) == 0
}

func (s *NotificationSettings) IsEmpty() bool {
	return len(s.GroupEmailAddresses) == 0 && s.MicrosoftTeamsWebhookUrl == "" && !s.SendMailToNewMembers && !s.SendNotificationToMembers && (s.SlackNotificationSetting == nil || s.SlackNotificationSetting.IsEmpty())
}

func (s *SlackNotificationSetting) IsEmpty() bool {
	return s.SlackChannelName == "" && s.SlackWebhookUrl == ""
}
