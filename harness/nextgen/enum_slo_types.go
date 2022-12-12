package nextgen

type SLOType string

var SLOTypes = struct {
	Simple        SLOType
	Composite     SLOType
}{
	Simple:       "Simple",
	Composite:    "Composite",
}

var SLOTypesSlice = []string{
	SLOTypes.Simple.String(),
	SLOTypes.Composite.String(),
}

func (c SLOType) String() string {
	return string(c)
}
