package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *MetricThresholdSpec) UnmarshalJSON(data []byte) error {

	type Alias HealthSource

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
	case MetricThresholdTypes.FailImmediately:
		err = json.Unmarshal(aux.Spec, &a.AppDynamics)
	case MetricThresholdTypes.IgnoreThreshold:
		err = json.Unmarshal(aux.Spec, &a.NewRelic)
	default:
		panic(fmt.Sprintf("unknown metric threshold type %s", a.Type_))
	}

	return err
}

func (a *MetricThresholdSpec) MarshalJSON() ([]byte, error) {
	type Alias MetricThreshold

	var spec []byte
	var err error

	switch a.Type_ {
	case MetricThresholdTypes.FailImmediately:
		spec, err = json.Marshal(a.FailImmediately)
	case MetricThresholdTypes.IgnoreThreshold:
		spec, err = json.Marshal(a.IgnoreThreshold)
	default:
		panic(fmt.Sprintf("unknown metric threshold type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
