package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ArtifactoryAuthentication) UnmarshalJSON(data []byte) error {

	type Alias ArtifactoryAuthentication

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
	case ArtifactoryAuthTypes.UsernamePassword:
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case ArtifactoryAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown artifactory auth method type %s", a.Type_))
	}

	return err
}

func (a *ArtifactoryAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ArtifactoryAuthentication

	var spec []byte
	var err error

	switch a.Type_ {
	case ArtifactoryAuthTypes.UsernamePassword:
		spec, err = json.Marshal(a.UsernamePassword)
	case ArtifactoryAuthTypes.Anonymous:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown artifactory auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
