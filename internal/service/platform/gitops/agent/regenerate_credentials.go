package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/harness/terraform-provider-harness/internal"
)

// regenerateAgentCredentials calls POST /gitops/api/v1/agents/{identifier}/credentials and returns the
// newly generated private key (the agent token). The generated SDK method does not accept the
// account/org/project scope query parameters, so the request is made directly.
func regenerateAgentCredentials(ctx context.Context, session *internal.Session, identifier, orgID, projectID string) (string, error) {
	q := url.Values{}
	q.Set("accountIdentifier", session.AccountId)
	if orgID != "" {
		q.Set("orgIdentifier", orgID)
	}
	if projectID != "" {
		q.Set("projectIdentifier", projectID)
	}
	endpoint := strings.TrimSuffix(session.Endpoint, "/")
	u := fmt.Sprintf("%s/gitops/api/v1/agents/%s/credentials?%s", endpoint, url.PathEscape(identifier), q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return "", fmt.Errorf("building regenerate credentials request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	if session.PLClient != nil && session.PLClient.ApiKey != "" {
		req.Header.Set("x-api-key", session.PLClient.ApiKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("regenerating agent credentials: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("regenerating agent credentials for %q failed with status %s: %s", identifier, resp.Status, string(body))
	}

	var out struct {
		Credentials struct {
			PrivateKey string `json:"privateKey"`
		} `json:"credentials"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("decoding regenerate credentials response: %w", err)
	}
	return out.Credentials.PrivateKey, nil
}
