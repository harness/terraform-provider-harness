package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GithubHttpCredentials) UnmarshalJSON(data []byte) error {

	type Alias GithubHttpCredentials

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
	case GithubHttpCredentialTypes.UsernameToken:
		err = json.Unmarshal(aux.Spec, &a.UsernameToken)
	case GithubHttpCredentialTypes.GithubApp:
		err = json.Unmarshal(aux.Spec, &a.GithubApp)
    case GithubHttpCredentialTypes.Anonymous:
    	// nothing to do
	default:
		panic(fmt.Sprintf("unknown http credentials method type %s", a.Type_))
	}

	return err
}

func (a *GithubHttpCredentials) MarshalJSON() ([]byte, error) {
	type Alias GithubHttpCredentials

	var spec []byte
	var err error

	switch a.Type_ {
	case GithubHttpCredentialTypes.UsernameToken:
		spec, err = json.Marshal(a.UsernameToken)
	case GithubHttpCredentialTypes.GithubApp:
		spec, err = json.Marshal(a.GithubApp)
    case GithubHttpCredentialTypes.Anonymous:
    	// nothing to do
	default:
		panic(fmt.Sprintf("unknown git auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
