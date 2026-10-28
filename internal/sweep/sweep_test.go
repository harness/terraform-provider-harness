package sweep_test

import (
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/application"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/cloudprovider"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/secrets"
	_ "github.com/harness/terraform-provider-harness/internal/service/split"
	"github.com/harness/terraform-provider-harness/internal/sweep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	// Sweepers need FirstGen ApiKey. `make test` / coverage CI have no credentials —
	// skip client setup so the package does not panic; `make testacc` sets TF_ACC=1.
	if os.Getenv(resource.TestEnvVar) == "" {
		os.Exit(m.Run())
	}

	client, err := cd.NewClient(cd.DefaultConfig())
	if err != nil {
		panic(err)
	}

	sweep.SweeperClient = client
	resource.TestMain(m)
}
