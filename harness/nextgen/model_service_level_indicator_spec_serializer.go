package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ServiceLevelIndicatorSpec) UnmarshalJSON(data []byte) error {

	type Alias ServiceLevelIndicatorSpec

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
	case SLIMetricTypes.Threshold:
		err = json.Unmarshal(aux.Spec, &a.Threshold)
	case SLIMetricTypes.Ratio:
		err = json.Unmarshal(aux.Spec, &a.Ratio)
	default:
		panic(fmt.Sprintf("unknown sli metric type %s", a.Type_))
	}

	return err
}

func (a *ServiceLevelIndicatorSpec) MarshalJSON() ([]byte, error) {
	type Alias ServiceLevelIndicatorSpec

	var spec []byte
	var err error

	switch a.Type_ {
	case SLIMetricTypes.Threshold:
		spec, err = json.Marshal(a.Threshold)
	case SLIMetricTypes.Ratio:
		spec, err = json.Marshal(a.Ratio)
	default:
		panic(fmt.Sprintf("unknown sli metric type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
