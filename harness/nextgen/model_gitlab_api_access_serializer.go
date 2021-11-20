package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GitlabApiAccess) UnmarshalJSON(data []byte) error {

	type Alias GitlabApiAccess

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
	case GitlabApiAuthTypes.Token:
		err = json.Unmarshal(aux.Spec, &a.Token)
	default:
		panic(fmt.Sprintf("unknown gitlab auth method type %s", a.Type_))
	}

	return err
}

func (a *GitlabApiAccess) MarshalJSON() ([]byte, error) {
	type Alias GitlabApiAccess

	var spec []byte
	var err error

	switch a.Type_ {
	case GitlabApiAuthTypes.Token:
		spec, err = json.Marshal(a.Token)
	default:
		panic(fmt.Sprintf("unknown gitlab auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
