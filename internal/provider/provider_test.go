package provider

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/micahlmartin/terraform-provider-harness/harness/api"
	"github.com/micahlmartin/terraform-provider-harness/harness/envvar"
	"github.com/stretchr/testify/require"
)

const (
	testAccSecretFileId = "2WnPVgLGSZW6KbApZuxeaw"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"harness": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestDefaultProvider_default_configuration(t *testing.T) {
	// Setup
	p := New("dev")()

	// verify
	// diag := p.Configure(context.TODO(), nil)
	// require.False(t, diag.HasError())
	require.NoError(t, p.InternalValidate())
}

func TestProvider_configure_url(t *testing.T) {

	// Setup provider
	const expectedEndpoint = "http://localhost:8200"
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{"endpoint": expectedEndpoint})
	p := New("dev")()
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
		os.Unsetenv(envvar.HarnessEndpoint)
	}()

	// Configure environment
	os.Setenv(envvar.HarnessEndpoint, expectedEndpoint)

	// Create provider
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{})
	p := New("dev")()
	diags := p.Configure(context.TODO(), rc)

	// Verify
	require.False(t, diags.HasError())
	require.NoError(t, p.InternalValidate())
	c := p.Meta().(*api.Client)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}

func testAccPreCheck(t *testing.T) {
	testAccProviderConfigure.Do(func() {
		testAccProvider = New("dev")()

		config := map[string]interface{}{
			"endpoint":     os.Getenv(envvar.HarnessEndpoint),
			"account_id":   os.Getenv(envvar.HarnessAccountId),
			"api_key":      os.Getenv(envvar.HarnessApiKey),
			"bearer_token": os.Getenv(envvar.HarnessBearerToken),
		}

		testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(config))
	})
}

var testAccProvider *schema.Provider
var testAccProviderConfigure sync.Once

func testAccGetResource(resourceName string, state *terraform.State) *terraform.ResourceState {
	rm := state.RootModule()
	return rm.Resources[resourceName]
}

func testAccGetApiClientFromProvider() *api.Client {
	return testAccProvider.Meta().(*api.Client)
}
