package cd

import (
	"encoding/json"
)

type APIVersionResponse struct {
	Metadata map[string]string `json:"metaData"`

	Resource struct {
		VersionInfo struct {
			Version     string `json:"version"`
			BuildNumber string `json:"buildNo"`
			GitCommit   string `json:"gitCommit"`
			GitBranch   string `json:"gitBranch"`
			Timestamp   string `json:"timestamp"`
		} `json:"versionInfo"`
		RuntimeInfo struct {
			Primary        bool   `json:"primary"`
			PrimaryVersion string `json:"primaryVersion"`
			DeployMode     string `json:"deployMode"`
		} `json:"runtimeInfo"`
	} `json:"resource"`

	ResponseMessages []string
}

// Returns the version of Harness that we're connecting to
func (client *ApiClient) GetAPIVersion() (*APIVersionResponse, error) {
	req, err := client.NewAuthorizedGetRequest("/api/version")

	if err != nil {
		return nil, err
	}

	resp, err := client.Configuration.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	apiVersion := &APIVersionResponse{}
	err = json.NewDecoder(resp.Body).Decode(apiVersion)

	if err != nil {
		return nil, err
	}

	return apiVersion, nil
}
