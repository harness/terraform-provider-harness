package acctest

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/terraform-provider-harness/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	TestAccSecretFileId = "2WnPVgLGSZW6KbApZuxeaw"
)

func TestAccConfigureProvider() {
	TestAccProviderConfigure.Do(func() {
		TestAccProvider = provider.New("dev")()

		config := map[string]interface{}{
			"endpoint":   os.Getenv(helpers.EnvVars.HarnessEndpoint.String()),
			"account_id": os.Getenv(helpers.EnvVars.HarnessAccountId.String()),
			"api_key":    os.Getenv(helpers.EnvVars.HarnessApiKey.String()),
			"ng_api_key": os.Getenv(helpers.EnvVars.HarnessNGApiKey.String()),
		}

		TestAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(config))
	})
}

func TestAccPreCheck(t *testing.T) {
	TestAccConfigureProvider()
}

var TestAccProvider *schema.Provider
var TestAccProviderConfigure sync.Once

func TestAccGetResource(resourceName string, state *terraform.State) *terraform.ResourceState {
	rm := state.RootModule()
	return rm.Resources[resourceName]
}

func TestAccGetApiClientFromProvider() *api.Client {
	return TestAccProvider.Meta().(*api.Client)
}

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var ProviderFactories = map[string]func() (*schema.Provider, error){
	"harness": func() (*schema.Provider, error) {
		return provider.New("dev")(), nil
	},
}
