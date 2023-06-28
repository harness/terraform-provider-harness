package nextgen

type RancherAuthType string

var RancherAuthTypes = struct {
	BearerToken RancherAuthType
}{
	BearerToken: "BearerToken",
}

func (e RancherAuthType) String() string {
	return string(e)
}
