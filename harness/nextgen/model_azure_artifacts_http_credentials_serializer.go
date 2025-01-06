package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureArtifactsHttpCredentials) UnmarshalJSON(data []byte) error {

	type Alias AzureArtifactsHttpCredentials

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
	case AzureArtifactsAuthTypes.PersonalAccessToken:
		err = json.Unmarshal(aux.Spec, &a.UserToken)
	default:
		panic(fmt.Sprintf("unknow azure artifacts auth method type %s", a.Type_))
	}

	return err
}

func (a *AzureArtifactsHttpCredentials) MarshalJSON() ([]byte, error) {
	type Alias AzureArtifactsHttpCredentials

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureArtifactsAuthTypes.PersonalAccessToken:
		spec, err = json.Marshal(a.UserToken)
	default:
		panic(fmt.Sprintf("unknow azure artifacts auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
