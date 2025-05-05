package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureRepoHttpCredentials) UnmarshalJSON(data []byte) error {

	type Alias AzureRepoHttpCredentials

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
	case AzureRepoHttpCredentialTypes.UsernameToken:
		err = json.Unmarshal(aux.Spec, &a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown azure repo http auth type %s", a.Type_))
	}

	return err
}

func (a *AzureRepoHttpCredentials) MarshalJSON() ([]byte, error) {
	type Alias AzureRepoHttpCredentials

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureRepoHttpCredentialTypes.UsernameToken:
		spec, err = json.Marshal(a.UsernameToken)
	default:
		panic(fmt.Sprintf("unknown azure repo http auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
