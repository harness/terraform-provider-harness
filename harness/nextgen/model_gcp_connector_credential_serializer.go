package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *GcpConnectorCredential) UnmarshalJSON(data []byte) error {

	type Alias GcpConnectorCredential

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
	case GcpAuthTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	case GcpAuthTypes.OidcAuthentication:
		err = json.Unmarshal(aux.Spec, &a.OidcConfig)
	case GcpAuthTypes.InheritFromDelegate:
		// do nothing
	default:
		panic(fmt.Sprintf("unknown gcp auth type %s", a.Type_))
	}

	return err
}

func (a *GcpConnectorCredential) MarshalJSON() ([]byte, error) {
	type Alias GcpConnectorCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case GcpAuthTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	case GcpAuthTypes.OidcAuthentication:
		spec, err = json.Marshal(&a.OidcConfig)
	case GcpAuthTypes.InheritFromDelegate:
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
