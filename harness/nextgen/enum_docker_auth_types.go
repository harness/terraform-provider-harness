package nextgen

type DockerAuthType string

var DockerAuthTypes = struct {
	Anonymous        DockerAuthType
	UsernamePassword DockerAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

func (e DockerAuthType) String() string {
	return string(e)
}
