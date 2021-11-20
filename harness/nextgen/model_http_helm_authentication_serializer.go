package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *HttpHelmAuthentication) UnmarshalJSON(data []byte) error {

	type Alias HttpHelmAuthentication

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
	case HttpHelmAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case HttpHelmAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown helm auth method type %s", a.Type_))
	}

	return err
}

func (a *HttpHelmAuthentication) MarshalJSON() ([]byte, error) {
	type Alias HttpHelmAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case HttpHelmAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case HttpHelmAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown helm auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
