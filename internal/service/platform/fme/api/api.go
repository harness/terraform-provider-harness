package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/davidji99/simpleresty"
)

const (
	DefaultAPIBaseURL        = "https://api.split.io/internal/api/v2"
	DefaultUserAgent         = "split-go"
	DefaultContentTypeHeader = "application/json"
	DefaultAcceptHeader      = "application/json"
	DefaultClientTimeout     = 300

	UserStatusPending     = "PENDING"
	UserStatusActive      = "ACTIVE"
	UserStatusDeactivated = "DEACTIVATED"

	timeoutError = "reached maximum client timeout"
)

type Client struct {
	http      *simpleresty.Client
	common    service
	config    *Config
	expiresAt time.Time

	ApiKeys                        *KeysService
	Environments                   *EnvironmentsService
	EnvironmentSegmentKeys         *EnvironmentSegmentKeysService
	FlagSets                       *FlagSetsService
	Segments                       *SegmentService
	SegmentEnvironmentAssociations *SegmentEnvironmentAssociationsService
	Splits                         *SplitsService
	SplitDefinitions               *SplitDefinitionsService
	TrafficTypes                   *TrafficTypesService
	TrafficTypeAttributes          *TrafficTypeAttributesService
	Workspaces                     *WorkspacesService
}

type service struct {
	client *Client
}

type GenericListResult struct {
	Offset     *int `json:"offset"`
	Limit      *int `json:"limit"`
	TotalCount *int `json:"totalCount"`
}

type GenericListQueryParams struct {
	Offset int `url:"offset,omitempty"`
	Limit  int `url:"limit,omitempty"`
}

func New(opts ...Option) (*Client, error) {
	config := &Config{
		APIBaseURL:        DefaultAPIBaseURL,
		UserAgent:         DefaultUserAgent,
		ContentTypeHeader: DefaultContentTypeHeader,
		AcceptHeader:      DefaultAcceptHeader,
		ClientTimeout:     DefaultClientTimeout,
		APIKey:            "",
	}

	if optErr := config.ParseOptions(opts...); optErr != nil {
		return nil, optErr
	}

	if config.APIKey == "" {
		return nil, errors.New("api key cannot be empty")
	}

	// Set HTTP client
	httpClient := simpleresty.NewWithBaseURL(config.APIBaseURL)

	// Set headers
	httpClient.SetHeader("Content-Type", config.ContentTypeHeader).
		SetHeader("Accept", config.AcceptHeader).
		SetHeader("User-Agent", config.UserAgent).
		SetHeader("X-API-Key", config.APIKey)

	// Set additional custom headers
	if len(config.CustomHTTPHeaders) != 0 {
		for k, v := range config.CustomHTTPHeaders {
			httpClient.SetHeader(k, v)
		}
	}

	// Set timeout
	httpClient.SetTimeout(time.Duration(config.ClientTimeout) * time.Second)

	log.Printf("[INFO] Split API Client configured")

	c := &Client{
		http:      httpClient,
		config:    config,
		expiresAt: time.Now().Add(time.Duration(config.ClientTimeout) * time.Second),
	}
	c.common.client = c
	c.ApiKeys = (*KeysService)(&c.common)
	c.Environments = (*EnvironmentsService)(&c.common)
	c.EnvironmentSegmentKeys = (*EnvironmentSegmentKeysService)(&c.common)
	c.FlagSets = (*FlagSetsService)(&c.common)
	c.Segments = (*SegmentService)(&c.common)
	c.SegmentEnvironmentAssociations = (*SegmentEnvironmentAssociationsService)(&c.common)
	c.Splits = (*SplitsService)(&c.common)
	c.SplitDefinitions = (*SplitDefinitionsService)(&c.common)
	c.TrafficTypes = (*TrafficTypesService)(&c.common)
	c.TrafficTypeAttributes = (*TrafficTypeAttributesService)(&c.common)
	c.Workspaces = (*WorkspacesService)(&c.common)

	return c, nil
}

func (c *Client) get(endpoint string, target interface{}) error {
	return c.executeRequest("GET", endpoint, nil, target)
}

func (c *Client) post(endpoint string, data, target interface{}) error {
	return c.executeRequest("POST", endpoint, data, target)
}

func (c *Client) put(endpoint string, data, target interface{}) error {
	return c.executeRequest("PUT", endpoint, data, target)
}

func (c *Client) patch(endpoint string, data, target interface{}) error {
	return c.executeRequest("PATCH", endpoint, data, target)
}

func (c *Client) delete(endpoint string) error {
	return c.executeRequest("DELETE", endpoint, nil, nil)
}

func (c *Client) executeRequest(method, endpoint string, data, target interface{}) error {
	return c.executeRequestWithRetry(method, endpoint, data, target, 0)
}

func (c *Client) executeRequestWithRetry(method, endpoint string, data, target interface{}, retryCount int) error {
	const maxRetries = 10

	if c.exceedsTimeout() {
		return errors.New(timeoutError)
	}

	// Log request details
	log.Printf("[DEBUG] API Request: %s %s", method, endpoint)
	if data != nil {
		if jsonBytes, err := json.Marshal(data); err == nil {
			log.Printf("[DEBUG] Request JSON: %s", string(jsonBytes))
		}
	}

	var resp *simpleresty.Response
	var err error

	switch method {
	case "GET":
		resp, err = c.http.Get(endpoint, target, nil)
	case "POST":
		resp, err = c.http.Post(endpoint, target, data)
	case "PUT":
		resp, err = c.http.Put(endpoint, target, data)
	case "PATCH":
		resp, err = c.http.Patch(endpoint, target, data)
	case "DELETE":
		resp, err = c.http.Delete(endpoint, nil, nil)
	}

	if err != nil {
		log.Printf("[DEBUG] API Request failed: %v", err)
		// Check if this is a 429 error embedded in the error message
		if strings.Contains(err.Error(), "429") {
			log.Printf("[DEBUG] 429 Rate Limited (from error) - Current retry count: %d, Max retries: %d", retryCount, maxRetries)

			if retryCount >= maxRetries {
				log.Printf("[ERROR] API rate limited - EXHAUSTED all %d retries for %s %s", maxRetries, method, endpoint)
				return fmt.Errorf("API rate limited after %d retries: %v", maxRetries, err)
			}

			// Fixed 5 second backoff for each retry
			backoffDuration := 5 * time.Second
			log.Printf("[DEBUG] API rate limited (retry %d/%d) - starting %v backoff for %s %s", retryCount+1, maxRetries, backoffDuration, method, endpoint)

			// Add timestamp logging
			startTime := time.Now()
			time.Sleep(backoffDuration)
			endTime := time.Now()
			actualWait := endTime.Sub(startTime)

			log.Printf("[DEBUG] Backoff completed - waited %v (expected %v) - retrying %s %s", actualWait, backoffDuration, method, endpoint)

			return c.executeRequestWithRetry(method, endpoint, data, target, retryCount+1)
		}
		return err
	}

	// Log response details
	log.Printf("[DEBUG] API Response: %d", resp.StatusCode)
	if resp.Body != "" {
		log.Printf("[DEBUG] Response JSON: %s", resp.Body)
	}

	if resp.StatusCode == 429 {
		log.Printf("[DEBUG] 429 Rate Limited - Current retry count: %d, Max retries: %d", retryCount, maxRetries)
		log.Printf("[DEBUG] 429 Response body: %s", resp.Body)

		if retryCount >= maxRetries {
			log.Printf("[ERROR] API rate limited - EXHAUSTED all %d retries for %s %s", maxRetries, method, endpoint)
			return fmt.Errorf("API rate limited after %d retries: status %d", maxRetries, resp.StatusCode)
		}

		// Fixed 5 second backoff for each retry
		backoffDuration := 5 * time.Second
		log.Printf("[DEBUG] API rate limited (retry %d/%d) - starting %v backoff for %s %s", retryCount+1, maxRetries, backoffDuration, method, endpoint)

		// Add timestamp logging
		startTime := time.Now()
		time.Sleep(backoffDuration)
		endTime := time.Now()
		actualWait := endTime.Sub(startTime)

		log.Printf("[DEBUG] Backoff completed - waited %v (expected %v) - retrying %s %s", actualWait, backoffDuration, method, endpoint)

		return c.executeRequestWithRetry(method, endpoint, data, target, retryCount+1)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) exceedsTimeout() bool {
	return time.Now().After(c.expiresAt)
}

func (c *Client) SetTimeout(t time.Time) {
	c.expiresAt = t
}

// buildURL constructs a URL with query parameters from GenericListQueryParams
func (c *Client) buildURL(baseURL string, opts *GenericListQueryParams) string {
	if opts == nil || (opts.Offset == 0 && opts.Limit == 0) {
		return baseURL
	}

	var params []string
	if opts.Offset > 0 {
		params = append(params, fmt.Sprintf("offset=%d", opts.Offset))
	}
	if opts.Limit > 0 {
		params = append(params, fmt.Sprintf("limit=%d", opts.Limit))
	}

	if len(params) > 0 {
		return baseURL + "?" + strings.Join(params, "&")
	}
	return baseURL
}
