package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsCredential) UnmarshalJSON(data []byte) error {

	type Alias AwsCredential

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
	case AwsAuthTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	case AwsAuthTypes.InheritFromDelegate:
	case AwsAuthTypes.Irsa:
		// do nothing
	default:
		panic(fmt.Sprintf("unknown aws auth type %s", a.Type_))
	}

	return err
}

func (a *AwsCredential) MarshalJSON() ([]byte, error) {
	type Alias AwsCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsAuthTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	case AwsAuthTypes.InheritFromDelegate:
	case AwsAuthTypes.Irsa:
		// do nothing
	default:
		panic(fmt.Sprintf("unknown aws auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
