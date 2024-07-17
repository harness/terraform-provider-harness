package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ChangeSourceDto) UnmarshalJSON(data []byte) error {

	type Alias ChangeSourceDto

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
	case ChangeSourceTypes.HarnessCDNextGen:
		err = json.Unmarshal(aux.Spec, &a.HarnessCDNextGen)
	case ChangeSourceTypes.PagerDuty:
		err = json.Unmarshal(aux.Spec, &a.PagerDuty)
	case ChangeSourceTypes.K8sCluster:
		err = json.Unmarshal(aux.Spec, &a.K8sCluster)
	case ChangeSourceTypes.HarnessCD:
		err = json.Unmarshal(aux.Spec, &a.HarnessCD)
	case ChangeSourceTypes.CustomDeploy:
		err = json.Unmarshal(aux.Spec, &a.CustomDeploy)
	case ChangeSourceTypes.CustomIncident:
		err = json.Unmarshal(aux.Spec, &a.CustomIncident)
	case ChangeSourceTypes.CustomFF:
		err = json.Unmarshal(aux.Spec, &a.CustomFF)
	default:
		panic(fmt.Sprintf("unknown change source type %s", a.Type_))
	}

	return err
}

func (a *ChangeSourceDto) MarshalJSON() ([]byte, error) {
	type Alias ChangeSourceDto

	var spec []byte
	var err error

	switch a.Type_ {
	case ChangeSourceTypes.HarnessCDNextGen:
		spec, err = json.Marshal(a.HarnessCDNextGen)
	case ChangeSourceTypes.PagerDuty:
		spec, err = json.Marshal(a.PagerDuty)
	case ChangeSourceTypes.K8sCluster:
		spec, err = json.Marshal(a.K8sCluster)
	case ChangeSourceTypes.HarnessCD:
		spec, err = json.Marshal(a.HarnessCD)
	case ChangeSourceTypes.CustomDeploy:
		spec, err = json.Marshal(a.CustomDeploy)
	case ChangeSourceTypes.CustomIncident:
		spec, err = json.Marshal(a.CustomIncident)
	case ChangeSourceTypes.CustomFF:
		spec, err = json.Marshal(a.CustomFF)
	default:
		panic(fmt.Sprintf("unknown change source type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
