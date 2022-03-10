package sweep_test

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/application"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/cloudprovider"
	_ "github.com/harness/terraform-provider-harness/internal/service/cd/secrets"
	"github.com/harness/terraform-provider-harness/internal/sweep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	client, err := cd.NewClient(cd.DefaultConfig())
	if err != nil {
		panic(err)
	}

	sweep.SweeperClient = client
	resource.TestMain(m)
}
