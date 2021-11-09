package nextgen

type DockerRegistry string

var DockerRegistryTypes = struct {
	DockerHub DockerRegistry
	Harbor    DockerRegistry
	Other     DockerRegistry
	Quay      DockerRegistry
}{
	DockerHub: "DockerHub",
	Harbor:    "Harbor",
	Other:     "Other",
	Quay:      "Quay",
}

var DockerRegistryTypesSlice = []string{
	DockerRegistryTypes.DockerHub.String(),
	DockerRegistryTypes.Harbor.String(),
	DockerRegistryTypes.Other.String(),
	DockerRegistryTypes.Quay.String(),
}

func (e DockerRegistry) String() string {
	return string(e)
}
