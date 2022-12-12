package nextgen

type SLOCalenderType string

var SLOCalenderTypes = struct {
	Weekly      SLOCalenderType
	Monthly     SLOCalenderType
	Quarterly   SLOCalenderType
}{
	Weekly:     "Weekly",
	Monthly:    "Monthly",
	Quarterly:  "Quarterly",
}

var SLOCalenderTypesSlice = []string{
	SLOCalenderTypes.Weekly.String(),
	SLOCalenderTypes.Monthly.String(),
	SLOCalenderTypes.Quarterly.String(),
}

func (c SLOCalenderType) String() string {
	return string(c)
}
