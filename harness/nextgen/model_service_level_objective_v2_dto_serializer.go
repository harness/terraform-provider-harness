package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ServiceLevelObjectiveV2Dto) UnmarshalJSON(data []byte) error {

	type Alias ServiceLevelObjectiveV2Dto

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
	case SLOTypes.Simple:
		err = json.Unmarshal(aux.Spec, &a.Simple)
	case SLOTypes.Composite:
		err = json.Unmarshal(aux.Spec, &a.Composite)
	default:
		panic(fmt.Sprintf("unknown slo type %s", a.Type_))
	}

	return err
}

func (a *ServiceLevelObjectiveV2Dto) MarshalJSON() ([]byte, error) {
	type Alias ServiceLevelObjectiveV2Dto

	var spec []byte
	var err error

	switch a.Type_ {
	case SLOTypes.Simple:
		spec, err = json.Marshal(a.Simple)
	case SLOTypes.Composite:
		spec, err = json.Marshal(a.Composite)
	default:
		panic(fmt.Sprintf("unknown slo type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
