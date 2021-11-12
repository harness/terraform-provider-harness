package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsCredentialDto) UnmarshalJSON(data []byte) error {

	type Alias AwsCredentialDto

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
	case AwsAuthTypes.ManualConfig.String():
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	case AwsAuthTypes.InheritFromDelegate.String():
	case AwsAuthTypes.Irsa.String():
		// do nothing
	default:
		panic(fmt.Sprintf("unknown aws auth type %s", a.Type_))
	}

	return err
}

func (a *AwsCredentialDto) MarshalJSON() ([]byte, error) {
	type Alias AwsCredentialDto

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsAuthTypes.ManualConfig.String():
		spec, err = json.Marshal(a.ManualConfig)
	case AwsAuthTypes.InheritFromDelegate.String():
	case AwsAuthTypes.Irsa.String():
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
