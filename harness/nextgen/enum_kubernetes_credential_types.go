package nextgen

type KubernetesCredentialType string

var KubernetesCredentialTypes = struct {
	InheritFromDelegate KubernetesCredentialType
	ManualConfig        KubernetesCredentialType
}{
	InheritFromDelegate: "InheritFromDelegate",
	ManualConfig:        "ManualConfig",
}

func (k KubernetesCredentialType) String() string {
	return string(k)
}
