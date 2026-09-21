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

// TestBuildApplicationSpecFromMapNilBlocks covers CDS-132670. An empty nested block such as
// `sync_policy { automated {} }` is expanded by Terraform SDK v2 into []interface{}{nil}, which
// passes a length check but is not a map, so unguarded assertions crashed the provider.
func TestBuildApplicationSpecFromMapNilBlocks(t *testing.T) {
	nilBlock := []interface{}{nil}

	cases := map[string]map[string]interface{}{
		"sync_policy":           {"sync_policy": nilBlock},
		"sync_policy.automated": {"sync_policy": []interface{}{map[string]interface{}{"automated": nilBlock}}},
		"sync_policy.retry":     {"sync_policy": []interface{}{map[string]interface{}{"retry": nilBlock}}},
		"retry.backoff":         {"sync_policy": []interface{}{map[string]interface{}{"retry": []interface{}{map[string]interface{}{"backoff": nilBlock}}}}},
		"destination":           {"destination": nilBlock},
		"source":                {"source": nilBlock},
		"sources":               {"sources": nilBlock},
		"ignore_difference":     {"ignore_difference": nilBlock},
		"source.helm":           {"source": []interface{}{map[string]interface{}{"helm": nilBlock}}},
		"source.kustomize":      {"source": []interface{}{map[string]interface{}{"kustomize": nilBlock}}},
		"source.ksonnet":        {"source": []interface{}{map[string]interface{}{"ksonnet": nilBlock}}},
		"source.directory":      {"source": []interface{}{map[string]interface{}{"directory": nilBlock}}},
		"directory.jsonnet":     {"source": []interface{}{map[string]interface{}{"directory": []interface{}{map[string]interface{}{"jsonnet": nilBlock}}}}},
		"source.plugin":         {"source": []interface{}{map[string]interface{}{"plugin": nilBlock}}},
	}

	for name, specData := range cases {
		t.Run(name, func(t *testing.T) {
			specData["project"] = "default"
			var spec nextgen.ApplicationsApplicationSpec
			require.NotPanics(t, func() {
				spec = BuildApplicationSpecFromMap(specData)
			})
			assert.Equal(t, "default", spec.Project)
		})
	}

	t.Run("nil entry is skipped not sent", func(t *testing.T) {
		spec := BuildApplicationSpecFromMap(map[string]interface{}{
			"project": "default",
			"sources": []interface{}{nil, map[string]interface{}{"repo_url": "https://example.com/repo.git"}},
		})
		require.Len(t, spec.Sources, 1)
		assert.Equal(t, "https://example.com/repo.git", spec.Sources[0].RepoURL)
	})

	t.Run("missing project", func(t *testing.T) {
		require.NotPanics(t, func() {
			BuildApplicationSpecFromMap(map[string]interface{}{})
		})
	})
}

// TestBuildApplicationSpecFromMapPopulatedBlocks asserts the CDS-132670 nil guards did not
// change how well formed blocks are expanded.
func TestBuildApplicationSpecFromMapPopulatedBlocks(t *testing.T) {
	spec := BuildApplicationSpecFromMap(map[string]interface{}{
		"project": "default",
		"sources": []interface{}{
			map[string]interface{}{"repo_url": "https://example.com/a.git", "target_revision": "main", "ref": "values"},
			map[string]interface{}{"repo_url": "https://example.com/b.git", "path": "chart"},
		},
		"destination": []interface{}{
			map[string]interface{}{"namespace": "default", "server": "https://kubernetes.default.svc"},
		},
		"sync_policy": []interface{}{
			map[string]interface{}{
				"automated": []interface{}{
					map[string]interface{}{"prune": true, "self_heal": true, "allow_empty": false},
				},
				"retry": []interface{}{
					map[string]interface{}{
						"limit":   "5",
						"backoff": []interface{}{map[string]interface{}{"duration": "5s", "factor": "2", "max_duration": "3m"}},
					},
				},
			},
		},
	})

	require.Len(t, spec.Sources, 2)
	assert.Equal(t, "https://example.com/a.git", spec.Sources[0].RepoURL)
	assert.Equal(t, "values", spec.Sources[0].Ref)
	assert.Equal(t, "chart", spec.Sources[1].Path)

	require.NotNil(t, spec.Destination)
	assert.Equal(t, "https://kubernetes.default.svc", spec.Destination.Server)

	require.NotNil(t, spec.SyncPolicy)
	require.NotNil(t, spec.SyncPolicy.Automated)
	assert.True(t, spec.SyncPolicy.Automated.Prune)
	assert.True(t, spec.SyncPolicy.Automated.SelfHeal)
	assert.False(t, spec.SyncPolicy.Automated.AllowEmpty)

	require.NotNil(t, spec.SyncPolicy.Retry)
	assert.Equal(t, "5", spec.SyncPolicy.Retry.Limit)
	require.NotNil(t, spec.SyncPolicy.Retry.Backoff)
	assert.Equal(t, "5s", spec.SyncPolicy.Retry.Backoff.Duration)
	assert.Equal(t, "3m", spec.SyncPolicy.Retry.Backoff.MaxDuration)
}
