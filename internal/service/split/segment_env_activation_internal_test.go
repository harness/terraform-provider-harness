package split

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitSegmentEnvKeysErrLooksNotFound(t *testing.T) {
	t.Parallel()
	require.False(t, splitSegmentEnvKeysErrLooksNotFound(nil))
	require.True(t, splitSegmentEnvKeysErrLooksNotFound(errors.New("get segment keys: 404 404 Not Found: {}")))
	require.True(t, splitSegmentEnvKeysErrLooksNotFound(errors.New("not found in workspace")))
	require.False(t, splitSegmentEnvKeysErrLooksNotFound(errors.New("get segment keys: 500 Internal Server Error")))
}
