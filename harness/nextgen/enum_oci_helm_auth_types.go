package nextgen

type OciHelmAuthType string

var OciHelmAuthTypes = struct {
	Anonymous        OciHelmAuthType
	UsernamePassword OciHelmAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

var OciHelmAuthTypeValues = []string{
	OciHelmAuthTypes.UsernamePassword.String(),
	OciHelmAuthTypes.Anonymous.String(),
}

func (e OciHelmAuthType) String() string {
	return string(e)
}
