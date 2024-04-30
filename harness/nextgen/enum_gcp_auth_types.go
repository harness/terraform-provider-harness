package nextgen

type GcpAuthType string

var GcpAuthTypes = struct {
	InheritFromDelegate GcpAuthType
	ManualConfig        GcpAuthType
	OidcAuthentication  GcpAuthType
}{
	InheritFromDelegate: "InheritFromDelegate",
	ManualConfig:        "ManualConfig",
	OidcAuthentication:  "OidcAuthentication",
}

func (g GcpAuthType) String() string {
	return string(g)
}
