package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GitConfigDto) UnmarshalJSON(data []byte) error {

	type Alias GitConfigDto

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
	case GitAuthTypes.Http.String():
		err = json.Unmarshal(aux.Spec, &a.Http)
	case GitAuthTypes.Ssh.String():
		err = json.Unmarshal(aux.Spec, &a.Ssh)
	default:
		panic(fmt.Sprintf("unknown git auth method type %s", a.Type_))
	}

	return err
}

func (a *GitConfigDto) MarshalJSON() ([]byte, error) {
	type Alias GitConfigDto

	var spec []byte
	var err error

	switch a.Type_ {
	case GitAuthTypes.Http.String():
		spec, err = json.Marshal(a.Http)
	case GitAuthTypes.Ssh.String():
		spec, err = json.Marshal(a.Ssh)
	default:
		panic(fmt.Sprintf("unknown git auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
