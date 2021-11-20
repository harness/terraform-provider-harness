package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsSecretManagerCredentialDto) UnmarshalJSON(data []byte) error {

	type Alias AwsSecretManagerCredentialDto

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
	case AwsSecretManagerAuthTypes.AssumeIAMRole:
		// noop
	case AwsSecretManagerAuthTypes.AssumeSTSRole:
		err = json.Unmarshal(aux.Spec, &a.AssumeStsRole)
	case AwsSecretManagerAuthTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	default:
		panic(fmt.Sprintf("unknown aws kms auth type %s", a.Type_))
	}

	return err
}

func (a *AwsSecretManagerCredentialDto) MarshalJSON() ([]byte, error) {
	type Alias AwsSecretManagerCredentialDto

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsSecretManagerAuthTypes.AssumeIAMRole:
		// noop
	case AwsSecretManagerAuthTypes.AssumeSTSRole:
		spec, err = json.Marshal(a.AssumeStsRole)
	case AwsSecretManagerAuthTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	default:
		panic(fmt.Sprintf("unknown aws kms auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
