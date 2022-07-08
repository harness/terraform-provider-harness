package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *BaseSshSpec) UnmarshalJSON(data []byte) error {

	type Alias BaseSshSpec

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
	case SSHSpecificationTypes.KerberosConfigDTO:
		err = json.Unmarshal(aux.Spec, &a.KerberosConfig)
	case SSHSpecificationTypes.SSHConfig:
		err = json.Unmarshal(aux.Spec, &a.SSHConfig)
	default:
		panic(fmt.Sprintf("unknown SSH specification type %s", a.Type_))
	}

	return err
}

func (a *BaseSshSpec) MarshalJSON() ([]byte, error) {
	type Alias BaseSshSpec

	var spec []byte
	var err error

	switch a.Type_ {
	case SSHSpecificationTypes.KerberosConfigDTO:
		spec, err = json.Marshal(a.KerberosConfig)
	case SSHSpecificationTypes.SSHConfig:
		spec, err = json.Marshal(a.SSHConfig)
	default:
		panic(fmt.Sprintf("unknown SSH specification type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
