package api

import (
	"fmt"
	"strings"
)

// SegmentEnvironmentAssociationsService handles operations related to Split segment environment associations
type SegmentEnvironmentAssociationsService service

// SegmentEnvironmentAssociation represents a segment environment association
type SegmentEnvironmentAssociation struct {
	SegmentName      *string `json:"segmentName"`
	EnvironmentID    *string `json:"environmentId"`
	EnvironmentName  *string `json:"environmentName"`
	IncludeInSegment *bool   `json:"includeInSegment"`
}

// SegmentEnvironmentAssociationCreateRequest represents the request for creating a segment environment association
type SegmentEnvironmentAssociationCreateRequest struct {
	IncludeInSegment bool `json:"includeInSegment"`
}

// SegmentEnvironmentAssociationUpdateRequest represents the request for updating a segment environment association
type SegmentEnvironmentAssociationUpdateRequest struct {
	IncludeInSegment bool `json:"includeInSegment"`
}

// SegmentEnvironmentAssociationListResult represents the response from the segment environment associations list endpoint with pagination
type SegmentEnvironmentAssociationListResult struct {
	GenericListResult
	Objects []*SegmentEnvironmentAssociation `json:"objects"`
}

// Create creates a new segment environment association (activates segment)
func (s *SegmentEnvironmentAssociationsService) Create(workspaceID, segmentName, environmentID string, req *SegmentEnvironmentAssociationCreateRequest) (*SegmentEnvironmentAssociation, error) {
	if req.IncludeInSegment {
		return s.Activate(environmentID, segmentName)
	}
	return s.Deactivate(environmentID, segmentName)
}

// Get retrieves a segment environment association
func (s *SegmentEnvironmentAssociationsService) Get(workspaceID, segmentName, environmentID string) (*SegmentEnvironmentAssociation, error) {
	// Check if segment is active in the environment by listing segments for the environment
	segments, err := s.client.Segments.ListAll(workspaceID)
	if err != nil {
		return nil, err
	}

	for _, segment := range segments {
		if segment.Name != nil && *segment.Name == segmentName {
			// Found the segment, now check if it's active in this environment
			return &SegmentEnvironmentAssociation{
				SegmentName:      &segmentName,
				EnvironmentID:    &environmentID,
				IncludeInSegment: &[]bool{true}[0], // Assume active if found
			}, nil
		}
	}

	return nil, fmt.Errorf("segment environment association not found")
}

// Update modifies an existing segment environment association
func (s *SegmentEnvironmentAssociationsService) Update(workspaceID, segmentName, environmentID string, req *SegmentEnvironmentAssociationUpdateRequest) (*SegmentEnvironmentAssociation, error) {
	if req.IncludeInSegment {
		return s.Activate(environmentID, segmentName)
	}
	return s.Deactivate(environmentID, segmentName)
}

// Delete removes a segment environment association (deactivates segment)
func (s *SegmentEnvironmentAssociationsService) Delete(workspaceID, segmentName, environmentID string) error {
	// Before deactivating, try to clear any existing keys to avoid 409 errors
	keys, err := s.client.Environments.GetSegmentKeys(environmentID, segmentName)
	if err != nil {
		fmt.Printf("DEBUG: Failed to get segment keys: %v\n", err)
	} else if len(keys) > 0 {
		fmt.Printf("DEBUG: Clearing %d existing keys from segment %s in environment %s before deactivation\n",
			len(keys), segmentName, environmentID)
		fmt.Printf("DEBUG: Keys to remove: %v\n", keys)
		clearErr := s.client.Environments.RemoveSegmentKeys(environmentID, segmentName, keys)
		if clearErr != nil {
			fmt.Printf("DEBUG: Failed to clear keys before deactivation: %v\n", clearErr)
			if !strings.Contains(clearErr.Error(), "404") {
				return clearErr // Return the error instead of ignoring it
			}
		} else {
			fmt.Printf("DEBUG: Successfully cleared keys, verifying removal...\n")
			// Verify keys were actually removed
			remainingKeys, verifyErr := s.client.Environments.GetSegmentKeys(environmentID, segmentName)
			if verifyErr != nil {
				fmt.Printf("DEBUG: Could not verify key removal: %v\n", verifyErr)
			} else {
				fmt.Printf("DEBUG: Remaining keys after removal: %d - %v\n", len(remainingKeys), remainingKeys)
			}
		}
	} else {
		fmt.Printf("DEBUG: No keys found in segment %s environment %s\n", segmentName, environmentID)
	}

	_, err = s.Deactivate(environmentID, segmentName)
	// Ignore 404 errors when deleting - segment may never have been activated
	if err != nil && strings.Contains(err.Error(), "404") {
		fmt.Printf("DEBUG: Ignoring 404 error when deactivating segment %s in environment %s\n", segmentName, environmentID)
		return nil
	}
	return err
}

// Activate activates a segment in an environment
func (s *SegmentEnvironmentAssociationsService) Activate(environmentID, segmentName string) (*SegmentEnvironmentAssociation, error) {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/%s/%s", environmentID, segmentName)
	fmt.Printf("DEBUG: Activating segment %s in environment %s\n", segmentName, environmentID)
	err := s.client.post(url, nil, nil)
	if err != nil {
		fmt.Printf("DEBUG: Segment activation failed: %v\n", err)
		return nil, err
	}
	fmt.Printf("DEBUG: Segment activation successful\n")

	return &SegmentEnvironmentAssociation{
		SegmentName:      &segmentName,
		EnvironmentID:    &environmentID,
		IncludeInSegment: &[]bool{true}[0],
	}, nil
}

// Deactivate deactivates a segment in an environment
func (s *SegmentEnvironmentAssociationsService) Deactivate(environmentID, segmentName string) (*SegmentEnvironmentAssociation, error) {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/%s/%s", environmentID, segmentName)
	err := s.client.delete(url)
	if err != nil {
		// Ignore 404 errors when deactivating - segment may not have been activated
		if strings.Contains(err.Error(), "404") {
			return &SegmentEnvironmentAssociation{
				SegmentName:      &segmentName,
				EnvironmentID:    &environmentID,
				IncludeInSegment: &[]bool{false}[0],
			}, nil
		}
		return nil, err
	}

	return &SegmentEnvironmentAssociation{
		SegmentName:      &segmentName,
		EnvironmentID:    &environmentID,
		IncludeInSegment: &[]bool{false}[0],
	}, nil
}

// List returns segment environment associations with optional pagination
func (s *SegmentEnvironmentAssociationsService) List(workspaceID, environmentID string, opts *GenericListQueryParams) (*SegmentEnvironmentAssociationListResult, error) {
	var result SegmentEnvironmentAssociationListResult
	baseURL := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/environments/%s/associations", workspaceID, environmentID)

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

// ListAll returns all segment environment associations by handling pagination automatically
func (s *SegmentEnvironmentAssociationsService) ListAll(workspaceID, environmentID string) ([]*SegmentEnvironmentAssociation, error) {
	var allAssociations []*SegmentEnvironmentAssociation
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

		allAssociations = append(allAssociations, result.Objects...)

		// If we got fewer results than the limit, we've reached the end
		if len(result.Objects) < limit {
			break
		}

		// If we have totalCount and we've retrieved all items, break
		if result.TotalCount != nil && len(allAssociations) >= *result.TotalCount {
			break
		}

		offset += limit
	}

	return allAssociations, nil
}