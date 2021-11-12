package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *KubernetesAuthDto) UnmarshalJSON(data []byte) error {

	type Alias KubernetesAuthDto

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
	case KubernetesAuthMethods.ClientKeyCert.String():
		err = json.Unmarshal(aux.Spec, &a.ClientKeyCert)
	case KubernetesAuthMethods.UsernamePassword.String():
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case KubernetesAuthMethods.ServiceAccount.String():
		err = json.Unmarshal(aux.Spec, &a.ServiceAccount)
	case KubernetesAuthMethods.OpenIdConnect.String():
		err = json.Unmarshal(aux.Spec, &a.OpenIdConnect)
	default:
		panic(fmt.Sprintf("unknown kubernetes auth method type %s", a.Type_))
	}

	return err
}

func (a *KubernetesAuthDto) MarshalJSON() ([]byte, error) {
	type Alias KubernetesAuthDto

	var spec []byte
	var err error

	switch a.Type_ {
	case KubernetesAuthMethods.ClientKeyCert.String():
		spec, err = json.Marshal(a.ClientKeyCert)
	case KubernetesAuthMethods.UsernamePassword.String():
		spec, err = json.Marshal(a.UsernamePassword)
	case KubernetesAuthMethods.ServiceAccount.String():
		spec, err = json.Marshal(a.ServiceAccount)
	case KubernetesAuthMethods.OpenIdConnect.String():
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
