package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *NexusAuthentication) UnmarshalJSON(data []byte) error {

	type Alias NexusAuthentication

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
	case NexusAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case NexusAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown nexus auth method type %s", a.Type_))
	}

	return err
}

func (a *NexusAuthentication) MarshalJSON() ([]byte, error) {
	type Alias NexusAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case NexusAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case NexusAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown nexus auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
