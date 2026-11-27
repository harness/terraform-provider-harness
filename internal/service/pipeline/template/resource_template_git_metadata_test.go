package template

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

const testRemoteTemplateYaml = `template:
  name: tmpl
  identifier: tmpl
  versionLabel: v1
  type: Step
  spec: {}
`

func TestResourceTemplateUpdateRemoteGitMetadata(t *testing.T) {
	tests := []struct {
		name         string
		org          string
		project      string
		metadataPath string
	}{
		{
			name:         "account scope",
			metadataPath: "/v1/templates/git-metadata/tmpl",
		},
		{
			name:         "org scope",
			org:          "org",
			metadataPath: "/v1/orgs/org/templates/git-metadata/tmpl",
		},
		{
			name:         "project scope",
			org:          "org",
			project:      "proj",
			metadataPath: "/v1/orgs/org/projects/proj/templates/git-metadata/tmpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var metadataPaths []string
			var metadataAccounts []string
			var metadataBodies []nextgen.TemplateUpdateGitMetadataRequest

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == http.MethodPut && strings.Contains(r.URL.Path, "/git-metadata/"):
					metadataPaths = append(metadataPaths, r.URL.Path)
					metadataAccounts = append(metadataAccounts, r.Header.Get("Harness-Account"))
					var body []nextgen.TemplateUpdateGitMetadataRequest
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					metadataBodies = append(metadataBodies, body...)
					_, _ = w.Write([]byte(`{"status":true}`))
				case r.Method == http.MethodPut:
					_ = json.NewEncoder(w).Encode(nextgen.TemplateResponse{Identifier: "tmpl"})
				case r.Method == http.MethodGet:
					_ = json.NewEncoder(w).Encode(nextgen.TemplateWithInputsResponse{Template: &nextgen.TemplateResponse{
						Identifier:   "tmpl",
						Name:         "tmpl",
						Org:          tt.org,
						Project:      tt.project,
						Yaml:         testRemoteTemplateYaml,
						VersionLabel: "v1",
						GitDetails: &nextgen.EntityGitDetails{
							BranchName: "main",
							FilePath:   ".harness/tmpl.yaml",
							RepoName:   "templates",
						},
					}})
				default:
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			t.Cleanup(server.Close)

			state, diff, session := remoteTemplateUpdate(t, tt.org, tt.project, server.URL)
			require.False(t, diff.Empty(), "changing connector_ref must produce an update")

			r := ResourceTemplate()
			_, diags := r.Apply(context.Background(), state, diff, session)
			require.False(t, diags.HasError(), "%v", diags)

			require.Equal(t, []string{tt.metadataPath}, metadataPaths)
			require.Equal(t, []string{"account-123"}, metadataAccounts)
			require.Len(t, metadataBodies, 1)
			require.Equal(t, "v1", metadataBodies[0].Version)
			require.Equal(t, "account.new_connector", metadataBodies[0].GitDetails.ConnectorRef)
			require.Equal(t, "templates", metadataBodies[0].GitDetails.RepoName)
			require.Equal(t, ".harness/tmpl.yaml", metadataBodies[0].GitDetails.FilePath)
		})
	}
}

func TestResourceTemplateUpdateRemoteGitMetadataError(t *testing.T) {
	tests := []struct {
		name          string
		status        int
		expectedError string
	}{
		{
			name:          "non-retryable API error",
			status:        http.StatusBadRequest,
			expectedError: "metadata update failed",
		},
		{
			name:          "retryable API error",
			status:        http.StatusInternalServerError,
			expectedError: "giving up after 1 attempt(s)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if strings.Contains(r.URL.Path, "/git-metadata/") {
					w.WriteHeader(tt.status)
					_, _ = w.Write([]byte(`{"message":"metadata update failed"}`))
					return
				}
				if r.Method == http.MethodPut {
					_ = json.NewEncoder(w).Encode(nextgen.TemplateResponse{Identifier: "tmpl"})
					return
				}
				_ = json.NewEncoder(w).Encode(nextgen.TemplateWithInputsResponse{Template: &nextgen.TemplateResponse{
					Identifier: "tmpl", Name: "tmpl", Yaml: testRemoteTemplateYaml, VersionLabel: "v1",
				}})
			}))
			t.Cleanup(server.Close)

			state, diff, session := remoteTemplateUpdate(t, "", "", server.URL)
			_, diags := ResourceTemplate().Apply(context.Background(), state, diff, session)
			require.True(t, diags.HasError(), "a failed Git metadata update must not be reported as a successful apply")
			require.Contains(t, fmt.Sprint(diags), tt.expectedError)
		})
	}
}

func remoteTemplateUpdate(t *testing.T, org, project, serverURL string) (*terraform.InstanceState, *terraform.InstanceDiff, *internal.Session) {
	t.Helper()

	r := ResourceTemplate()
	stateData := schema.TestResourceDataRaw(t, r.Schema, remoteTemplateConfig(org, project, "account.old_connector"))
	stateData.SetId("tmpl")
	state := stateData.State()

	diff, err := r.SimpleDiff(
		context.Background(),
		state,
		terraform.NewResourceConfigRaw(remoteTemplateConfig(org, project, "account.new_connector")),
		nil,
	)
	require.NoError(t, err)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = serverURL
	cfg.AccountId = "account-123"
	cfg.ApiKey = "test-key"
	cfg.HTTPClient.RetryMax = 0
	session := &internal.Session{AccountId: "account-123", Client: nextgen.NewAPIClient(cfg)}
	return state, diff, session
}

func remoteTemplateConfig(org, project, connectorRef string) map[string]interface{} {
	return map[string]interface{}{
		"identifier":    "tmpl",
		"name":          "tmpl",
		"org_id":        org,
		"project_id":    project,
		"version":       "v1",
		"is_stable":     false,
		"template_yaml": testRemoteTemplateYaml,
		"git_details": []interface{}{map[string]interface{}{
			"branch_name":    "main",
			"file_path":      ".harness/tmpl.yaml",
			"commit_message": "update template",
			"base_branch":    "main",
			"connector_ref":  connectorRef,
			"store_type":     "REMOTE",
			"repo_name":      "templates",
			"last_object_id": "object-id",
			"last_commit_id": "commit-id",
		}},
	}
}
