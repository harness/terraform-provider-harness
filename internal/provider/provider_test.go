package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/stretchr/testify/require"
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
	c := p.Meta().(*client.ApiClient)
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
	c := p.Meta().(*client.ApiClient)
	require.Equal(t, expectedEndpoint, c.Endpoint)
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func testProvider() string {
	return fmt.Sprintf(`
provider "harness" {
  endpoint 		= %[1]q
  api_key    	= %[2]q
  account_id 	= %[3]q
}
`, os.Getenv(envvar.HarnessEndpoint), os.Getenv(envvar.HarnessApiKey), os.Getenv(envvar.HarnessAccountId))
}
