package sweep

import (
	"testing"

	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	acctest.TestAccConfigureProvider()
	resource.TestMain(m)
}
