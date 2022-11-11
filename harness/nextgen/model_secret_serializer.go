package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *Secret) UnmarshalJSON(data []byte) error {

	type Alias Secret

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
	case SecretTypes.SecretFile:
		err = json.Unmarshal(aux.Spec, &a.File)
	case SecretTypes.SSHKey:
		err = json.Unmarshal(aux.Spec, &a.SSHKey)
	case SecretTypes.SecretText:
		err = json.Unmarshal(aux.Spec, &a.Text)
	case SecretTypes.WinRmCredentials:
		err = json.Unmarshal(aux.Spec, &a.WinRmCredentials)
	default:
		panic(fmt.Sprintf("unknown secret type %s", a.Type_))
	}

	return err
}

func (a *Secret) MarshalJSON() ([]byte, error) {
	type Alias Secret

	var spec []byte
	var err error

	switch a.Type_ {
	case SecretTypes.SecretFile:
		spec, err = json.Marshal(a.File)
	case SecretTypes.SSHKey:
		spec, err = json.Marshal(a.SSHKey)
		// noop
	case SecretTypes.SecretText:
		spec, err = json.Marshal(a.Text)
	case SecretTypes.WinRmCredentials:
		spec, err = json.Marshal(a.WinRmCredentials)
	default:
		panic(fmt.Sprintf("unknown secret type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
