package connectors

var connectorConfigNames = []string{
	"aws",
	"docker_registry",
	"gcp",
	"k8s_cluster",
}

func getConflictsWithSlice(self string) []string {
	tmp := make([]string, len(connectorConfigNames))
	copy(tmp, connectorConfigNames)

	for i, v := range tmp {
		if v == self {
			return append(tmp[:i], tmp[i+1:]...)
		}
	}
	return tmp
}
