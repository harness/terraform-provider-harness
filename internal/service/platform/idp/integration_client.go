package idp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/harness/terraform-provider-harness/internal"
)

// integrationsClient calls the integration-manager integrations API.
type integrationsClient struct {
	base      string
	accountID string
	apiKey    string
	http      *http.Client
}

// integrationsAPIError is a non-2xx response from the integrations API.
type integrationsAPIError struct {
	Status    int    `json:"-"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Details   []struct {
		Field   string `json:"field"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"details"`
}

func (e *integrationsAPIError) Error() string {
	msg := fmt.Sprintf("integrations API: %d %s: %s", e.Status, e.Code, e.Message)
	for _, d := range e.Details {
		msg += fmt.Sprintf("; %s: %s", d.Field, d.Message)
	}
	if e.RequestID != "" {
		msg += " (request " + e.RequestID + ")"
	}
	return msg
}

func isIntegrationsStatus(err error, status int) bool {
	apiErr, ok := err.(*integrationsAPIError)
	return ok && apiErr.Status == status
}

func newIntegrationsClient(s *internal.Session) *integrationsClient {
	apiKey := ""
	if s.IDPClient != nil {
		apiKey = s.IDPClient.ApiKey
	}
	return &integrationsClient{
		base:      strings.TrimRight(s.Endpoint, "/") + "/integration/api/v1",
		accountID: s.AccountId,
		apiKey:    apiKey,
		http:      http.DefaultClient,
	}
}

// scopePath is the account, org or project prefix of every integrations route.
func (c *integrationsClient) scopePath(org, project string) string {
	p := "/accounts/" + url.PathEscape(c.accountID)
	if org != "" {
		p += "/orgs/" + url.PathEscape(org)
		if project != "" {
			p += "/projects/" + url.PathEscape(project)
		}
	}
	return p
}

// getRetries is how often a GET is retried on a transient server error, e.g. a space created concurrently.
const getRetries = 3

// do sends a request and decodes a JSON response into out; it returns the response headers.
func (c *integrationsClient) do(
	ctx context.Context, method, path string, query url.Values, ifMatch *int64, body, out any,
) (http.Header, error) {
	for attempt := 0; ; attempt++ {
		h, err := c.doOnce(ctx, method, path, query, ifMatch, body, out)
		if method != http.MethodGet || attempt == getRetries || !isTransient(err) {
			return h, err
		}
		select {
		case <-ctx.Done():
			return h, err
		case <-time.After(time.Duration(attempt+1) * retryBackoff):
		}
	}
}

var retryBackoff = 500 * time.Millisecond

func isTransient(err error) bool {
	apiErr, ok := err.(*integrationsAPIError)
	return ok && (apiErr.Status == http.StatusInternalServerError || apiErr.Status == http.StatusBadGateway ||
		apiErr.Status == http.StatusServiceUnavailable || apiErr.Status == http.StatusGatewayTimeout)
}

func (c *integrationsClient) doOnce(
	ctx context.Context, method, path string, query url.Values, ifMatch *int64, body, out any,
) (http.Header, error) {
	u := c.base + path
	if len(query) > 0 {
		// The gateway doesn't decode + as a space.
		u += "?" + strings.ReplaceAll(query.Encode(), "+", "%20")
	}
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Harness-Account", c.accountID)
	if ifMatch != nil {
		req.Header.Set("If-Match", strconv.Quote(strconv.FormatInt(*ifMatch, 10)))
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		apiErr := &integrationsAPIError{Status: resp.StatusCode}
		if json.Unmarshal(raw, apiErr) != nil || apiErr.Message == "" {
			apiErr.Message = strings.TrimSpace(string(raw))
		}
		return resp.Header, apiErr
	}
	if out != nil && len(raw) > 0 {
		if err := json.Unmarshal(raw, out); err != nil {
			return resp.Header, fmt.Errorf("decoding %s %s: %w", method, path, err)
		}
	}
	return resp.Header, nil
}

func integrationPath(id string) string {
	return "/integrations/" + url.PathEscape(id)
}

func (c *integrationsClient) getIntegration(ctx context.Context, org, project, id string) (map[string]any, error) {
	var out map[string]any
	_, err := c.do(ctx, http.MethodGet, c.scopePath(org, project)+integrationPath(id), nil, nil, nil, &out)
	return out, err
}

func (c *integrationsClient) createIntegration(
	ctx context.Context, org, project string, body map[string]any,
) (map[string]any, error) {
	var out map[string]any
	_, err := c.do(ctx, http.MethodPost, c.scopePath(org, project)+"/integrations", nil, nil, body, &out)
	return out, err
}

func (c *integrationsClient) replaceIntegration(
	ctx context.Context, org, project, id string, revision int64, body map[string]any,
) (map[string]any, error) {
	var out map[string]any
	_, err := c.do(ctx, http.MethodPut, c.scopePath(org, project)+integrationPath(id), nil, &revision, body, &out)
	return out, err
}

func (c *integrationsClient) deleteIntegration(ctx context.Context, org, project, id string, revision int64) error {
	_, err := c.do(ctx, http.MethodDelete, c.scopePath(org, project)+integrationPath(id), nil, &revision, nil, nil)
	return err
}

// listIntegrations pages through every integration matching query.
func (c *integrationsClient) listIntegrations(
	ctx context.Context, org, project string, query url.Values,
) ([]map[string]any, error) {
	const limit = 200
	var all []map[string]any
	for offset := 0; ; offset += limit {
		q := url.Values{}
		for k, v := range query {
			q[k] = v
		}
		q.Set("limit", strconv.Itoa(limit))
		q.Set("offset", strconv.Itoa(offset))
		var page struct {
			Items []map[string]any `json:"items"`
			Page  struct {
				Total int `json:"total"`
			} `json:"page"`
		}
		if _, err := c.do(ctx, http.MethodGet, c.scopePath(org, project)+"/integrations", q, nil, nil, &page); err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if len(page.Items) < limit || len(all) >= page.Page.Total {
			return all, nil
		}
	}
}

func mappingPath(id, kind string) string {
	return integrationPath(id) + "/kinds/" + url.PathEscape(kind) + "/mapping"
}

func (c *integrationsClient) getMapping(ctx context.Context, org, project, id, kind string) (map[string]any, error) {
	var out map[string]any
	_, err := c.do(ctx, http.MethodGet, c.scopePath(org, project)+mappingPath(id, kind), nil, nil, nil, &out)
	return out, err
}

func (c *integrationsClient) resetMapping(ctx context.Context, org, project, id, kind string, revision int64) error {
	_, err := c.do(ctx, http.MethodDelete, c.scopePath(org, project)+mappingPath(id, kind), nil, &revision, nil, nil)
	return err
}

func (c *integrationsClient) replaceMapping(
	ctx context.Context, org, project, id, kind string, revision int64, body map[string]any,
) (map[string]any, error) {
	var out map[string]any
	_, err := c.do(ctx, http.MethodPut, c.scopePath(org, project)+mappingPath(id, kind), nil, &revision, body, &out)
	return out, err
}

func (c *integrationsClient) getCatalog(ctx context.Context, org, project, id string) (map[string]any, error) {
	var out map[string]any
	q := url.Values{"include": []string{"json_schema"}}
	_, err := c.do(ctx, http.MethodGet, c.scopePath(org, project)+integrationPath(id)+"/catalog", q, nil, nil, &out)
	return out, err
}

func (c *integrationsClient) getIntegrationType(ctx context.Context, typ string) (map[string]any, error) {
	var out map[string]any
	path := c.scopePath("", "") + "/integration-types/" + url.PathEscape(typ)
	_, err := c.do(ctx, http.MethodGet, path, nil, nil, nil, &out)
	return out, err
}
