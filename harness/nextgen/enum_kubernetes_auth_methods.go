package nextgen

type KubernetesAuthMethod string

var KubernetesAuthMethods = struct {
	UsernamePassword KubernetesAuthMethod
	ServiceAccount   KubernetesAuthMethod
	OpenIdConnect    KubernetesAuthMethod
	ClientKeyCert    KubernetesAuthMethod
}{
	UsernamePassword: "UsernamePassword",
	ServiceAccount:   "ServiceAccount",
	OpenIdConnect:    "OpenIdConnect",
	ClientKeyCert:    "ClientKeyCert",
}

func (k KubernetesAuthMethod) String() string {
	return string(k)
}
