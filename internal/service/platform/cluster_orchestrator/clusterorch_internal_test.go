package cluster_orchestrator

// White-box coverage for the not-found detection that makes destroy idempotent
// when a cluster orchestrator was deleted out of band (CCM-34967). This file is
// in package cluster_orchestrator rather than cluster_orchestrator_test because
// isClusterOrchestratorNotFound is unexported; it cannot move into
// clusterorch_test.go, which imports internal/acctest and would form an import
// cycle. Deterministic; no live API / TF_ACC required.

import (
	"errors"
	"net/http"
	"testing"
)

func TestIsClusterOrchestratorNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		httpResp *http.Response
		want     bool
	}{
		{
			name:     "no error",
			err:      nil,
			httpResp: &http.Response{StatusCode: http.StatusOK},
			want:     false,
		},
		{
			name:     "http 404",
			err:      errors.New("404 Not Found"),
			httpResp: &http.Response{StatusCode: http.StatusNotFound},
			want:     true,
		},
		{
			// GET details joins the API's errors array into the error message.
			name:     "http 500 with invalid cluster id message",
			err:      errors.New("Could not find cluster orchestrator details. invalid cluster id"),
			httpResp: &http.Response{StatusCode: http.StatusInternalServerError},
			want:     true,
		},
		{
			name:     "message casing is ignored",
			err:      errors.New("INVALID CLUSTER ID"),
			httpResp: &http.Response{StatusCode: http.StatusInternalServerError},
			want:     true,
		},
		{
			// Retries strip the response and body, leaving nothing to match on, so
			// callers must confirm with clusterOrchestratorMissing instead.
			name:     "retries exhausted",
			err:      errors.New("DELETE https://app.harness.io/gateway/lw/api/accounts/acc/clusters/orchestrator/orch-1 giving up after 11 attempt(s)"),
			httpResp: nil,
			want:     false,
		},
		{
			name:     "unrelated server error",
			err:      errors.New("500 Internal Server Error"),
			httpResp: &http.Response{StatusCode: http.StatusInternalServerError},
			want:     false,
		},
		{
			name:     "permission error is not treated as gone",
			err:      errors.New("403 Forbidden"),
			httpResp: &http.Response{StatusCode: http.StatusForbidden},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isClusterOrchestratorNotFound(tt.err, tt.httpResp); got != tt.want {
				t.Errorf("isClusterOrchestratorNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
