package nextgen

import (
	"encoding/json"
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

	err = json.Unmarshal(aux.Spec, &a.UsernamePassword)

	return err
}

func (a *JdbcAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias JdbcAuthenticationDto

	var spec []byte
	var err error

	spec, err = json.Marshal(a.UsernamePassword)

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
