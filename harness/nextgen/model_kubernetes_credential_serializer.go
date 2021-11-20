package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *KubernetesCredential) UnmarshalJSON(data []byte) error {

	type Alias KubernetesCredential

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
	case KubernetesCredentialTypes.InheritFromDelegate:
		// do nothing
	case KubernetesCredentialTypes.ManualConfig:
		err = json.Unmarshal(a.Spec, &a.ManualConfig)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	return err
}

func (a *KubernetesCredential) MarshalJSON() ([]byte, error) {
	type Alias KubernetesCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case KubernetesCredentialTypes.InheritFromDelegate:
		// do nothing
	case KubernetesCredentialTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
