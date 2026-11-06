package pipeline

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestBuildPipelineRequestSendsYamlVersion(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourcePipeline().Schema, map[string]interface{}{
		"identifier":   "id",
		"name":         "name",
		"yaml":         "version: 1\n",
		"yaml_version": "1",
	})
	require.NoError(t, d.Set("tags", schema.NewSet(schema.HashString, []interface{}{})))

	createBody, diags := buildCreatePipeline(d)
	require.False(t, diags.HasError())
	require.Equal(t, "1", createBody.Version)

	createJSON, err := json.Marshal(createBody)
	require.NoError(t, err)
	require.Contains(t, string(createJSON), `"version":"1"`)

	updateBody := buildUpdatePipeline(d)
	require.Equal(t, "1", updateBody.Version)

	updateJSON, err := json.Marshal(updateBody)
	require.NoError(t, err)
	require.Contains(t, string(updateJSON), `"version":"1"`)
}

func TestBuildPipelineRequestDefaultsYamlVersionToZero(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourcePipeline().Schema, map[string]interface{}{
		"identifier": "id",
		"name":       "name",
		"yaml":       "pipeline:\n  identifier: id\n",
	})
	require.NoError(t, d.Set("tags", schema.NewSet(schema.HashString, []interface{}{})))

	createBody, diags := buildCreatePipeline(d)
	require.False(t, diags.HasError())
	require.Equal(t, "0", createBody.Version)

	createJSON, err := json.Marshal(createBody)
	require.NoError(t, err)
	require.Contains(t, string(createJSON), `"version":"0"`)

	updateJSON, err := json.Marshal(buildUpdatePipeline(d))
	require.NoError(t, err)
	require.Contains(t, string(updateJSON), `"version":"0"`)
}
