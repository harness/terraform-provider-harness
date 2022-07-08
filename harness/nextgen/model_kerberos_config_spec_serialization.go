package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *KerberosConfig) UnmarshalJSON(data []byte) error {

	type Alias KerberosConfig

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.TgtGenerationMethod {
	case TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO:
		err = json.Unmarshal(aux.Spec, &a.KeyTabFilePathSpec)
	case TgtGenerationMethodTypes.TGTPasswordSpecDTO:
		err = json.Unmarshal(aux.Spec, &a.PasswordSpec)
	default:
		panic(fmt.Sprintf("unknown Tgt generation method type %s", a.TgtGenerationMethod))
	}

	return err
}

func (a *KerberosConfig) MarshalJSON() ([]byte, error) {
	type Alias KerberosConfig

	var spec []byte
	var err error

	switch a.TgtGenerationMethod {
	case TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO:
		spec, err = json.Marshal(a.KeyTabFilePathSpec)
	case TgtGenerationMethodTypes.TGTPasswordSpecDTO:
		spec, err = json.Marshal(a.PasswordSpec)
	default:
		panic(fmt.Sprintf("unknown Tgt generation method type %s", a.TgtGenerationMethod))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
