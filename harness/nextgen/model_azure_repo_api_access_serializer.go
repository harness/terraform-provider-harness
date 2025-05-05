package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureRepoApiAccess) UnmarshalJSON(data []byte) error {

	type Alias AzureRepoApiAccess

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
	case AzureRepoApiAuthTypes.Token:
		err = json.Unmarshal(aux.Spec, &a.Token)
	default:
		panic(fmt.Sprintf("unknown azure repo auth method type %s", a.Type_))
	}

	return err
}

func (a *AzureRepoApiAccess) MarshalJSON() ([]byte, error) {
	type Alias AzureRepoApiAccess

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureRepoApiAuthTypes.Token:
		spec, err = json.Marshal(a.Token)
	default:
		panic(fmt.Sprintf("unknown azure repo auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
