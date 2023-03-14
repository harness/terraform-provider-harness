package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *TerraformCloudCredential) UnmarshalJSON(data []byte) error {

	type Alias TerraformCloudCredential

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
	case TerraformCloudAuthTypes.ApiToken:
		err = json.Unmarshal(aux.Spec, &a.ApiToken)
	default:
		panic(fmt.Sprintf("unknown terraform cloud auth type %s", a.Type_))
	}

	return err
}

func (a *TerraformCloudCredential) MarshalJSON() ([]byte, error) {
	type Alias TerraformCloudCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case TerraformCloudAuthTypes.ApiToken:
		spec, err = json.Marshal(a.ApiToken)
	default:
		panic(fmt.Sprintf("unknown terraform cloud auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
