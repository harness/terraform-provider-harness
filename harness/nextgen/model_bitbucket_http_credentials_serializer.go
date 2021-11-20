package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *BitbucketHttpCredentials) UnmarshalJSON(data []byte) error {

	type Alias BitbucketHttpCredentials

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
	case BitBucketHttpCredentialTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	default:
		panic(fmt.Sprintf("unknown http credentials method type %s", a.Type_))
	}

	return err
}

func (a *BitbucketHttpCredentials) MarshalJSON() ([]byte, error) {
	type Alias BitbucketHttpCredentials

	var spec []byte
	var err error

	switch a.Type_ {
	case BitBucketHttpCredentialTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	default:
		panic(fmt.Sprintf("unknown git auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
