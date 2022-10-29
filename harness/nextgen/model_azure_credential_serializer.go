package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureCredential) UnmarshalJSON(data []byte) error {
	type Alias AzureCredential

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
	case AzureCredentialTypes.InheritFromDelegate:
		err = json.Unmarshal(aux.Spec, &a.AzureInheritFromDelegateDetails)
	case AzureCredentialTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.AzureManualDetails)
	default:
		panic(fmt.Sprintf("unknown azure credential type %s", a.Type_))
	}

	return err
}

func (a *AzureCredential) MarshalJSON() ([]byte, error) {
	type Alias AzureCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureCredentialTypes.InheritFromDelegate:
		spec, err = json.Marshal(a.AzureInheritFromDelegateDetails)
	case AzureCredentialTypes.ManualConfig:
		spec, err = json.Marshal(a.AzureManualDetails)
	default:
		panic(fmt.Sprintf("unknown azure credential type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
