package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *BitbucketApiAccess) UnmarshalJSON(data []byte) error {

	type Alias BitbucketApiAccess

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
	case BitBucketApiAccessTypes.UsernameToken:
		err = json.Unmarshal(aux.Spec, &a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown BitBucket api auth method type %s", a.Type_))
	}

	return err
}

func (a *BitbucketApiAccess) MarshalJSON() ([]byte, error) {
	type Alias BitbucketApiAccess

	var spec []byte
	var err error

	switch a.Type_ {
	case BitBucketApiAccessTypes.UsernameToken:
		spec, err = json.Marshal(a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown BitBucket api auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
