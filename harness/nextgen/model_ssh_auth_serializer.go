package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *SshAuth) UnmarshalJSON(data []byte) error {

	type Alias SshAuth

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
	case SSHAuthenticationTypes.Kerberos:
		err = json.Unmarshal(aux.Spec, &a.KerberosConfig)
	case SSHAuthenticationTypes.SSH:
		err = json.Unmarshal(aux.Spec, &a.SSHConfig)
	default:
		panic(fmt.Sprintf("unknown SSH authentication type %s", a.Type_))
	}

	return err
}

func (a *SshAuth) MarshalJSON() ([]byte, error) {
	type Alias SshAuth

	var spec []byte
	var err error

	switch a.Type_ {
	case SSHAuthenticationTypes.Kerberos:
		spec, err = json.Marshal(a.KerberosConfig)
	case SSHAuthenticationTypes.SSH:
		spec, err = json.Marshal(a.SSHConfig)
	default:
		panic(fmt.Sprintf("unknown SSH authentication type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
