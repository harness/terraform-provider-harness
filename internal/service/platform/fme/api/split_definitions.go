package api

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// SplitDefinitionsService handles communication with the split definitions related
// methods of the Split.io APIv2.
type SplitDefinitionsService service

// SplitDefinition represents a split definition in Split.io
type SplitDefinition struct {
	Name              *string                       `json:"name"`
	Environment       *SplitDefinitionEnvironment   `json:"environment"`
	Treatments        []SplitDefinitionTreatment    `json:"treatments"`
	Rules             []SplitDefinitionRule         `json:"rules"`
	DefaultRule       []SplitDefinitionDefaultRule  `json:"defaultRule"` // Array in response too
	BaselineTreatment *string                       `json:"baselineTreatment"`
	DefaultTreatment  *string                       `json:"defaultTreatment"`
	TrafficAllocation *int                          `json:"trafficAllocation"`
	LastUpdateTime    *int64                        `json:"lastUpdateTime"`
	CreationTime      *int64                        `json:"creationTime"`
}

// SplitDefinitionEnvironment represents the environment info in a split definition
type SplitDefinitionEnvironment struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

// SplitDefinitionTreatment represents a treatment in a split definition
type SplitDefinitionTreatment struct {
	Name           *string   `json:"name"`
	Description    *string   `json:"description"`
	Configurations *string   `json:"configurations"` // JSON string, not array
	Keys           []string  `json:"keys,omitempty"`
	Segments       []string  `json:"segments,omitempty"`
}

// SplitDefinitionConfiguration represents a configuration in a treatment
type SplitDefinitionConfiguration struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// SplitDefinitionRule represents a rule in a split definition
type SplitDefinitionRule struct {
	Buckets   []SplitDefinitionBucket      `json:"buckets"`
	Condition *SplitDefinitionCondition    `json:"condition"`
}

// SplitDefinitionBucket represents a bucket in a rule
type SplitDefinitionBucket struct {
	Treatment *string `json:"treatment"`
	Size      *int    `json:"size"`
}

// SplitDefinitionCondition represents a condition in a rule
type SplitDefinitionCondition struct {
	Combiner *string                      `json:"combiner"`
	Matchers []SplitDefinitionMatcher     `json:"matchers"`
}

// SplitDefinitionMatcher represents a matcher in a condition
type SplitDefinitionMatcher struct {
	Type      *string  `json:"type"`
	Attribute *string  `json:"attribute"`
	Strings   []string `json:"strings,omitempty"`
	Negate    *bool    `json:"negate,omitempty"`
}

// SplitDefinitionWhitelistMatcherData represents whitelist matcher data
type SplitDefinitionWhitelistMatcherData struct {
	Whitelist []string `json:"whitelist"`
}

// SplitDefinitionDefaultRule represents the default rule in a split definition
type SplitDefinitionDefaultRule struct {
	Treatment   *string `json:"treatment"`
	Size        *int    `json:"size"`
}

// SplitDefinitionCreateRequest represents a request to create a split definition
type SplitDefinitionCreateRequest struct {
	Treatments        []SplitDefinitionTreatment    `json:"treatments"`
	Rules             []SplitDefinitionRule         `json:"rules"`
	DefaultRule       []SplitDefinitionDefaultRule  `json:"defaultRule"` // Array, not pointer
	BaselineTreatment *string                       `json:"baselineTreatment,omitempty"`
	DefaultTreatment  *string                       `json:"defaultTreatment"`
	TrafficAllocation *int                          `json:"trafficAllocation,omitempty"`
	Title             *string                       `json:"title,omitempty"`
	Comment           *string                       `json:"comment,omitempty"`
}

// SplitDefinitionUpdateRequest represents a request to update a split definition
type SplitDefinitionUpdateRequest struct {
	Treatments        []SplitDefinitionTreatment    `json:"treatments"`
	Rules             []SplitDefinitionRule         `json:"rules"`
	DefaultRule       []SplitDefinitionDefaultRule  `json:"defaultRule"` // Array, not pointer
	BaselineTreatment *string                       `json:"baselineTreatment,omitempty"`
	DefaultTreatment  *string                       `json:"defaultTreatment"`
	TrafficAllocation *int                          `json:"trafficAllocation,omitempty"`
	Title             *string                       `json:"title,omitempty"`
	Comment           *string                       `json:"comment,omitempty"`
}

// SplitDefinitionListResult represents the response from the split definitions list endpoint with pagination
type SplitDefinitionListResult struct {
	GenericListResult
	Objects []*SplitDefinition `json:"objects"`
}

// Get fetches a split definition by split name and environment ID
func (s *SplitDefinitionsService) Get(workspaceID, environmentID, splitName string) (*SplitDefinition, error) {
	var result SplitDefinition
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s/environments/%s", workspaceID, splitName, environmentID)
	err := s.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new split definition
func (s *SplitDefinitionsService) Create(workspaceID, environmentID, splitName string, opts *SplitDefinitionCreateRequest) (*SplitDefinition, error) {
	var result SplitDefinition
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s/environments/%s", workspaceID, splitName, environmentID)

	// Debug logging - print the request JSON
	if jsonBytes, err := json.Marshal(opts); err == nil {
		log.Printf("[DEBUG] Split definition create request JSON: %s", string(jsonBytes))
		fmt.Printf("DEBUG: Split definition create request JSON: %s\n", string(jsonBytes))
	}

	// Try POST first to create/activate the split in environment (normal case for new splits)
	err := s.client.post(url, opts, &result)
	if err != nil {
		// If POST fails with 409 (split already exists in environment), try PUT to update it
		if strings.Contains(err.Error(), "409") {
			log.Printf("[DEBUG] POST failed with 409 (split already in environment), trying PUT to update")
			fmt.Printf("DEBUG: POST failed with 409, trying PUT to update split definition\n")

			err = s.client.put(url, opts, &result)
			if err != nil {
				return nil, err
			}
			log.Printf("[DEBUG] PUT succeeded")
			fmt.Printf("DEBUG: PUT succeeded\n")
		} else {
			return nil, err
		}
	}
	return &result, nil
}

// Update modifies an existing split definition
func (s *SplitDefinitionsService) Update(workspaceID, environmentID, splitName string, opts *SplitDefinitionUpdateRequest) (*SplitDefinition, error) {
	var result SplitDefinition
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s/environments/%s", workspaceID, splitName, environmentID)
	err := s.client.put(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a split definition (resets to default state)
func (s *SplitDefinitionsService) Delete(workspaceID, environmentID, splitName string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s/environments/%s", workspaceID, splitName, environmentID)
	err := s.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List returns split definitions for a workspace and environment with optional pagination
func (s *SplitDefinitionsService) List(workspaceID, environmentID string, opts *GenericListQueryParams) (*SplitDefinitionListResult, error) {
	var result SplitDefinitionListResult
	baseURL := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/environments/%s", workspaceID, environmentID)

	// Enforce max limit of 50
	if opts != nil && opts.Limit > 50 {
		opts.Limit = 50
	}

	finalURL := s.client.buildURL(baseURL, opts)
	err := s.client.get(finalURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAll returns all split definitions for a workspace and environment by handling pagination automatically
func (s *SplitDefinitionsService) ListAll(workspaceID, environmentID string) ([]*SplitDefinition, error) {
	var allSplitDefinitions []*SplitDefinition
	offset := 0
	limit := 50 // Use max limit

	for {
		opts := &GenericListQueryParams{
			Offset: offset,
			Limit:  limit,
		}

		result, err := s.List(workspaceID, environmentID, opts)
		if err != nil {
			return nil, err
		}

		allSplitDefinitions = append(allSplitDefinitions, result.Objects...)

		// If we got fewer results than the limit, we've reached the end
		if len(result.Objects) < limit {
			break
		}

		// If we have totalCount and we've retrieved all items, break
		if result.TotalCount != nil && len(allSplitDefinitions) >= *result.TotalCount {
			break
		}

		offset += limit
	}

	return allSplitDefinitions, nil
}

