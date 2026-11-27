package pipeline

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestResourcePipelineDescriptionValidation(t *testing.T) {
	tests := []struct {
		name            string
		description     string
		yaml            string
		omitDescription bool
		omitYAML        bool
		errorContains   []string
	}{
		{
			name:        "different non-empty descriptions fail at plan time",
			description: "Terraform description",
			yaml:        "pipeline:\n  description: YAML description\n",
			errorContains: []string{
				"pipeline description mismatch",
				"top-level description",
				"pipeline.description",
				"Terraform description",
				"YAML description",
				"update",
				"match",
			},
		},
		{
			name:        "matching descriptions pass",
			description: "Shared description",
			yaml:        "pipeline:\n  description: Shared description\n",
		},
		{
			name:        "empty top-level description retains backend behavior",
			description: "",
			yaml:        "pipeline:\n  description: YAML description\n",
		},
		{
			name:            "omitted top-level description retains backend behavior",
			omitDescription: true,
			yaml:            "pipeline:\n  description: YAML description\n",
		},
		{
			name:        "empty YAML description retains backend behavior",
			description: "Terraform description",
			yaml:        "pipeline:\n  description: \"\"\n",
		},
		{
			name:        "omitted YAML description retains backend behavior",
			description: "Terraform description",
			yaml:        "pipeline:\n  name: Test Pipeline\n",
		},
		{
			name:        "empty YAML retains backend behavior",
			description: "Terraform description",
			yaml:        "",
		},
		{
			name:        "omitted YAML retains backend behavior",
			description: "Terraform description",
			omitYAML:    true,
		},
		{
			name:          "invalid YAML produces a clear error even without a top-level description",
			description:   "",
			yaml:          "pipeline: [\n",
			errorContains: []string{"invalid pipeline YAML"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := map[string]interface{}{
				"identifier": "test_pipeline",
				"name":       "Test Pipeline",
				"org_id":     "default",
				"project_id": "test_project",
			}
			if !tt.omitDescription {
				config["description"] = tt.description
			}
			if !tt.omitYAML {
				config["yaml"] = tt.yaml
			}

			_, err := ResourcePipeline().Diff(
				context.Background(),
				nil,
				terraform.NewResourceConfigRaw(config),
				nil,
			)

			if len(tt.errorContains) == 0 {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			for _, message := range tt.errorContains {
				require.ErrorContains(t, err, message)
			}
		})
	}
}
