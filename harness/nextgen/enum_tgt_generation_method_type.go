package nextgen

type TgtGenerationMethodType string

var TgtGenerationMethodTypes = struct {
	TGTKeyTabFilePathSpecDTO TgtGenerationMethodType
	TGTPasswordSpecDTO       TgtGenerationMethodType
}{
	TGTKeyTabFilePathSpecDTO: "KeyTabFilePath",
	TGTPasswordSpecDTO:       "Password",
}

var TgtGenerationMethodTypeValues = []string{
	TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO.String(),
	TgtGenerationMethodTypes.TGTPasswordSpecDTO.String(),
}

func (e TgtGenerationMethodType) String() string {
	return string(e)
}
