package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSSOProviderById_LDAP(t *testing.T) {
	client := getClient()
	id := "NUXZisowSVSkhBSUxo9CSQ"
	provider, err := client.SSO().GetSSOProviderById(id)
	require.NoError(t, err, "Failed to get SSO provider by id: %s", err)
	require.NotNil(t, provider, "SSO provider should not be nil")
}

func TestGetSSOProviderById_SAML(t *testing.T) {
	client := getClient()
	id := "livHmA12TZSo5k1RwcdjOw"
	provider, err := client.SSO().GetSSOProviderById(id)
	require.NoError(t, err, "Failed to get SSO provider by id: %s", err)
	require.NotNil(t, provider, "SSO provider should not be nil")
}

func TestGetSSOProviderByName_LDAP(t *testing.T) {
	client := getClient()
	name := "ldap-test"

	provider, err := client.SSO().GetSSOProviderByName(name)
	require.NoError(t, err, "Failed to get SSO provider by name: %s", err)
	require.NotNil(t, provider, "SSO provider should not be nil")
}

func TestGetSSOProviderByName_SAML(t *testing.T) {
	client := getClient()
	name := "saml-test"

	provider, err := client.SSO().GetSSOProviderByName(name)
	require.NoError(t, err, "Failed to get SSO provider by name: %s", err)
	require.NotNil(t, provider, "SSO provider should not be nil")
}

func TestListSSOProviders(t *testing.T) {
	client := getClient()
	limit := 10
	offset := 0
	hasMore := true

	for hasMore {
		providers, pagination, err := client.SSO().ListSSOProviders(limit, offset)
		require.NoError(t, err, "Failed to list SSO providers: %s", err)
		require.NotEmpty(t, providers, "No SSO providers found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(providers) == limit
		offset += limit
	}
}
