package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *SpotCredential) UnmarshalJSON(data []byte) error {

	type Alias SpotCredential

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
	case SpotAuthTypes.PermanentTokenConfig:
		err = json.Unmarshal(aux.Spec, &a.PermanentTokenConfig)
	default:
		panic(fmt.Sprintf("unknown spot auth type %s", a.Type_))
	}

	return err
}

func (a *SpotCredential) MarshalJSON() ([]byte, error) {
	type Alias SpotCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case SpotAuthTypes.PermanentTokenConfig:
		spec, err = json.Marshal(a.PermanentTokenConfig)
	default:
		panic(fmt.Sprintf("unknown spot auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
