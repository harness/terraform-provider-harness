package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsKmsConnectorCredentialDto) UnmarshalJSON(data []byte) error {

	type Alias AwsKmsConnectorCredentialDto

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
	case AwsKmsAuthTypes.AssumeIAMRole:
		err = json.Unmarshal(aux.Spec, &a.AssumeIamRole)
	case AwsKmsAuthTypes.AssumeSTSRole:
		err = json.Unmarshal(aux.Spec, &a.AssumeStsRole)
	case AwsKmsAuthTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	default:
		panic(fmt.Sprintf("unknown aws kms auth type %s", a.Type_))
	}

	return err
}

func (a *AwsKmsConnectorCredentialDto) MarshalJSON() ([]byte, error) {
	type Alias AwsKmsConnectorCredentialDto

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsKmsAuthTypes.AssumeIAMRole:
		spec, err = json.Marshal(a.AssumeIamRole)
	case AwsKmsAuthTypes.AssumeSTSRole:
		spec, err = json.Marshal(a.AssumeStsRole)
	case AwsKmsAuthTypes.ManualConfig:
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
