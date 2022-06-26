package nextgen

type SSHSpecificationType string

var SSHSpecificationTypes = struct {
	KerberosConfigDTO SSHSpecificationType
	SSHConfig         SSHSpecificationType
}{
	KerberosConfigDTO: "KerberosConfigDTO",
	SSHConfig:         "SSHConfig",
}

var SSHSpecificationTypeValues = []string{
	SSHSpecificationTypes.KerberosConfigDTO.String(),
	SSHSpecificationTypes.SSHConfig.String(),
}

func (e SSHSpecificationType) String() string {
	return string(e)
}
