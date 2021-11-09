package nextgen

type HttpHelmAuthType string

var HttpHelmAuthTypes = struct {
	Anonymous        HttpHelmAuthType
	UsernamePassword HttpHelmAuthType
}{
	Anonymous:        "Anonymous",
	UsernamePassword: "UsernamePassword",
}

func (e HttpHelmAuthType) String() string {
	return string(e)
}
