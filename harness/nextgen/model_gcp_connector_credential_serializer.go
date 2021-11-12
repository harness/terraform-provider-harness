package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GcpConnectorCredentialDto) UnmarshalJSON(data []byte) error {

	type Alias GcpConnectorCredentialDto

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
	case GcpAuthTypes.ManualConfig.String():
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	case GcpAuthTypes.InheritFromDelegate.String():
		// do nothing
	default:
		panic(fmt.Sprintf("unknown gcp auth type %s", a.Type_))
	}

	return err
}

func (a *GcpConnectorCredentialDto) MarshalJSON() ([]byte, error) {
	type Alias GcpConnectorCredentialDto

	var spec []byte
	var err error

	switch a.Type_ {
	case GcpAuthTypes.ManualConfig.String():
		spec, err = json.Marshal(a.ManualConfig)
	case GcpAuthTypes.InheritFromDelegate.String():
		// do nothing
	default:
		panic(fmt.Sprintf("unknown gcp auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
