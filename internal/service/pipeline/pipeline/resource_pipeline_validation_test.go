package pipeline

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestResourcePipelineCustomizeDiffValidatesYamlFields(t *testing.T) {
	tests := []struct {
		name         string
		identifier   string
		pipelineName string
		description  string
		yaml         string
		wantError    string
	}{
		{
			name:         "matching fields pass",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			description:  "Pipeline description",
			yaml: `pipeline:
  identifier: pipeline_id
  name: Pipeline name
  description: Pipeline description
`,
		},
		{
			name:         "description mismatch fails",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			description:  "Terraform description",
			yaml: `pipeline:
  identifier: pipeline_id
  name: Pipeline name
  description: YAML description
`,
			wantError: `description in schema ("Terraform description") does not match pipeline.description in YAML ("YAML description"); update description or pipeline.description so they match`,
		},
		{
			name:         "identifier mismatch fails",
			identifier:   "terraform_id",
			pipelineName: "Pipeline name",
			yaml: `pipeline:
  identifier: yaml_id
  name: Pipeline name
`,
			wantError: `identifier in schema ("terraform_id") does not match pipeline.identifier in YAML ("yaml_id"); update identifier or pipeline.identifier so they match`,
		},
		{
			name:         "name mismatch fails",
			identifier:   "pipeline_id",
			pipelineName: "Terraform name",
			yaml: `pipeline:
  identifier: pipeline_id
  name: YAML name
`,
			wantError: `name in schema ("Terraform name") does not match pipeline.name in YAML ("YAML name"); update name or pipeline.name so they match`,
		},
		{
			name:         "empty schema description retains YAML behavior",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			yaml: `pipeline:
  identifier: pipeline_id
  name: Pipeline name
  description: YAML description
`,
		},
		{
			name:         "omitted YAML description retains schema behavior",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			description:  "Terraform description",
			yaml: `pipeline:
  identifier: pipeline_id
  name: Pipeline name
`,
		},
		{
			name:         "empty YAML description retains schema behavior",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			description:  "Terraform description",
			yaml: `pipeline:
  identifier: pipeline_id
  name: Pipeline name
  description: ""
`,
		},
		{
			name:         "empty YAML retains existing behavior",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			description:  "Terraform description",
		},
		{
			name:         "invalid YAML fails clearly",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			yaml:         "pipeline: [",
			wantError:    "failed to parse YAML",
		},
		{
			name:         "missing pipeline root fails clearly",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			yaml:         "other: {}\n",
			wantError:    "YAML must contain a 'pipeline' key",
		},
		{
			name:         "missing required YAML identifier fails clearly",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			yaml: `pipeline:
  name: Pipeline name
`,
			wantError: "YAML must contain 'pipeline.identifier' field",
		},
		{
			name:         "missing required YAML name fails clearly",
			identifier:   "pipeline_id",
			pipelineName: "Pipeline name",
			yaml: `pipeline:
  identifier: pipeline_id
`,
			wantError: "YAML must contain 'pipeline.name' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := terraform.NewResourceConfigRaw(map[string]interface{}{
				"identifier":  tt.identifier,
				"name":        tt.pipelineName,
				"description": tt.description,
				"org_id":      "org",
				"project_id":  "project",
				"yaml":        tt.yaml,
			})

			diff, err := ResourcePipeline().Diff(context.Background(), nil, config, nil)
			if tt.wantError != "" {
				require.ErrorContains(t, err, tt.wantError)
				return
			}
			require.NoError(t, err)
			if tt.yaml != "" {
				require.NotNil(t, diff)
				require.Equal(t, tt.yaml, diff.Attributes["yaml"].New, "CustomizeDiff must not rewrite pipeline YAML")
			}
		})
	}
}

func TestResourcePipelineCustomizeDiffDoesNotSetDescriptionFromYaml(t *testing.T) {
	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"identifier": "pipeline_id",
		"name":       "Pipeline name",
		"org_id":     "org",
		"project_id": "project",
		"yaml": `pipeline:
  identifier: pipeline_id
  name: Pipeline name
  description: YAML description
`,
	})

	diff, err := ResourcePipeline().Diff(context.Background(), nil, config, nil)
	require.NoError(t, err)
	require.NotNil(t, diff)
	require.NotContains(t, diff.Attributes, "description")
}
