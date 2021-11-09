package acctest

import (
	"context"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/terraform-provider-harness/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestDefaultProvider_default_configuration(t *testing.T) {
	// Setup
	p := provider.New("dev")()

	// verify
	// diag := p.Configure(context.TODO(), nil)
	// require.False(t, diag.HasError())
	require.NoError(t, p.InternalValidate())
}

func TestProvider_configure_url(t *testing.T) {

	// Setup provider
	const expectedEndpoint = "http://localhost:8200"
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{"endpoint": expectedEndpoint})
	p := provider.New("dev")()
	diags := p.Configure(context.TODO(), rc)

	// Verify
	require.False(t, diags.HasError())
	require.NoError(t, p.InternalValidate())
	c := p.Meta().(*api.Client)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}

func TestProvider_configure_url_env(t *testing.T) {

	// Setup
	const expectedEndpoint = "http://localhost:8200"

	// Cleanup function
	defer func() {
		os.Unsetenv(helpers.EnvVars.HarnessEndpoint.String())
	}()

	// Configure environment
	os.Setenv(helpers.EnvVars.HarnessEndpoint.String(), expectedEndpoint)

	// Create provider
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{})
	p := provider.New("dev")()
	diags := p.Configure(context.TODO(), rc)

	// Verify
	require.False(t, diags.HasError())
	require.NoError(t, p.InternalValidate())
	c := p.Meta().(*api.Client)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}
