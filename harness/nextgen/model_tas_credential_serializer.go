package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *TasCredential) UnmarshalJSON(data []byte) error {
	type Alias TasCredential

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
	case TasCredentialTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.TasManualDetails)
	default:
		panic(fmt.Sprintf("unknown tas credential type %s", a.Type_))
	}

	return err
}

func (a *TasCredential) MarshalJSON() ([]byte, error) {
	type Alias TasCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case TasCredentialTypes.ManualConfig:
		spec, err = json.Marshal(a.TasManualDetails)
	default:
		panic(fmt.Sprintf("unknown tas credential type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
