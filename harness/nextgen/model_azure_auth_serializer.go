package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AzureAuth) UnmarshalJSON(data []byte) error {
	type Alias AzureAuth

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
	case AzureAuthTypes.Certificate.String():
		err = json.Unmarshal(aux.Spec, &a.AzureClientKeyCert)
	case AzureAuthTypes.SecretKey.String():
		err = json.Unmarshal(aux.Spec, &a.AzureClientSecretKey)
	default:
		panic(fmt.Sprintf("unknown azure auth type %s", a.Type_))
	}

	return err
}

func (a *AzureAuth) MarshalJSON() ([]byte, error) {
	type Alias AzureAuth

	var spec []byte
	var err error

	switch a.Type_ {
	case AzureAuthTypes.Certificate.String():
		spec, err = json.Marshal(a.AzureClientKeyCert)
	case AzureAuthTypes.SecretKey.String():
		spec, err = json.Marshal(a.AzureClientSecretKey)
	default:
		panic(fmt.Sprintf("unknown azure auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
