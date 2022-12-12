package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ServiceDependencyDto) UnmarshalJSON(data []byte) error {

	type Alias ServiceDependencyDto

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
	case DependencyMetadataTypes.KUBERNETES:
		err = json.Unmarshal(aux.DependencyMetadata, &a.KUBERNETES)
	default:
		panic(fmt.Sprintf("unknown dependency metadata type %s", a.Type_))
	}

	return err
}

func (a *ServiceDependencyDto) MarshalJSON() ([]byte, error) {
	type Alias ServiceDependencyDto

	var spec []byte
	var err error

	switch a.Type_ {
	case DependencyMetadataTypes.KUBERNETES:
		spec, err = json.Marshal(a.KUBERNETES)
	default:
		panic(fmt.Sprintf("unknown dependency metadata type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.DependencyMetadata = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
