package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *AwsSdkClientBackoffStrategy) UnmarshalJSON(data []byte) error {

	type Alias AwsSdkClientBackoffStrategy

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
	case AwsSdkClientBackOffStrategyTypes.EqualJitterBackoffStrategy:
		err = json.Unmarshal(aux.Spec, &a.EqualJitter)
	case AwsSdkClientBackOffStrategyTypes.FullJitterBackoffStrategy:
		err = json.Unmarshal(aux.Spec, &a.FullJitter)
	case AwsSdkClientBackOffStrategyTypes.FixedDelayBackoffStrategy:
		err = json.Unmarshal(aux.Spec, &a.FixedDelay)
	default:
		panic(fmt.Sprintf("unknown aws backoff strategy type %s", a.Type_))
	}

	return err
}

func (a *AwsSdkClientBackoffStrategy) MarshalJSON() ([]byte, error) {
	type Alias AwsSdkClientBackoffStrategy

	var spec []byte
	var err error

	switch a.Type_ {
	case AwsSdkClientBackOffStrategyTypes.EqualJitterBackoffStrategy:
		spec, err = json.Marshal(a.EqualJitter)
	case AwsSdkClientBackOffStrategyTypes.FullJitterBackoffStrategy:
		spec, err = json.Marshal(a.FullJitter)
	case AwsSdkClientBackOffStrategyTypes.FixedDelayBackoffStrategy:
		spec, err = json.Marshal(a.FixedDelay)
	default:
		panic(fmt.Sprintf("unknown aws backoff strategy type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
