package nextgen

type NexusVersion string

var NexusVersions = struct {
	V2X NexusVersion
	V3X NexusVersion
}{
	V2X: "2.x",
	V3X: "3.x",
}

var NexusVersionSlice = []string{
	NexusVersions.V2X.String(),
	NexusVersions.V3X.String(),
}

func (e NexusVersion) String() string {
	return string(e)
}
