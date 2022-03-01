package utils_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestGetConflictsWithSlice(t *testing.T) {
	source := []string{"one", "two", "three", "four"}

	expectedItemCount := len(source)

	result := utils.GetConflictsWithSlice(source, "three")
	require.Len(t, result, expectedItemCount-1)
	require.Contains(t, result, "one")
	require.Contains(t, result, "two")
	require.Contains(t, result, "four")

	result = utils.GetConflictsWithSlice(source, "one")
	require.Len(t, result, expectedItemCount-1)
	require.Contains(t, result, "two")
	require.Contains(t, result, "three")
	require.Contains(t, result, "four")

	require.Len(t, source, expectedItemCount)
}
