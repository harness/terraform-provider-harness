package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *KubernetesAuth) UnmarshalJSON(data []byte) error {

	type Alias KubernetesAuth

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
	case KubernetesAuthTypes.ClientKeyCert:
		err = json.Unmarshal(aux.Spec, &a.ClientKeyCert)
	case KubernetesAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case KubernetesAuthTypes.ServiceAccount:
		err = json.Unmarshal(aux.Spec, &a.ServiceAccount)
	case KubernetesAuthTypes.OpenIdConnect:
		err = json.Unmarshal(aux.Spec, &a.OpenIdConnect)
	default:
		panic(fmt.Sprintf("unknown kubernetes auth method type %s", a.Type_))
	}

	return err
}

func (a *KubernetesAuth) MarshalJSON() ([]byte, error) {
	type Alias KubernetesAuth

	var spec []byte
	var err error

	switch a.Type_ {
	case KubernetesAuthTypes.ClientKeyCert:
		spec, err = json.Marshal(a.ClientKeyCert)
	case KubernetesAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case KubernetesAuthTypes.ServiceAccount:
		spec, err = json.Marshal(a.ServiceAccount)
	case KubernetesAuthTypes.OpenIdConnect:
		spec, err = json.Marshal(a.OpenIdConnect)
	default:
		panic(fmt.Sprintf("unknown kubernetes auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
