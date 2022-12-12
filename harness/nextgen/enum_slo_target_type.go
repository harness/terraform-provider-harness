package nextgen

type SLOTargetType string

var SLOTargetTypes = struct {
	Rolling      SLOTargetType
	Calender     SLOTargetType
}{
	Rolling:     "Rolling",
	Calender:    "Calender",
}

var SLOTargetTypesSlice = []string{
	SLOTargetTypes.Rolling.String(),
	SLOTargetTypes.Calender.String(),
}

func (c SLOTargetType) String() string {
	return string(c)
}
