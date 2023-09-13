package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *CvngNotificationChannel) UnmarshalJSON(data []byte) error {

	type Alias CvngNotificationChannel

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
	case CVNGNotificationChannelTypes.Email:
		err = json.Unmarshal(aux.Spec, &a.Email)
	case CVNGNotificationChannelTypes.Slack:
		err = json.Unmarshal(aux.Spec, &a.Slack)
	case CVNGNotificationChannelTypes.PagerDuty:
		err = json.Unmarshal(aux.Spec, &a.PagerDuty)
	case CVNGNotificationChannelTypes.MsTeams:
		err = json.Unmarshal(aux.Spec, &a.MsTeams)
	default:
		panic(fmt.Sprintf("unknown notification channel type type %s", a.Type_))
	}

	return err
}

func (a *CvngNotificationChannel) MarshalJSON() ([]byte, error) {
	type Alias CvngNotificationChannel

	var spec []byte
	var err error

	switch a.Type_ {
	case CVNGNotificationChannelTypes.Email:
		spec, err = json.Marshal(a.Email)
	case CVNGNotificationChannelTypes.Slack:
		spec, err = json.Marshal(a.Slack)
	case CVNGNotificationChannelTypes.PagerDuty:
		spec, err = json.Marshal(a.PagerDuty)
	case CVNGNotificationChannelTypes.MsTeams:
		spec, err = json.Marshal(a.MsTeams)
	default:
		panic(fmt.Sprintf("unknown notification channel type type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
