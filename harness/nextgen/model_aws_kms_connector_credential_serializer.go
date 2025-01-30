package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsKmsConnectorCredential) UnmarshalJSON(data []byte) error {

	type Alias AwsKmsConnectorCredential

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
	case AwsKmsAuthTypes.OidcAuthentication:
		err = json.Unmarshal(aux.Spec, &a.OidcConfig)
	default:
		panic(fmt.Sprintf("unknown aws kms auth type %s", a.Type_))
	}

	return err
}

func (a *AwsKmsConnectorCredential) MarshalJSON() ([]byte, error) {
	type Alias AwsKmsConnectorCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsKmsAuthTypes.AssumeIAMRole:
		spec, err = json.Marshal(a.AssumeIamRole)
	case AwsKmsAuthTypes.AssumeSTSRole:
		spec, err = json.Marshal(a.AssumeStsRole)
	case AwsKmsAuthTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	case AwsKmsAuthTypes.OidcAuthentication:
		spec, err = json.Marshal(a.OidcConfig)
	default:
		panic(fmt.Sprintf("unknown aws kms auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
