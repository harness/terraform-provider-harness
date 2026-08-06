package provider

import (
	"context"
	"testing"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestGetSplitClient_BasePathUsesHarnessEndpoint(t *testing.T) {
	providerSchema := Provider("test")().Schema

	tests := []struct {
		name     string
		endpoint string
		override string
		want     string
	}{
		{
			name:     "default gateway",
			endpoint: "https://app.harness.io/gateway",
			want:     "https://app.harness.io/fme",
		},
		{
			name:     "trailing slash",
			endpoint: "https://qa.harness.io/gateway/",
			want:     "https://qa.harness.io/fme",
		},
		{
			name:     "override wins",
			endpoint: "https://qa.harness.io/gateway",
			override: "https://fme.example.com/custom/",
			want:     "https://fme.example.com/custom/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, providerSchema, map[string]interface{}{
				"endpoint":               tt.endpoint,
				"fme_admin_api_endpoint": tt.override,
				"account_id":             "acc",
				"platform_api_key":       "key",
			})

			basePath, err := resolveFMEAdminAPIEndpoint(tt.endpoint, tt.override)
			require.NoError(t, err)
			require.Equal(t, tt.want, basePath)

			client := getSplitClient(d, "test", basePath)
			require.NotNil(t, client)
			require.Equal(t, tt.want, client.BasePath)
		})
	}
}

func TestProviderConfigure_SplitClientUsesHarnessEndpoint(t *testing.T) {
	t.Setenv("FME_ADMIN_API_ENDPOINT", "")

	p := Provider("test")()
	diags := p.Configure(context.Background(), terraform.NewResourceConfigRaw(map[string]interface{}{
		"endpoint":         "https://app.harness.io/gateway",
		"account_id":       "acc",
		"platform_api_key": "key",
	}))
	require.Empty(t, diags)

	sess := p.Meta().(*internal.Session)
	require.NotNil(t, sess.SplitClient)
	require.Equal(t, "https://app.harness.io/fme", sess.SplitClient.BasePath)
}

func TestProviderConfigure_SplitClientUsesFMEEndpointFromEnvironment(t *testing.T) {
	t.Setenv("FME_ADMIN_API_ENDPOINT", "https://fme.example.com/custom")

	p := Provider("test")()
	diags := p.Configure(context.Background(), terraform.NewResourceConfigRaw(map[string]interface{}{
		"endpoint":         "https://app.harness.io/api",
		"account_id":       "acc",
		"platform_api_key": "key",
	}))
	require.Empty(t, diags)

	sess := p.Meta().(*internal.Session)
	require.NotNil(t, sess.SplitClient)
	require.Equal(t, "https://fme.example.com/custom", sess.SplitClient.BasePath)
}

func TestProviderConfigure_InvalidFMEEndpointDefersError(t *testing.T) {
	t.Setenv("FME_ADMIN_API_ENDPOINT", "")

	p := Provider("test")()
	diags := p.Configure(context.Background(), terraform.NewResourceConfigRaw(map[string]interface{}{
		"endpoint":         "https://app.harness.io/api",
		"account_id":       "acc",
		"platform_api_key": "key",
	}))
	require.Empty(t, diags)

	sess := p.Meta().(*internal.Session)
	require.Nil(t, sess.SplitClient)
	require.Error(t, sess.FMEAdminAPIEndpointError)
	require.Contains(t, sess.FMEAdminAPIEndpointError.Error(), "must end in /gateway")
}
