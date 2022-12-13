package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *JenkinsAuthentication) UnmarshalJSON(data []byte) error {
	type Alias JenkinsAuthentication

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
	case "JenkinsBearerToken":
		err = json.Unmarshal(aux.Spec, &a.JenkinsBearerToken)
	case "JenkinsUserNamePassword":
		err = json.Unmarshal(aux.Spec, &a.JenkinsUserNamePassword)
	default:
		panic(fmt.Sprintf("unknown authentication type %s", a.Type_))
	}

	return err
}

func (a *JenkinsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias JenkinsAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case "JenkinsBearerToken":
		spec, err = json.Marshal(a.JenkinsBearerToken)
	case "JenkinsUserNamePassword":
		spec, err = json.Marshal(a.JenkinsUserNamePassword)
	default:
		panic(fmt.Sprintf("unknown authentication type %s", a.Type_))
	}
	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
