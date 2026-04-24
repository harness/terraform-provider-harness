package split

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitDefinitionJSONSemanticallyEqual_keyOrder(t *testing.T) {
	t.Parallel()
	// Same fields, different key order (like HCL jsonencode vs encoding/json field order).
	a := `{"defaultTreatment":"off","treatments":[{"name":"on"},{"name":"off"}],"defaultRule":[{"treatment":"off","size":100}],"trafficAllocation":100,"rules":[]}`
	b := `{"treatments":[{"name":"on"},{"name":"off"}],"rules":[],"defaultRule":[{"treatment":"off","size":100}],"defaultTreatment":"off","trafficAllocation":100}`
	require.True(t, splitDefinitionJSONSemanticallyEqual(a, b))
}

func TestSplitDefinitionJSONSemanticallyEqual_notEqual(t *testing.T) {
	t.Parallel()
	a := `{"treatments":[{"name":"on"}],"defaultRule":[],"defaultTreatment":"off","trafficAllocation":100}`
	b := `{"treatments":[{"name":"on"}],"defaultRule":[],"defaultTreatment":"on","trafficAllocation":100}`
	require.False(t, splitDefinitionJSONSemanticallyEqual(a, b))
}

func TestSplitDefinitionMergePresentationFromPrior_titleComment(t *testing.T) {
	t.Parallel()
	api := `{"treatments":[{"name":"on"},{"name":"off"}],"defaultRule":[{"treatment":"off","size":100}],"defaultTreatment":"off","trafficAllocation":100,"rules":[]}`
	prior := `{"title":"T","comment":"C","treatments":[{"name":"on"},{"name":"off"}],"defaultRule":[{"treatment":"off","size":100}],"defaultTreatment":"off","trafficAllocation":100,"rules":[]}`
	out, err := splitDefinitionMergePresentationFromPrior(api, prior)
	require.NoError(t, err)
	require.True(t, splitDefinitionJSONSemanticallyEqual(prior, out), "merged should match prior semantically")
	var req struct {
		Title, Comment string
	}
	require.NoError(t, json.Unmarshal([]byte(out), &req))
	require.Equal(t, "T", req.Title)
	require.Equal(t, "C", req.Comment)
}

func TestSplitDefinitionMergePresentationFromPrior_malformedPrior(t *testing.T) {
	t.Parallel()
	api := `{"treatments":[],"defaultRule":[],"defaultTreatment":"off","trafficAllocation":100}`
	out, err := splitDefinitionMergePresentationFromPrior(api, "{")
	require.NoError(t, err)
	require.Equal(t, api, out)
}

func TestSplitDefinitionRequestFromString_invalidJSON(t *testing.T) {
	t.Parallel()
	_, err := splitDefinitionRequestFromString("{")
	if err == nil {
		t.Fatal("expected error for truncated JSON")
	}
	_, err = splitDefinitionRequestFromString("clearly-not-json")
	if err == nil {
		t.Fatal("expected error for non-JSON")
	}
}

func TestFmeFeatureFlagDefinitionID(t *testing.T) {
	t.Parallel()
	got := fmeFeatureFlagDefinitionID("org1", "proj2", "env3", "flag4")
	want := "org1/proj2/env3/flag4"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
