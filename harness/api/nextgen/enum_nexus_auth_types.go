package nextgen

type NexusAuthType string

var NexusAuthTypes = struct {
	Anonymous        NexusAuthType
	UsernamePassword NexusAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

func (e NexusAuthType) String() string {
	return string(e)
}
