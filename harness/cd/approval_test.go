package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApprovalDetails(t *testing.T) {
	c := getClient()
	approval, err := c.ApprovalClient.GetApprovalDetails("q_nsFwL5TUq4Jsq00vNGRQ-gf39sw", "ntUxF8b_SRuSWLckvPkODw")
	require.NoError(t, err)
	require.NotNil(t, approval)
}
