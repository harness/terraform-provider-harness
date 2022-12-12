package nextgen

type DependencyMetadataType string

var DependencyMetadataTypes = struct {
	KUBERNETES    DependencyMetadataType
}{
	KUBERNETES:   "KUBERNETES",
}

var DependencyMetadataTypesSlice = []string{
	DependencyMetadataTypes.KUBERNETES.String(),
}

func (c DependencyMetadataType) String() string {
	return string(c)
}
