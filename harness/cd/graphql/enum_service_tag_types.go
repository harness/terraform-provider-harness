package graphql

type ServiceTagType string

var ServiceTagTypes = struct {
	Application ServiceTagType
	Service     ServiceTagType
}{
	Application: "APPLICATION",
	Service:     "SERVICE",
}

var ServiceTagTypeValues = []string{
	ServiceTagTypes.Application.String(),
	ServiceTagTypes.Service.String(),
}

func (e ServiceTagType) String() string {
	return string(e)
}
