package graphql

type Delegate struct {
	AccountId          string       `json:"accountId,omitempty"`
	DelegateName       string       `json:"delegateName,omitempty"`
	DelegateProfileId  string       `json:"delegateProfileId,omitempty"`
	DelegateType       DelegateType `json:"delegateType,omitempty"`
	Description        string       `json:"description,omitempty"`
	HostName           string       `json:"hostName,omitempty"`
	Ip                 string       `json:"ip,omitempty"`
	LastHeartBeat      string       `json:"lastHeartBeat,omitempty"`
	PollingModeEnabled bool         `json:"pollingModeEnabled,omitempty"`
	Status             string       `json:"status,omitempty"`
	UUID               string       `json:"uuid,omitempty"`
	Version            string       `json:"version,omitempty"`
}
