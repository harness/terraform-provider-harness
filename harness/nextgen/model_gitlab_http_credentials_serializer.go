package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GitlabHttpCredentials) UnmarshalJSON(data []byte) error {

	type Alias GitlabHttpCredentials

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
	case GitlabHttpCredentialTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case GitlabHttpCredentialTypes.UsernameToken:
		err = json.Unmarshal(aux.Spec, &a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown gitlab http auth type %s", a.Type_))
	}

	return err
}

func (a *GitlabHttpCredentials) MarshalJSON() ([]byte, error) {
	type Alias GitlabHttpCredentials

	var spec []byte
	var err error

	switch a.Type_ {
	case GitlabHttpCredentialTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case GitlabHttpCredentialTypes.UsernameToken:
		spec, err = json.Marshal(a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown gitlab http auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
