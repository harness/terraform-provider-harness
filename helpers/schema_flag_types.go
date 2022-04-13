package helpers

type SchemaFlagType string

var SchemaFlagTypes = struct {
	Required SchemaFlagType
	Optional SchemaFlagType
	Computed SchemaFlagType
}{
	Required: "Required",
	Optional: "Optional",
	Computed: "Computed",
}

var SchemaFlagTypeValues = []string{
	SchemaFlagTypes.Required.String(),
	SchemaFlagTypes.Optional.String(),
	SchemaFlagTypes.Computed.String(),
}

func (e SchemaFlagType) String() string {
	return string(e)
}
