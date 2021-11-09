package nextgen

type ArtifactoryAuthType string

var ArtifactoryAuthTypes = struct {
	Anonymous        ArtifactoryAuthType
	UsernamePassword ArtifactoryAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

func (e ArtifactoryAuthType) String() string {
	return string(e)
}
