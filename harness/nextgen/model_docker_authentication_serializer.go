package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *DockerAuthentication) UnmarshalJSON(data []byte) error {

	type Alias DockerAuthentication

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
	case DockerAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case DockerAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown docker auth method type %s", a.Type_))
	}

	return err
}

func (a *DockerAuthentication) MarshalJSON() ([]byte, error) {
	type Alias DockerAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case DockerAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case DockerAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown docker auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
