package nextgen

type OciHelmAuthType string

var OciHelmAuthTypes = struct {
	Anonymous        OciHelmAuthType
	UsernamePassword OciHelmAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

func (e OciHelmAuthType) String() string {
	return string(e)
}
