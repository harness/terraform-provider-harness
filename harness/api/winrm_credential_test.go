package api

import (
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/stretchr/testify/require"
)

func TestGetWinRMCredentialById(t *testing.T) {
	client := getClient()
	expectedId := "CpiGVJOHSEmzSc39KFVBJg"
	s, err := client.Secrets().GetWinRMCredentialById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
	require.Equal(t, graphql.WinRMAuthenticationSchemes.NTLM, s.AuthenticationScheme)
}

func TestGetWinRMCredentialByName(t *testing.T) {
	client := getClient()
	expectedName := "winrm_ntlm"
	s, err := client.Secrets().GetWinRMCredentialByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.WinRMAuthenticationSchemes.NTLM, s.AuthenticationScheme)
}
