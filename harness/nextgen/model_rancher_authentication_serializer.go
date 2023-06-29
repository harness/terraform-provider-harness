package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *RancherAuthentication) UnmarshalJSON(data []byte) error {
	type Alias RancherAuthentication
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
	case RancherAuthTypes.BearerToken:
		err = json.Unmarshal(aux.Spec, &a.BearerTokenConfig)
	default:
		panic(fmt.Sprintf("unknown rancher auth type %s", a.Type_))
	}

	return err
}

func (a *RancherAuthentication) MarshalJSON() ([]byte, error) {
	type Alias RancherAuthentication
	var spec []byte
	var err error

	switch a.Type_ {
	case RancherAuthTypes.BearerToken:
		spec, err = json.Marshal(a.BearerTokenConfig)
	default:
		panic(fmt.Sprintf("unknown rancher auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
