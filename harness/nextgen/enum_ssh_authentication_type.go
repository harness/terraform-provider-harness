package nextgen

type SSHAuthenticationType string

var SSHAuthenticationTypes = struct {
	Kerberos SSHAuthenticationType
	SSH      SSHAuthenticationType
}{
	Kerberos: "Kerberos",
	SSH:      "SSH",
}

var SSHAuthenticationTypeValues = []string{
	SSHAuthenticationTypes.Kerberos.String(),
	SSHAuthenticationTypes.SSH.String(),
}

func (e SSHAuthenticationType) String() string {
	return string(e)
}
