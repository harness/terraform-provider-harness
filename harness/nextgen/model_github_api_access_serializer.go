package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GithubApiAccess) UnmarshalJSON(data []byte) error {

	type Alias GithubApiAccess

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
	case GithubApiAccessTypes.GithubApp:
		err = json.Unmarshal(aux.Spec, &a.GithubApp)
	case GithubApiAccessTypes.Token:
		err = json.Unmarshal(aux.Spec, &a.Token)
	default:
		panic(fmt.Sprintf("unknown github auth method type %s", a.Type_))
	}

	return err
}

func (a *GithubApiAccess) MarshalJSON() ([]byte, error) {
	type Alias GithubApiAccess

	var spec []byte
	var err error

	switch a.Type_ {
	case GithubApiAccessTypes.GithubApp:
		spec, err = json.Marshal(a.GithubApp)
	case GithubApiAccessTypes.Token:
		spec, err = json.Marshal(a.Token)
	default:
		panic(fmt.Sprintf("unknown github auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
