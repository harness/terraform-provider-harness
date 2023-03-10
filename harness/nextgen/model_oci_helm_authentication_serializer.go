package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *OciHelmAuthentication) UnmarshalJSON(data []byte) error {

	type Alias OciHelmAuthentication

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
	case OciHelmAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case OciHelmAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown helm auth method type %s", a.Type_))
	}

	return err
}

func (a *OciHelmAuthentication) MarshalJSON() ([]byte, error) {
	type Alias OciHelmAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case OciHelmAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case OciHelmAuthTypes.Anonymous:
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
