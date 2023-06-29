package nextgen

type RancherConfigType string

var RancherConfigTypes = struct {
	ManualConfig RancherConfigType
}{
	ManualConfig: "ManualConfig",
}

func (e RancherConfigType) String() string {
	return string(e)
}
