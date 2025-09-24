package api

import (
	"fmt"
)

// SegmentService handles operations related to Split segments
type SegmentService service

// Segment represents a Split segment
type Segment struct {
	Name            *string `json:"name"`
	Description     *string `json:"description,omitempty"`
	WorkspaceID     *string `json:"workspaceId"`
	TrafficTypeID   *string `json:"trafficTypeId"`
	CreationTime    *int64  `json:"creationTime,omitempty"`
	LastUpdateTime  *int64  `json:"lastUpdateTime,omitempty"`
}

// SegmentListResult represents the response from the segments list endpoint
type SegmentListResult struct {
	GenericListResult
	Objects []*Segment `json:"objects"`
}

// SegmentCreateRequest represents the request body for creating a segment
type SegmentCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Create creates a new segment in the specified workspace
func (s *SegmentService) Create(workspaceID, trafficTypeID string, req *SegmentCreateRequest) (*Segment, error) {
	var result Segment
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/trafficTypes/%s", workspaceID, trafficTypeID)
	err := s.client.post(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a segment by workspace ID and segment name using List endpoint
func (s *SegmentService) Get(workspaceID, segmentName string) (*Segment, error) {
	segments, err := s.ListAll(workspaceID)
	if err != nil {
		return nil, err
	}

	for _, segment := range segments {
		if segment.Name != nil && *segment.Name == segmentName {
			return segment, nil
		}
	}

	return nil, fmt.Errorf("segment with name %s not found in workspace %s", segmentName, workspaceID)
}

// Delete removes a segment by workspace ID and segment name
func (s *SegmentService) Delete(workspaceID, segmentName string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s", workspaceID, segmentName)
	err := s.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List retrieves segments for a workspace with optional pagination
func (s *SegmentService) List(workspaceID string, opts *GenericListQueryParams) (*SegmentListResult, error) {
	var result SegmentListResult
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s", workspaceID)

	// Enforce max limit of 50
	if opts != nil && opts.Limit > 50 {
		opts.Limit = 50
	}

	// Build URL with query parameters and use retry-enabled client
	finalURL := s.client.buildURL(url, opts)
	err := s.client.get(finalURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAll retrieves all segments for a workspace by handling pagination automatically
func (s *SegmentService) ListAll(workspaceID string) ([]*Segment, error) {
	var allSegments []*Segment
	offset := 0
	limit := 50 // Use max limit

	for {
		opts := &GenericListQueryParams{
			Offset: offset,
			Limit:  limit,
		}

		result, err := s.List(workspaceID, opts)
		if err != nil {
			return nil, err
		}

		allSegments = append(allSegments, result.Objects...)

		// If we got fewer results than the limit, we've reached the end
		if len(result.Objects) < limit {
			break
		}

		// If we have totalCount and we've retrieved all items, break
		if result.TotalCount != nil && len(allSegments) >= *result.TotalCount {
			break
		}

		offset += limit
	}

	return allSegments, nil
}