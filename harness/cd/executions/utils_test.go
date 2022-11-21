package executions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutionItem_IsEmpty(t *testing.T) {
	eItem := ExecutionItem{}
	require.Equal(t, eItem.IsEmpty(), false)

	eNilItem := Response{
		Resource: nil,
	}
	require.Equal(t, eNilItem.Resource.IsEmpty(), true)
}

func TestResponseMessage_ToError(t *testing.T) {
	resp := Response{
		Metadata:         nil,
		Resource:         nil,
		ResponseMessages: nil,
	}
	require.Equal(t, resp.IsEmpty(), true)

	resp = Response{
		Metadata:         &ResponseMetadata{},
		Resource:         &ExecutionItem{},
		ResponseMessages: make([]ResponseMessage, 0),
	}
	require.Equal(t, resp.IsEmpty(), false)
}

func TestResponse_IsEmpty(t *testing.T) {
	resp := ResponseMessage{
		Code:    "404",
		Level:   "Warning",
		Message: "Not found",
	}

	err := resp.ToError()
	require.Equal(t, err.Error(), "404: Not found")
}
