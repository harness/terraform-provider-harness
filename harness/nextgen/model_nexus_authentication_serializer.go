package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *NexusAuthenticationDto) UnmarshalJSON(data []byte) error {

	type Alias NexusAuthenticationDto

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
	case NexusAuthTypes.UsernamePassword.String():
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case NexusAuthTypes.Anonymous.String():
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown nexus auth method type %s", a.Type_))
	}

	return err
}

func (a *NexusAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias NexusAuthenticationDto

	var spec []byte
	var err error

	switch a.Type_ {
	case NexusAuthTypes.UsernamePassword.String():
		spec, err = json.Marshal(a.UsernamePassword)
	case NexusAuthTypes.Anonymous.String():
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
