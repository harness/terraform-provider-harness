// Unit tests for OPA policy denial detection in template create/update.
// Tests isOPADeniedResponse and formatOPADenyMessages without live Harness credentials,
// ensuring CI coverage for the deny-message extraction logic (including EqualFold status matching).
//
// Run: go test -run "TestFormatOPADenyMessages|TestIsOPADeniedResponse" ./internal/service/pipeline/template/ -v
package template

import (
	"testing"

	nextgen "github.com/harness/harness-openapi-go-client/nextgen"
)

func TestFormatOPADenyMessages(t *testing.T) {
	tests := []struct {
		name     string
		resp     nextgen.TemplateResponse
		expected string
	}{
		{
			name: "error status policy returns formatted deny message",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							Deny: true,
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName:   "Enforce Version Format",
									Status:       "error",
									DenyMessages: []string{"Version must start with 'v'"},
								},
							},
						},
					},
				},
			},
			expected: "[Enforce Version Format] Version must start with 'v'",
		},
		{
			name: "ERROR uppercase status also matched via EqualFold",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName:   "Policy A",
									Status:       "ERROR",
									DenyMessages: []string{"denied"},
								},
							},
						},
					},
				},
			},
			expected: "[Policy A] denied",
		},
		{
			name: "mixed case Error status also matched via EqualFold",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName:   "Policy B",
									Status:       "Error",
									DenyMessages: []string{"denied"},
								},
							},
						},
					},
				},
			},
			expected: "[Policy B] denied",
		},
		{
			name: "pass status policy returns empty string",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: false,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName: "Passing Policy",
									Status:     "pass",
								},
							},
						},
					},
				},
			},
			expected: "",
		},
		{
			name:     "nil GovernanceMetadata returns empty string",
			resp:     nextgen.TemplateResponse{},
			expected: "",
		},
		{
			name: "multiple policies joined with semicolons",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName:   "Policy A",
									Status:       "error",
									DenyMessages: []string{"msg1"},
								},
								{
									PolicyName:   "Policy B",
									Status:       "error",
									DenyMessages: []string{"msg2", "msg3"},
								},
							},
						},
					},
				},
			},
			expected: "[Policy A] msg1; [Policy B] msg2; [Policy B] msg3",
		},
		{
			name: "passing policies filtered out, only error policies included",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
					Details: []nextgen.TemplatePolicySetMetadata{
						{
							PolicyMetadata: []nextgen.TemplatePolicyMetadata{
								{
									PolicyName: "Passing Policy",
									Status:     "pass",
								},
								{
									PolicyName:   "Failing Policy",
									Status:       "error",
									DenyMessages: []string{"denied"},
								},
							},
						},
					},
				},
			},
			expected: "[Failing Policy] denied",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatOPADenyMessages(tt.resp)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestIsOPADeniedResponse(t *testing.T) {
	tests := []struct {
		name       string
		templateID string
		resp       nextgen.TemplateResponse
		expected   bool
	}{
		{
			name:       "GovernanceMetadata.Deny true",
			templateID: "",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: true,
				},
			},
			expected: true,
		},
		{
			name:       "GovernanceMetadata.Deny false",
			templateID: "some_id",
			resp: nextgen.TemplateResponse{
				GovernanceMetadata: &nextgen.TemplateGovernanceMetadata{
					Deny: false,
				},
			},
			expected: false,
		},
		{
			name:       "nil GovernanceMetadata with populated fields",
			templateID: "test_template",
			resp: nextgen.TemplateResponse{
				Identifier: "test_template",
				Name:       "Test Template",
				Yaml:       "yaml: content",
			},
			expected: false,
		},
		{
			name:       "nil GovernanceMetadata with non-empty templateID",
			templateID: "existing_id",
			resp:       nextgen.TemplateResponse{},
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOPADeniedResponse(tt.templateID, tt.resp)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
