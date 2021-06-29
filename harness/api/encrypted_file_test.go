package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEncryptedFileName(t *testing.T) {
	client := getClient()
	expectedName := "secretfile"
	s, err := client.Secrets().GetEncryptedFileByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
}

func TestGetEncryptedFileById(t *testing.T) {
	client := getClient()
	expectedId := secretFileId
	s, err := client.Secrets().GetEncryptedFileById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
}
