package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ArtifactoryAuthenticationDto) UnmarshalJSON(data []byte) error {

	type Alias ArtifactoryAuthenticationDto

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
	case ArtifactoryAuthTypes.UsernamePassword.String():
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case ArtifactoryAuthTypes.Anonymous.String():
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown artifactory auth method type %s", a.Type_))
	}

	return err
}

func (a *ArtifactoryAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias ArtifactoryAuthenticationDto

	var spec []byte
	var err error

	switch a.Type_ {
	case ArtifactoryAuthTypes.UsernamePassword.String():
		spec, err = json.Marshal(a.UsernamePassword)
	case ArtifactoryAuthTypes.Anonymous.String():
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
