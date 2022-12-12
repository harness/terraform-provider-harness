package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *SloTargetDto) UnmarshalJSON(data []byte) error {

	type Alias SloTargetDto

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
	case SLOTargetTypes.Rolling:
		err = json.Unmarshal(aux.Spec, &a.Rolling)
	case SLOTargetTypes.Calender:
		err = json.Unmarshal(aux.Spec, &a.Calender)
	default:
		panic(fmt.Sprintf("unknown slo target type %s", a.Type_))
	}

	return err
}

func (a *SloTargetDto) MarshalJSON() ([]byte, error) {
	type Alias SloTargetDto

	var spec []byte
	var err error

	switch a.Type_ {
	case SLOTargetTypes.Rolling:
		spec, err = json.Marshal(a.Rolling)
	case SLOTargetTypes.Calender:
		spec, err = json.Marshal(a.Calender)
	default:
		panic(fmt.Sprintf("unknown slo target type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
