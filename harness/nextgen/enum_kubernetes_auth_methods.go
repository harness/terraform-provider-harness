package nextgen

type KubernetesAuthType string

var KubernetesAuthTypes = struct {
	UsernamePassword KubernetesAuthType
	ServiceAccount   KubernetesAuthType
	OpenIdConnect    KubernetesAuthType
	ClientKeyCert    KubernetesAuthType
}{
	UsernamePassword: "UsernamePassword",
	ServiceAccount:   "ServiceAccount",
	OpenIdConnect:    "OpenIdConnect",
	ClientKeyCert:    "ClientKeyCert",
}

func (k KubernetesAuthType) String() string {
	return string(k)
}
