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

	var dependencyMetadata []byte
	var err error

	switch a.Type_ {
	case DependencyMetadataTypes.KUBERNETES:
		dependencyMetadata, err = json.Marshal(a.KUBERNETES)
	default:
		panic(fmt.Sprintf("unknown dependency metadata type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.DependencyMetadata = dependencyMetadata

	return json.Marshal((*Alias)(a))
}
