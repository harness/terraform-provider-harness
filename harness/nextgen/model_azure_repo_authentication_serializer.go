package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureRepoAuthentication) UnmarshalJSON(data []byte) error {

	type Alias AzureRepoAuthentication

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
	case GitAuthTypes.Http:
		err = json.Unmarshal(aux.Spec, &a.Http)
	case GitAuthTypes.Ssh:
		err = json.Unmarshal(aux.Spec, &a.Ssh)
	default:
		panic(fmt.Sprintf("unknown azure repo auth method type %s", a.Type_))
	}

	return err
}

func (a *AzureRepoAuthentication) MarshalJSON() ([]byte, error) {
	type Alias AzureRepoAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case GitAuthTypes.Http:
		spec, err = json.Marshal(a.Http)
	case GitAuthTypes.Ssh:
		spec, err = json.Marshal(a.Ssh)
	default:
		panic(fmt.Sprintf("unknown azure repo auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
