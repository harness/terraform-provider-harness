package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ServiceNowAuthentication) UnmarshalJSON(data []byte) error {

	type Alias ServiceNowAuthentication

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
	case ServiceNowAuthTypes.ServiceNowUserNamePassword:
		err = json.Unmarshal(aux.Spec, &a.ServiceNowUserNamePassword)
	case ServiceNowAuthTypes.ServiceNowAdfs:
		err = json.Unmarshal(aux.Spec, &a.ServiceNowAdfs)
	case ServiceNowAuthTypes.ServiceNowRefreshToken:
		err = json.Unmarshal(aux.Spec, &a.ServiceNowRefreshToken)		
	default:
		panic(fmt.Sprintf("unknown service now auth method type %s", a.Type_))
	}

	return err
}

func (a *ServiceNowAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ServiceNowAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case ServiceNowAuthTypes.ServiceNowUserNamePassword:
		spec, err = json.Marshal(a.ServiceNowUserNamePassword)
	case ServiceNowAuthTypes.ServiceNowAdfs:
		spec, err = json.Marshal(a.ServiceNowAdfs)
	case ServiceNowAuthTypes.ServiceNowRefreshToken:
		spec, err = json.Marshal(a.ServiceNowRefreshToken)		
	default:
		panic(fmt.Sprintf("unknown service now auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
