package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *Provider) UnmarshalJSON(data []byte) error {

	type Alias Provider

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
	case ProviderTypes.BitbucketServer:
		err = json.Unmarshal(aux.Spec, &a.BitbucketServerSpec)
	default:
		panic(fmt.Sprintf("unknown provider type %s", a.Type_))
	}

	return err
}

func (a *Provider) MarshalJSON() ([]byte, error) {
	type Alias Provider

	var spec []byte
	var err error

	switch a.Type_ {
	case ProviderTypes.BitbucketServer:
		spec, err = json.Marshal(a.BitbucketServerSpec)
	default:
		panic(fmt.Sprintf("unknown secret type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
