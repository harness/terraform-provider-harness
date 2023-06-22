package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *JiraAuthentication) UnmarshalJSON(data []byte) error {

	type Alias JiraAuthentication

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
	case JiraAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case JiraAuthTypes.PersonalAccessToken:
		err = json.Unmarshal(aux.Spec, &a.PersonalAccessToken)		
	default:
		panic(fmt.Sprintf("unknown jira auth method type %s", a.Type_))
	}

	return err
}

func (a *JiraAuthentication) MarshalJSON() ([]byte, error) {
	type Alias JiraAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case JiraAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case JiraAuthTypes.PersonalAccessToken:
		spec, err = json.Marshal(a.PersonalAccessToken)		
	default:
		panic(fmt.Sprintf("unknown jira auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
