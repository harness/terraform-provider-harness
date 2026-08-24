package applications

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBuildApplicationSpecFromMapIgnoreDifference covers CDS-130253, where spec
// ignore_difference was accepted by the schema but never included in the payload.
func TestBuildApplicationSpecFromMapIgnoreDifference(t *testing.T) {
	spec := BuildApplicationSpecFromMap(map[string]interface{}{
		"project": "default",
		"ignore_difference": []interface{}{
			map[string]interface{}{
				"kind":          "Service",
				"json_pointers": []interface{}{"/spec/selector/rollouts-pod-template-hash"},
			},
		},
	})

	require.Len(t, spec.IgnoreDifferences, 1)
	assert.Equal(t, "Service", spec.IgnoreDifferences[0].Kind)
	assert.Equal(t, []string{"/spec/selector/rollouts-pod-template-hash"}, spec.IgnoreDifferences[0].JsonPointers)
}

// TestBuildAppSpecMapIgnoreDifference asserts ignore_difference returned by the API is
// written back into the spec map for state.
func TestBuildAppSpecMapIgnoreDifference(t *testing.T) {
	specMap := BuildAppSpecMap(&nextgen.ApplicationsApplicationSpec{
		Project: "default",
		IgnoreDifferences: []nextgen.ApplicationsResourceIgnoreDifferences{
			{
				Group:             "apps",
				Kind:              "Deployment",
				Name:              "my-deploy",
				Namespace:         "default",
				JqPathExpressions: []string{".spec.replicas"},
			},
		},
	})

	ignoreDiff, ok := specMap["ignore_difference"]
	require.True(t, ok, "ignore_difference must be present in spec map")
	ignoreList, ok := ignoreDiff.([]interface{})
	require.True(t, ok)
	require.Len(t, ignoreList, 1)

	diffMap, ok := ignoreList[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "apps", diffMap["group"])
	assert.Equal(t, "Deployment", diffMap["kind"])
	assert.Equal(t, "my-deploy", diffMap["name"])
	assert.Equal(t, "default", diffMap["namespace"])
	assert.Equal(t, []string{".spec.replicas"}, diffMap["jq_path_expressions"])
}

// TestBuildApplicationSpecFromMapWithoutIgnoreDifference guards against sending an empty
// ignore_difference block when the optional field is omitted or empty.
func TestBuildApplicationSpecFromMapWithoutIgnoreDifference(t *testing.T) {
	t.Run("absent", func(t *testing.T) {
		spec := BuildApplicationSpecFromMap(map[string]interface{}{
			"project": "default",
		})
		assert.Nil(t, spec.IgnoreDifferences)
	})

	t.Run("empty list", func(t *testing.T) {
		spec := BuildApplicationSpecFromMap(map[string]interface{}{
			"project":           "default",
			"ignore_difference": []interface{}{},
		})
		assert.Nil(t, spec.IgnoreDifferences)
	})
}

// TestBuildAppSpecMapWithoutIgnoreDifference guards against writing an empty ignore_difference
// block into state when the API returns no ignore differences.
func TestBuildAppSpecMapWithoutIgnoreDifference(t *testing.T) {
	specMap := BuildAppSpecMap(&nextgen.ApplicationsApplicationSpec{
		Project: "default",
	})

	_, ok := specMap["ignore_difference"]
	assert.False(t, ok, "ignore_difference must not be set when IgnoreDifferences is empty")
}
