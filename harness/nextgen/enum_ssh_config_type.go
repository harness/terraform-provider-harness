package nextgen

type SSHConfigType string

var SSHConfigTypes = struct {
	Password     SSHConfigType
	KeyPath      SSHConfigType
	KeyReference SSHConfigType
}{
	Password:     "Password",
	KeyPath:      "KeyPath",
	KeyReference: "KeyReference",
}

var SSHConfigTypeValues = []string{
	SSHConfigTypes.Password.String(),
	SSHConfigTypes.KeyPath.String(),
	SSHConfigTypes.KeyReference.String(),
}

func (e SSHConfigType) String() string {
	return string(e)
}
