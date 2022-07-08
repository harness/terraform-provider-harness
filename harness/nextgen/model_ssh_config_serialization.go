package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *SshConfig) UnmarshalJSON(data []byte) error {

	type Alias SshConfig

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.CredentialType {
	case SSHConfigTypes.KeyPath:
		err = json.Unmarshal(aux.Spec, &a.KeyPathCredential)
	case SSHConfigTypes.KeyReference:
		err = json.Unmarshal(aux.Spec, &a.KeyReferenceCredential)
	case SSHConfigTypes.Password:
		err = json.Unmarshal(aux.Spec, &a.PasswordCredential)
	default:
		panic(fmt.Sprintf("unknown SSH config type %s", a.Type_))
	}

	return err
}

func (a *SshConfig) MarshalJSON() ([]byte, error) {
	type Alias SshConfig

	var spec []byte
	var err error

	switch a.CredentialType {
	case SSHConfigTypes.KeyPath:
		spec, err = json.Marshal(a.KeyPathCredential)
	case SSHConfigTypes.KeyReference:
		spec, err = json.Marshal(a.KeyReferenceCredential)
	case SSHConfigTypes.Password:
		spec, err = json.Marshal(a.PasswordCredential)
	default:
		panic(fmt.Sprintf("unknown SSH config type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
