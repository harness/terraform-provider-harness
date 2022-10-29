package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureMsiAuth) UnmarshalJSON(data []byte) error {

	type Alias AzureMsiAuth

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
	case AzureMsiAuthTypes.SystemAssignedManagedIdentity.String():
		// err = json.Unmarshal(aux.Spec, &a.AzureMSIAuthSA)
	case AzureMsiAuthTypes.UserAssignedManagedIdentity.String():
		err = json.Unmarshal(aux.Spec, &a.AzureMSIAuthUA)
	default:
		panic(fmt.Sprintf("unknown azure inherit from delegate type %s", a.Type_))
	}

	return err
}

func (a *AzureMsiAuth) MarshalJSON() ([]byte, error) {
	type Alias AzureMsiAuth

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureMsiAuthTypes.SystemAssignedManagedIdentity.String():
		spec, err = json.Marshal(a.AzureMSIAuthSA)
	case AzureMsiAuthTypes.UserAssignedManagedIdentity.String():
		spec, err = json.Marshal(a.AzureMSIAuthUA)
	default:
		panic(fmt.Sprintf("unknown azure inherit from delegate type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
