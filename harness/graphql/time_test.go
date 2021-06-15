package graphql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type testTimeType struct {
	CreatedAt Time `json:"createdAt"`
}

func TestTimeUnmarshalJSON(t *testing.T) {

	ts := &testTimeType{}
	jsonPayload := []byte(`{"createdAt": 1623202185}`)

	err := json.Unmarshal(jsonPayload, &ts)
	require.NoError(t, err)

}
