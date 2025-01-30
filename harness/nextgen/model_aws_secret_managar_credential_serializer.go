package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsSecretManagerCredential) UnmarshalJSON(data []byte) error {

	type Alias AwsSecretManagerCredential

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
	case AwsSecretManagerAuthTypes.ManualConfig:
		err = json.Unmarshal(aux.Spec, &a.ManualConfig)
	case AwsSecretManagerAuthTypes.AssumeIAMRole:
		err = json.Unmarshal(aux.Spec, &a.AssumeIamRole)
	case AwsSecretManagerAuthTypes.AssumeSTSRole:
		err = json.Unmarshal(aux.Spec, &a.AssumeStsRole)
	case AwsSecretManagerAuthTypes.OidcAuthentication:
		err = json.Unmarshal(aux.Spec, &a.OidcConfig)
	default:
		panic(fmt.Sprintf("unknown aws auth type %s", a.Type_))
	}

	return err
}

func (a *AwsSecretManagerCredential) MarshalJSON() ([]byte, error) {
	type Alias AwsSecretManagerCredential

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsSecretManagerAuthTypes.ManualConfig:
		spec, err = json.Marshal(a.ManualConfig)
	case AwsSecretManagerAuthTypes.AssumeIAMRole:
		// spec, err = json.Marshal(a.AssumeIamRole)
		// noop
	case AwsSecretManagerAuthTypes.AssumeSTSRole:
		spec, err = json.Marshal(a.AssumeStsRole)
	case AwsSecretManagerAuthTypes.OidcAuthentication:
		spec, err = json.Marshal(a.OidcConfig)
	default:
		panic(fmt.Sprintf("unknown aws auth type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
