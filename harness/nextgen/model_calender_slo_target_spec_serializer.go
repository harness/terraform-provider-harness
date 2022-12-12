package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *CalenderSloTargetSpec) UnmarshalJSON(data []byte) error {

	type Alias CalenderSloTargetSpec

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case SLOCalenderTypes.Weekly:
		err = json.Unmarshal(aux.Spec, &a.Weekly)
	case SLOCalenderTypes.Monthly:
		err = json.Unmarshal(aux.Spec, &a.Monthly)
	case SLOCalenderTypes.Quarterly:
		err = json.Unmarshal(aux.Spec, &a.Quarterly)
	default:
		panic(fmt.Sprintf("unknown slo calender type %s", a.Type_))
	}

	return err
}

func (a *CalenderSloTargetSpec) MarshalJSON() ([]byte, error) {
	type Alias CalenderSloTargetSpec

	var spec []byte
	var err error

	switch a.Type_ {
	case SLOCalenderTypes.Weekly:
		spec, err = json.Marshal(a.Weekly)
	case SLOCalenderTypes.Monthly:
		spec, err = json.Marshal(a.Monthly)
	case SLOCalenderTypes.Quarterly:
		spec, err = json.Marshal(a.Quarterly)
	default:
		panic(fmt.Sprintf("unknown slo calender type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
