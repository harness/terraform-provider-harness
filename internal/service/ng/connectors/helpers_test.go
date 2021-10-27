package connectors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_conflictsWith(t *testing.T) {
	result := getConflictsWithSlice("gcp")
	require.Len(t, result, 3)
	require.Contains(t, result, "aws")
	require.Contains(t, result, "docker_registry")
	require.Contains(t, result, "k8s_cluster")
	require.NotContains(t, result, "gcp")

	result = getConflictsWithSlice("aws")
	require.Contains(t, result, "docker_registry")
	require.Contains(t, result, "k8s_cluster")
	require.Contains(t, result, "gcp")
	require.NotContains(t, result, "aws")

	require.Len(t, connectorConfigNames, 4)
}
