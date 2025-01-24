package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *JdbcAuthenticationDto) UnmarshalJSON(data []byte) error {

	type Alias JdbcAuthenticationDto

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
	case JDBCAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case JDBCAuthTypes.ServiceAccount:
		err = json.Unmarshal(aux.Spec, &a.ServiceAccount)
	default:
		panic(fmt.Sprintf("unknown jdbc auth method type %s", a.Type_))
	}

	return err
}

func (a *JdbcAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias JdbcAuthenticationDto

	var spec []byte
	var err error

	switch a.Type_ {
	case JDBCAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case JDBCAuthTypes.ServiceAccount:
		spec, err = json.Marshal(a.ServiceAccount)
	default:
		panic(fmt.Sprintf("unknown jdbc auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
