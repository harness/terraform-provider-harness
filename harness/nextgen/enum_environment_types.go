package nextgen

type EnvironmentType string

var EnvironmentTypes = struct {
	PreProduction EnvironmentType
	Production    EnvironmentType
}{
	PreProduction: "PreProduction",
	Production:    "Production",
}

var EnvironmentTypeValues = []string{
	EnvironmentTypes.PreProduction.String(),
	EnvironmentTypes.Production.String(),
}

func (e EnvironmentType) String() string {
	return string(e)
}
