package api

import (
	"fmt"
)

// EnvironmentsService handles communication with the environments related
// methods of the Split.io APIv2.
//
// Reference: https://docs.split.io/reference#environments-overview
type EnvironmentsService service

// Environment reflects a stage in the development process, such as your production application or your internal staging
// environment. During the feature release process, Splits can be promoted through the various environments; allowing for
// a targeted roll out throughout the development process.
type Environment struct {
	ID         *string `json:"id"`
	Name       *string `json:"name"`
	Production *bool   `json:"production"`
}

// EnvironmentSegment represents a segment in an environment.
type EnvironmentSegment struct {
	ID             *string      `json:"id"`
	OrgID          *string      `json:"orgId"`
	Environment    *string      `json:"environment"`
	Name           *string      `json:"name"`
	TrafficTypeID  *string      `json:"trafficTypeId"`
	Description    *string      `json:"description"`
	Status         *string      `json:"status"`
	CreationTime   *int64       `json:"creationTime"`
	LastUpdateTime *int64       `json:"lastUpdateTime"`
	TrafficTypeURN *TrafficType `json:"trafficTypeURN"`
	Creator        *User        `json:"creator"`
}

// EnvironmentRequest represents a request modify an environment.
type EnvironmentRequest struct {
	Name       *string `json:"name,omitempty"`
	Production *bool   `json:"production,omitempty"`
}

// EnvironmentSegmentKeysRequest represents a request to add/remove segment keys in an environment.
type EnvironmentSegmentKeysRequest struct {
	Keys    []string `json:"keys"`
	Comment string   `json:"comment,omitempty"`
	Title   string   `json:"title,omitempty"`
}


// Get fetches an environment by ID by searching through the list
func (e *EnvironmentsService) Get(workspaceID, environmentID string) (*Environment, error) {
	envs, err := e.List(workspaceID)
	if err != nil {
		return nil, err
	}
	for _, env := range envs {
		if env.ID != nil && *env.ID == environmentID {
			return env, nil
		}
	}
	return nil, fmt.Errorf("environment with ID %s not found", environmentID)
}

// List all environments in a workspace.
func (e *EnvironmentsService) List(workspaceID string) ([]*Environment, error) {
	var result []*Environment
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/environments/ws/%s", workspaceID)
	err := e.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new environment.
func (e *EnvironmentsService) Create(workspaceID string, opts *EnvironmentRequest) (*Environment, error) {
	var result Environment
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/environments/ws/%s", workspaceID)
	err := e.client.post(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing environment.
func (e *EnvironmentsService) Update(workspaceID, environmentID string, opts *EnvironmentRequest) (*Environment, error) {
	var result Environment
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/environments/ws/%s/%s", workspaceID, environmentID)
	err := e.client.patch(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes an environment.
func (e *EnvironmentsService) Delete(workspaceID, environmentID string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/environments/ws/%s/%s", workspaceID, environmentID)
	err := e.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// SegmentKeysResponse represents the API response structure for segment keys
type SegmentKeysResponse struct {
	Keys interface{} `json:"keys"`
}

// SegmentKeysListResponse represents a possible alternative response structure
type SegmentKeysListResponse struct {
	Objects []SegmentKeyObject `json:"objects"`
}

type SegmentKeyObject struct {
	Key string `json:"key"`
}

// GetSegmentKeys retrieves segment keys for a specific environment and segment
func (e *EnvironmentsService) GetSegmentKeys(environmentID, segmentName string) ([]string, error) {
	var result SegmentKeysResponse
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/%s/%s/keys", environmentID, segmentName)
	err := e.client.get(url, &result)
	if err != nil {
		return nil, err
	}

	// Convert the interface{} keys to []string
	return convertKeysToStringSlice(result.Keys), nil
}

// convertKeysToStringSlice converts various possible key formats to []string
func convertKeysToStringSlice(keys interface{}) []string {
	if keys == nil {
		return []string{}
	}

	// If it's already a []string
	if stringSlice, ok := keys.([]string); ok {
		return stringSlice
	}

	// If it's a []interface{} (common in JSON unmarshaling)
	if interfaceSlice, ok := keys.([]interface{}); ok {
		result := make([]string, 0, len(interfaceSlice))
		for _, v := range interfaceSlice {
			// Check if it's a string directly
			if str, ok := v.(string); ok {
				result = append(result, str)
			} else if keyObj, ok := v.(map[string]interface{}); ok {
				// Check if it's an object with "key" field like {"key": "user123"}
				if keyValue, exists := keyObj["key"]; exists {
					if str, ok := keyValue.(string); ok {
						result = append(result, str)
					}
				}
			}
		}
		return result
	}

	// If it's some other format, return empty slice and log
	fmt.Printf("DEBUG: Unexpected keys format: %T = %+v\n", keys, keys)
	return []string{}
}

// AddSegmentKeys adds new keys to a segment in an environment
func (e *EnvironmentsService) AddSegmentKeys(environmentID, segmentName string, replace bool, keys []string) ([]string, error) {
	// Use the proper struct matching Split.io provider
	request := &EnvironmentSegmentKeysRequest{
		Keys:    keys,
		Comment: "modified by Terraform",
		Title:   "modified by Terraform",
	}

	var result SegmentKeysResponse
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/%s/%s/uploadKeys?replace=%v", environmentID, segmentName, replace)
	err := e.client.put(url, request, &result)
	if err != nil {
		return nil, err
	}
	return convertKeysToStringSlice(result.Keys), nil
}

// RemoveSegmentKeys removes keys from a segment in an environment
func (e *EnvironmentsService) RemoveSegmentKeys(environmentID, segmentName string, keys []string) error {
	// Use the proper struct matching Split.io provider
	request := &EnvironmentSegmentKeysRequest{
		Keys:    keys,
		Comment: "modified by Terraform",
		Title:   "modified by Terraform",
	}

	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/%s/%s/removeKeys", environmentID, segmentName)
	err := e.client.put(url, request, nil)
	if err != nil {
		return err
	}
	return nil
}