package graphql

import (
	"encoding/json"

	"github.com/harness/harness-go-sdk/harness/time"
)

type Delegate struct {
	AccountId          string       `json:"accountId,omitempty"`
	DelegateName       string       `json:"delegateName,omitempty"`
	DelegateProfileId  string       `json:"delegateProfileId,omitempty"`
	DelegateType       DelegateType `json:"delegateType,omitempty"`
	Description        string       `json:"description,omitempty"`
	HostName           string       `json:"hostName,omitempty"`
	Ip                 string       `json:"ip,omitempty"`
	LastHeartBeat      *time.Time   `json:"lastHeartBeat,omitempty"`
	PollingModeEnabled bool         `json:"pollingModeEnabled,omitempty"`
	Status             string       `json:"status,omitempty"`
	UUID               string       `json:"uuid,omitempty"`
	Version            string       `json:"version,omitempty"`
}

func (a *Delegate) UnmarshalJSON(data []byte) error {

	type Alias Delegate

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	return err
}

func (a *Delegate) MarshalJSON() ([]byte, error) {
	type Alias Delegate

	return json.Marshal((*Alias)(a))
}
