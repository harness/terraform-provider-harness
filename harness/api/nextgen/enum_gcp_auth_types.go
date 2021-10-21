package nextgen

type GcpAuthType string

var GcpAuthTypes = struct {
	InheritFromDelegate GcpAuthType
	ManualConfig        GcpAuthType
}{
	InheritFromDelegate: "InheritFromDelegate",
	ManualConfig:        "ManualConfig",
}

func (g GcpAuthType) String() string {
	return string(g)
}
